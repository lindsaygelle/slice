package slice_test

import (
	"reflect"
	"testing"

	"github.com/lindsaygelle/slice"
)

func TestAppend(t *testing.T) {
	// Test case 1: Append values to an empty slice.
	s := &slice.Slice[int]{}
	s.Append(1, 2, 3)

	expected := &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Append values to a non-empty slice.
	s = &slice.Slice[int]{1, 2}
	s.Append(3, 4, 5)

	expected = &slice.Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestAppendFunc(t *testing.T) {
	// Test case 1: Append values to an empty slice based on a function.
	s := &slice.Slice[int]{}
	fn := func(i int, value int) bool {
		return value%2 == 0
	}
	s.AppendFunc([]int{1, 2, 3, 4, 5}, fn)

	expected := &slice.Slice[int]{2, 4}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Append values to a non-empty slice based on a function.
	s = &slice.Slice[int]{1, 2, 3}
	fn = func(i int, value int) bool {
		return value > 1
	}
	s.AppendFunc([]int{1, 2, 3, 4, 5}, fn)

	expected = &slice.Slice[int]{1, 2, 3, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestAppendLength(t *testing.T) {
	// Test case 1: Append values to an empty slice and check the length.
	s := &slice.Slice[int]{}
	length := s.AppendLength(1, 2, 3)

	expected := &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
	if length != 3 {
		t.Errorf("Expected length 3, but got %d", length)
	}

	// Test case 2: Append values to a non-empty slice and check the length.
	s = &slice.Slice[int]{1, 2}
	length = s.AppendLength(3, 4, 5)

	expected = &slice.Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
	if length != 5 {
		t.Errorf("Expected length 5, but got %d", length)
	}
}

func TestBounds(t *testing.T) {
	// Test case 1: Check index within bounds.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	result := s.Bounds(2)
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Check index out of bounds.
	result = s.Bounds(10)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestClone(t *testing.T) {
	// Test case 1: Clone an empty slice.
	s := &slice.Slice[int]{}
	clone := s.Clone()

	if !reflect.DeepEqual(s, clone) {
		t.Errorf("Expected %v, but got %v", s, clone)
	}

	// Test case 2: Clone a non-empty slice.
	s = &slice.Slice[int]{1, 2, 3}
	clone = s.Clone()

	if !reflect.DeepEqual(s, clone) {
		t.Errorf("Expected %v, but got %v", s, clone)
	}
}

func TestConcatenate(t *testing.T) {
	// Test case 1: Concatenate with nil slice.
	s := &slice.Slice[int]{1, 2, 3}
	result := s.Concatenate(nil)

	expected := &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
	if result != s {
		t.Errorf("Expected result to be the same slice, but got a different slice")
	}

	// Test case 2: Concatenate with a non-empty slice.
	s = &slice.Slice[int]{1, 2}
	other := &slice.Slice[int]{3, 4, 5}
	result = s.Concatenate(other)

	expected = &slice.Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
	if result != s {
		t.Errorf("Expected result to be the same slice, but got a different slice")
	}
}

func TestConcatenateFunc(t *testing.T) {
	// Test case 1: Concatenate with nil slice using a function.
	s := &slice.Slice[int]{1, 2, 3}
	fn := func(i int, value int) bool {
		return value%2 == 0
	}
	result := s.ConcatenateFunc(nil, fn)

	expected := &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
	if result != s {
		t.Errorf("Expected result to be the same slice, but got a different slice")
	}

	// Test case 2: Concatenate with a non-empty slice using a function.
	s = &slice.Slice[int]{1, 2}
	other := &slice.Slice[int]{3, 4, 5}
	fn = func(i int, value int) bool {
		return value > 3
	}
	result = s.ConcatenateFunc(other, fn)

	expected = &slice.Slice[int]{1, 2, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
	if result != s {
		t.Errorf("Expected result to be the same slice, but got a different slice")
	}
}

func TestConcatenateLength(t *testing.T) {
	// Test case 1: Concatenate with nil slice and check the length.
	s := &slice.Slice[int]{1, 2, 3}
	length := s.ConcatenateLength(nil)

	expected := &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
	if length != 3 {
		t.Errorf("Expected length 3, but got %d", length)
	}

	// Test case 2: Concatenate with a non-empty slice and check the length.
	s = &slice.Slice[int]{1, 2}
	other := &slice.Slice[int]{3, 4, 5}
	length = s.ConcatenateLength(other)

	expected = &slice.Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
	if length != 5 {
		t.Errorf("Expected length 5, but got %d", length)
	}
}

func TestContains(t *testing.T) {
	// Test case 1: Check if value is present in the slice.
	s := &slice.Slice[string]{"apple", "banana", "orange"}
	result := s.Contains("banana")
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Check if value is not present in the slice.
	result = s.Contains("grape")
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestContainsMany(t *testing.T) {
	// Test case 1: Check multiple values present in the slice.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	result := s.ContainsMany(2, 4, 6)

	expected := &slice.Slice[bool]{true, true, false}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case 2: Check multiple values, some not present in the slice.
	result = s.ContainsMany(3, 5, 7)

	expected = &slice.Slice[bool]{true, true, false}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestDeduplicate(t *testing.T) {
	// Test case: Deduplicate a slice with duplicate values.
	s := &slice.Slice[int]{1, 2, 2, 3, 4, 4, 5}
	s.Deduplicate()

	expected := &slice.Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestDelete(t *testing.T) {
	// Test case: Delete a value at a specific index.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Delete(2)

	expected := &slice.Slice[int]{1, 2, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestDeleteFunc(t *testing.T) {
	// Test case: Delete values based on a function.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	fn := func(i int, value int) bool {
		return value%2 == 0
	}
	s.DeleteFunc(fn)

	expected := &slice.Slice[int]{1, 3, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestDeleteLength(t *testing.T) {
	// Test case: Delete a value at a specific index and check the length.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	length := s.DeleteLength(2)

	expected := &slice.Slice[int]{1, 2, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
	if length != 4 {
		t.Errorf("Expected length 4, but got %d", length)
	}
}

func TestDeleteOK(t *testing.T) {
	// Test case 1: Delete a value at a specific index successfully.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	result := s.DeleteOK(2)
	if !result {
		t.Errorf("Expected true, but got false")
	}

	expected := &slice.Slice[int]{1, 2, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Delete a value at an out-of-bounds index.
	result = s.DeleteOK(10)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestDeleteUnsafe(t *testing.T) {
	// Test case: Delete a value at a specific index without bounds checking.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.DeleteUnsafe(2)

	expected := &slice.Slice[int]{1, 2, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestEach(t *testing.T) {
	// Test case: Apply a function to each element in the slice.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	var sum int
	fn := func(i int, value int) {
		sum += value
	}
	s.Each(fn)

	if sum != 15 {
		t.Errorf("Expected sum 15, but got %d", sum)
	}
}

func TestEachBreak(t *testing.T) {
	// Test case 1: Apply a function to each element until the function returns false.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	var sum int
	fn := func(i int, value int) bool {
		sum += value
		return value%2 != 0
	}
	s.EachBreak(fn)

	if sum != 3 {
		t.Errorf("Expected sum 3, but got %d", sum)
	}

	// Test case 2: Apply a function to each element without breaking.
	sum = 0
	fn = func(i int, value int) bool {
		sum += value
		return true
	}
	s.EachBreak(fn)

	if sum != 15 {
		t.Errorf("Expected sum 15, but got %d", sum)
	}
}

func TestEachOK(t *testing.T) {
	// Test case 1: Apply a function to each element until the function returns false and check the result.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	var sum int
	fn := func(i int, value int) bool {
		sum += value
		return value%2 != 0
	}
	result := s.EachOK(fn)

	if sum != 3 {
		t.Errorf("Expected sum 3, but got %d", sum)
	}
	if result {
		t.Errorf("Expected false, but got true")
	}

	// Test case 2: Apply a function to each element without breaking and check the result.
	sum = 0
	fn = func(i int, value int) bool {
		sum += value
		return true
	}
	result = s.EachOK(fn)

	if sum != 15 {
		t.Errorf("Expected sum 15, but got %d", sum)
	}
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

func TestEachReverse(t *testing.T) {
	// Test case: Apply a function to each element in reverse order.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	var sum int
	fn := func(i int, value int) {
		sum += value
	}
	s.EachReverse(fn)

	if sum != 15 {
		t.Errorf("Expected sum 15, but got %d", sum)
	}
}

func TestEachReverseBreak(t *testing.T) {
	// Test case 1: Apply a function to each element in reverse order until the function returns false.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	var sum int
	fn := func(i int, value int) bool {
		sum += value
		return value%2 != 0
	}
	s.EachReverseBreak(fn)

	if sum != 9 {
		t.Errorf("Expected sum 9, but got %d", sum)
	}

	// Test case 2: Apply a function to each element in reverse order without breaking.
	sum = 0
	fn = func(i int, value int) bool {
		sum += value
		return true
	}
	s.EachReverseBreak(fn)

	if sum != 15 {
		t.Errorf("Expected sum 15, but got %d", sum)
	}
}

func TestEachReverseOK(t *testing.T) {
	// Test case 1: Apply a function to each element in reverse order until the function returns false and check the result.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	var sum int
	fn := func(i int, value int) bool {
		sum += value
		return value%2 != 0
	}
	result := s.EachReverseOK(fn)

	if sum != 9 {
		t.Errorf("Expected sum 9, but got %d", sum)
	}
	if result {
		t.Errorf("Expected false, but got true")
	}

	// Test case 2: Apply a function to each element in reverse order without breaking and check the result.
	sum = 0
	fn = func(i int, value int) bool {
		sum += value
		return true
	}
	result = s.EachReverseOK(fn)

	if sum != 15 {
		t.Errorf("Expected sum 15, but got %d", sum)
	}
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

func TestEqual(t *testing.T) {
	// Test case 1: Check equality of two equal slices.
	s1 := &slice.Slice[int]{1, 2, 3}
	s2 := &slice.Slice[int]{1, 2, 3}
	result := s1.Equal(s2)
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Check equality of two different slices.
	s1 = &slice.Slice[int]{1, 2, 3}
	s2 = &slice.Slice[int]{3, 2, 1}
	result = s1.Equal(s2)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestEqualFunc(t *testing.T) {
	// Test case 1: Check equality of two slices using a function.
	s1 := &slice.Slice[int]{1, 2, 3}
	s2 := &slice.Slice[int]{2, 4, 6}
	fn := func(i int, a int, b int) bool {
		return a*2 == b
	}
	result := s1.EqualFunc(s2, fn)
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Check equality of two slices using a different function.
	s1 = &slice.Slice[int]{1, 2, 3}
	s2 = &slice.Slice[int]{2, 4, 5}
	fn = func(i int, a int, b int) bool {
		return a*2 == b
	}
	result = s1.EqualFunc(s2, fn)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestEqualLength(t *testing.T) {
	// Test case 1: Check equality of two slices with equal lengths.
	s1 := &slice.Slice[int]{1, 2, 3}
	s2 := &slice.Slice[int]{1, 2, 3}
	result := s1.EqualLength(s2)
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Check equality of two slices with different lengths.
	s1 = &slice.Slice[int]{1, 2, 3}
	s2 = &slice.Slice[int]{1, 2}
	result = s1.EqualLength(s2)
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestFetch(t *testing.T) {
	// Test case 1: Fetch a value at a specific index successfully.
	s := &slice.Slice[string]{"apple", "banana", "orange"}
	value := s.Fetch(1)
	if value != "banana" {
		t.Errorf("Expected 'banana', but got '%s'", value)
	}

	// Test case 2: Fetch a value at an out-of-bounds index.
	value = s.Fetch(10)
	if value != "" {
		t.Errorf("Expected '', but got '%s'", value)
	}
}

func TestFetchLength(t *testing.T) {
	// Test case 1: Fetch a value at a specific index successfully and check the length.
	s := &slice.Slice[string]{"apple", "banana", "orange"}
	value, length := s.FetchLength(1)
	if value != "banana" {
		t.Errorf("Expected 'banana', but got '%s'", value)
	}
	if length != 3 {
		t.Errorf("Expected length 3, but got %d", length)
	}

	// Test case 2: Fetch a value at an out-of-bounds index.
	value, length = s.FetchLength(10)
	if value != "" {
		t.Errorf("Expected empty string, but got '%s'", value)
	}
	if length != 3 {
		t.Errorf("Expected length 3, but got %d", length)
	}
}

func TestIsEmpty(t *testing.T) {
	// Test case 1: Check if an empty slice is empty.
	s := &slice.Slice[int]{}
	result := s.IsEmpty()
	if !result {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Check if a non-empty slice is empty.
	s = &slice.Slice[int]{1, 2, 3}
	result = s.IsEmpty()
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestIsPopulated(t *testing.T) {
	// Test case 1: Check if an empty slice is not populated.
	s := &slice.Slice[int]{}
	result := s.IsPopulated()
	if result {
		t.Errorf("Expected false, but got true")
	}

	// Test case 2: Check if a non-empty slice is populated.
	s = &slice.Slice[int]{1, 2, 3}
	result = s.IsPopulated()
	if !result {
		t.Errorf("Expected true, but got false")
	}
}

func TestLength(t *testing.T) {
	// Test case 1: Get the length of an empty slice.
	s := &slice.Slice[int]{}
	length := s.Length()
	if length != 0 {
		t.Errorf("Expected length 0, but got %d", length)
	}

	// Test case 2: Get the length of a non-empty slice.
	s = &slice.Slice[int]{1, 2, 3, 4, 5}
	length = s.Length()
	if length != 5 {
		t.Errorf("Expected length 5, but got %d", length)
	}
}

func TestMake(t *testing.T) {
	// Test case 1: Make a slice with a specific length.
	s := &slice.Slice[int]{}
	s.Make(5)

	expected := &slice.Slice[int]{0, 0, 0, 0, 0}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Make a slice with a different length.
	s = &slice.Slice[int]{}
	s.Make(3)

	expected = &slice.Slice[int]{0, 0, 0}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestMakeEach(t *testing.T) {
	// Test case 1: Make a slice with specific values.
	s := &slice.Slice[int]{}
	s.MakeEach(1, 2, 3)

	expected := &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Make a slice with different values.
	s = &slice.Slice[int]{}
	s.MakeEach(4, 5, 6)

	expected = &slice.Slice[int]{4, 5, 6}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestMakeEachReverse(t *testing.T) {
	// Test case 1: Make a slice with specific values in reverse order.
	s := &slice.Slice[int]{}
	s.MakeEachReverse(1, 2, 3)

	expected := &slice.Slice[int]{3, 2, 1}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Make a slice with different values in reverse order.
	s = &slice.Slice[int]{}
	s.MakeEachReverse(4, 5, 6)

	expected = &slice.Slice[int]{6, 5, 4}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestMap(t *testing.T) {
	// Test case: Apply a mapping function to each element in the slice.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	fn := func(i int, value int) int {
		return value * 2
	}
	s = s.Map(fn)

	expected := &slice.Slice[int]{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestMapReverse(t *testing.T) {
	// Test case: Apply a mapping function to each element in reverse order.
	s := &slice.Slice[int]{5, 4, 3, 2, 1}
	fn := func(i int, value int) int {
		return value * 2
	}
	s = s.MapReverse(fn)

	expected := &slice.Slice[int]{10, 8, 6, 4, 2}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestModify(t *testing.T) {
	// Test case: Modify each element in the slice using a modifying function.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	fn := func(i int, value int) int {
		return value * 2
	}
	s.Modify(fn)

	expected := &slice.Slice[int]{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestModifyReverse(t *testing.T) {
	// Test case: Modify each element in the slice in reverse order using a modifying function.
	s := &slice.Slice[int]{5, 4, 3, 2, 1}
	fn := func(i int, value int) int {
		return value * 2
	}
	s.ModifyReverse(fn)

	expected := &slice.Slice[int]{10, 8, 6, 4, 2}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestPoll(t *testing.T) {
	// Test case 1: Poll a value from the end of the slice successfully.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	value := s.Poll()
	if value != 1 {
		t.Errorf("Expected 1, but got %d", value)
	}

	// Test case 2: Poll a value from an empty slice.
	s = &slice.Slice[int]{}
	value = s.Poll()
	if value != 0 {
		t.Errorf("Expected 0, but got %d", value)
	}
}

func TestPollLength(t *testing.T) {
	// Test case 1: Poll a value from the end of the slice successfully and check the length.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	value, length := s.PollLength()
	if value != 1 {
		t.Errorf("Expected 1, but got %d", value)
	}
	if length != 4 {
		t.Errorf("Expected length 4, but got %d", length)
	}

	// Test case 2: Poll a value from an empty slice and check the length.
	s = &slice.Slice[int]{}
	value, length = s.PollLength()
	if value != 0 {
		t.Errorf("Expected 0, but got %d", value)
	}
	if length != 0 {
		t.Errorf("Expected length 0, but got %d", length)
	}
}

func TestPollOK(t *testing.T) {
	// Test case 1: Poll a value from the end of the slice successfully and check the result.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	value, ok := s.PollOK()
	if value != 1 {
		t.Errorf("Expected 5, but got %d", value)
	}
	if !ok {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: Poll a value from an empty slice and check the result.
	s = &slice.Slice[int]{}
	value, ok = s.PollOK()
	if value != 0 {
		t.Errorf("Expected 0, but got %d", value)
	}
	if ok {
		t.Errorf("Expected false, but got true")
	}
}

func TestPop(t *testing.T) {
	// Test case 1: Pop a value from the end of the slice successfully.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	value := s.Pop()
	if value != 5 {
		t.Errorf("Expected 5, but got %d", value)
	}
	expected := &slice.Slice[int]{1, 2, 3, 4}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Pop a value from an empty slice.
	s = &slice.Slice[int]{}
	value = s.Pop()
	if value != 0 {
		t.Errorf("Expected 0, but got %d", value)
	}
	expected = &slice.Slice[int]{}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestPopLength(t *testing.T) {
	// Test case 1: Pop a value from the end of the slice successfully and check the length.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	value, length := s.PopLength()
	if value != 5 {
		t.Errorf("Expected 5, but got %d", value)
	}
	if length != 4 {
		t.Errorf("Expected length 4, but got %d", length)
	}
	expected := &slice.Slice[int]{1, 2, 3, 4}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Pop a value from an empty slice and check the length.
	s = &slice.Slice[int]{}
	value, length = s.PopLength()
	if value != 0 {
		t.Errorf("Expected 0, but got %d", value)
	}
	if length != 0 {
		t.Errorf("Expected length 0, but got %d", length)
	}
	expected = &slice.Slice[int]{}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestPopOK(t *testing.T) {
	// Test case 1: Pop a value from the end of the slice successfully and check the result.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	value, ok := s.PopOK()
	if value != 5 {
		t.Errorf("Expected 5, but got %d", value)
	}
	if !ok {
		t.Errorf("Expected true, but got false")
	}
	expected := &slice.Slice[int]{1, 2, 3, 4}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Pop a value from an empty slice and check the result.
	s = &slice.Slice[int]{}
	value, ok = s.PopOK()
	if value != 0 {
		t.Errorf("Expected 0, but got %d", value)
	}
	if ok {
		t.Errorf("Expected false, but got true")
	}
	expected = &slice.Slice[int]{}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestPrecatenate(t *testing.T) {
	// Test case 1: Precatenate with an empty slice.
	s := &slice.Slice[int]{1, 2, 3}
	other := &slice.Slice[int]{}
	s.Precatenate(other)

	expected := &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Precatenate with a non-empty slice.
	s = &slice.Slice[int]{3, 4, 5}
	other = &slice.Slice[int]{1, 2}
	s.Precatenate(other)

	expected = &slice.Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestPrecatenateFunc(t *testing.T) {
	// Test case 1: Precatenate with an empty slice using a function.
	s := &slice.Slice[int]{1, 2, 3}
	other := &slice.Slice[int]{}
	fn := func(i int, value int) bool {
		return value > 0
	}
	s.PrecatenateFunc(other, fn)

	expected := &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Precatenate with a non-empty slice using a function.
	s = &slice.Slice[int]{3, 2, 1}
	other = &slice.Slice[int]{4, 5}
	fn = func(i int, value int) bool {
		return value%2 == 0
	}
	s.PrecatenateFunc(other, fn)

	expected = &slice.Slice[int]{4, 3, 2, 1}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestPrecatenateLength(t *testing.T) {
	// Test case 1: Precatenate with an empty slice and check the length.
	s := &slice.Slice[int]{1, 2, 3}
	other := &slice.Slice[int]{}
	length := s.PrecatenateLength(other)
	if length == 0 {
		t.Errorf("Expected length 3, but got %d", length)
	}

	// Test case 2: Precatenate with a non-empty slice and check the length.
	s = &slice.Slice[int]{1, 2, 3}
	other = &slice.Slice[int]{4, 5}
	length = s.PrecatenateLength(other)
	if length == 2 {
		t.Errorf("Expected length 5, but got %d", length)
	}
}

func TestPrepend(t *testing.T) {
	// Test case 1: Prepend values to an empty slice.
	s := &slice.Slice[int]{}
	s.Prepend(3, 2, 1)

	expected := &slice.Slice[int]{3, 2, 1}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Prepend values to a non-empty slice.
	s = &slice.Slice[int]{4, 5}
	s.Prepend(1, 2, 3)

	expected = &slice.Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestPrependFunc(t *testing.T) {
	// Test case 1: Prepend values to an empty slice using a function.
	s := &slice.Slice[int]{}
	values := []int{3, 2, 1}
	fn := func(i int, value int) bool {
		return value > 0
	}
	s.PrependFunc(values, fn)

	expected := &slice.Slice[int]{3, 2, 1}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Prepend values to a non-empty slice using a function.
	s = &slice.Slice[int]{4, 5}
	values = []int{1, 2, 3}
	fn = func(i int, value int) bool {
		return true
	}
	s.PrependFunc(values, fn)

	expected = &slice.Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestPrependLength(t *testing.T) {
	// Test case 1: Prepend values to an empty slice and check the length.
	s := &slice.Slice[int]{}
	values := []int{3, 2, 1}
	length := s.PrependLength(values...)
	if length != 3 {
		t.Errorf("Expected length 3, but got %d", length)
	}

	// Test case 2: Prepend values to a non-empty slice and check the length.
	s = &slice.Slice[int]{4, 5}
	values = []int{3, 2, 1}
	length = s.PrependLength(values...)
	if length == 2 {
		t.Errorf("Expected length 2, but got %d", length)
	}
}

func TestReduce(t *testing.T) {
	// Test case: Reduce the elements of the slice using a reducing function.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	fn := func(i int, currentValue int, resultValue int) int {
		return resultValue + currentValue
	}
	result := s.Reduce(fn)

	if result != 15 {
		t.Errorf("Expected result 15, but got %d", result)
	}
}

func TestReduceReverse(t *testing.T) {
	// Test case: Reduce the elements of the slice in reverse order using a reducing function.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	fn := func(i int, currentValue int, resultValue int) int {
		return resultValue + currentValue
	}
	result := s.ReduceReverse(fn)

	if result != 15 {
		t.Errorf("Expected result 15, but got %d", result)
	}
}

func TestReplace(t *testing.T) {
	// Test case 1: Replace an element at a specific index successfully.
	s := &slice.Slice[string]{"apple", "banana", "orange"}
	result := s.Replace(1, "grape")

	expected := &slice.Slice[string]{"apple", "grape", "orange"}
	if !result || !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Replace an element at an out-of-bounds index.
	result = s.Replace(10, "melon")
	if result {
		t.Errorf("Expected false, but got true")
	}
}

func TestReverse(t *testing.T) {
	// Test case: Reverse the elements of the slice.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Reverse()

	expected := &slice.Slice[int]{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestShuffle(t *testing.T) {
	// Test case: Shuffle the elements of the slice.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Shuffle()

	// It's not practical to test the exact output of a shuffle, so we can only check if the length is the same.
	if s.Length() != 5 {
		t.Errorf("Expected length 5, but got %d", s.Length())
	}
}

func TestSlice(t *testing.T) {
	// Test case: Slice the elements of the slice from index 1 to 3.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	sliced := s.Slice(1, 3)

	expected := &slice.Slice[int]{2, 3, 4}
	if !reflect.DeepEqual(sliced, expected) {
		t.Errorf("Expected %v, but got %v", expected, sliced)
	}
}

func TestSortFunc(t *testing.T) {
	// Test case: Sort the elements of the slice using a sorting function.
	s := &slice.Slice[int]{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fn := func(i int, j int, a int, b int) bool {
		return a < b
	}
	s.SortFunc(fn)

	expected := &slice.Slice[int]{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestSplice(t *testing.T) {
	// Test case: Splice elements from index 1 to 4.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Splice(1, 4)

	expected := &slice.Slice[int]{2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}

func TestSplit(t *testing.T) {
	// Test case: Split the elements of the slice from index 2.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	part1, part2 := s.Split(2)

	expectedPart1 := &slice.Slice[int]{1, 2}
	expectedPart2 := &slice.Slice[int]{3, 4, 5}
	if !reflect.DeepEqual(part1, expectedPart1) || !reflect.DeepEqual(part2, expectedPart2) {
		t.Errorf("Expected %v and %v, but got %v and %v", expectedPart1, expectedPart2, part1, part2)
	}
}

func TestSplitAtIndex(t *testing.T) {
	// Test case 1: Split the elements of the slice from index 2 successfully.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	part1, part2, ok := s.SplitOK(2)

	expectedPart1 := &slice.Slice[int]{1, 2}
	expectedPart2 := &slice.Slice[int]{3, 4, 5}
	if !ok || !reflect.DeepEqual(part1, expectedPart1) || !reflect.DeepEqual(part2, expectedPart2) {
		t.Errorf("Expected %v and %v with ok true, but got %v and %v with ok %v", expectedPart1, expectedPart2, part1, part2, ok)
	}

	// Test case 2: Split the elements of the slice from an out-of-bounds index.
	part1, part2, ok = s.SplitOK(10)
	if ok {
		t.Errorf("Expected ok false, but got true")
	}
}

func TestSplitFunc(t *testing.T) {
	// Test case: Split the elements of the slice using a function.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	fn := func(i int, value int) bool {
		return value%2 == 0
	}
	part1, part2 := s.SplitFunc(fn)

	expectedPart1 := &slice.Slice[int]{2, 4}
	expectedPart2 := &slice.Slice[int]{1, 3, 5}
	if !reflect.DeepEqual(part1, expectedPart1) || !reflect.DeepEqual(part2, expectedPart2) {
		t.Errorf("Expected %v and %v, but got %v and %v", expectedPart1, expectedPart2, part1, part2)
	}
}

func TestSwap(t *testing.T) {
	// Test case 1: Swap elements at valid indices.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Swap(1, 3)

	expected := &slice.Slice[int]{1, 4, 3, 2, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}

	// Test case 2: Swap elements at out-of-bounds indices.
	s.Swap(10, 20)
	// No changes should occur.
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected %v, but got %v", expected, s)
	}
}
