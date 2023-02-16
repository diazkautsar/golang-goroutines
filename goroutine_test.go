package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	// ketika di tambahkan go. akan running secara async. Hatihati gunakan goroutine ketika aplikasi selelsai. akan berhenti juga
	// kurang cocok mengunakan function yang punya sebuah return value. karena tidak akan ter catch
	go RunHelloWorld()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	// secara result ini hasilnya akan berubah-ubah tidak berurutan
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
