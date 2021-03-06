package main

import (
	"fmt"
)

func worker(shared chan string) {
	// Sends 100 numerated messages on the shared channel
	for i := 0; i < 100; i++ {
		shared <- fmt.Sprintf("This is message number: %d", i)
	}
	close(shared) // Close the channel before returning
}

func main() {
	// Creates the shared channel
	channel := make(chan string, 10)
	// Starts the "worker" process
	go worker(channel)
	// Receives until the channel is closed by worker
	for msg := range channel {
		fmt.Printf("Received message: '%s'\n", msg)
	}
}
