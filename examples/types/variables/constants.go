package main

import (
	"fmt"
	"math"
)

/* CONSTANTS
Constants are individual fixed values
const Pi = 3.14
- Constants are immutable (once declared their values cannot be changed)
- Constants without an explicit type are "untyped" and can be used with various types depending on context
*/

const s string = "constant"

func constants() {

	fmt.Println(s)
	//constant

	const n = 500000000 // const statement can appear anywhere a var statement can

	const d = 3e20 / n // constants perform arimetrric with arbitraty precision
	fmt.Println(d)
	//6e+11

	fmt.Println(int64(d)) // numeric constant has no type until it's given one; for example by explicit conversion
	//600000000000

	//a number can be given type by using it in a context that requires one such as variable assignment or functions call
	fmt.Println(math.Sin(n)) //  math.Sin expects a float64
	//-0.20470407323754404

	const ( // multiple constants can be declared at once using parentheses
		A = 1
		B = 2
	)

	// related to Enums
	const ( // iota - is a special constant used for enumerated constants, which incremenets automatically
		C0 = iota // C0 == 0
		C1 = iota // C1 == 1
		C2 = iota // C2 == 2
	)
}
