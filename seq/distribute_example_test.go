package seq_test

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/xuender/kit/v2/seq"
)

func ExampleDistribute() {
	seqs := seq.Distribute(seq.Range(100), 5)
	wait := sync.WaitGroup{}

	var count int32

	wait.Add(5)

	for _, input := range seqs {
		go func() {
			for range input {
				atomic.AddInt32(&count, 1)
			}

			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println(count)

	// Output:
	// 100
}

func ExampleDistribute2() {
	seqs := seq.Distribute2(seq.Range2(100), 5)
	wait := sync.WaitGroup{}

	var count int32

	wait.Add(5)

	for _, input := range seqs {
		go func() {
			for range input {
				atomic.AddInt32(&count, 1)
			}

			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println(count)

	// Output:
	// 100
}
