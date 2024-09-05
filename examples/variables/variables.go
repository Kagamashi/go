package main

import "fmt"

/* VARIABLES
- Go infer type if not explicitly states
- Shorthand syntax inside functions ':='
- Declaration of multiple variables at once:
var x, y int = 1, 2
x, y := 1, 2
- Variable names starting with a capital letter are exported and accessible from other packages;
Lowercase ones are package-private
*/

func variables() {

	var a = "initial" // var declares 1 or more variables
	fmt.Println(a)
	//initial

	var b, c int = 1, 2 // you can declare multiple variables at once
	fmt.Println(b, c)
	//1 2

	var d = true // go will infer the type of initialized variables
	fmt.Println(d)
	//true

	var e int // variables declared without initilization are zero-valued
	fmt.Println(e)
	//0

	f := "apple"   // := syntax is shorthand for declaring and initializing a variable
	fmt.Println(f) // e.g. f string = "apple" - this syntax is only available inside functions
	//apple
}
