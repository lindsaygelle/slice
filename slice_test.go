package slice_test

import (
	"testing"

	"github.com/lindsaygelle/slice"
)

// TestAppend tests Slice.Append.
func TestAppend(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Append(1)
	if ok := len(*s) == 1; !ok {
		t.Fatalf("len(*Slice) != 1")
	}
	s = &slice.Slice[int]{}
	s.Append(1, 2)
	if ok := len(*s) == 2; !ok {
		t.Fatalf("len(*Slice) != 2")
	}
}

// TestAppendFunc tests Slice.AppendFunc.
func TestAppendFunc(t *testing.T) {
	s := &slice.Slice[int]{}
	s.AppendFunc(func(i int, value int) bool {
		return value%2 == 0
	}, 1, 2, 3)
	if ok := len(*s) == 1; !ok {
		t.Fatalf("len(*Slice) != 1")
	}
}

// TestAppendLength tests Slice.AppendLength.
func TestAppendLength(t *testing.T) {
	s := &slice.Slice[int]{}
	if n := s.AppendLength(1, 2, 3, 4); n != len(*s) {
		t.Fatalf("len(*Slice) != 4")
	}
}

// TestBounds tests Slice.Bounds.
func TestBounds(t *testing.T) {
	s := &slice.Slice[int]{}
	if ok := s.Bounds(0); ok {
		t.Fatalf("*Slice.Bounds() != false")
	}
	s.Append(1)
	if ok := s.Bounds(0); !ok {
		t.Fatalf("*Slice.Bounds() != true")
	}
}

// TestConcatenate tests Slice.Concatenate.
func TestConcatenate(t *testing.T) {
	s := &slice.Slice[int]{}
	s.Append(1)
	s.Concatenate(&slice.Slice[int]{2, 3})
	if ok := (*s)[0] == 1; !ok {
		t.Fatalf("*Slice[0] != 1")
	}
	if ok := (*s)[1] == 2; !ok {
		t.Fatalf("*Slice[1] != 2")
	}
	if ok := (*s)[2] == 3; !ok {
		t.Fatalf("*Slice[2] != 3")
	}
}

// TestConcatenateFunc tests Slice.ConcatenateFunc
func TestConcatenateFunc(t *testing.T) {
	a := &slice.Slice[int]{}
	b := &slice.Slice[int]{2, 4, 6}
	a.ConcatenateFunc(b, func(i int, value int) bool {
		return value%2 == 0
	})
	if ok := len(*a) == 3; !ok {
		t.Fatalf("len(*Slice) != %d", len(*b))
	}
}

// TestConcatenateLength tests Slice.ConcatenateLength.
func TestConcatenateLength(t *testing.T) {
	s := &slice.Slice[int]{}
	if n := s.ConcatenateLength(&slice.Slice[int]{1}); n != len(*s) {
		t.Fatalf("len(*Slice) != %d", len(*s))
	}
}

// TestContains tests Slice.Contains.
func TestContains(t *testing.T) {
	// Create a test slice with some data
	s := &slice.Slice[int]{1, 2, 3, 4, 5}

	// Test for a value that exists in the slice
	existingValue := 3
	if !s.Contains(existingValue) {
		t.Fatalf("Contains(%d) returned false, expected true", existingValue)
	}

	// Test for a value that does not exist in the slice
	nonExistentValue := 6
	if s.Contains(nonExistentValue) {
		t.Fatalf("Contains(%d) returned true, expected false", nonExistentValue)
	}

	// Test for a value using a custom type (e.g., a struct)
	type CustomStruct struct {
		Name string
		Age  int
	}

	customSlice := &slice.Slice[CustomStruct]{
		{"Alice", 25},
		{"Bob", 30},
	}

	// Define a custom value to search for
	customValue := CustomStruct{"Alice", 25}
	if !customSlice.Contains(customValue) {
		t.Fatalf("Contains(%+v) returned false, expected true", customValue)
	}

	// Define a different custom value that does not exist in the slice
	nonExistentCustomValue := CustomStruct{"Eve", 35}
	if customSlice.Contains(nonExistentCustomValue) {
		t.Fatalf("Contains(%+v) returned true, expected false", nonExistentCustomValue)
	}
}

