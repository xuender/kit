package seq_test

import (
	"fmt"
	"iter"
	"sync"

	"github.com/xuender/kit/v2/cont"
	"github.com/xuender/kit/v2/seq"
)

func ExampleClone() {
	seq1, seq2 := seq.Clone(seq.Range(1, 10))
	group := sync.WaitGroup{}
	set1 := cont.NewSet[int]()
	set2 := cont.NewSet[int]()

	group.Add(2)

	call := func(set cont.Set[int], seq iter.Seq[int]) {
		for item := range seq {
			set.Add(item)
		}

		group.Done()
	}

	go call(set1, seq1)
	go call(set2, seq2)

	group.Wait()

	fmt.Println(set1.Equal(set2))

	// Output:
	// true
}
