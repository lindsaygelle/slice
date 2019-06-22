package slice_test

import (
	"fmt"

	"github.com/gellel/slice"
)

func ExampleNew() {
	fmt.Println(slice.New("a", 1, false))
	// Output: &[a 1 false]
}
