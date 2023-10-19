package slice_test

import (
	"reflect"
	"testing"

	"github.com/lindsaygelle/slice"
)

// TestAppend tests Slice.Append.
func TestAppend(t *testing.T) {
	// Test case 1: Append values to an empty slice.
	input := &slice.Slice[int]{}
	expectedOutput := &slice.Slice[int]{1, 2, 3}
	input.Append(1, 2, 3)

	if !reflect.DeepEqual(input, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, input)
	}

	// Test case 2: Append values to a non-empty slice.
	input = &slice.Slice[int]{1, 2, 3}
	expectedOutput = &slice.Slice[int]{1, 2, 3, 4, 5}
	input.Append(4, 5)

	if !reflect.DeepEqual(input, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, input)
	}
}

// TestAppendFunc tests Slice.AppendFunc.
func TestAppendFunc(t *testing.T) {
	// Test case 1: Append even numbers.
	input := &slice.Slice[int]{}
	input.AppendFunc(func(i int, value int) bool {
		return value%2 == 0 // Append even numbers to the Slice.
	}, 1, 2, 3, 4, 5)

	expectedOutput := &slice.Slice[int]{2, 4}
	if !reflect.DeepEqual(input, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, input)
	}

	// Test case 2: Append odd numbers.
	input = &slice.Slice[int]{}
	input.AppendFunc(func(i int, value int) bool {
		return value%2 != 0 // Append odd numbers to the Slice.
	}, 1, 2, 3, 4, 5)

	expectedOutput = &slice.Slice[int]{1, 3, 5}
	if !reflect.DeepEqual(input, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, input)
	}
}

// TestAppendLength tests Slice.AppendLength.
func TestAppendLength(t *testing.T) {
	// Test case 1: Append values to an empty slice.
	input := &slice.Slice[int]{}
	length := input.AppendLength(1, 2, 3)

	expectedOutput := 3
	if length != expectedOutput {
		t.Errorf("Expected length %d, but got %d", expectedOutput, length)
	}

	// Test case 2: Append values to a non-empty slice.
	input = &slice.Slice[int]{1, 2, 3}
	length = input.AppendLength(4, 5)

	expectedOutput = 5
	if length != expectedOutput {
		t.Errorf("Expected length %d, but got %d", expectedOutput, length)
	}

	// Test case 3: Append no values (length should remain the same).
	input = &slice.Slice[int]{1, 2, 3}
	length = input.AppendLength()

	expectedOutput = 3
	if length != expectedOutput {
		t.Errorf("Expected length %d, but got %d", expectedOutput, length)
	}
}

// TestBounds tests Slice.Bounds.
func TestBounds(t *testing.T) {
	// Test case 1: Index within bounds.
	input := &slice.Slice[int]{1, 2, 3}
	inBounds := input.Bounds(1)
	expectedInBounds := true
	if inBounds != expectedInBounds {
		t.Errorf("Expected %t, but got %t", expectedInBounds, inBounds)
	}

	// Test case 2: Index out of bounds (negative index).
	outOfBounds := input.Bounds(-1)
	expectedOutOfBounds := false
	if outOfBounds != expectedOutOfBounds {
		t.Errorf("Expected %t, but got %t", expectedOutOfBounds, outOfBounds)
	}

	// Test case 3: Index out of bounds (index greater than the length of the slice).
	outOfBounds = input.Bounds(5)
	if outOfBounds != expectedOutOfBounds {
		t.Errorf("Expected %t, but got %t", expectedOutOfBounds, outOfBounds)
	}

	// Test case 4: Index within bounds (edge case: index equal to the length of the slice).
	inBounds = input.Bounds(3)
	expectedInBounds = false // Index 3 is out of bounds for a slice with length 3.
	if inBounds != expectedInBounds {
		t.Errorf("Expected %t, but got %t", expectedInBounds, inBounds)
	}
}

// TestClone tests Slice.Clone.
func TestClone(t *testing.T) {
	// Test case 1: Clone an empty slice.
	emptySlice := &slice.Slice[int]{}
	clonedEmptySlice := emptySlice.Clone()

	if !reflect.DeepEqual(emptySlice, clonedEmptySlice) {
		t.Errorf("Clone of an empty slice should be empty, but got %v", clonedEmptySlice)
	}

	// Test case 2: Clone a non-empty slice.
	numbers := &slice.Slice[int]{1, 2, 3, 4, 5}
	clonedSlice := numbers.Clone()

	// Expected cloned slice: [1, 2, 3, 4, 5].
	expected := []int{1, 2, 3, 4, 5}

	// Modify the cloned slice to check if it is independent of the original slice.
	clonedSlice.Append(6)

	// Check if the original slice remains unchanged.
	if reflect.DeepEqual(numbers, expected) {
		t.Errorf("Original slice should remain unchanged after cloning, but got %v", numbers)
	}
}

// TestConcatenate tests Slice.Concatenate.
func TestConcatenate(t *testing.T) {
	// Test case 1: Concatenating with a non-nil slice.
	s1 := &slice.Slice[int]{1, 2, 3}
	s2 := &slice.Slice[int]{4, 5}
	s1.Concatenate(s2)

	expectedOutput := &slice.Slice[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s1, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s1)
	}

	// Test case 2: Concatenating with a nil slice (no modification should occur).
	s1 = &slice.Slice[int]{1, 2, 3}
	s2 = &slice.Slice[int]{} // nil slice.
	s1.Concatenate(s2)

	expectedOutput = &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s1, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s1)
	}

	// Test case 3: Concatenating with an empty slice.
	s1 = &slice.Slice[int]{1, 2, 3}
	s2 = &slice.Slice[int]{} // empty slice.
	s1.Concatenate(s2)

	expectedOutput = &slice.Slice[int]{1, 2, 3} // no modification should occur.
	if !reflect.DeepEqual(s1, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s1)
	}
}

// TestConcatenateFunc tests Slice.ConcatenateFunc.
func TestConcatenateFunc(t *testing.T) {
	// Test case 1: Concatenating with a non-nil slice, filtering even numbers.
	s1 := &slice.Slice[int]{1, 2, 3}
	s2 := &slice.Slice[int]{4, 5, 6}
	s1.ConcatenateFunc(s2, func(i int, value int) bool {
		return value%2 == 0
	})

	expectedOutput := &slice.Slice[int]{1, 2, 3, 4, 6}
	if !reflect.DeepEqual(s1, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s1)
	}

	// Test case 2: Concatenating with a nil slice (no modification should occur).
	s1 = &slice.Slice[int]{1, 2, 3}
	s2 = &slice.Slice[int]{} // nil slice.
	s1.ConcatenateFunc(s2, func(i int, value int) bool {
		return true
	})

	expectedOutput = &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s1, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s1)
	}

	// Test case 3: Concatenating with an empty slice (no modification should occur).
	s1 = &slice.Slice[int]{1, 2, 3}
	s2 = &slice.Slice[int]{} // empty slice.
	s1.ConcatenateFunc(s2, func(i int, value int) bool {
		return true
	})

	expectedOutput = &slice.Slice[int]{1, 2, 3}
	if !reflect.DeepEqual(s1, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s1)
	}
}

