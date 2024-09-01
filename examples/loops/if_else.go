package main

import "fmt"

func if_else() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { //if can exist without an else
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 { //logical operators
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 { // variables declared in statement are available in the current/subsequent branches
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// 9 has 1 digit

	// parentheses around conditions are not needed in Go
	// but braces are required
}
