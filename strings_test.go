package slice_test

import (
	"reflect"
	"testing"

	"github.com/gellel/slice"
)

var (
	strsSlice *slice.Strings
)

func TestStringsSlice(t *testing.T) {

	strsSlice = slice.NewStrings()

	ok := ss != nil && reflect.ValueOf(strsSlice).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(slice.Strings) != reflect.Ptr")
	}
}

func TestStringsFlatten(t *testing.T) {

	strsSlice.Assign(slice.NewStringSlice("a", "b"), slice.NewStringSlice("c"), slice.NewStringSlice("d", "e"))

	if ok := strsSlice.Flatten().Join("") == "abcde"; ok != true {
		t.Fatalf("stringsSlice.Flatten() did not return a flattened string slice")
	}
}
