package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	cond.Wait() // ketika ini di aktifkan dan tidak di kasih signal, maka dia akan diem saja.
	fmt.Println("Done", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		time.Sleep(1 * time.Second)
	//		cond.Signal() // SIgnal ini satu satu
	//	}
	//}()

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	cond.Broadcast() // ini untuk return langsung semuanya
	//}()

	group.Wait()
}
