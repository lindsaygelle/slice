package slice_test

import (
	"fmt"

	"github.com/gellel/slice"
)

func ExampleNewString() {

	fmt.Println(slice.NewString())
	// Output: &[]
}

func ExampleNewStringSlice() {
	fmt.Println(slice.NewStringSlice("a", "b", "c"))
	// Output: &[a b c]
}
