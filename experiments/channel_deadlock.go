package main

import (
	"fmt"
	"os"
)

func send(channel chan int, value int) {
	channel <- value
}

func main() {
	ch := make(chan int)

	// Check for command-line arguments
	goroutineFlag := false
	if len(os.Args) > 1 && os.Args[1] == "goroutine" {
		goroutineFlag = true
	}

	// Choose whether to use a goroutine or use the current routine
	if goroutineFlag {
		go send(ch, 1)
	} else {
		send(ch, 1)
	}

	// Receive results and print them
	fmt.Println(<-ch)
}

/*
Output:
- CASE 1: send to channel from an external goroutine (run with `go run . goroutine`)
1

- CASE 2: send to channel from the main routine (run with `go run .`):
fatal error: all goroutines are asleep - deadlock!

Explanations:
- If the value is sent to the channel from a separate thread (child goroutine), it is blocked until the main routine has
a receiver ready for it. Both routines stop until both are ready, and no error occurs.
- If the value is sent to the channel from the main routine (no child goroutine), the main thread is blocked and waits
for a receiver to be ready. Since the receiver is defined later in that same blocked thread, it cannot proceed and we
get a deadlock error.
=> Using a channel necessarily requires having at least 2 threads (one to send, the other to receive).
*/
