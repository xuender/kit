package seq

import (
	"iter"
	"slices"
	"sync"

	"github.com/xuender/kit/v2/cont"
)

func SplitGroup[S ~[]V, K comparable, V any](slice S, getKey func(V) K) (iter.Seq[V], func(V)) {
	set := cont.NewSyncSet[K]()
	cond := sync.NewCond(&sync.Mutex{})
	group := sync.WaitGroup{}
	split := func(yield func(V) bool) {
		defer func() {
			group.Wait()
		}()

		for len(slice) > 0 {
			if idx := slices.IndexFunc(slice, func(val V) bool {
				return !set.Has(getKey(val))
			}); idx >= 0 {
				item := slice[idx]

				set.Add(getKey(item))
				group.Add(1)

				if !yield(item) {
					return
				}

				slice = append(slice[:idx], slice[idx+1:]...)

				continue
			}

			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
		}
	}
	done := func(val V) {
		set.Del(getKey(val))
		cond.L.Lock()
		cond.Signal()
		cond.L.Unlock()
		group.Done()
	}

	return split, done
}
