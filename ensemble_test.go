package slice_test

import (
	"fmt"
	"testing"

	"github.com/gellel/slice"
)

func TestEnsemble(t *testing.T) {

	a := slice.NewEnsemble()
	b := slice.NewEnsemble()

	a.Append(slice.New("1"))
	b.Assign(slice.New("2"), slice.New("3"))

	fmt.Println(a.Fetch(0))
	fmt.Println(b.Fetch(0), b.Fetch(1))

	a.Precatenate(b)

	a.Each(func(i int, slice *slice.Slice) {

		fmt.Println(i, slice)
	})
}
