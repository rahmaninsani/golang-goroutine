package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for timeValue := range ticker.C {
		fmt.Println(timeValue)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for timeValue := range channel {
		fmt.Println(timeValue)
	}
}