func TestContainsMany(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	results := s.ContainsMany(2, 4)
	for i, value := range *results {
		if !value {
			t.Fatalf("results[%d] != true", i)
		}
	}
	results = s.ContainsMany(0, 10, 99)
	for i, value := range *results {
		if value {
			t.Fatalf("results[%d] != false", i)
		}
	}
}

// TestDelete tests Slice.Delete.
func TestDelete(t *testing.T) {
	s := &slice.Slice[int]{1}
	s.Delete(0)
	if ok := len(*s) == 0; !ok {
		t.Fatalf("len(*Slice) != 0")
	}
}

// TestDeleteLength tests Slice.DeleteLength.
func TestDeleteLength(t *testing.T) {
	s := &slice.Slice[int]{1}
	if n := s.DeleteLength(0); n != len(*s) {
		t.Fatalf("len(*Slice) != 0")
	}
}

// TestEach tests Slice.Each.
func TestEach(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Each(func(i int, value int) {
		if ok := (*s)[i] == value; !ok {
			t.Fatalf("*Slice[%d] != %d", i, value)
		}
	})
}

// TestEachBreak tests Slice.EachBreak.
func TestEachBreak(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	count := 0
	s.EachBreak(func(i int, value int) bool {
		count = count + 1
		return false
	})
	if ok := count == 1; !ok {
		t.Fatalf("count != 1")
	}
}

// TestEachReverse tests Slice.EachReverse.
func TestEachReverse(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.EachReverse(func(i int, value int) {
		if ok := (*s)[i] == value; !ok {
			t.Fatalf("*Slice[%d] != %d", i, value)
		}
	})
}

// TestEachReverseBreak tests Slice.EachReverseBreak.
func TestEachReverseBreak(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	count := 0
	s.EachReverseBreak(func(i int, value int) bool {
		count = count + 1
		return false
	})
	if ok := count == 1; !ok {
		t.Fatalf("count != 1")
	}
}

// TestFetch tests Slice.Fetch.
func TestFetch(t *testing.T) {
	s := &slice.Slice[int]{1}
	for i := range *s {
		value := s.Fetch(i)
		if ok := value == (*s)[i]; !ok {
			t.Fatalf("%d != %d", value, (*s)[i])
		}
	}
	// Deliberately empty and check default is returned.
	s = &slice.Slice[int]{}
	value := s.Fetch(1)
	if ok := value == 0; !ok {
		t.Fatalf("%d != 0", value)
	}
}

// TestFetchLength tests Slice.FetchLength.
func TestFetchLength(t *testing.T) {
	s := &slice.Slice[int]{1, 2}
	for i := range *s {
		_, value := s.FetchLength(i)
		if ok := value == len(*s); !ok {
			t.Fatalf("%d != %d", value, len(*s))
		}
	}
}

// TestFindIndex tests Slice.FindIndex.
func TestFindIndex(t *testing.T) {
	// Create a test slice with some data
	s := &slice.Slice[int]{1, 2, 3, 4, 5}

	// Test for a predicate function that matches an element
	index, found := s.FindIndex(func(value int) bool {
		return value == 3
	})
	expectedIndex := 2
	if !found || index != expectedIndex {
		t.Fatalf("FindIndex returned (%d, %v), expected (%d, true)", index, found, expectedIndex)
	}

	// Test for a predicate function that doesn't match any element
	index, found = s.FindIndex(func(value int) bool {
		return value == 6
	})
	expectedIndex = -1
	if found || index != expectedIndex {
		t.Fatalf("FindIndex returned (%d, %v), expected (%d, false)", index, found, expectedIndex)
	}

	// Test for a custom type (e.g., a struct)
	type CustomStruct struct {
		Name string
		Age  int
	}

	customSlice := &slice.Slice[CustomStruct]{
		{"Alice", 25},
		{"Bob", 30},
	}

	// Define a custom predicate function
	predicate := func(value CustomStruct) bool {
		return value.Name == "Bob"
	}

	// Test for a predicate function that matches an element in the custom slice
	customIndex, customFound := customSlice.FindIndex(predicate)
	expectedCustomIndex := 1
	if !customFound || customIndex != expectedCustomIndex {
		t.Fatalf("FindIndex returned (%d, %v), expected (%d, true)", customIndex, customFound, expectedCustomIndex)
	}

	// Test for a predicate function that doesn't match any element in the custom slice
	index, found = customSlice.FindIndex(func(value CustomStruct) bool {
		return value.Name == "Eve"
	})
	expectedIndex = -1
	if found || index != expectedIndex {
		t.Fatalf("FindIndex returned (%d, %v), expected (%d, false)", index, found, expectedIndex)
	}
}

