package main

import "fmt"

func for_loop() {

	i := 1
	for i <= 3 { // single condition
		fmt.Println(i)
		i = i + 1
	}
	/*
		1
		2
		3
	*/

	for j := 0; j < 3; j++ { // initial/condition/after for loop
		fmt.Println(j)
	}
	/*
		0
		1
		2
	*/

	for i := range 3 { // range over an integer - "do this N times" iteration
		fmt.Println("range", i)
	}
	/*
		range 0
		range 1
		range 2
	*/

	for { // for without condition will loop repeatedly until we break out of it or return from enclosing function
		fmt.Println("loop")
		break
	}
	/*
		loop
	*/

	for n := range 6 { // we can also continue to the next iteration of the loop
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	/*
		1
		3
		5
	*/

}
