package main

import "fmt"

func main() {
	fmt.Println("-- First case:")
	a := []int{0, 1, 2, 3, 4}
	var b = append(a, 5)
	b[0] = 10
	fmt.Println("a", a) // a [0 1 2 3 4]
	fmt.Println("b", b) // b [10 1 2 3 4 5]

	fmt.Println("\n-- Second case:")
	var c = append(b, 6)
	c[1] = 11
	fmt.Println("a", a) // a [0 1 2 3 4]
	fmt.Println("b", b) // b [10 11 2 3 4 5]
	fmt.Println("c", c) // c [10 11 2 3 4 5 6]

	fmt.Println("\n-- Extra info:")
	fmt.Print("a: capacity=", cap(a), ", length=", len(a), ", equal?", len(a) == cap(a), "\n")
	fmt.Print("b: capacity=", cap(b), ", length=", len(b), ", equal?", len(b) == cap(b), "\n")
	fmt.Print("c: capacity=", cap(c), ", length=", len(c), ", equal?", len(c) == cap(c), "\n")

	/* Explanations:
	-- First case ('b' is a copy of 'a'):
	Upon creation, the length and capacity of 'a' are equal.
	Therefore, when appending a new element, the slice is copied into a new one whose capacity is the double (10 vs 5).
	Since 'b' is a new slice, modifying one of its elements does not change the content of 'a'.


	-- Second case ('c' is a reference to 'b'):
	The length of 'b' is smaller than its capacity.
	Therefore, when appending a new element, the original slice ('b') is mutated directly to add the extra element.
	Since 'c' is a reference to 'b', modifying one element also modifies the same element in 'b'.

	If 'c' is a reference to 'b', why does 'b' show only 6 elements and not 7 like 'c' ?
	Because a slice stores a pointer to the array where the elements are stored AND the length (i.e. the number of
	elements the slice contains) (see https://stackoverflow.com/questions/39993688/are-slices-passed-by-value).
	Somehow, a slice is thus a view of its underlying array, from the first item to item nº length_slice.

	See also: https://stackoverflow.com/questions/20195296/golang-append-an-item-to-a-slice
	*/
}