// TestGet tests Slice.Get.
func TestGet(t *testing.T) {
	s := &slice.Slice[int]{1}
	for i := range *s {
		value, ok := s.Get(i)
		if value != (*s)[i] {
			t.Fatalf("%d != %d", value, (*s)[i])
		}
		if !ok {
			t.Fatalf("%t != true", ok)
		}
	}
}

// TestGetLength tests Slice.GetLength.
func TestGetLength(t *testing.T) {
	s := &slice.Slice[int]{1}
	for i := range *s {
		value, length, ok := s.GetLength(i)
		if value != (*s)[i] {
			t.Fatalf("%d != %d", value, (*s)[i])
		}
		if length != len(*s) {
			t.Fatalf("%d = %d", length, len(*s))
		}
		if !ok {
			t.Fatalf("%t != true", ok)
		}
	}
}

// TestIsEmpty tests Slice.IsEmpty.
func TestIsEmpty(t *testing.T) {
	if ok := (&slice.Slice[int]{}).IsEmpty(); !ok {
		t.Fatal("IsEmpty did not return true")
	}
}

func TestIsPopulated(t *testing.T) {
	if ok := (&slice.Slice[int]{1}).IsPopulated(); !ok {
		t.Fatal("IsPopulated did not return true")
	}
}

// TestLength tests Slice.Length.
func TestLength(t *testing.T) {
	s := &slice.Slice[int]{}
	if ok := s.Length() == len(*s); !ok {
		t.Fatalf("len(*Slice) != %d", len(*s))
	}
}

// TestMake tests Slice.Make.
func TestMake(t *testing.T) {
	s := &slice.Slice[int]{}
	size := 10
	s.Make(size)
	if ok := len(*s) == size; !ok {
		t.Fatalf("len(*Slice) != %d", size)
	}
}

// TestMakeEach tests Slice.MakeEach.
func TestMakeEach(t *testing.T) {
	s := &slice.Slice[int]{}
	s.MakeEach(1, 2, 3, 4)
	for i, value := range *s {
		if ok := (*s)[i] == value; !ok {
			t.Fatalf("(*Slice)[%d] != %d", i, value)
		}
	}
}

// TestMakeEachReverse tests Slice.MakeEachReverse.
func TestMakeEachReverse(t *testing.T) {
	s := &slice.Slice[int]{}
	s.MakeEachReverse(1, 2, 3, 4)
	for i, value := range *s {
		if ok := (*s)[i] == value; !ok {
			t.Fatalf("(*Slice)[%d] != %d", i, value)
		}
	}
}

// TestMap tests Slice.Map.
func TestMap(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	x := []int{1, 2, 3, 4, 5}
	s.Map(func(_ int, value int) int {
		return value * 2
	})
	for i, value := range x {
		if ok := value*2 == (*s)[i]; !ok {
			t.Fatalf("%d != %d", value*2, (*s)[i])
		}
	}
}

// TestPoll tests Slice.Poll.
func TestPoll(t *testing.T) {
	s := &slice.Slice[int]{1}
	if value := s.Poll(); value != 1 {
		t.Fatalf("%d != 1", value)
	}
}

// TestPollLength tests Slice.PollLength.
func TestPollLength(t *testing.T) {
	s := &slice.Slice[int]{1}
	value, length := s.PollLength()
	if ok := value == 1; !ok {
		t.Fatalf("%d != 1", value)
	}
	if ok := length == 0; !ok {
		t.Fatalf("%d != 0", length)
	}
}

// TestPollOK tests Slice.PollOK.
func TestPollOK(t *testing.T) {
	s := &slice.Slice[int]{1, 2}
	value, ok := s.PollOK()
	if value != 1 {
		t.Fatalf("%d != 1", value)
	}
	if !ok {
		t.Fatalf("%t != true", ok)
	}
}

