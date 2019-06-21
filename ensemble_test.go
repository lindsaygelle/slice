package slice_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/gellel/slice"
)

var (
	e *slice.Ensemble
)

func TestEnsembleSlice(t *testing.T) {

	e = slice.NewEnsemble()

	ok := e != nil && reflect.ValueOf(e).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(slice.Ensemble) != reflect.Ptr")
	}
}

func TestEnsembleFlatten(t *testing.T) {

	a := &slice.Slice{}

	b := &slice.Slice{}

	aN := rand.Intn(10)

	bN := rand.Intn(10)

	for i := 0; i < aN; i++ {
		a.Append(i)
	}
	for i := 0; i < bN; i++ {
		b.Append(i)
	}
	e.Assign(a, b)

	c := e.Flatten()

	if ok := c.Len() == (aN + bN); ok != true {
		t.Fatalf("ensembleSlice.Flatten() did not return a slice containing all values from the ensemble")
	}
}
