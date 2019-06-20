package slice_test

import (
	"math/rand"
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

func TestEach(t *testing.T) {

	currentLength := s.Len()

	i := 0

	s.Each(func(j int, value interface{}) {
		if (*s)[i] != value {
			t.Fatalf("slice.Each(func(i int, value interface{}) value != (*s)[i]")
		}
		i = i + 1
	})

	if i != currentLength {
		t.Fatalf("slice.Each(func(i int, value interface{}) final i count is not equal to slice length")
	}
}

func TestFetch(t *testing.T) {

	i := rand.Intn(s.Len() - 1)

	if s.Fetch(i) == nil {
		t.Fatalf("slice.Fetch(i int) did not return a non-nil property")
	}
}

func TestGet(t *testing.T) {

	i := rand.Intn(s.Len() - 1)

	if _, ok := s.Get(i); ok != true {
		t.Fatalf("slice.Get(i int) did not return an interface and a true boolean")
	}
	if i, ok := s.Get(-1); ok != false || i != nil {
		t.Fatalf("slice.Get(i int) did not return nil or a false boolean")
	}
}