// TestConcatenateLength tests Slice.ConcatenateLength.
func TestConcatenateLength(t *testing.T) {
	// Test case 1: Concatenating with a non-nil slice.
	s1 := &slice.Slice[int]{1, 2, 3}
	s2 := &slice.Slice[int]{4, 5}
	length := s1.ConcatenateLength(s2)

	expectedOutput := 5
	if length != expectedOutput {
		t.Errorf("Expected length %d, but got %d", expectedOutput, length)
	}

	// Test case 2: Concatenating with a nil slice (length should remain the same).
	s1 = &slice.Slice[int]{1, 2, 3}
	s2 = &slice.Slice[int]{} // nil slice.
	length = s1.ConcatenateLength(s2)

	expectedOutput = 3
	if length != expectedOutput {
		t.Errorf("Expected length %d, but got %d", expectedOutput, length)
	}

	// Test case 3: Concatenating with an empty slice (length should remain the same).
	s1 = &slice.Slice[int]{1, 2, 3}
	s2 = &slice.Slice[int]{} // empty slice.
	length = s1.ConcatenateLength(s2)

	expectedOutput = 3
	if length != expectedOutput {
		t.Errorf("Expected length %d, but got %d", expectedOutput, length)
	}
}

// TestContains tests Slice.Contains.
func TestContains(t *testing.T) {
	// Test case 1: Value is present in the slice.
	s := &slice.Slice[string]{"apple", "banana", "cherry"}
	containsBanana := s.Contains("banana")
	expectedResult := true

	if containsBanana != expectedResult {
		t.Errorf("Expected %t, but got %t", expectedResult, containsBanana)
	}

	// Test case 2: Value is not present in the slice.
	containsGrapes := s.Contains("grapes")
	expectedResult = false

	if containsGrapes != expectedResult {
		t.Errorf("Expected %t, but got %t", expectedResult, containsGrapes)
	}

	// Test case 3: Value is present in an empty slice.
	emptySlice := &slice.Slice[string]{}
	containsEmpty := emptySlice.Contains("")
	expectedResult = false

	if containsEmpty != expectedResult {
		t.Errorf("Expected %t, but got %t", expectedResult, containsEmpty)
	}
}

// TestContainsMany tests Slice.ContainsMany.
func TestContainsMany(t *testing.T) {
	// Test case 1: Multiple values are present in the slice.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	result := s.ContainsMany(2, 4, 6)
	expectedResult := &slice.Slice[bool]{true, true, false}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}

	// Test case 2: All values are present in the slice.
	result = s.ContainsMany(2, 4, 1)
	expectedResult = &slice.Slice[bool]{true, true, true}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}

	// Test case 3: None of the values are present in the slice.
	result = s.ContainsMany(6, 7, 8)
	expectedResult = &slice.Slice[bool]{false, false, false}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}

	// Test case 4: Empty slice (no values should be found).
	emptySlice := &slice.Slice[int]{}
	result = emptySlice.ContainsMany(1, 2, 3)
	expectedResult = &slice.Slice[bool]{false, false, false}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected %v, but got %v", expectedResult, result)
	}
}

// TestDelete tests Slice.Delete.
func TestDelete(t *testing.T) {
	// Test case 1: Delete element within bounds.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Delete(2)

	expectedOutput := &slice.Slice[int]{1, 2, 4, 5}
	if !reflect.DeepEqual(s, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s)
	}

	// Test case 2: Delete element out of bounds (no modification should occur).
	s = &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Delete(5) // Index out of bounds.

	expectedOutput = &slice.Slice[int]{1, 2, 3, 4, 5} // No modification expected.
	if !reflect.DeepEqual(s, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s)
	}

	// Test case 3: Delete element out of bounds (negative index, no modification should occur).
	s = &slice.Slice[int]{1, 2, 3, 4, 5}
	s.Delete(-1) // Index out of bounds.

	expectedOutput = &slice.Slice[int]{1, 2, 3, 4, 5} // No modification expected.
	if !reflect.DeepEqual(s, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s)
	}
}

// TestDeleteFunc tests Slice.DeleteFunc.
func TestDeleteFunc(t *testing.T) {
	// Test case 1: Delete even numbers.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	even := func(i int, value int) bool { return value%2 == 0 }
	s.DeleteFunc(even)

	expectedOutput := &slice.Slice[int]{1, 3, 5}
	if !reflect.DeepEqual(s, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s)
	}

	// Test case 2: Delete odd numbers.
	s = &slice.Slice[int]{1, 2, 3, 4, 5}
	odd := func(i int, value int) bool { return value%2 != 0 }
	s.DeleteFunc(odd)

	expectedOutput = &slice.Slice[int]{2, 4}
	if !reflect.DeepEqual(s, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s)
	}

	// Test case 3: Delete all elements (empty slice).
	s = &slice.Slice[int]{1, 2, 3, 4, 5}
	all := func(i int, value int) bool { return true }
	s.DeleteFunc(all)

	expectedOutput = &slice.Slice[int]{}
	if !reflect.DeepEqual(s, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s)
	}
}

// TestDeleteLength tests Slice.DeleteLength.
func TestDeleteLength(t *testing.T) {
	// Test case 1: Delete element within bounds.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	newLength := s.DeleteLength(2)

	expectedOutput := 4
	if newLength != expectedOutput {
		t.Errorf("Expected new length %d, but got %d", expectedOutput, newLength)
	}

	// Test case 2: Delete element out of bounds (no modification should occur).
	s = &slice.Slice[int]{1, 2, 3, 4, 5}
	newLength = s.DeleteLength(5) // Index out of bounds.

	expectedOutput = 5 // Length remains the same.
	if newLength != expectedOutput {
		t.Errorf("Expected new length %d, but got %d", expectedOutput, newLength)
	}

	// Test case 3: Delete element out of bounds (negative index, no modification should occur).
	s = &slice.Slice[int]{1, 2, 3, 4, 5}
	newLength = s.DeleteLength(-1) // Index out of bounds.

	expectedOutput = 5 // Length remains the same.
	if newLength != expectedOutput {
		t.Errorf("Expected new length %d, but got %d", expectedOutput, newLength)
	}
}

