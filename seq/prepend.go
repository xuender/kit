package seq

import (
	"iter"

	"github.com/xuender/kit/v2/types"
)

// Prepend adds items to the beginning of the sequence.
//
// It function is useful for prepending items to a sequence in a functional manner.
//
// Play: https://go.dev/play/p/F-H1tK8saqh
func Prepend[V any](input iter.Seq[V], items ...V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}

		for item := range input {
			if !yield(item) {
				return
			}
		}
	}
}

// Prepend2 adds types.Tuples to the beginning of the sequence.
//
// It function is useful for prepending types.Tuples to a sequence in a functional manner.
//
// Play: https://go.dev/play/p/twjNH_C90Nt
func Prepend2[K, V any](input iter.Seq2[K, V], items ...types.Tuple[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, item := range items {
			if !yield(item.K, item.V) {
				return
			}
		}

		for key, val := range input {
			if !yield(key, val) {
				return
			}
		}
	}
}
