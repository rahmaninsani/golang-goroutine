package golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
Same as TestRaceCondition
*/
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)

			for j := 1; j <= 100; j++ {
				//x += 1 // race condition
				atomic.AddInt64(&x, 1)
			}

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", x)
}
