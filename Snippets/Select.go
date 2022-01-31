package main

import "math/rand"

func responder(channel chan int) {
	// ...
	channel <- rand.Int()
	// ...
}

func main() {
	// Creates the channels
	chanA := make(chan int)
	chanB := make(chan int)

	// Starts the worker processes
	go responder(chanA)
	go responder(chanB)

	// Select from both channels
	select {
	case <-chanA:
		<-chanB
		// ...
	case <-chanB:
		<-chanA
		// ...
	}
}
