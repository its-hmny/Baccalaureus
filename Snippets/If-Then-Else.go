package main

import (
	"math/rand"
	"time"
)

func dummy(channel chan int) {
	channel <- rand.Int() // Send some random message
	close(channel)        // Closes the cannel before returning
}

func main() {
	// Creates the shared channel
	A, B := make(chan int), make(chan int)

	// Starts the "dummy" processes
	go dummy(A)
	go dummy(B)

	// Sleeps in order to let the other Goroutine send messages
	time.Sleep(5 * time.Second)

	if true { // Since the conditions is always verified
		<-A // Always receive from A
	}

	<-B // Receives a message from B
}
