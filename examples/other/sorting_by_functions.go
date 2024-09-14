package main

import (
	"cmp"
	"fmt"
	"slices"
)

/*
Sometimes we want to sort collection by something other than its natural order.
For example sort strings by their length instead of alphabetically.
 */

func sorting_by_functions() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {  // Comparison function for string lengths
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)  // Now we call slices.SortFunc with custom comporison function to sort fruits slice by name length
	fmt.Println(fruits)
	// [kiwi peach banana]

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people, // Sort people by age
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
	// [{TJ 25} {Jax 37} {Alex 72}]

}
