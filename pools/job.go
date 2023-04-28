package pools

import "sync"

type job[I, O any] struct {
	wgp    *sync.WaitGroup
	input  I
	output O
	index  int
}
