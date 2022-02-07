package main

import (
	"math/rand"
	"time"
)

func dummy(channel chan int) {
	channel <- rand.Int()
	close(channel) // Closes the cannel before returning
}

func main() {
	// Creates the shared channel
	A, B, C := make(chan int), make(chan int), make(chan int)

	// Starts the worker processes
	go dummy(A)
	go dummy(B)
	go dummy(C)

	// Sleeps in order to let the other Goroutine send messages
	time.Sleep(5 * time.Second)

	if rand.Int() >= 10 {
		<-A
	} else if rand.Int() >= 10 {
		<-B
	}

	if rand.Int() <= 10 {
		<-C
	}
}
