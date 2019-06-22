package slice_test

import (
	"reflect"
	"testing"

	"github.com/gellel/slice"
)

var (
	slices *slice.Slices
)

func TestSlices(t *testing.T) {

	slices = slice.NewSlices()

	ok := slices != nil && reflect.ValueOf(slices).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(slice.Slices) != reflect.Ptr")
	}
}

func TestSlicesAppend(t *testing.T) {

	slices.Append(slice.New(1))

	if s, ok := slices.Get(0); ok != true || s == nil {
		t.Fatalf("slices.Append(slice *slice.Slice) did not append a new slice pointer")
	}
}

func TestSlicesAssign(t *testing.T) {

	slices.Assign(slice.New(2), slice.New(3))

	for i := 1; i < 3; i++ {
		if s, ok := slices.Get(i); ok != true || s == nil {
			t.Fatalf("slices.Assign(slices ...*slice.Slice) did not append a new slice pointer")
		}
	}
}

func TestSlicesConcatenate(t *testing.T) {

	slices = slice.NewSlices()

	x := slice.NewSlices()

	x.Append(slice.New(1))

	slices.Concatenate(x)

	if ok := slices.Fetch(0) != nil; ok != true {
		t.Fatalf("slices.Concatenate(slices *slice.Slices) did not concatenate the argument slices slice")
	}
}

func TestSlicesEach(t *testing.T) {

	slices.Each(func(_ int, slice *slice.Slice) {})
}

func TestSlicesFlatten(t *testing.T) {

	n := 0
	slices.Each(func(i int, slice *slice.Slice) {
		n = n + slice.Len()
	})
	if ok := slices.Flatten().Len() == n; ok != true {
		t.Fatalf("slices.Flatten() did not flatten the slices slice into a single slice of the same length")
	}
}

func TestSlicesGet(t *testing.T) {

	if slice, ok := slices.Get(0); ok != true || slice == nil {
		t.Fatalf("slices.Get(i int) did not return a slice and true")
	}
	if slice, ok := slices.Get(-1); ok != false || slice != nil {
		t.Fatalf("slices.Get(i int) did not return nil and false")
	}
}

func TestSlicesPoll(t *testing.T) {

	slices = slice.NewSlicesSlice(slice.New(), slice.New())

	for slices.Len() != 0 {
		if slices.Poll() == nil {
			t.Fatalf("slices.Poll() returned nil and not a slice")
		}
	}
}

func TestSlicesPop(t *testing.T) {

	slices = slice.NewSlicesSlice(slice.New(), slice.New())

	for slices.Len() != 0 {
		if slices.Pop() == nil {
			t.Fatalf("slices.Pop() return nil and not a slice")
		}
	}
}

func TestSlicesPreassign(t *testing.T) {

	slices = slice.NewSlicesSlice(slice.New(1))

	slices.Preassign(slice.New(2))

	if ok := slices.Fetch(0).Fetch(0) == 2; ok != true {
		t.Fatalf("slices.Preassign(slices ...*slices.Slice) did not push the argument slice to the beginning of the slices")
	}
}

func TestSlicesPrecatenate(t *testing.T) {

	slices = slice.NewSlices()

	x := slice.NewSlicesSlice(slice.New(1))

	slices.Precatenate(x)

	if s, ok := slices.Get(0); ok != true || s == nil || s.Fetch(0).(int) != 1 {
		t.Fatalf("slices.Precatenate(s *slice.Slices) did not prepend the argument slices to the receiver slice")
	}
}
