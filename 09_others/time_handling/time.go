package main

import (
	"fmt"
	"time"
)

/* TIME
time package in Go provides functions to work with date and time, including getting the current time, formatting, and time arithmetic

time.Now() - returns current local time
time.Add(d time.Duration) - adds duration to a time.Time value
time.Sub(t) - computes the duration between two time.Time values
time.Sleep(d time.Duration) - pauses execution for a specified duration
*/

func time_def() {
	p := fmt.Println

	now := time.Now()
	p(now)
	// 2024-09-14 21:47:57.504068 +0200 CEST m=+0.000124335

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	// 2009-11-17 20:34:58.651387237 +0000 UTC

	p(then.Year())       // 2009
	p(then.Month())      // November
	p(then.Day())        // 17
	p(then.Hour())       // 20
	p(then.Minute())     // 34
	p(then.Second())     // 58
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // UTC

	p(then.Weekday()) // Tuesday

	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	diff := now.Sub(then)
	p(diff)
	// 129959h12m58.852680763s

	p(diff.Hours())       // 129959.21634796688
	p(diff.Minutes())     // 7.797552980878012e+06
	p(diff.Seconds())     // 4.6785317885268074e+08
	p(diff.Nanoseconds()) // 467853178852680763

	p(then.Add(diff))  // 2024-09-14 19:47:57.504068 +0000 UTC
	p(then.Add(-diff)) // 1995-01-20 21:21:59.798706474 +0000 UTC

}
