package slice

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

// Slice represents a generic slice of any data type.
type Slice[T any] []T

// Append appends the given values to the slice and returns the modified slice.
func (slice *Slice[T]) Append(values ...T) *Slice[T] {
	*slice = append(*slice, values...)
	return slice
}

// AppendFunc appends values to the slice based on the provided function and returns the modified slice.
func (slice *Slice[T]) AppendFunc(values []T, fn func(i int, value T) bool) *Slice[T] {
	for i, value := range values {
		if fn(i, value) {
			slice.Append(value)
		}
	}
	return slice
}

// AppendLength appends values to the slice and returns the length of the modified slice.
func (slice *Slice[T]) AppendLength(values ...T) int {
	return slice.Append(values...).Length()
}

// Bounds checks if the given index is within the bounds of the slice.
func (slice *Slice[T]) Bounds(i int) bool {
	return i >= 0 && i < slice.Length()
}

// Clone creates a new slice that is a copy of the original slice.
func (slice *Slice[T]) Clone() *Slice[T] {
	s := *slice
	return &s
}

// Concatenate concatenates the given slice with the original slice and returns the modified slice.
func (slice *Slice[T]) Concatenate(s *Slice[T]) *Slice[T] {
	if s != nil {
		slice.Append((*s)...)
	}
	return slice
}

// ConcatenateFunc concatenates values based on the provided function and returns the modified slice.
func (slice *Slice[T]) ConcatenateFunc(s *Slice[T], fn func(i int, value T) bool) *Slice[T] {
	if s != nil {
		slice.AppendFunc(*s, fn)
	}
	return slice
}

// ConcatenateLength concatenates the given slice with the original slice and returns the length of the modified slice.
func (slice *Slice[T]) ConcatenateLength(s *Slice[T]) int {
	return slice.Concatenate(s).Length()
}

// Contains checks if the slice contains the given value.
func (slice *Slice[T]) Contains(value T) bool {
	var ok bool
	slice.EachBreak(func(i int, v T) bool {
		ok = reflect.DeepEqual(v, value)
		return !ok
	})
	return ok
}

// ContainsMany checks if the slice contains any of the given values and returns a boolean slice indicating the results.
func (slice *Slice[T]) ContainsMany(values ...T) *Slice[bool] {
	s := &Slice[bool]{}
	for _, value := range values {
		s.Append(slice.Contains(value))
	}
	return s
}

// Deduplicate removes duplicate values from the slice and returns the modified slice.
func (slice *Slice[T]) Deduplicate() *Slice[T] {
	uniqueValues := make(map[string]bool)
	var uniqueSlice Slice[T]

	slice.Each(func(_ int, value T) {
		if !uniqueValues[fmt.Sprintf("%v", value)] {
			uniqueValues[fmt.Sprintf("%v", value)] = true
			uniqueSlice.Append(value)
		}
	})
	*slice = uniqueSlice
	return slice
}

// Delete removes the element at the specified index from the slice and returns the modified slice.
func (slice *Slice[T]) Delete(i int) *Slice[T] {
	if slice.Bounds(i) {
		slice.DeleteUnsafe(i)
	}
	return slice
}

// DeleteFunc removes elements from the slice based on the provided function and returns the modified slice.
func (slice *Slice[T]) DeleteFunc(fn func(i int, value T) bool) *Slice[T] {
	length := slice.Length()
	for i := 0; i < length; i++ {
		if fn(i, slice.Fetch(i)) {
			slice.DeleteUnsafe(i)
			i--      // Adjust the index after deletion.
			length-- // Update the length.
		}
	}
	return slice
}

// DeleteLength removes the element at the specified index from the slice and returns the length of the modified slice.
func (slice *Slice[T]) DeleteLength(i int) int {
	return slice.Delete(i).Length()
}

// DeleteOK removes the element at the specified index from the slice if it exists and returns true, or false otherwise.
func (slice *Slice[T]) DeleteOK(i int) bool {
	if slice.Bounds(i) {
		slice.Delete(i)
		return true
	}
	return false
}

// DeleteUnsafe removes the element at the specified index from the slice without performing bounds checks and returns the modified slice.
func (slice *Slice[T]) DeleteUnsafe(i int) *Slice[T] {
	*slice = append((*slice)[:i], (*slice)[i+1:]...)
	return slice
}

