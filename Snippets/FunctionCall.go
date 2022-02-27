package main

import "fmt"

func baz(channel chan string) {
	channel <- "Hello from dummy" //Sends a message on the shared channel
}

func f(channel chan string) {
	channel <- "Hello from nested" // Send a message on channel
	go baz(channel)                // Spawns a new "dummy" Goroutine
	fmt.Println(<-channel)         // Receives the message sent by itself
}

func main() {
	// Creates the shared channel
	channel := make(chan string, 1)
	// Call the "f" function
	f(channel)
	// Receives something from "channel"
	fmt.Println(<-channel)
}
