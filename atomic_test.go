package golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var groupTestAtomic = &sync.WaitGroup{}
var counterTestAtomic int64 = 0

func TestAtomic(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go func() {
			groupTestAtomic.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counterTestAtomic, 1)
			}
			groupTestAtomic.Done()
		}()
	}

	groupTestAtomic.Wait()
	fmt.Println(counterTestAtomic)
}
