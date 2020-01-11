package slice_test

import (
	"testing"

	"github.com/gellel/slice"
)

func Test(t *testing.T) {

	var (
		s = slice.NewStringer()
	)
	s.Bounds(0)
}
