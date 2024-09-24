package main

import "fmt"

/* FUNCTIONS
Functions are reusable blocks of code that perform a specific task.
func functionName(parameters) returnType {
    // function body
}

- Parameters are passed by value unless pointers are used
- Go supports multiple return values from a function
- Functions can be anonymous - defined without a name often for short-term purposes
- 'defer' keyword delays function execution until the surrounding function completes
- Return values can be named in the function signature
*/

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int { //Go requires explicit returns i.e. it wont automatially return the value of the last expression
	return a + b + c
}

func vals() (int, int) { //(int, int) means this function returns 2 ints
	return 3, 7
}

func return_values() {

	res := plus(1, 2) //call a function with name(args)
	fmt.Println("1+2 =", res)
	// 1+2 = 3

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	// 1+2+3 = 6

	/* MULTIPLE RETURN VALUES */
	a, b := vals() //2 different return values from the call with multiple assignment
	fmt.Println(a)
	// 3
	fmt.Println(b)
	// 7

	_, c := vals() //if we need only a subset of returned values, use blank identifier _
	fmt.Println(c)
	// 7

}
