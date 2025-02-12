package main

import (
	"fmt"
	"time"
)

/* TIME FORMATTING
time package in Go allows formatting and parsing of time values into different string representartions and vice versa
- Go unique reference layout to format time:
Mon Jan 2 15:04:05 MST 2006
- Common Layouts:
2006-01-02 : YYYY-MM-DD (data format)
15:04:05 : HH:MM (24-hour time)
Mon, 02 Jan 2006 15:04:05 MST : full timestamp with day and timezone


- time.Format(layout string) - formats time.Time value into a string according to specified layout
now := time.Now()
formatted := now.Format("2006-01-02 15:04:05") // Format as "YYYY-MM-DD HH:MM:SS"
fmt.Println("Formatted time:", formatted)

- time.Parse(layout, value) - parses a string into a time.Time value based on layout
layout := "2006-01-02"
t, err := time.Parse(layout, "2024-09-12")
if err == nil {
    fmt.Println("Parsed time:", t)
}


*/

func time_formatting() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339)) // Format time according to RFC3339 (ISO 8601 standard)
	// 2024-09-14T21:58:26+02:00

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	// 2012-11-01 22:08:41 +0000 +0000

	p(t.Format("3:04PM")) // Format current time as 12-hour clock with AM/PM notation
	// 9:58PM

	p(t.Format("Mon Jan _2 15:04:05 2006")) // Format current time in more human-readable format
	// Sat Sep 14 21:58:26 2024

	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	// 2024-09-14T21:58:26.704653+02:00

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM") // Define custom time format and parse a string using that format
	p(t2)
	// 0000-01-01 20:41:00 +0000 UTC

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	// 2024-09-14T21:58:26-00:00

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
	// parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"

}
