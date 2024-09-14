package main

import (
	"fmt"
	"time"
)

/* SELECT
Select allows a goroutine to wait on multiple channel operations, proceeding with the first one ready to communicate
select {
case val := <-ch1:
    // Handle data from ch1
case ch2 <- 42:
    // Send data to ch2
default:
    // Optional: handle if no channels are ready
}

- Using a default case makes select non-blocking, allowing the goroutine to move on if no channels are ready
- If multiple channels are ready, Go picks one randomly to avoid bias
- Combining select and time.After enables graceful handling of timeouts when waiting on channels
*/

func select_ch() {

	c1 := make(chan string) //we create two new channels
	c2 := make(chan string)

	go func() { // each channel will receive a value after some amount of time
		time.Sleep(1 * time.Second) // to simulate blocking RPC operations executing in concurrent goroutines
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ { //we will use select to await both of these values simultaneously, printing each one as it arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	// received one
	// received two

	// execution time is only 2 seconds since both the 1 and 2 second Sleeps execute concurrently
}