// TestDeleteOK tests Slice.DeleteOK.
func TestDeleteOK(t *testing.T) {
	// Test case 1: Delete element within bounds (successful deletion).
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	deleted := s.DeleteOK(2)

	expectedOutput := true
	if deleted != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, deleted)
	}

	expectedSlice := &slice.Slice[int]{1, 2, 4, 5}
	if !s.Equal(expectedSlice) {
		t.Errorf("Expected %v, but got %v", expectedSlice, s)
	}

	// Test case 2: Delete element out of bounds (no modification should occur).
	s = &slice.Slice[int]{1, 2, 3, 4, 5}
	deleted = s.DeleteOK(5) // Index out of bounds.

	expectedOutput = false
	if deleted != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, deleted)
	}

	expectedSlice = &slice.Slice[int]{1, 2, 3, 4, 5} // No modification expected.
	if !s.Equal(expectedSlice) {
		t.Errorf("Expected %v, but got %v", expectedSlice, s)
	}

	// Test case 3: Delete element out of bounds (negative index, no modification should occur).
	s = &slice.Slice[int]{1, 2, 3, 4, 5}
	deleted = s.DeleteOK(-1) // Index out of bounds.

	expectedOutput = false
	if deleted != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, deleted)
	}

	expectedSlice = &slice.Slice[int]{1, 2, 3, 4, 5} // No modification expected.
	if !s.Equal(expectedSlice) {
		t.Errorf("Expected %v, but got %v", expectedSlice, s)
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

// TestEqual tests Slice.Equal.
func TestEqual(t *testing.T) {
	// Test case 1: Equal slices.
	slice1 := &slice.Slice[int]{1, 2, 3}
	slice2 := &slice.Slice[int]{1, 2, 3}

	equal := slice1.Equal(slice2)
	expectedOutput := true
	if equal != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, equal)
	}

	// Test case 2: Unequal slices (different length).
	slice3 := &slice.Slice[int]{1, 2}
	equal = slice1.Equal(slice3)
	expectedOutput = false
	if equal != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, equal)
	}

	// Test case 3: Unequal slices (different elements).
	slice4 := &slice.Slice[int]{1, 2, 4}
	equal = slice1.Equal(slice4)
	expectedOutput = false
	if equal != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, equal)
	}

	// Test case 4: Equal slices with different types (should return false).
	slice5 := &slice.Slice[interface{}]{1, uint(2), 3}
	slice6 := &slice.Slice[interface{}]{1, 2, 3}
	equal = slice5.Equal(slice6)
	expectedOutput = false
	if equal != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, equal)
	}
}

// TestEqualFunc test Slice.EqualFunc.
func TestEqualFunc(t *testing.T) {
	// Test case 1: Equal slices using a custom comparison function.
	slice1 := &slice.Slice[int]{1, 2, 3}
	slice2 := &slice.Slice[int]{2, 4, 6}
	customEqual := func(i int, value1, value2 int) bool {
		return value1*2 == value2
	}

	equal := slice1.EqualFunc(slice2, customEqual)
	expectedOutput := true
	if equal != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, equal)
	}

	// Test case 2: Unequal slices using a custom comparison function.
	slice3 := &slice.Slice[int]{1, 3, 5}
	equal = slice1.EqualFunc(slice3, customEqual)
	expectedOutput = false
	if equal != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, equal)
	}

	// Test case 3: Equal slices with different values.
	slice4 := &slice.Slice[int]{9, 12, 51}
	equal = slice1.EqualFunc(slice4, func(i int, value1, value2 int) bool {
		return value1*2 == value2
	})
	expectedOutput = false
	if equal != expectedOutput {
		t.Errorf("Expected %t, but got %t", expectedOutput, equal)
	}
}

// TestFetch tests Slice.Fetch.
func TestFetch(t *testing.T) {
	// Test case: Fetch element at a valid index.
	s := &slice.Slice[string]{"apple", "banana", "cherry"}
	fruit := s.Fetch(1)
	expectedFruit := "banana"

	if fruit != expectedFruit {
		t.Errorf("Expected %s, but got %s", expectedFruit, fruit)
	}

	// Test case: Fetch element at an invalid index (out of bounds).
	invalidIndex := -1
	fruit = s.Fetch(invalidIndex) // This line should return default empty due to an out-of-bounds index.
	if fruit != "" {
		t.Errorf("Expected %s, but go %s", "", fruit)
	}
}

// TestFetchLength tests Slice.FetchLength.
func TestFetchLength(t *testing.T) {
	// Test case: Fetch element at a valid index and check length.
	s := &slice.Slice[int]{10, 20, 30, 40, 50}
	index := 2
	expectedValue := 30
	expectedLength := 5

	value, length := s.FetchLength(index)

	if value != expectedValue {
		t.Errorf("Expected value %d, but got %d", expectedValue, value)
	}

	if length != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, length)
	}
}

// TestFilter tests Slice.Filter.
func TestFilter(t *testing.T) {
	// Test case: Filter even numbers.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}
	filtered := s.Filter(func(i int, value int) bool {
		return value%2 == 0 // Keep only even numbers.
	})

	expectedFilteredSlice := &slice.Slice[int]{2, 4}
	if !reflect.DeepEqual(filtered, expectedFilteredSlice) {
		t.Errorf("Expected filtered slice %v, but got %v", expectedFilteredSlice, filtered)
	}

	// Test case: Filter using a custom predicate (e.g., greater than 2).
	customPredicate := func(i int, value int) bool {
		return value > 2
	}
	customFiltered := s.Filter(customPredicate)
	expectedCustomFilteredSlice := &slice.Slice[int]{3, 4, 5}
	if !reflect.DeepEqual(customFiltered, expectedCustomFilteredSlice) {
		t.Errorf("Expected filtered slice %v, but got %v", expectedCustomFilteredSlice, customFiltered)
	}
}

