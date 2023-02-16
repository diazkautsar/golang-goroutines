package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// mengirim
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "DJALAL KURNIA"
}

// menerima
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// untuk kirim data ke channel
	channel <- "DJALAL"

	// untuk menerima data. dari channel ke variable
	// cara pertama
	var dataSatu string
	dataSatu = <-channel
	// cara kedua
	dataDua := <-channel

	fmt.Println(dataSatu, dataDua)
	// atau bisa kirim langsung sebagai parameter
	fmt.Println(<-channel)

	go func() {
		time.Sleep(2 * time.Second)
		// pastikan ada yang kirim
		channel <- "Djalal Kurnia"
		fmt.Println("Selesai Kirim Data")
	}()

	// pastikan ada yang nerima
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func TestBufferredChannel(t *testing.T) {
	channel := make(chan string, 2) // angka 2 disini untuk set bufferednya.
	defer close(channel)

	channel <- "DJALAL"
	channel <- "KURNIA"
	fmt.Println("SELESAI")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel) // wajib set close ini. agar for rangenya tidak deadlock
	}()

	for data := range channel {
		fmt.Println("menerima data", data)
	}

	fmt.Println("SELESAI")
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "DJALAL KURNIA"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// dengan tanpa for ini akan menerima data yang sudah ada aja
	var counter int = 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// dengan tanpa for ini akan menerima data yang sudah ada aja
	var counter int = 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default: // set default. do something ketika datanya belum masuk ke select
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
