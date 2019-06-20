package slice_test

import (
	"reflect"
	"testing"

	"github.com/gellel/slice"
)

var (
	ss *slice.String
)

func TestStringSlice(t *testing.T) {

	ss = slice.NewString()

	ok := ss != nil && reflect.ValueOf(ss).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(slice.String) != reflect.Ptr")
	}
}
