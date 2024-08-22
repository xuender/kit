package pools

import "sync"

// Pool Goroutine 池.
type Pool[I, O any] struct {
	queue chan *job[I, O]
	yield func(I, int) O
	jobs  *SyncPool[*job[I, O]]
}

// New 新建 Goroutine 池.
func New[I, O any](size int, yield func(I, int) O) *Pool[I, O] {
	pool := &Pool[I, O]{
		make(chan *job[I, O], size),
		yield,
		NewSyncPool(
			func() *job[I, O] { return &job[I, O]{} },
			func(j *job[I, O]) {
				j.callback = nil
				j.wgp = nil
				j.index = 0
			},
		),
	}

	for num := range size {
		go pool.run(num)
	}

	return pool
}

// Run 执行单个任务.
func (p *Pool[I, O]) Run(input I) O {
	job := p.jobs.Get()
	defer p.jobs.Put(job)

	job.callback = make(chan O)
	job.input = input

	p.queue <- job

	return <-job.callback
}

// Post 批量任务处理.
func (p *Pool[I, O]) Post(inputs []I) []O {
	var (
		jobs = make([]*job[I, O], len(inputs))
		wgp  = sync.WaitGroup{}
	)

	wgp.Add(len(inputs))

	for index, input := range inputs {
		job := p.jobs.Get()

		job.wgp = &wgp
		job.index = index
		job.input = input
		jobs[index] = job

		p.queue <- job
	}

	wgp.Wait()

	res := make([]O, len(inputs))

	for _, job := range jobs {
		res[job.index] = job.output

		p.jobs.Put(job)
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
