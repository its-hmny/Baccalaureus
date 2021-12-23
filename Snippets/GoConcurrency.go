package main

import (
	"fmt"
	"time"
)

func fuzzer(channel chan int, timeout time.Duration) {
	for i := 0; i <= 10; i++ {
		channel <- i
		time.Sleep(timeout * time.Second)
	}

	// From now on sending on this channel will cause an error
	close(channel) 
}

func main() {
	// Buffered channel => asynchronous communication
	a := make(chan int, 10)
	// Unbuffered channel => synchronous communication
	b := make(chan int)

	// Starts two "fuzzer" processes
	go fuzzer(a, 4)
	go fuzzer(b, 7)

	for { // Iterates until both channels are closed
		select { 
		case data := <-a:
			fmt.Printf("Received from a %d\n", data)
		case data := <-b:
			fmt.Printf("Received from b: %d\n", data)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
