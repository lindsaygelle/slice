package slice_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/gellel/slice"
)

var (
	is *slice.Int
)

func TestIntSlice(t *testing.T) {

	is = slice.NewInt()

	ok := is != nil && reflect.ValueOf(is).Kind() == reflect.Ptr

	if ok != true {
		t.Fatalf("reflect.ValueOf(slice.Int) != reflect.Ptr")
	}
}

func TestIntMax(t *testing.T) {

	for i := 0; i < 10; i++ {
		is.Append(rand.Intn(10))
	}

	is.Replace(rand.Intn(is.Len()-1), 1000)

	if ok := is.Max() == 1000; ok != true {
		t.Fatalf("intSlice.Max() did not return 1000 as the max value")
	}
}

func TestIntMin(t *testing.T) {

	is.Replace(rand.Intn(is.Len()-1), -1000)

	if ok := is.Min() == -1000; ok != true {
		t.Fatalf("intSlice.Min() did not return -1000 as the min value")
	}
}

func TestIntSort(t *testing.T) {

	is.Sort().Each(func(i, _ int) {
		if i != 0 {
			if ok := is.Fetch(i-1) <= is.Fetch(i); ok != true {
				t.Fatalf("intSlice.Sort() did not sort the int slice")
			}
		}
	})
}

func TestInSum(t *testing.T) {

	is = slice.NewInt()

	n := rand.Intn(1000)

	for i := 0; i < n; i++ {
		is.Append(1)
	}

	if x := is.Sum(); x != n {
		t.Fatalf("intSlice.Sum() did not return the sum of " + string(n) + " as the total of all slice values; returned " + string(x))
	}
}
