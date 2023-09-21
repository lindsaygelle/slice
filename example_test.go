package slice_test

import (
	"fmt"

	"github.com/lindsaygelle/slice"
)

func ExampleNew() {
	// Create a new Slice
	s := slice.New(1, 2, 3, 4, 5)

	// Append values to the Slice
	s.Append(6, 7, 8)
	fmt.Println("Appended:", s)

	// Append values based on a custom condition
	s.AppendFunc(func(i int, value int) bool {
		return value%2 == 0
	}, 10, 11, 12)
	fmt.Println("Appended based on condition:", s)

	// Get the length of the Slice
	fmt.Println("Length of Slice:", s.Length())

	// Check if the Slice contains a specific value
	fmt.Println("Contains 3:", s.Contains(3)) // Should print true
	fmt.Println("Contains 9:", s.Contains(9)) // Should print false

	// Create a new Slice of booleans indicating if values are present in the original Slice
	contains := s.ContainsMany(2, 5, 8, 10)
	fmt.Println("Contains:", contains)

	// Delete an element at a specific index
	s.Delete(2)
	fmt.Println("Deleted at index 2:", s)

	// Iterate over the Slice and apply a function to each element
	s.Each(func(i int, value int) {
		fmt.Printf("Index %d: %d\n", i, value)
	})

	// Find the index of an element that satisfies a condition
	index, found := s.FindIndex(func(value int) bool {
		return value == 4
	})
	if found {
		fmt.Println("Index of 4:", index) // Should print 3
	} else {
		fmt.Println("4 not found")
	}

	// Get an element at a specific index
	value, ok := s.Get(1)
	if ok {
		fmt.Println("Value at index 1:", value) // Should print 2
	} else {
		fmt.Println("Index 1 not found")
	}

	// Make a new Slice with specific values
	s = s.MakeEach(5, 4, 3, 2, 1)
	fmt.Println("New Slice:", s)

	// Reverse the Slice
	s.Reverse()
	fmt.Println("Reversed Slice:", s)

	// Shuffle the Slice
	s.Shuffle()
	fmt.Println("Shuffled Slice:", s)

	// Slice the Slice from index 2 to 5 (exclusive)
	s.Slice(2, 5)
	fmt.Println("Sliced Slice:", s)

}
