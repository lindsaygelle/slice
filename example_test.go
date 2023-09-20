package slice_test

import (
	"fmt"

	"github.com/lindsaygelle/slice"
)

func ExampleSlice() {

	// Create a Slice of integers.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}

	// Append values to the Slice.
	s.Append(6, 7, 8)

	// Check if a value exists in the Slice.
	contains := s.Contains(3) // contains is true

	fmt.Println("Contains:", contains)

	// Find the index of the first element that satisfies a condition.
	index, found := s.FindIndex(func(value int) bool {
		return value > 4
	})

	fmt.Println("Index:", index, "Found:", found)

	// Delete an element by index.
	s.Delete(2) // Removes the element at index 2

	// Iterate over the elements and perform an operation on each.
	s.Each(func(i int, value int) {
		fmt.Printf("Element %d: %d\n", i, value)
	})
}

func ExampleAppend() {
	s := &slice.Slice[int]{1, 2, 3}
	s.Append(4, 5) // s is now [1, 2, 3, 4, 5]
	fmt.Println(s)
}
