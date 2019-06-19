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

	x := s.Append(1)

	if x != s {
		t.Fatalf("slice.Append(i interface{}) did not return the same slice")
	}
}
