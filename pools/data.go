package pools

import "sync"

type data[I, O any] struct {
	chans chan *job[I, O]
	yield func(I, int) O
}

func (p *data[I, O]) run(num int) {
	for input := range p.chans {
		input.output = p.yield(input.input, num)
		input.Done()
	}
}

func (p *data[I, O]) Run(input I) O {
	jobs := &job[I, O]{input: input, callback: make(chan O)}

	p.chans <- jobs

	return <-jobs.callback
}

// Post 批量任务处理.
func (p *data[I, O]) Post(inputs []I) []O {
	jobs := make([]*job[I, O], len(inputs))
	wgp := sync.WaitGroup{}

	wgp.Add(len(inputs))

	for index, input := range inputs {
		jobs[index] = &job[I, O]{
			wgp:   &wgp,
			input: input,
			index: index,
		}

		p.chans <- jobs[index]
	}

	wgp.Wait()

	res := make([]O, len(inputs))

	for _, job := range jobs {
		res[job.index] = job.output
	}

	return res
}
