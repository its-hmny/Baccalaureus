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

func worker(incoming, outgoing chan msg) {
	for msg := range incoming {
		// ...do something with the message then sends
		// back the results on the out channel
		outgoing <- doSometing(msg)
	}
}

func main() {
	// Creates the channles needed
	in, out := make(chan msg, 10), make(chan msg, 10)
	// Starts the worker processes
	go worker(in, out)
	go worker(in, out)
	// Infinite loop
	for {
		newMsg := msg{rand.Int(), time.Now().Unix()}
		in <- newMsg
		res := <-out
		fmt.Printf("Received %d at :%d", res.data, res.timestamp)
	}
}
