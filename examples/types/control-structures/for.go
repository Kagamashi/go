package main

import "fmt"

/* FOR
For is the only looping construct in Go
for i := 0; i < 10; i++ {
	// loop body
}
	
- Loop can be infinite
for {}
- Use 'break' to exit the loop
- Use 'continue' to skip the current iteration
- Multiple variables can be updated in for loop
for i, j := 0, 0; i < 5 && j < 5; j = i+1, j+2 {
	// loop body
}
*/

func for_loop() {
	
	i := 1
	for i <= 3 { // single condition - similar to while loop
		fmt.Println(i)
		i = i + 1
	}
	/* 	1
	2
	3	 */

	for j := 0; j < 3; j++ { // initial/condition/after for loop
		fmt.Println(j)
	}
	/*	0
		1
		2	 */

	for i := range 3 { // range over an integer - "do this N times" iteration
		fmt.Println("range", i)
	}
	/*	range 0
		range 1
		range 2	 */

	for { // for without condition will loop repeatedly until we break out of it or return from enclosing function
		fmt.Println("loop")
		break
	}
	// loop

	for n := range 6 { // we can also continue to the next iteration of the loop
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	/*	1
		3
		5  */

}
