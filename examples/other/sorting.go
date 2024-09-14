package main

import (
	"fmt"
	"slices"
)

/* SORTING
Sorting is used to arrange data in specific order (ascending or descending).
- Go provides sort package for sorting built-in types
- Simplifies sorting of common data types without requiring custom logic
*/

func sorting() {

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)
	// Strings: [a b c]

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)
	// Ints: [2 4 7]

	s := slices.IsSorted(ints) // Check if slice is already sorted
	fmt.Println("Sorted: ", s)
	// Sorted: true

}