// TestFindIndex tests Slice.FindIndex.
func TestFindIndex(t *testing.T) {
	// Create a test slice with some data.
	s := &slice.Slice[int]{1, 2, 3, 4, 5}

	// Test for a predicate function that matches an element.
	index, found := s.FindIndex(func(value int) bool {
		return value == 3
	})
	expectedIndex := 2
	if !found || index != expectedIndex {
		t.Fatalf("FindIndex returned (%d, %v), expected (%d, true)", index, found, expectedIndex)
	}

	// Test for a predicate function that doesn't match any element.
	index, found = s.FindIndex(func(value int) bool {
		return value == 6
	})
	expectedIndex = -1
	if found || index != expectedIndex {
		t.Fatalf("FindIndex returned (%d, %v), expected (%d, false)", index, found, expectedIndex)
	}

	// Test for a custom type (e.g., a struct).
	type CustomStruct struct {
		Name string
		Age  int
	}

	customSlice := &slice.Slice[CustomStruct]{
		{"Alice", 25},
		{"Bob", 30},
	}

	// Define a custom predicate function.
	predicate := func(value CustomStruct) bool {
		return value.Name == "Bob"
	}

	// Test for a predicate function that matches an element in the custom slice.
	customIndex, customFound := customSlice.FindIndex(predicate)
	expectedCustomIndex := 1
	if !customFound || customIndex != expectedCustomIndex {
		t.Fatalf("FindIndex returned (%d, %v), expected (%d, true)", customIndex, customFound, expectedCustomIndex)
	}

	// Test for a predicate function that doesn't match any element in the custom slice.
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
	// Test case: Get element at a valid index.
	s := &slice.Slice[float64]{3.14, 2.71, 1.61}
	index := 1
	expectedValue := 2.71
	expectedOK := true

	value, ok := s.Get(index)

	if value != expectedValue {
		t.Errorf("Expected value %f, but got %f", expectedValue, value)
	}

	if ok != expectedOK {
		t.Errorf("Expected OK %t, but got %t", expectedOK, ok)
	}

	// Test case: Get element at an invalid index (out of bounds).
	invalidIndex := -1
	expectedOK = false

	value, ok = s.Get(invalidIndex)

	if value != 0 {
		t.Errorf("Expected value 0, but got %f", value)
	}

	if ok != expectedOK {
		t.Errorf("Expected OK %t, but got %t", expectedOK, ok)
	}
}

// TestGetLength tests Slice.GetLength.
func TestGetLength(t *testing.T) {
	// Test case: Get element at a valid index and check length.
	s := &slice.Slice[int]{10, 20, 30, 40, 50}
	index := 2
	expectedValue := 30
	expectedLength := 5
	expectedOK := true

	value, length, ok := s.GetLength(index)

	if value != expectedValue {
		t.Errorf("Expected value %d, but got %d", expectedValue, value)
	}

	if length != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, length)
	}

	if ok != expectedOK {
		t.Errorf("Expected OK %t, but got %t", expectedOK, ok)
	}

	// Test case: Get element at an invalid index (out of bounds).
	invalidIndex := -1
	expectedOK = false

	value, length, ok = s.GetLength(invalidIndex)

	if value != 0 {
		t.Errorf("Expected value 0, but got %d", value)
	}

	if length != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, length)
	}

	if ok != expectedOK {
		t.Errorf("Expected OK %t, but got %t", expectedOK, ok)
	}
}

// TestIsEmpty tests Slice.IsEmpty.
func TestIsEmpty(t *testing.T) {
	// Test case: Empty slice.
	emptySlice := &slice.Slice[int]{}
	expectedEmptyResult := true
	isEmpty := emptySlice.IsEmpty()

	if isEmpty != expectedEmptyResult {
		t.Errorf("Expected IsEmpty() to return %t for an empty slice, but got %t", expectedEmptyResult, isEmpty)
	}

	// Test case: Non-empty slice.
	nonEmptySlice := &slice.Slice[int]{1, 2, 3}
	expectedNonEmptyResult := false
	isEmpty = nonEmptySlice.IsEmpty()

	if isEmpty != expectedNonEmptyResult {
		t.Errorf("Expected IsEmpty() to return %t for a non-empty slice, but got %t", expectedNonEmptyResult, isEmpty)
	}
}

// TestIsPopulated tests Slice.IsPopulated.
func TestIsPopulated(t *testing.T) {
	// Test case: Populated slice.
	populatedSlice := &slice.Slice[int]{1, 2, 3}
	expectedPopulatedResult := true
	isPopulated := populatedSlice.IsPopulated()

	if isPopulated != expectedPopulatedResult {
		t.Errorf("Expected IsPopulated() to return %t for a populated slice, but got %t", expectedPopulatedResult, isPopulated)
	}

	// Test case: Empty slice.
	emptySlice := &slice.Slice[int]{}
	expectedEmptyResult := false
	isPopulated = emptySlice.IsPopulated()

	if isPopulated != expectedEmptyResult {
		t.Errorf("Expected IsPopulated() to return %t for an empty slice, but got %t", expectedEmptyResult, isPopulated)
	}
}

// TestLength tests Slice.Length.
func TestLength(t *testing.T) {
	// Test case: Populated slice.
	populatedSlice := &slice.Slice[int]{1, 2, 3}
	expectedLength := 3
	length := populatedSlice.Length()

	if length != expectedLength {
		t.Errorf("Expected Length() to return %d for a populated slice, but got %d", expectedLength, length)
	}

	// Test case: Empty slice.
	emptySlice := &slice.Slice[int]{}
	expectedEmptyLength := 0
	emptyLength := emptySlice.Length()

	if emptyLength != expectedEmptyLength {
		t.Errorf("Expected Length() to return %d for an empty slice, but got %d", expectedEmptyLength, emptyLength)
	}
}

// TestMake tests Slice.Make.
func TestMake(t *testing.T) {
	// Test case: Make a new empty slice with length 3.
	s := &slice.Slice[int]{10, 20, 30}
	expectedLength := 3
	sliceLength := s.Make(expectedLength).Length()

	if sliceLength != expectedLength {
		t.Errorf("Expected Make() to create a new empty slice with length %d, but got length %d", expectedLength, sliceLength)
	}

	// Test case: Make a new empty slice with length 0.
	s = &slice.Slice[int]{10, 20, 30}
	expectedEmptyLength := 0
	emptySliceLength := s.Make(expectedEmptyLength).Length()

	if emptySliceLength != expectedEmptyLength {
		t.Errorf("Expected Make() to create a new empty slice with length %d, but got length %d", expectedEmptyLength, emptySliceLength)
	}
}

// TestMakeEach tests Slice.MakeEach.
func TestMakeEach(t *testing.T) {
	// Test case: Make a new slice with provided values.
	s := &slice.Slice[int]{}
	expectedValues := []int{10, 20, 30}
	s.MakeEach(expectedValues...)
	sliceLength := s.Length()

	if sliceLength != len(expectedValues) {
		t.Errorf("Expected MakeEach() to create a new slice with length %d, but got length %d", len(expectedValues), sliceLength)
	}

	// Test case: Make a new empty slice with no provided values.
	emptySlice := &slice.Slice[int]{}
	emptySlice.MakeEach()
	emptySliceLength := emptySlice.Length()

	if emptySliceLength != 0 {
		t.Errorf("Expected MakeEach() with no values to create an empty slice, but got length %d", emptySliceLength)
	}
}

