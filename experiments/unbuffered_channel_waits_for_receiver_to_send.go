package main

import (
	"fmt"
	"time"
)

func send(channel chan int, value int) {
	fmt.Printf("Attempting to send %d...\n", value)
	channel <- value
	time.Sleep(1 * time.Millisecond) // Added to allow switch to other goroutines for logging order purposes
	fmt.Printf("Sent %d\n", value)
}

func main() {
	ch := make(chan int)
	go send(ch, 1)
	go send(ch, 2)

	// Ensure the goroutines have had time to attempt the send
	time.Sleep(1 * time.Second)

	fmt.Println("Ready to receive...")

	x := <-ch
	time.Sleep(1 * time.Millisecond) // Added to allow switch to other goroutines for logging order purposes
	fmt.Printf("Received %d\n", x)

	y := <-ch
	time.Sleep(1 * time.Millisecond) // Added to allow switch to other goroutines for logging order purposes
	fmt.Printf("Received %d\n", y)
}

/*
Output:

Attempting to send 1...
Attempting to send 2...
Ready to receive...
Sent 1
Received 1
Sent 2
Received 2

Explanations:
- What matters: the `Ready to receive...` occurs before `Sent X`. This is because sending to an unbuffered channel can
only be done when there is a receiver ready. Until the receiver is ready, the goroutine is blocked.
- NB: Sometimes, the order between Sent/Received is in the other way round.
This is due to how Go schedules goroutines: Since receiving from the channel and printing "Received" happens almost
back-to-back in the main goroutine, these actions can complete before the scheduler switches back to the sending
goroutine to allow it to execute the print statement for "Sent".
*/
