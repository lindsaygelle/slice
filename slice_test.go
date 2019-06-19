package slice_test

import (
	"reflect"
	"testing"

	"github.com/gellel/slice"
)

var (
	s *slice.Slice
)

func Test(t *testing.T) {

	s = slice.New()

	ok := reflect.ValueOf(s).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(slice) != reflect.Ptr")
	}
}

func TestAppend(t *testing.T) {

	previousLength := s.Len()

	x := s.Append(1)

	if x != s {
		t.Fatalf("slice.Append(i interface{}) did not return the same slice")
	}

	currentLength := x.Len()

	if previousLength == currentLength {
		t.Fatalf("slice.Append(i interface{}) did not append a new element to the slice pointer")
	}
}

func TestAssign(t *testing.T) {

	previousLength := s.Len()

	x := s.Assign(1, 2, 3)

	if x != s {
		t.Fatalf("slice.Assign(values ...interface{}) did not return the same slice")
	}

	currentLength := x.Len()

	if previousLength == currentLength {
		t.Fatalf("slice.Assign(values ...interface{}) did not append 3 new elements to the slice pointer")
	}
}

func TestBounds(t *testing.T) {

	if ok := s.Bounds(-1) || s.Bounds(s.Len()+1); ok == true {
		t.Fatalf("slice.Bounds(i int) did not return false for an out of bounds integer")
	}
	if ok := s.Bounds(s.Len() / 2); ok != true {
		t.Fatalf("slice.Bounds(i int) did not return true for an in bounds integer")
	}
}

func TestConcatenate(t *testing.T) {

	x := slice.New("x")

	s.Concatenate(x)

	i := (*s)[s.Len()-1].(string)

	if i != "x" {
		t.Fatalf("slice.Concatenate(s *Slice) did not append the contents of a sibling slice to the reference slice")
	}
}
