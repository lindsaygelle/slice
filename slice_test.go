package slice_test

import (
	"testing"

	"github.com/lindsaygelle/slice"
)

// TestAppend tests Slice.Append.
func TestAppend(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Append(1)
	if ok := len(*s) == 1; !ok {
		t.Fatalf("len(*Slice) != 1")
	}
	s = &slice.Slice[int]{}
	s.Append(1, 2)
	if ok := len(*s) == 2; !ok {
		t.Fatalf("len(*Slice) != 2")
	}
}
