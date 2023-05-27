package main

import (
	"fmt"
	"github.com/theorx/go_routine_pool"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	rp := routinepool.NewPool(256)
	finished := &atomic.Int64{}

	wg := &sync.WaitGroup{}

	for i := 0; i < 1024; i++ {

		wg.Add(1)
		go func(wg *sync.WaitGroup, job int) {
			rp.Allocate()
			defer wg.Done()
			defer rp.Release()

			//Perform expensive work here
			time.Sleep(time.Millisecond * 30)

			if job%128 == 0 {
				fmt.Printf("%d Workers busy out of %d. Tasks completed: %d\n", rp.Used(), rp.Size(), finished.Load())
			}

			finished.Add(1)
		}(wg, i)
	}

	fmt.Println("Work has been scheduled")

	wg.Wait()
	fmt.Printf("%d Workers busy out of %d. Tasks completed: %d\n", rp.Used(), rp.Size(), finished.Load())
}

/*

go run example/main.go
Work has been scheduled
256 Workers busy out of 256. Tasks completed: 3
256 Workers busy out of 256. Tasks completed: 32
256 Workers busy out of 256. Tasks completed: 135
256 Workers busy out of 256. Tasks completed: 279
256 Workers busy out of 256. Tasks completed: 499
256 Workers busy out of 256. Tasks completed: 684
256 Workers busy out of 256. Tasks completed: 681
176 Workers busy out of 256. Tasks completed: 848
0 Workers busy out of 256. Tasks completed: 1024

*/
