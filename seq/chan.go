package seq

import (
	"iter"
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
