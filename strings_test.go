package slice_test

import (
	"testing"

	"github.com/gellel/slice"
)

var (
	strsSlice *slice.Strings
)

func TestStringsSlice(t *testing.T) {

	strsSlice = slice.NewStringsSlice()

	strsSlice.Append(slice.NewStringSlice("a"))

}
