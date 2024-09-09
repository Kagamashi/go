package main

import (
    "fmt"
    "time"
)

/* TIMERS
Fire once after a delay.

- Creating a timer
timer := time.NewTimer(2 * time.Second)
- Receiving from timer
<-timer.C
- time.After shortcut returns a channel that sends the current time after a duration (timeouts)
<-time.After(2 * time.Second)
- Stopping a timer
timer.Stop()
*/

func timers() {

    timer1 := time.NewTimer(2 * time.Second)  // timers represent a single event in the future
		// we tell the timer how long we want to wait, and it provides a channel that will be notified at that time

    <-timer1.C  // <-timer1.C blocks on the timer's channel C until it sends a value indicating that the timer fired
    fmt.Println("Timer 1 fired")

    timer2 := time.NewTimer(time.Second)  // if we just want to wait we can use time.Sleep
    go func() {  //t timer can be useful because it can be canceled before it fires
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    time.Sleep(2 * time.Second) // give the timer2 enough time to fire, if it ever was going to show it is in fact stopped
		// Timer 1 fired
		// Timer 2 stopped

		/*
		The first timer will fire 2s after we start the program, 
		but the second should be stopeed before it has a chance to fire.
		*/

}
