package main

import "math/rand"

func foo(channel chan int) {
	// ... do something
	channel <- rand.Int()
	// ... do something
}

func main() {
	// Creates the channels
	chanA, chanB := make(chan int), make(chan int)

	// Starts the "responder" processes
	go foo(chanA)
	go foo(chanB)

	// Select from both channels
	select {
	case <-chanA:
		<-chanB
		// ... do something
	case <-chanB:
		<-chanA
		// ... do something
	}
}
