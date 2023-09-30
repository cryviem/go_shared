package concurrency

import "sync"

func PackageInfo() string {
	return "concurrency now has: Runnable!"
}

type Runnable struct {
	wg sync.WaitGroup
}

func (r *Runnable) Run(f func()) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		f()
	}()
}

func (r *Runnable) Wait() {
	r.wg.Wait()
}
