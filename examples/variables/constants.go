package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func constants() {

	//*** CONSTANTS
	fmt.Println(s) //constant

	const n = 500000000 //const statement can appear anywhere a var statement can

	const d = 3e20 / n //constants perform arimetrric with arbitraty precision
	fmt.Println(d)     //6e+11

	// numeric constant has no type until it's given one; for example by explicit conversion
	fmt.Println(int64(d)) //600000000000

	//a number can be given type by using it in a context that requires one such as variable assignment or functions call
	fmt.Println(math.Sin(n)) //-0.20470407323754404
	//math.Sin expects a float64

}
