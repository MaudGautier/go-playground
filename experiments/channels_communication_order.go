package main

import (
	"fmt"
	"time"
)

func sendToChannel(channel chan int, value int, delay time.Duration) {
	time.Sleep(delay)
	channel <- value
	fmt.Printf("Sent %d after sleeping for %v\n", value, delay)
}

func main() {
	channel := make(chan int)

	// Start 2 go routines concurrently
	go sendToChannel(channel, 1, 2*time.Second)
	go sendToChannel(channel, 2, 1*time.Second)

	// Receive results
	firstResult := <-channel
	secondResult := <-channel

	// Print
	fmt.Printf("First result: %d, Second result: %d\n", firstResult, secondResult)
}

/*
Output:

Sent 2 after sleeping for 1s
Sent 1 after sleeping for 2s
First result: 2, Second result: 1

Explanations:
- Timing affects ordering in channel communication!
- Messages are sent to the channel in the order in which the goroutines finish their executions and send their results
to the channel.
*/
