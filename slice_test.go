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

// TestAppendLength tests Slice.AppendLength.
func TestAppendLength(t *testing.T) {
	s := &slice.Slice[int]{}
	if n := s.AppendLength(1, 2, 3, 4); n != len(*s) {
		t.Fatalf("len(*Slice) != 4")
	}
}

// TestBounds tests Slice.Bounds.
func TestBounds(t *testing.T) {
	s := &slice.Slice[int]{}
	if ok := s.Bounds(0); ok {
		t.Fatalf("*Slice.Bounds() != false")
	}
	s.Append(1)
	if ok := s.Bounds(0); !ok {
		t.Fatalf("*Slice.Bounds() != true")
	}
}

// TestConcatenate tests Slice.Concatenate.
func TestConcatenate(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Append(1)
	s.Concatenate(&slice.Slice[int]{2, 3})
	if ok := (*s)[0] == 1; !ok {
		t.Fatalf("*Slice[0] != 1")
	}
	if ok := (*s)[1] == 2; !ok {
		t.Fatalf("*Slice[1] != 2")
	}
	if ok := (*s)[2] == 3; !ok {
		t.Fatalf("*Slice[2] != 3")
	}
}

// TestConcatenateLength tests Slice.ConcatenateLength.
func TestConcatenateLength(t *testing.T) {
	s := &slice.Slice[int]{}
	if n := s.ConcatenateLength(&slice.Slice[int]{1}); n != len(*s) {
		t.Fatalf("len(*Slice) != %d", len(*s))
	}
}

// TestDelete tests Slice.Delete.
func TestDelete(t *testing.T) {
	s := &slice.Slice[int]{1}
	s.Delete()
	if ok := len(*s) == 0; !ok {
		t.Fatalf("len(*Slice) != 0")
	}
}

// TestDeleteLength tests Slice.DeleteLength.
func TestDeleteLength(t *testing.T) {
	s := &slice.Slice[int]{1}
	if n := s.DeleteLength(); n != len(*s) {
		t.Fatalf("len(*Slice) != 0")
	}
}

// TestEach tests Slice.Each.
func TestEach(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Each(func(i int, value int) {
		if ok := (*s)[i] == value; !ok {
			t.Fatalf("*Slice[%d] != %d", i, value)
		}
	})
}
