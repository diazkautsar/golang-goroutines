package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// struct mutex di package sync.
// untuk handle race condition.
// lock dan unlock dengan mutex.

func TestRaceCondition(t *testing.T) {
	var x int = 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter ", x)
}
