package seq

import (
	"iter"
	"time"

	"github.com/xuender/kit/v2/types"
)

// Chan converts a channel into an iterator sequence.
func Chan[V any](input chan V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range input {
			if !yield(item) {
				return
			}
		}
	}
}

func Chan2[K, V any](input chan types.Tuple[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for item := range input {
			if !yield(item.K, item.V) {
				return
			}
		}
	}
}

func ToChans[V any](input iter.Seq[V], size int) []chan V {
	chans := make([]chan V, size)
	for idx := range size {
		chans[idx] = make(chan V)
	}

	go chanRun(input, chans)

	return chans
}

func ToChans2[K, V any](input iter.Seq2[K, V], size int) []chan types.Tuple[K, V] {
	chans := make([]chan types.Tuple[K, V], size)
	for idx := range size {
		chans[idx] = make(chan types.Tuple[K, V])
	}

	go chanRun2(input, chans)

	return chans
}

func chanRun[V any](input iter.Seq[V], chans []chan V) {
	isClose := make([]bool, len(chans))

	for item := range input {
		if chanSend(item, chans, isClose) {
			break
		}
	}

	for idx, cha := range chans {
		if isClose[idx] {
			continue
		}

		close(cha)
	}
}

func chanRun2[K, V any](input iter.Seq2[K, V], chans []chan types.Tuple[K, V]) {
	isClose := make([]bool, len(chans))

	for key, val := range input {
		if chanSend(types.T(key, val), chans, isClose) {
			break
		}
	}

	for idx, cha := range chans {
		if isClose[idx] {
			continue
		}

		close(cha)
	}
}

func allClose(isClose []bool) bool {
	for _, clo := range isClose {
		if !clo {
			return false
		}
	}

	return true
}

func chanSend[V any](item V, chans []chan V, isClose []bool) bool {
	var closeIdx int
	// nolint
	defaultDuration := time.Duration(50) * time.Millisecond

	defer func() {
		if err := recover(); err != nil {
			isClose[closeIdx] = true
		}
	}()

	for {
		if allClose(isClose) {
			return true
		}

		for idx, cha := range chans {
			if isClose[idx] {
				continue
			}

			closeIdx = idx

			select {
			case cha <- item:
				return false
			default:
				continue
			}
		}

		time.Sleep(defaultDuration)
	}
}
