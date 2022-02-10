package main

import "math/rand"

func responder(channel chan int) {
	// ... do something else
	channel <- rand.Int()
	// ... do something else
}

func main() {
	// Creates the channels
	chanA, chanB := make(chan int), make(chan int)

	// Starts the "responder" processes
	go responder(chanA)
	go responder(chanB)

	// Select from both channels
	select {
	case <-chanA:
		<-chanB
		// ... do something else
	case <-chanB:
		<-chanA
		// ... do something else
	}
}
