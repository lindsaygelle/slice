package slice_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/gellel/slice"
)

var (
	sss *slice.Slices
)

func TestSlicesSlice(t *testing.T) {

	sss = slice.NewSlices()

	ok := sss != nil && reflect.ValueOf(sss).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(slice.Slices) != reflect.Ptr")
	}
}

func TestSlicesFlatten(t *testing.T) {

	a := &slice.Slice{}

	b := &slice.Slice{}

	aN := rand.Intn(10)

	bN := rand.Intn(10)

	for i := 0; i < aN; i++ {
		a.Append(i)
	}
	for i := 0; i < bN; i++ {
		b.Append(i)
	}

	c := sss.Assign(a, b).Flatten()

	if ok := c.Len() == (aN + bN); ok != true {
		t.Fatalf("slicesSlice.Flatten() did not return a slice containing all values from the slice")
	}
}
