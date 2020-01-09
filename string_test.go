package slice_test

import (
	"fmt"
	"testing"

	"github.com/gellel/slice"
)

func TestString(t *testing.T) {

	var (
		s = slice.String("a", "b", "c", "d")
	)
	s.Each(func(_ int, x string) {
		fmt.Println(x)
	})
	fmt.Println(s.Pop())

}
