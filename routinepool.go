package routinepool

type routinePool struct {
	workers chan struct{}
	size    int
}

func NewPool(workers int) *routinePool {
	if workers < 1 {
		workers = 1
	}

	return &routinePool{
		workers: make(chan struct{}, workers),
		size:    workers,
	}
}

func (r *routinePool) Allocate() {
	r.workers <- struct{}{}
}

func (r *routinePool) Release() {
	<-r.workers
}

func (r *routinePool) Size() int {
	return r.size
}

func (r *routinePool) Used() int {
	return len(r.workers)
}