// TestMakeEachReverse tests Slice.MakeEachReverse.
func TestMakeEachReverse(t *testing.T) {
	// Test case 1: MakeEachReverse with values.
	s := &slice.Slice[int]{}
	s.MakeEachReverse(10, 20, 30)

	expectedOutput := &slice.Slice[int]{30, 20, 10}

	if !reflect.DeepEqual(s, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s)
	}

	// Test case 2: MakeEachReverse with empty slice.
	s = &slice.Slice[int]{}
	s.MakeEachReverse()

	expectedOutput = &slice.Slice[int]{}

	if !reflect.DeepEqual(s, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s)
	}

	// Test case 3: MakeEachReverse with single value.
	s = &slice.Slice[int]{}
	s.MakeEachReverse(42)

	expectedOutput = &slice.Slice[int]{42}

	if !reflect.DeepEqual(s, expectedOutput) {
		t.Errorf("Expected %v, but got %v", expectedOutput, s)
	}
}

// TestMap tests Slice.Map.
func TestMap(t *testing.T) {
	// Test case: Map a function to square the elements in the slice.
	numbers := &slice.Slice[int]{1, 2, 3, 4, 5}

	// Clone the original slice for comparison.
	originalSlice := *numbers

	squared := numbers.Map(func(i int, value int) int {
		return value * value
	})

	expectedSquared := []int{1, 4, 9, 16, 25}

	if squared.Length() != len(expectedSquared) {
		t.Errorf("Expected squared slice length to be %d, but got %d", len(expectedSquared), squared.Length())
	}

	for i, val := range expectedSquared {
		retrievedVal, ok := squared.Get(i)
		if !ok {
			t.Errorf("Expected squared slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected squared slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Check that the original slice is not modified.
	if originalSlice.Equal(squared) {
		t.Errorf("Original slice is modified after Map operation")
	}

	// Test case: Map a function to double the elements in the slice.
	numbers = &slice.Slice[int]{2, 4, 6, 8, 10}

	// Clone the original slice for comparison.
	originalSlice = *numbers

	doubled := numbers.Map(func(i int, value int) int {
		return value * 2
	})

	expectedDoubled := []int{4, 8, 12, 16, 20}

	if doubled.Length() != len(expectedDoubled) {
		t.Errorf("Expected doubled slice length to be %d, but got %d", len(expectedDoubled), doubled.Length())
	}

	for i, val := range expectedDoubled {
		retrievedVal, ok := doubled.Get(i)
		if !ok {
			t.Errorf("Expected doubled slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected doubled slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Check that the original slice is not modified.
	if !originalSlice.Equal(numbers) {
		t.Errorf("Original slice is modified after Map operation")
	}
}

// TestMapReverse tests Slice.MapReverse.
func TestMapReverse(t *testing.T) {
	// Test case: Map a function to square the elements in reverse order in the slice.
	numbers := &slice.Slice[int]{1, 2, 3, 4, 5}

	// Clone the original slice for comparison.
	originalSlice := *numbers

	reversedSquared := numbers.MapReverse(func(i int, value int) int {
		return value * value
	})

	expectedReversedSquared := []int{25, 16, 9, 4, 1}

	if reversedSquared.Length() != len(expectedReversedSquared) {
		t.Errorf("Expected reversed squared slice length to be %d, but got %d", len(expectedReversedSquared), reversedSquared.Length())
	}

	for i, val := range expectedReversedSquared {
		retrievedVal, ok := reversedSquared.Get(i)
		if !ok {
			t.Errorf("Expected reversed squared slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected reversed squared slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Check that the original slice is not modified.
	if !numbers.Equal(&originalSlice) {
		t.Errorf("Original slice is modified after MapReverse operation")
	}
}

// TestModify tests Slice.Modify.
func TestModify(t *testing.T) {
	// Test case: Modify the elements in the slice by doubling them.
	numbers := &slice.Slice[int]{1, 2, 3, 4, 5}

	// Modify the slice by doubling the elements.
	modifiedSlice := numbers.Modify(func(i int, value int) int {
		return value * 2
	})

	expectedModifiedSlice := []int{2, 4, 6, 8, 10}

	if modifiedSlice.Length() != len(expectedModifiedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedModifiedSlice), modifiedSlice.Length())
	}

	for i, val := range expectedModifiedSlice {
		retrievedVal, ok := modifiedSlice.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: Modify the elements in the slice by squaring them.
	numbers = &slice.Slice[int]{2, 4, 6, 8, 10}

	// Modify the slice by squaring the elements.
	modifiedSlice = numbers.Modify(func(i int, value int) int {
		return value * value
	})

	expectedModifiedSlice = []int{4, 16, 36, 64, 100}

	if modifiedSlice.Length() != len(expectedModifiedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedModifiedSlice), modifiedSlice.Length())
	}

	for i, val := range expectedModifiedSlice {
		retrievedVal, ok := modifiedSlice.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: Modify an empty slice, expect the length to remain 0.
	emptySlice := &slice.Slice[int]{}
	emptySlice.Modify(func(i int, value int) int {
		return value * 2
	})

	if emptySlice.Length() != 0 {
		t.Errorf("Expected modified empty slice length to be 0, but got %d", emptySlice.Length())
	}
}

// TestPoll tests Slice.Poll.
func TestPoll(t *testing.T) {
	// Test case: Poll the first element from the non-empty slice.
	numbers := &slice.Slice[int]{10, 20, 30}

	polledValue := numbers.Poll()
	expectedPolledValue := 10

	if polledValue != expectedPolledValue {
		t.Errorf("Expected polled value to be %d, but got %d", expectedPolledValue, polledValue)
	}

	// Check the modified slice.
	expectedModifiedSlice := []int{20, 30}

	if numbers.Length() != len(expectedModifiedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedModifiedSlice), numbers.Length())
	}

	for i, val := range expectedModifiedSlice {
		retrievedVal, ok := numbers.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: Poll from an empty slice, expect the returned value to be zero.
	emptySlice := &slice.Slice[int]{}
	polledValue = emptySlice.Poll()
	expectedPolledValue = 0

	if polledValue != expectedPolledValue {
		t.Errorf("Expected polled value from an empty slice to be %d, but got %d", expectedPolledValue, polledValue)
	}
}

// TestPollLength tests Slice.PollLength.
func TestPollLength(t *testing.T) {
	// Test case: Poll the first element from the slice.
	numbers := &slice.Slice[int]{1, 2, 3}

	polledValue, newLength := numbers.PollLength()
	expectedPolledValue := 1
	expectedNewLength := 2

	if polledValue != expectedPolledValue {
		t.Errorf("Expected polled value to be %d, but got %d", expectedPolledValue, polledValue)
	}

	if newLength != expectedNewLength {
		t.Errorf("Expected new length to be %d, but got %d", expectedNewLength, newLength)
	}

	// Check the modified slice.
	expectedModifiedSlice := []int{2, 3}

	if numbers.Length() != len(expectedModifiedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedModifiedSlice), numbers.Length())
	}

	for i, val := range expectedModifiedSlice {
		retrievedVal, ok := numbers.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: Poll from an empty slice, expect the returned value to be zero and new length to be zero.
	emptySlice := &slice.Slice[int]{}
	polledValue, newLength = emptySlice.PollLength()
	expectedPolledValue = 0
	expectedNewLength = 0

	if polledValue != expectedPolledValue {
		t.Errorf("Expected polled value from an empty slice to be %d, but got %d", expectedPolledValue, polledValue)
	}

	if newLength != expectedNewLength {
		t.Errorf("Expected new length of an empty slice to be %d, but got %d", expectedNewLength, newLength)
	}
}

// TestPollOK tests Slice.PollOK.
func TestPollOK(t *testing.T) {
	// Test case: PollOK from a non-empty slice.
	numbers := &slice.Slice[int]{1, 2, 3}

	polledValue, ok := numbers.PollOK()
	expectedPolledValue := 1

	if !ok {
		t.Errorf("Expected PollOK to return true for a non-empty slice, but got false")
	}

	if polledValue != expectedPolledValue {
		t.Errorf("Expected polled value to be %d, but got %d", expectedPolledValue, polledValue)
	}

	// Check the modified slice.
	expectedModifiedSlice := []int{2, 3}

	if numbers.Length() != len(expectedModifiedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedModifiedSlice), numbers.Length())
	}

	for i, val := range expectedModifiedSlice {
		retrievedVal, ok := numbers.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: PollOK from an empty slice.
	emptySlice := &slice.Slice[int]{}
	polledValue, ok = emptySlice.PollOK()

	if ok {
		t.Errorf("Expected PollOK to return false for an empty slice, but got true")
	}

	if polledValue != 0 {
		t.Errorf("Expected polled value from an empty slice to be 0, but got %d", polledValue)
	}
}

// TestPop tests Slice.Pop.
func TestPop(t *testing.T) {
	// Test case: Pop from a non-empty slice.
	numbers := &slice.Slice[int]{1, 2, 3}

	poppedValue := numbers.Pop()
	expectedPoppedValue := 3

	if poppedValue != expectedPoppedValue {
		t.Errorf("Expected popped value to be %d, but got %d", expectedPoppedValue, poppedValue)
	}

	// Check the modified slice.
	expectedModifiedSlice := []int{1, 2}

	if numbers.Length() != len(expectedModifiedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedModifiedSlice), numbers.Length())
	}

	for i, val := range expectedModifiedSlice {
		retrievedVal, ok := numbers.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: Pop from an empty slice.
	emptySlice := &slice.Slice[int]{}
	poppedValue = emptySlice.Pop()

	if poppedValue != 0 {
		t.Errorf("Expected popped value from an empty slice to be 0, but got %d", poppedValue)
	}
}

// TestPopLength tests Slice.PopLength.
func TestPopLength(t *testing.T) {
	// Test case: PopLength from a non-empty slice.
	numbers := &slice.Slice[int]{10, 20, 30}

	poppedValue, newLength := numbers.PopLength()
	expectedPoppedValue := 30
	expectedNewLength := 2

	if poppedValue != expectedPoppedValue {
		t.Errorf("Expected popped value to be %d, but got %d", expectedPoppedValue, poppedValue)
	}

	if newLength != expectedNewLength {
		t.Errorf("Expected new length to be %d, but got %d", expectedNewLength, newLength)
	}

	// Check the modified slice.
	expectedModifiedSlice := []int{10, 20}

	if numbers.Length() != len(expectedModifiedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedModifiedSlice), numbers.Length())
	}

	for i, val := range expectedModifiedSlice {
		retrievedVal, ok := numbers.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: PopLength from an empty slice.
	emptySlice := &slice.Slice[int]{}
	poppedValue, newLength = emptySlice.PopLength()

	if poppedValue != 0 {
		t.Errorf("Expected popped value from an empty slice to be 0, but got %d", poppedValue)
	}

	if newLength != 0 {
		t.Errorf("Expected new length of an empty slice to be 0, but got %d", newLength)
	}
}

// TestPopOK tests Slice.PopOK.
func TestPopOK(t *testing.T) {
	// Test case: PopOK from a non-empty slice.
	numbers := &slice.Slice[int]{10, 20, 30}

	poppedValue, ok := numbers.PopOK()
	expectedPoppedValue := 30

	if !ok {
		t.Errorf("Expected PopOK to return true for a non-empty slice, but got false")
	}

	if poppedValue != expectedPoppedValue {
		t.Errorf("Expected popped value to be %d, but got %d", expectedPoppedValue, poppedValue)
	}

	// Check the modified slice.
	expectedModifiedSlice := []int{10, 20}

	if numbers.Length() != len(expectedModifiedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedModifiedSlice), numbers.Length())
	}

	for i, val := range expectedModifiedSlice {
		retrievedVal, ok := numbers.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: PopOK from an empty slice.
	emptySlice := &slice.Slice[int]{}
	poppedValue, ok = emptySlice.PopOK()

	if ok {
		t.Errorf("Expected PopOK to return false for an empty slice, but got true")
	}

	if poppedValue != 0 {
		t.Errorf("Expected popped value from an empty slice to be 0, but got %d", poppedValue)
	}
}

// TestPrecatenate tests Slice.Precatenate.
func TestPrecatenate(t *testing.T) {
	// Test case: Precatenate with a non-empty slice.
	slice1 := &slice.Slice[int]{1, 2, 3}
	slice2 := &slice.Slice[int]{4, 5}

	slice1.Precatenate(slice2)

	expectedSlice := []int{4, 5, 1, 2, 3}

	if slice1.Length() != len(expectedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedSlice), slice1.Length())
	}

	for i, val := range expectedSlice {
		retrievedVal, ok := slice1.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: Precatenate with a nil slice.
	slice3 := &slice.Slice[int]{1, 2, 3}
	slice3.Precatenate(nil)

	// The slice should remain unchanged as there is nothing to prepend.
	expectedSliceUnchanged := []int{1, 2, 3}

	if slice3.Length() != len(expectedSliceUnchanged) {
		t.Errorf("Expected slice length to be %d after Precatenate with nil, but got %d", len(expectedSliceUnchanged), slice3.Length())
	}

	for i, val := range expectedSliceUnchanged {
		retrievedVal, ok := slice3.Get(i)
		if !ok {
			t.Errorf("Expected slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}
}

// TestPrecatenateFunc tests Slice.PrecatenateFunc.
func TestPrecatenateFunc(t *testing.T) {
	// Test case: PrecatenateFunc with a non-empty slice and even number predicate.
	slice1 := &slice.Slice[int]{1, 2, 3}
	slice2 := &slice.Slice[int]{4, 5, 6}

	// Prepend elements from slice2 if they are even.
	result := slice1.PrecatenateFunc(slice2, func(i int, value int) bool {
		return value%2 == 0
	})

	expectedSlice := []int{4, 6, 1, 2, 3}

	if result.Length() != len(expectedSlice) {
		t.Errorf("Expected modified slice length to be %d, but got %d", len(expectedSlice), result.Length())
	}

	for i, val := range expectedSlice {
		retrievedVal, ok := result.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}
}

// TestPrecatenateLength tests Slice.PrecatenateLength.
func TestPrecatenateLength(t *testing.T) {
	// Test case: PrecatenateLength with a non-empty slice.
	slice1 := &slice.Slice[int]{1, 2, 3}
	slice2 := &slice.Slice[int]{4, 5}

	// Prepend elements from slice2.
	length := slice1.PrecatenateLength(slice2)

	expectedLength := 5

	if length != expectedLength {
		t.Errorf("Expected modified slice length to be %d, but got %d", expectedLength, length)
	}

	expectedSlice := []int{4, 5, 1, 2, 3}

	for i, val := range expectedSlice {
		retrievedVal, ok := slice1.Get(i)
		if !ok {
			t.Errorf("Expected modified slice to contain value %d at index %d, but it was not found", val, i)
		}
		if retrievedVal != val {
			t.Errorf("Expected modified slice to contain value %d at index %d, but got %d", val, i, retrievedVal)
		}
	}

	// Test case: PrecatenateLength with an empty slice.
	slice3 := &slice.Slice[int]{}
	length2 := slice3.PrecatenateLength(nil)

	// The slice should remain empty as there are no elements to prepend.
	if length2 != 0 {
		t.Errorf("Expected modified empty slice length to be 0, but got %d", length2)
	}
}

// TestPrepend tests Slice.Prepend.
func TestPrepend(t *testing.T) {
	// Test case: Prepend to an empty slice.
	slice1 := &slice.Slice[int]{}
	slice1.Prepend(1)

	expectedSlice1 := []int{1}

	if len(*slice1) != len(expectedSlice1) {
		t.Errorf("Expected modified slice to be %v, but got %v", expectedSlice1, *slice1)
	}

	// Test case: Prepend to a non-empty slice.
	slice2 := &slice.Slice[int]{2, 3}
	slice2.Prepend(1)

	expectedSlice2 := []int{1, 2, 3}

	if len(*slice2) != len(expectedSlice2) {
		t.Errorf("Expected modified slice to be %v, but got %v", expectedSlice2, *slice2)
	}
}

// TestPrependFunc tests Slice.PrependFunc.
func TestPrependFunc(t *testing.T) {
	// Test case: PrependFunc to an empty slice.
	slice1 := &slice.Slice[int]{}
	slice1.PrependFunc(func(i int, value int) bool {
		return value%2 == 0
	}, 4, 5, 6)

	expectedSlice1 := []int{6, 4}

	if len(*slice1) != len(expectedSlice1) {
		t.Errorf("Expected modified slice to be %v, but got %v", expectedSlice1, *slice1)
	}

	// Test case: PrependFunc to a non-empty slice.
	slice2 := &slice.Slice[int]{1, 2, 3}
	slice2.PrependFunc(func(i int, value int) bool {
		return value%2 == 0
	}, 4, 5, 6)

	expectedSlice2 := []int{6, 4, 2, 1, 3}

	if len(*slice2) != len(expectedSlice2) {
		t.Errorf("Expected modified slice to be %v, but got %v", expectedSlice2, *slice2)
	}
}

// TestPrependLength tests Slice.PrependLength.
func TestPrependLength(t *testing.T) {
	// Test case: PrependLength to an empty slice.
	slice1 := &slice.Slice[int]{}
	length1 := slice1.PrependLength(1, 0)

	expectedSlice1 := []int{1, 0}

	if len(*slice1) != len(expectedSlice1) {
		t.Errorf("Expected modified slice to be %v, but got %v", expectedSlice1, *slice1)
	}

	if length1 != len(expectedSlice1) {
		t.Errorf("Expected length to be %d, but got %d", len(expectedSlice1), length1)
	}

	// Test case: PrependLength to a non-empty slice.
	slice2 := &slice.Slice[int]{2, 3}
	length2 := slice2.PrependLength(1, 0)

	expectedSlice2 := []int{1, 0, 2, 3}

	if len(*slice2) != len(expectedSlice2) {
		t.Errorf("Expected modified slice to be %v, but got %v", expectedSlice2, *slice2)
	}

	if length2 != len(expectedSlice2) {
		t.Errorf("Expected length to be %d, but got %d", len(expectedSlice2), length2)
	}
}

func TestReduce(t *testing.T) {
	// Test case: Reduce sum of integers in the slice.
	numbers := &slice.Slice[int]{1, 2, 3, 4, 5}
	sum := numbers.Reduce(func(i int, currentValue int, resultValue int) int {
		return resultValue + currentValue
	})

	expectedSum := 15

	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}

	// Test case: Reduce concatenation of strings in the slice.
	words := &slice.Slice[string]{"Hello", " ", "World", "!"}
	concatenated := words.Reduce(func(i int, currentValue string, resultValue string) string {
		return resultValue + currentValue
	})

	expectedConcatenated := "Hello World!"

	if concatenated != expectedConcatenated {
		t.Errorf("Expected concatenated string to be %s, but got %s", expectedConcatenated, concatenated)
	}
}

func TestReduceReverse(t *testing.T) {
	// Test case: Reduce sum of integers in the slice in reverse order.
	numbers := &slice.Slice[int]{1, 2, 3, 4, 5}
	sum := numbers.ReduceReverse(func(i int, currentValue int, resultValue int) int {
		return currentValue + resultValue
	})

	expectedSum := 15

	if sum != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, sum)
	}

	// Test case: Reduce concatenation of strings in the slice in reverse order.
	words := &slice.Slice[string]{"!", "World", " ", "Hello"}
	concatenated := words.ReduceReverse(func(i int, currentValue string, resultValue string) string {
		return resultValue + currentValue
	})

	expectedConcatenated := "Hello World!"

	if concatenated != expectedConcatenated {
		t.Errorf("Expected concatenated string to be %s, but got %s", expectedConcatenated, concatenated)
	}
}

