# go-routine-pool

*Simple routine pool using counting semaphore for goroutines*

## Author
* Lauri Orgla

## Usage

* `go get github.com/theorx/go_routine_pool`
* `import "github.com/theorx/go_routine_pool"`
* Main concept of controlling the number of goroutines processing is using `Allocate()` and `Release()` calls
    * Both of them are blocking in case the pool is full (Completely exhausted, busy)

### Example

```go
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

```