// TestPop tests Slice.Pop.
func TestPop(t *testing.T) {
	s := &slice.Slice[int]{1, 2}
	value := s.Pop()
	if ok := value == 2; !ok {
		t.Fatalf("%d != 2", value)
	}
}

// TestPopLength tests Slice.PopLength.
func TestPopLength(t *testing.T) {
	s := &slice.Slice[int]{1, 2}
	value, length := s.PopLength()
	if ok := value == 2; !ok {
		t.Fatalf("%d != 2", value)
	}
	if ok := length == len(*s); !ok {
		t.Fatalf("len(*Slice) != %d", len(*s))
	}
}

// TestPopOK tests Slice.PopOK.
func TestPopOK(t *testing.T) {
	s := &slice.Slice[int]{1, 2}
	value, ok := s.PopOK()
	if value != 2 {
		t.Fatalf("%d != 2", value)
	}
	if !ok {
		t.Fatalf("%t != true", ok)
	}
}

// TestPrecatenate tests Slice.Precatenate.
func TestPrecatenate(t *testing.T) {
	s := &slice.Slice[int]{}
	value := 1
	s.Precatenate(&slice.Slice[int]{value})
	if ok := (*s)[0] == 1; !ok {
		t.Fatalf("(*Slice)[0] != %d", value)
	}
}

// TestPrecatenateLength tests Slice.PrecatenateLength.
func TestPrecatenateLength(t *testing.T) {
	s := &slice.Slice[int]{}
	value := 1
	length := s.PrecatenateLength(&slice.Slice[int]{value})
	if ok := length == len(*s); !ok {
		t.Fatalf("len(*Slice) != %d", length)
	}
}

// TestPrepend tests Slice.Prepend.
func TestPrepend(t *testing.T) {
	s := &slice.Slice[int]{2}
	value := 1
	s.Prepend(value)
	if ok := (*s)[0] == value; !ok {
		t.Fatalf("(*Slice)[0] != %d", value)
	}
}

// TestPrependLength tests Slice.PrependLength.
func TestPrependLength(t *testing.T) {
	s := &slice.Slice[int]{}
	length := s.PrependLength(1, 2, 3, 4, 5)
	if ok := length == len(*s); !ok {
		t.Fatalf("%d != %d", length, len(*s))
	}
}

// TestReplace tests Slice.Replace.
func TestReplace(t *testing.T) {
	s := &slice.Slice[int]{1}
	s.Replace(0, 2)
	if ok := (*s)[0] == 2; !ok {
		t.Fatalf("%d != 2", (*s)[0])
	}
}

// TestReverse tests Slice.Reverse.
func TestReverse(t *testing.T) {
	s := &slice.Slice[int]{1, 2}
	s.Reverse()
	if ok := (*s)[0] == 2; !ok {
		t.Fatalf("(*Slice)[0] != %d", 2)
	}
}

// TestReduce tests Slice.Reduce.
func TestReduce(t *testing.T) {
	a := &slice.Slice[int]{1, 2, 3, 4}
	b := a.Reduce(func(i int, value int) bool { return value%2 == 0 })
	if ok := len(*a)/2 == len(*b); !ok {
		t.Fatal("len(*Slice) != 2")
	}
}

// TestSet tests Slice.Set.
func TestSet(t *testing.T) {
	s := &slice.Slice[int]{2, 2, 3, 3}
	s.Set()
	values := map[int]bool{}
	for _, value := range *s {
		if _, ok := values[value]; ok {
			t.Fatalf("Slice contains duplicate value %d", value)
		}
		values[value] = true
	}
}

// TestSlice tests Slice.Slice.
func TestSlice(t *testing.T) {
	s := &slice.Slice[int]{1, 2, 3}
	s = s.Slice(0, 2)
	if ok := len(*s) == 2; !ok {
		t.Fatalf("len(*Slice) != %d", 2)
	}

}

// TestSwap tests Slice.Swap.
func TestSwap(t *testing.T) {
	a := 1
	b := 2
	s := &slice.Slice[int]{a, b}
	s.Swap(0, 1)
	if ok := (*s)[0] == b; !ok {
		t.Fatalf("(*Slice)[0] != %d", b)
	}
	if ok := (*s)[1] == a; !ok {
		t.Fatalf("(*Slice)[1] != %d", a)
	}
}
