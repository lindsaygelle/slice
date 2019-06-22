package slice_test

import (
	"reflect"
	"testing"

	"github.com/gellel/slice"
)

var (
	sss *slice.Slices
)

func TestSlices(t *testing.T) {

	sss = slice.NewSlices()

	ok := sss != nil && reflect.ValueOf(sss).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(slice.Slices) != reflect.Ptr")
	}
}

func TestSlicesAppend(t *testing.T) {

	sss.Append(slice.New(1))

	if s, ok := sss.Get(0); ok != true || s == nil {
		t.Fatalf("slices.Append(slice *slice.Slice) did not append a new slice pointer")
	}
}

func TestSlicesAssign(t *testing.T) {

	sss.Assign(slice.New(2), slice.New(3))

	for i := 1; i < 3; i++ {
		if s, ok := sss.Get(i); ok != true || s == nil {
			t.Fatalf("slices.Assign(slices ...*slice.Slice) did not append a new slice pointer")
		}
	}
}

func TestSlicesConcatenate(t *testing.T) {

	sss = slice.NewSlices()

	x := slice.NewSlices()

	x.Append(slice.New(1))

	sss.Concatenate(x)

	if ok := sss.Fetch(0) != nil; ok != true {
		t.Fatalf("slices.Concatenate(slices *slice.Slices) did not concatenate the argument slices slice")
	}
}