// TestReplace tests Slice.Replace.
func TestReplace(t *testing.T) {
	// Test case: Replace element in bounds.
	s := &slice.Slice[int]{1, 2, 3}
	ok := s.Replace(1, 4)

	expectedSlice := &slice.Slice[int]{1, 4, 3}

	if !ok {
		t.Errorf("Expected Replace method to return true for in-bounds index")
	}
	if !s.Equal(expectedSlice) {
		t.Errorf("Expected slice after replacement to be %v, but got %v", expectedSlice, s)
	}

	// Test case: Replace element out of bounds.
	s = &slice.Slice[int]{1, 2, 3}
	ok = s.Replace(3, 4)

	expectedSlice = &slice.Slice[int]{1, 2, 3}

	if ok {
		t.Errorf("Expected Replace method to return false for out-of-bounds index")
	}
	if !s.Equal(expectedSlice) {
		t.Errorf("Expected slice to remain unchanged as %v, but got %v", expectedSlice, s)
	}
}

// TestReverse tests Slice.Reverse.
func TestReverse(t *testing.T) {
	// Test case: Reverse a slice of integers.
	s := &slice.Slice[int]{1, 2, 3}
	s.Reverse()

	expectedSlice := &slice.Slice[int]{3, 2, 1}

	if !s.Equal(expectedSlice) {
		t.Errorf("Expected reversed slice to be %v, but got %v", expectedSlice, s)
	}

	// Test case: Reverse an empty slice.
	emptySlice := &slice.Slice[int]{}
	emptySlice.Reverse()

	if !emptySlice.IsEmpty() {
		t.Errorf("Expected reversed empty slice to remain empty")
	}

	// Test case: Reverse a slice with a single element.
	singleElementSlice := &slice.Slice[int]{42}
	singleElementSlice.Reverse()

	if length := singleElementSlice.Length(); length != 1 {
		t.Errorf("Expected reversed single-element slice to have length 1, but got %d", length)
	}
	if value, _ := singleElementSlice.Get(0); value != 42 {
		t.Errorf("Expected reversed single-element slice to have value 42, but got %v", value)
	}
}

