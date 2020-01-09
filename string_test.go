package slice_test

import (
	"fmt"
	"testing"

	"github.com/gellel/slice"
)

func TestStringer(t *testing.T) {

	var (
		stringer = slice.NewStringer()
	)
	fmt.Println(stringer.Append("a", "b", "c"))
}
