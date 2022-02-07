package main

func fuzzer(channel chan string) {
	channel <- "Hello from fuzzzer"
	close(channel) // Closes the cannel before returning
}

func nested(channel chan string) {
	channel <- "Hello from nested"
	go fuzzer(channel)
	<-channel
}

func main() {
	// Creates the shared channel
	channel := make(chan string)

	nested(channel)
	go fuzzer(channel)

	<-channel
}
