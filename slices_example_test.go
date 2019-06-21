package slice_test

import "github.com/gellel/slice"

func ExampleNewSlices() {
	slice.NewSlices()
	// Output: &slice.Slices{}
}

func ExampleNewSlicesSlice() {
	slice.NewSlicesSlice(&slice.Slice{1}, &slice.Slice{2})
	// Output: &slice.Slices{&slices.Slice{1}, &slices.Slice{2}}
}
