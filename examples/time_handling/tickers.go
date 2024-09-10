package main

import (
	"fmt"
	"time"
)

/* TICKERS
Trigger events at regular intervals.

- Creating ticker
ticker := time.NewTicker(500 * time.Millisecond)
- Receiving from a ticker
for t := range ticker.C {
	fmt.Println("Tick at", t)
}
- Stopping a ticker
ticker.Stop()
*/

func tickers() {

	ticker := time.NewTicker(500 * time.Millisecond) // tickers use simillar mechanism to timers
	done := make(chan bool)                          // a channel that is sent values

	go func() {
		for {
			select { // we use select on the channel to await the values as they arrive every 500ms
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond) //tickers can be stopped like timers
	ticker.Stop()                       // once a ticker is stopped it won't receive any more valiues on its channel
	done <- true
	fmt.Println("Ticker stopped")
	// Tick at 2012-09-23 11:29:56.487625 -0700 PDT
	// Tick at 2012-09-23 11:29:56.988063 -0700 PDT
	// Tick at 2012-09-23 11:29:57.488076 -0700 PDT
	// Ticker stopped

	/* When we run the program the ticker should tick 3 times before we stop it */
}