// TestSet tests Slice.Set.
func TestSet(t *testing.T) {
	// Test case: Remove duplicate elements from a slice of integers.
	s := &slice.Slice[int]{1, 2, 2, 3, 3, 3}
	s.Set()

	expectedSlice := &slice.Slice[int]{1, 2, 3}

	if !s.Equal(expectedSlice) {
		t.Errorf("Expected unique elements slice to be %v, but got %v", expectedSlice, s)
	}

	// Test case: Remove duplicates from an empty slice.
	emptySlice := &slice.Slice[int]{}
	emptySlice.Set()

	if !emptySlice.IsEmpty() {
		t.Errorf("Expected unique elements empty slice to remain empty")
	}

	// Test case: Remove duplicates from a slice with a single element.
	singleElementSlice := &slice.Slice[int]{42}
	singleElementSlice.Set()

	if length := singleElementSlice.Length(); length != 1 {
		t.Errorf("Expected unique elements single-element slice to have length 1, but got %d", length)
	}
	if value, _ := singleElementSlice.Get(0); value != 42 {
		t.Errorf("Expected unique elements single-element slice to have value 42, but got %v", value)
	}
}

func TestShuffle(t *testing.T) {
	// Seed the random number generator with the current timestamp.

	// Test case: Shuffling a slice of integers.
	originalSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
	shuffledSlice := originalSlice.Clone().Shuffle()

	// Check if the length remains the same after shuffling.
	if originalSlice.Length() != shuffledSlice.Length() {
		t.Errorf("Length of the shuffled slice (%d) is not equal to the original slice (%d)", shuffledSlice.Length(), originalSlice.Length())
	}

	// Check if all elements from the original slice exist in the shuffled slice after shuffling.
	for i := 0; i < originalSlice.Length(); i++ {
		originalElement, _ := originalSlice.Get(i)
		if !shuffledSlice.Contains(originalElement) {
			t.Errorf("Shuffled slice does not contain the element %v from the original slice", originalElement)
		}
	}
}

