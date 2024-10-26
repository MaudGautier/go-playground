package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func send(wg *sync.WaitGroup, channel chan int, value int) {
	defer wg.Done()
	fmt.Printf("Attempting to send %d...\n", value)
	channel <- value
	fmt.Printf("Sent %d\n", value)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Check for command-line arguments
	waitFlag := false
	if len(os.Args) > 1 && os.Args[1] == "wait" {
		waitFlag = true
	}

	wg.Add(2) // Adding count for two goroutines
	go send(&wg, ch, 1)
	go send(&wg, ch, 2)

	// Ensure the goroutines have time to attempt the send
	time.Sleep(1 * time.Second)
	fmt.Println("Ready to receive...")

	x := <-ch
	fmt.Printf("Received %d\n", x)

	y := <-ch
	fmt.Printf("Received %d\n", y)

	// Conditionally wait based on the command-line argument
	if waitFlag {
		wg.Wait() // Wait for all goroutines to finish
		fmt.Println("Waited for all goroutines.")
	} else {
		fmt.Println("Did not wait for all goroutines.")
	}
}

/*
Output:
- CASE 1 without the wait (Run with `go run .`):
Attempting to send 1...
Attempting to send 2...
Ready to receive...
Received 1
Received 2
Did not wait for all goroutines.
Sent 1

- CASE 2 with the wait (Run with `go run . wait`)
Attempting to send 1...
Attempting to send 2...
Ready to receive...
Received 1
Received 2
Sent 1
Sent 2
Waited for all goroutines.

Explanations:
- Without the wait, one `Sent` line is not printed, but both are printed when we use the wait group.
- As soon as the main goroutine finishes, all children goroutines are killed (even if they haven't finished executing).
- To prevent that, we need to use a `wait group` (`wg.Wait()` to wait for all goroutines to complete)
*/
