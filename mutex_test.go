package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
TestMutex is Race Condition Solution
*/
func TestMutex(t *testing.T) {
	var mutex sync.Mutex
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter:", x)
}