// TestSlice tests Slice.Slice.
func TestSlice(t *testing.T) {
	// Test case: Slicing a slice of integers.
	numbers := &slice.Slice[int]{1, 2, 3, 4, 5}
	subset := numbers.Slice(1, 3)

	// Check if the length of the subset is correct.
	expectedLength := 3 // Elements at indexes 1 to 3 are [2, 3, 4].
	if subset.Length() != expectedLength {
		t.Errorf("Length of the subset (%d) is not equal to the expected length (%d)", subset.Length(), expectedLength)
	}

	// Check if the elements in the subset are correct.
	expectedElements := []int{2, 3, 4}
	for i, expected := range expectedElements {
		element, _ := subset.Get(i)
		if element != expected {
			t.Errorf("Element at index %d in the subset is %v, expected %v", i, element, expected)
		}
	}
}

// TestSplice tests Slice.Splice.
func TestSplice(t *testing.T) {
	// Test case: Slicing a slice of integers.
	numbers := &slice.Slice[int]{1, 2, 3, 4, 5}
	numbers.Splice(1, 3)

	// Check if the length of the modified slice is correct.
	expectedLength := 3 // Elements at indexes 1 to 3 are [2, 3, 4].
	if numbers.Length() != expectedLength {
		t.Errorf("Length of the modified slice (%d) is not equal to the expected length (%d)", numbers.Length(), expectedLength)
	}

	// Check if the elements in the modified slice are correct.
	expectedElements := []int{2, 3, 4}
	for i, expected := range expectedElements {
		element, _ := numbers.Get(i)
		if element != expected {
			t.Errorf("Element at index %d in the modified slice is %v, expected %v", i, element, expected)
		}
	}
}

// TestSortFunc tests Slice.SortFunc.
func TestSortFunc(t *testing.T) {
	// Test case: Sorting a slice of integers in ascending order.
	numbers := &slice.Slice[int]{5, 2, 8, 1, 9}
	numbers.SortFunc(func(i, j int, a, b int) bool {
		return a < b
	})

	// Check if the elements in the sorted slice are correct.
	expectedSorted := []int{1, 2, 5, 8, 9}
	for i, expected := range expectedSorted {
		element, _ := numbers.Get(i)
		if element != expected {
			t.Errorf("Element at index %d in the sorted slice is %v, expected %v", i, element, expected)
		}
	}
}

// TestSwap tests Slice.Swap.
func TestSwap(t *testing.T) {
	// Test case: Swapping elements in a slice of integers.
	numbers := &slice.Slice[int]{1, 2, 3}
	numbers.Swap(0, 2)

	// Check if the elements in the slice are swapped correctly.
	expected := []int{3, 2, 1}
	for i, val := range expected {
		element, _ := numbers.Get(i)
		if element != val {
			t.Errorf("Element at index %d is %v, expected %v", i, element, val)
		}
	}
}