// Each applies the given function to each element of the slice and returns the modified slice.
func (slice *Slice[T]) Each(fn func(i int, value T)) *Slice[T] {
	slice.EachBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachBreak applies the given function to each element of the slice until the function returns false and returns the modified slice.
func (slice *Slice[T]) EachBreak(fn func(i int, value T) bool) *Slice[T] {
	for i, v := range *slice {
		if !fn(i, v) {
			break
		}
	}
	return slice
}

// EachOK applies the given function to each element of the slice and returns true if the function returned true for all elements, or false otherwise.
func (slice *Slice[T]) EachOK(fn func(i int, value T) bool) bool {
	var ok bool
	slice.EachBreak(func(i int, value T) bool {
		ok = fn(i, value)
		return ok
	})
	return ok
}

// EachReverse applies the given function to each element of the slice in reverse order and returns the modified slice.
func (slice *Slice[T]) EachReverse(fn func(i int, value T)) *Slice[T] {
	slice.EachReverseBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachReverseBreak applies the given function to each element of the slice in reverse order until the function returns false and returns the modified slice.
func (slice *Slice[T]) EachReverseBreak(fn func(i int, value T) bool) *Slice[T] {
	for i := slice.Length() - 1; i >= 0; i-- {
		if !fn(i, (*slice)[i]) {
			break
		}
	}
	return slice
}

// EachReverseOK applies the given function to each element of the slice in reverse order and returns true if the function returned true for all elements, or false otherwise.
func (slice *Slice[T]) EachReverseOK(fn func(i int, value T) bool) bool {
	var ok bool
	slice.EachReverseBreak(func(i int, value T) bool {
		ok = fn(i, value)
		return ok
	})
	return ok
}

// Equal checks if the slice is equal to another slice based on the element values.
func (slice *Slice[T]) Equal(s *Slice[T]) bool {
	return slice.EqualFunc(s, func(i int, value1, value2 T) bool {
		return reflect.DeepEqual(value1, value2)
	})
}

// EqualFunc checks if the slice is equal to another slice based on the element values and a provided comparison function.
func (slice *Slice[T]) EqualFunc(s *Slice[T], fn func(i int, a T, b T) bool) bool {
	if !slice.EqualLength(s) {
		return false
	}
	var ok bool
	slice.EachBreak(func(i int, value T) bool {
		ok = fn(i, slice.Fetch(i), s.Fetch(i))
		return ok
	})
	return ok
}

// EqualLength checks if the length of the slice is equal to the length of another slice.
func (slice *Slice[T]) EqualLength(s *Slice[T]) bool {
	return slice.Length() == s.Length()
}

// Fetch returns the element at the specified index and a boolean indicating whether the index is valid.
func (slice *Slice[T]) Fetch(i int) T {
	value, _ := slice.Get(i)
	return value
}

// FetchLength returns the element at the specified index, a boolean indicating whether the index is valid, and the length of the slice.
func (slice *Slice[T]) FetchLength(i int) (T, int) {
	return slice.Fetch(i), slice.Length()
}

// Filter returns a new slice containing elements that satisfy the provided function.
func (slice *Slice[T]) Filter(fn func(i int, value T) bool) *Slice[T] {
	s := &Slice[T]{}
	slice.Each(func(i int, value T) {
		if fn(i, value) {
			s.Append(value)
		}
	})
	return s
}

// FindIndex finds the index of the first element that satisfies the provided function and returns the index and true,
// or -1 and false if no such element is found.
func (slice *Slice[T]) FindIndex(fn func(value T) bool) (int, bool) {
	for i, value := range *slice {
		if fn(value) {
			return i, true
		}
	}
	return -1, false
}

// Get returns the element at the specified index and a boolean indicating whether the index is valid.
func (slice *Slice[T]) Get(i int) (T, bool) {
	var value T
	if !slice.Bounds(i) {
		return value, false
	}
	return (*slice)[i], true
}

// GetLength returns the element at the specified index, a boolean indicating whether the index is valid, and the length of the slice.
func (slice *Slice[T]) GetLength(i int) (T, bool, int) {
	value, ok := slice.Get(i)
	return value, ok, slice.Length()
}

// IsEmpty returns true if the slice is empty, or false otherwise.
func (slice *Slice[T]) IsEmpty() bool {
	return slice.Length() == 0
}

// IsPopulated returns true if the slice is not empty, or false otherwise.
func (slice *Slice[T]) IsPopulated() bool {
	return !slice.IsEmpty()
}

// Length returns the length of the slice.
func (slice *Slice[T]) Length() int {
	return len(*slice)
}

// Make creates a new slice with the specified length and capacity and returns the modified slice.
func (slice *Slice[T]) Make(i int) *Slice[T] {
	*slice = make(Slice[T], i)
	return slice
}

// MakeEach creates a new slice with the specified elements and returns the modified slice.
func (slice *Slice[T]) MakeEach(v ...T) *Slice[T] {
	return slice.Make(len(v)).Each(func(i int, _ T) {
		slice.Replace(i, v[i])
	})
}

// MakeEachReverse creates a new slice with the specified elements in reverse order and returns the modified slice.
func (slice *Slice[T]) MakeEachReverse(values ...T) *Slice[T] {
	currentOffset := 0
	return slice.Make(len(values)).EachReverse(func(i int, _ T) {
		slice.Replace(currentOffset, values[i])
		currentOffset++
	})
}

// Map applies the given function to each element of the slice and returns a new slice with the modified elements.
func (slice *Slice[T]) Map(fn func(i int, value T) T) *Slice[T] {
	s := make(Slice[T], slice.Length())
	slice.Each(func(i int, value T) {
		s.Replace(i, fn(i, value))
	})
	return &s
}

// MapReverse applies the given function to each element of the slice in reverse order and returns a new slice with the modified elements.
func (slice *Slice[T]) MapReverse(fn func(i int, value T) T) *Slice[T] {
	s := make(Slice[T], slice.Length())
	slice.EachReverse(func(i int, value T) {
		s.Replace(i, fn(i, value))
	})
	return &s
}

// Modify applies the given function to each element of the slice and modifies the original slice with the modified elements.
func (slice *Slice[T]) Modify(fn func(i int, value T) T) *Slice[T] {
	slice.Each(func(i int, value T) {
		slice.Replace(i, fn(i, value))
	})
	return slice
}

// ModifyReverse applies the given function to each element of the slice in reverse order and modifies the original slice with the modified elements.
func (slice *Slice[T]) ModifyReverse(fn func(i int, value T) T) *Slice[T] {
	slice.EachReverse(func(i int, value T) {
		slice.Replace(i, fn(i, value))
	})
	return slice
}

// Poll removes and returns the first element of the slice, or a zero value if the slice is empty.
func (slice *Slice[T]) Poll() T {
	var value T
	if !slice.IsEmpty() {
		value = (*slice)[0]
		*slice = (*slice)[1:]
	}
	return value
}

// PollLength removes and returns the first element of the slice and the length of the modified slice.
func (slice *Slice[T]) PollLength() (T, int) {
	return slice.Poll(), slice.Length()
}

// PollOK removes and returns the first element of the slice and true if the slice is not empty, or a zero value and false otherwise.
func (slice *Slice[T]) PollOK() (T, bool) {
	if slice.IsEmpty() {
		var value T
		return value, false
	}
	return slice.Poll(), true
}

// Pop removes and returns the last element of the slice, or a zero value if the slice is empty.
func (slice *Slice[T]) Pop() T {
	var value T
	if slice.IsPopulated() {
		length := slice.Length()
		value = (*slice)[length-1]
		*slice = (*slice)[:length-1]
	}
	return value
}

// PopLength removes and returns the last element of the slice and the length of the modified slice.
func (slice *Slice[T]) PopLength() (T, int) {
	return slice.Pop(), slice.Length()
}

// PopOK removes and returns the last element of the slice and true if the slice is not empty, or a zero value and false otherwise.
func (slice *Slice[T]) PopOK() (T, bool) {
	if slice.IsEmpty() {
		var value T
		return value, false
	}
	return slice.Pop(), true
}

// Precatenate concatenates the given slice with the original slice and returns the modified slice.
func (slice *Slice[T]) Precatenate(s *Slice[T]) *Slice[T] {
	if s != nil {
		slice.Prepend((*s)...)
	}
	return slice
}

// PrecatenateFunc concatenates values based on the provided function and returns the modified slice.
func (slice *Slice[T]) PrecatenateFunc(s *Slice[T], fn func(i int, value T) bool) *Slice[T] {
	return slice.PrependFunc(*s, fn)
}

// PrecatenateLength concatenates the given slice with the original slice and returns the length of the modified slice.
func (slice *Slice[T]) PrecatenateLength(s *Slice[T]) int {
	return slice.Precatenate(s).Length()
}

// Prepend adds the given values to the beginning of the slice and returns the modified slice.
func (slice *Slice[T]) Prepend(values ...T) *Slice[T] {
	*slice = append(values, *slice...)
	return slice
}

// PrependFunc adds values to the beginning of the slice based on the provided function and returns the modified slice.
func (slice *Slice[T]) PrependFunc(values []T, fn func(i int, value T) bool) *Slice[T] {
	s := &Slice[T]{}
	for i, value := range values {
		if fn(i, value) {
			s.Append(value)
		}
	}
	*slice = *s.Concatenate(slice)
	return slice
}

// PrependLength adds values to the beginning of the slice and returns the length of the modified slice.
func (slice *Slice[T]) PrependLength(values ...T) int {
	return slice.Prepend(values...).Length()
}

// Reduce applies the given function to each element of the slice and returns the reduced result value.
func (slice *Slice[T]) Reduce(fn func(i int, currentValue T, resultValue T) T) T {
	var resultValue T
	slice.Each(func(i int, currentValue T) {
		resultValue = fn(i, currentValue, resultValue)
	})
	return resultValue
}

// ReduceReverse applies the given function to each element of the slice in reverse order and returns the reduced result value.
func (slice *Slice[T]) ReduceReverse(fn func(i int, currentValue T, resultValue T) T) T {
	var resultValue T
	slice.EachReverse(func(i int, currentValue T) {
		resultValue = fn(i, currentValue, resultValue)
	})
	return resultValue
}

// Replace replaces the element at the specified index with the given value and returns true if the index is valid, or false otherwise.
func (slice *Slice[T]) Replace(i int, value T) bool {
	if slice.Bounds(i) {
		(*slice)[i] = value
		return true
	}
	return false
}

// Reverse reverses the elements of the slice and returns the modified slice.
func (slice *Slice[T]) Reverse() *Slice[T] {
	for i, j := 0, slice.Length()-1; i < j; i, j = i+1, j-1 {
		slice.Swap(i, j)
	}
	return slice
}

// Shuffle shuffles the elements of the slice and returns the modified slice.
func (slice *Slice[T]) Shuffle() *Slice[T] {
	rand.Shuffle(slice.Length(), func(i, j int) {
		slice.Swap(i, j)
	})
	return slice
}

// Slice returns a new slice containing the elements from index i to j (inclusive) of the original slice.
func (slice *Slice[T]) Slice(i int, j int) *Slice[T] {
	if j < i {
		i, j = j, i
	}
	s := (*slice)[i : j+1]
	return &s
}

// SortFunc sorts the elements of the slice based on the provided comparison function and returns the modified slice.
func (slice *Slice[T]) SortFunc(fn func(i int, j int, a T, b T) bool) *Slice[T] {
	s := *slice // Copy the slice to a new variable.
	sort.Slice(s, func(i int, j int) bool {
		return fn(i, j, s[i], s[j]) // Use the copied slice (v) instead of slice.
	})
	*slice = s // Update the original slice with the sorted copy.
	return slice
}

// Splice removes elements from the slice starting from index i to j (inclusive) and returns the modified slice.
func (slice *Slice[T]) Splice(i int, j int) *Slice[T] {
	if j < i {
		i, j = j, i
	}
	if slice.Bounds(i) && slice.Bounds(j) {
		*slice = (*slice)[i : j+1]
	}
	return slice
}

// Split divides the slice into two slices at the specified index and returns the two new slices.
func (slice *Slice[T]) Split(i int) (*Slice[T], *Slice[T]) {
	if !slice.Bounds(i) {
		return nil, nil
	}
	firstSlice, secondSlice := (*slice)[:i], (*slice)[i:]
	return &firstSlice, &secondSlice
}

// SplitFunc divides the slice into two slices based on the provided function and returns the two new slices.
func (slice *Slice[T]) SplitFunc(fn func(i int, value T) bool) (*Slice[T], *Slice[T]) {
	firstSlice, secondSlice := &Slice[T]{}, &Slice[T]{}
	slice.Each(func(i int, value T) {
		if fn(i, value) {
			firstSlice.Append(value)
		} else {
			secondSlice.Append(value)
		}
	})
	return firstSlice, secondSlice
}

// SplitOK divides the slice into two slices at the specified index and returns the two new slices, or false if the index is invalid.
func (slice *Slice[T]) SplitOK(i int) (*Slice[T], *Slice[T], bool) {
	if !slice.Bounds(i) {
		return nil, nil, false
	}
	firstSlice, secondSlice := (*slice)[:i], (*slice)[i:]
	return &firstSlice, &secondSlice, true
}

// Swap swaps the elements at the specified indices in the slice.
func (slice *Slice[T]) Swap(i int, j int) {
	if slice.Bounds(i) && slice.Bounds(j) {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

// New creates a new instance of the Slice[T] type and initializes it with the provided values.
func New[T any](values ...T) *Slice[T] {
	return (&Slice[T]{}).Append(values...)
}
