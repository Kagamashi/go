package main

import (
	"fmt"
	"time"
)

/* EPOCH
Unix epoch (1 January 1970) is used as a reference point for representing time as an integer (seconds or milliseconds since the epoch)
time.Now().Unix() - returns current Unix time in seconds
time.Now().UnixMilli() - returns current Unix time in Milliseconds

- Converting Epoch to time.Time:
time.Unix(seconds, nanoseconds)
- Useful for working with timestamps in systems that store time as Unix timestamps (databases, APIs)
*/

func epoch_def() {

	now := time.Now()
	fmt.Println(now)
	// 2024-09-14 21:49:30.421022 +0200 CEST m=+0.000124460

	fmt.Println(now.Unix())      // 1726343370
	fmt.Println(now.UnixMilli()) // 1726343370421
	fmt.Println(now.UnixNano())  // 1726343370421022000

	// Converting integer seconds/nanoseconds since the epoch into corresponding time
	fmt.Println(time.Unix(now.Unix(), 0))     // 2024-09-14 21:49:30 +0200 CEST
	fmt.Println(time.Unix(0, now.UnixNano())) // 2024-09-14 21:49:30.421022 +0200 CEST

}
