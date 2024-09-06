package main

import "fmt"

/* GENERICS
Generics use type paarameters that allow to specify types at the time of calling a function or creating an instance of a type
func Print[T any](item T) {
	fmt.Println(item)
}
T is a type parameter, any is a type constraint

- Type constraints define the types that can be passed as type parameters
	- 'any' keyword is a predefined type that allows any type
func Add[T int | float64](a, b T) T {
	return a + b
}
*/

/*
SliceIndex takes a slice of any comparable type and an element of that type
and returns the index of the first occurence of v in s, or -1 if not present.

The comparable constraint means that we can compare values of this type with the == and != operators (numbers, strings, pointers)
https://go.dev/blog/deconstructing-type-parameters
*/
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct { // example of generic type, List is a singly-linked list with values of any type
	head, tail *element[T]
}

type element[T any] struct {  // represents a single element in the list
	next *element[T]  // points to the next element in the list
	val  T  // value of the element
}

func (lst *List[T]) Push(v T) { // we can define methods on generic types just like we do on regular types
	if lst.tail == nil { // but we have to keep the type parameters in place. Type is List[T], not List
		lst.head = &element[T]{val: v}  // new element becomes the head and tail of the list
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T { // AllElements returns all the List elements as a slice
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func generics() {
	var s = []string{"foo", "bar", "zoo"}

	fmt.Println("index of zoo:", SlicesIndex(s, "zoo")) // when invoking generic functions, we can rely on type interference
	// index of zoo: 2

	_ = SlicesIndex[[]string, string](s, "zoo") 
	// _ czyli 'blank identifier' means that the result of the function call will be ignored

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
	// list: [10 13 23]

}
