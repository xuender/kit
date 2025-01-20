package seq

import (
	"iter"

	"github.com/xuender/kit/v2/types"
)

// Send sends items from the input sequence to the provided output channel.
//
// Play: https://go.dev/play/p/WTYoREBeZd3
func Send[V any](input iter.Seq[V], output chan<- V) {
	defer func() {
		_ = recover()
	}()

	for item := range input {
		output <- item
	}
}

// Send2 sends items from the input sequence to the provided output channel.
func Send2[K, V any](input iter.Seq2[K, V], output chan types.Tuple[K, V]) {
	defer func() {
		_ = recover()
	}()

	for key, val := range input {
		output <- types.T(key, val)
	}
}
