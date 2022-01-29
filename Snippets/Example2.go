package main

import (
	"fmt"
	"math/rand"
	"time"
)

type msg struct {
	data      int
	timestamp int64
}

// Receives the message and sends it back
func worker(from, to chan msg) {
	for num := range from {
		to <- msg{
			data:      num.data + 1,
			timestamp: time.Now().Unix()}
	}
}

func main() {
	// Creates the channels
	in := make(chan msg, 10)
	out := make(chan msg, 10)

	// Starts the worker processes
	go worker(in, out)
	go worker(in, out)

	// Infinite loop
	for {
		in <- msg{rand.Int(), -1}
		res := <-out
		fmt.Printf("Received %d at %d \n", res.data, res.timestamp)
	}
}
