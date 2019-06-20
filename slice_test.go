package slice_test

import (
	"math/rand"
	"reflect"
	"strings"
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

func TestJoin(t *testing.T) {

	type X struct{}

	x := []string{}
	y := []interface{}{}
	z := []X{X{}}

	s.Append(x).Append(y).Append(z)

	if len(strings.Split(s.Join("-"), "-")) != s.Len() {
		t.Fatalf("slice.Join(s string) is not the same length as the slice length")
	}
}

func TestLess(t *testing.T) {

	s.Append("a").Append("b")

	i := s.Len() - 2
	j := s.Len() - 1

	if ok := s.Less(i, j); ok != true {
		t.Fatalf("slice.Less(i, j int) did not return true (a < b)")
	}

	s.Append(1).Append(2)

	i = s.Len() - 2
	j = s.Len() - 1

	if ok := s.Less(i, j); ok != true {
		t.Fatalf("slice.Less(i, j int) did not return true (1 < 2)")
	}
}

func TestMap(t *testing.T) {

	s.Map(func(i int, value interface{}) interface{} {

		value = i

		return value
	})

	s.Each(func(i int, value interface{}) {
		if _, ok := value.(int); ok != true {
			t.Fatalf("slice.Map(func(i int, value interface{}) interface{}) did not mutate the slice at position " + string(i))
		}
	})
}

func TestPoll(t *testing.T) {

	x := s.Fetch(0).(int)

	i := s.Poll()

	if i == nil {
		t.Fatalf("slice.Poll() did not return an interface")
	}
	if _, ok := i.(int); ok != true {
		t.Fatalf("slice.Poll() did not return an integer")
	}
	if i.(int) != x {
		t.Fatalf("slice.Poll() return the wrong integer; did not return " + string(x) + " returned " + string(i.(int)))
	}
}

func TestPop(t *testing.T) {

	x := s.Fetch(s.Len() - 1).(int)

	i := s.Pop()

	if i == nil {
		t.Fatalf("slice.Pop() did not return an interface")
	}
	if _, ok := i.(int); ok != true {
		t.Fatalf("slice.Pop() did not return an integer")
	}
	if i.(int) != x {
		t.Fatalf("slice.Pop() return the wrong integer; did not return " + string(x) + " returned " + string(i.(int)))
	}
}

func TestPreassign(t *testing.T) {

	s.Preassign(0, 1, 2)

	for i := 0; i < 3; i++ {

		if s.Fetch(i).(int) != i {
			t.Fatalf("slice.Preassign(values ...interface{}) did not push new elements to the beginning of slice")
		}
	}
}

func TestPrecatenate(t *testing.T) {

	x := slice.New("a", "b", "c")

	s.Precatenate(x)

	x.Each(func(i int, value interface{}) {
		if s.Fetch(i) != value {
			t.Fatalf("slice.Precatenate(slice *Slice) did not push the argument slice elements to the beginning of the receiver slice")
		}
	})
}

func TestPrepend(t *testing.T) {

	if s.Prepend(-1).Fetch(0).(int) != -1 {
		t.Fatalf("slice.Prepend(value interface{}) did not push the new element to the beginning of the slice")
	}
}

func TestReplace(t *testing.T) {

	if ok := s.Replace(0, "N"); ok != true {
		t.Fatalf("slice.Replace(i int, value interface{}) did not replace the element at the target index")
	}
}
