package pools

import "sync"

// Pool Goroutine 池.
type Pool[I, O any] struct {
	queue chan *job[I, O]
	yield func(I, int) O
}

// New 新建 Goroutine 池.
func New[I, O any](size int, yield func(I, int) O) *Pool[I, O] {
	pool := &Pool[I, O]{
		make(chan *job[I, O], size),
		yield,
	}

	for num := 0; num < size; num++ {
		go pool.run(num)
	}

	return pool
}

// Run 执行单个任务.
func (p *Pool[I, O]) Run(input I) O {
	jobs := &job[I, O]{input: input, callback: make(chan O)}

	p.queue <- jobs

	return <-jobs.callback
}

// Post 批量任务处理.
func (p *Pool[I, O]) Post(inputs []I) []O {
	jobs := make([]*job[I, O], len(inputs))
	wgp := sync.WaitGroup{}

	wgp.Add(len(inputs))

	for index, input := range inputs {
		jobs[index] = &job[I, O]{
			wgp:   &wgp,
			input: input,
			index: index,
		}

		p.queue <- jobs[index]
	}

	wgp.Wait()

	res := make([]O, len(inputs))

	for _, elem := range jobs {
		res[elem.index] = elem.output
	}

	return res
}

// Close 关闭协程池.
func (p *Pool[I, O]) Close() {
	close(p.queue)
}

func (p *Pool[I, O]) run(num int) {
	for elem := range p.queue {
		elem.output = p.yield(elem.input, num)
		elem.Done()
	}
}
