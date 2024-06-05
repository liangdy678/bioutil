package bioutil

import (
	"fmt"
	"runtime"
	"sync"
)

type Parallel[T any] struct {
	ch  chan T
	max int
}

func NewParallel[T any]() *Parallel[T] {
	p := new(Parallel[T])
	p.ch = make(chan T)
	p.max = runtime.NumCPU()

	return p
}

func (p *Parallel[T]) SetMax(max int) {

	if max <= 0 {

		return
	}
	p.max = max
}

func (p *Parallel[T]) Run(send func(chan T), handle func(T)) {

	wg := new(sync.WaitGroup)
	wg.Add(p.max)

	fn := func(wg *sync.WaitGroup) {
		for v := range p.ch {
			fmt.Println("处理数据:", v)
			handle(v)
		}
		wg.Done()
	}

	for i := 0; i < p.max; i++ {
		go fn(wg)
	}

	send(p.ch)
	close(p.ch)

	wg.Wait()
}
