package seq

import "iter"

// Clone creates two independent Seq instances by duplicating the elements from the input Seq.
// It uses goroutines and channels to concurrently copy items to two new channels.
// The original sequence is read only once, but its items are sent to both output sequences.
func Clone[T any](input iter.Seq[T]) (iter.Seq[T], iter.Seq[T]) {
	chan1 := make(chan T)
	chan2 := make(chan T)

	go func() {
		for item := range input {
			chan1 <- item
			chan2 <- item
		}

		defer func() { _ = recover() }()

		close(chan1)
		close(chan2)
	}()

	return Chan(chan1), Chan(chan2)
}
