package pools

import "sync"

type job[I, O any] struct {
	wgp      *sync.WaitGroup
	input    I
	output   O
	index    int
	callback chan O
}

func (p *job[I, O]) Done() {
	if p.wgp != nil {
		p.wgp.Done()
	}

	if p.callback != nil {
		p.callback <- p.output
	}
}
