package slice

import (
	"fmt"
	"math/rand"
	"reflect"
)

// Slice is a generic list-like struct that allows for the manipulation and traversal of elements by numeric index.
// It is parameterized with the type T, allowing it to work with slices of any data type.
//
// Example usage:
//
//	// Create a Slice of integers.
//	s := &Slice[int]{1, 2, 3, 4, 5}
//
//	// Append values to the Slice.
//	s.Append(6, 7, 8)
//
//	// Check if a value exists in the Slice.
//	contains := s.Contains(3) // contains is true
//
//	// Find the index of the first element that satisfies a condition.
//	index, found := s.FindIndex(func(value int) bool {
//	    return value > 4
//	})
//
//	// Delete an element by index.
//	s.Delete(2) // Removes the element at index 2
//
//	// Iterate over the elements and perform an operation on each.
//	s.Each(func(i int, value int) {
//	    fmt.Printf("Element %d: %d\n", i, value)
//	})
//
// This type provides a set of methods for common Slice operations and is designed to be generic for flexibility.
type Slice[T any] []T

// Append appends the specified values to the end of the Slice and returns a pointer to the modified Slice.
// It extends the Slice by adding the provided values to the end, updating the Slice in place.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3}
//	s.Append(4, 5) // s is now [1, 2, 3, 4, 5]
func (slice *Slice[T]) Append(values ...T) *Slice[T] {
	*slice = append(*slice, values...)
	return slice
}

// AppendFunc appends selected elements to the end of the Slice based on a provided condition.
// It iterates over the specified values, invoking the provided function for each element.
// If the function returns true for an element, that element is added to the end of the Slice.
//
// Example:
//
//	s := &slice.Slice[int]{}
//	s.AppendFunc(func(i int, value int) bool {
//	  return value%2 == 0 // Append even numbers to the Slice.
//	}, 1, 2, 3, 4, 5)
//
// After this operation, s will contain [2, 4].
//
// This method modifies the original Slice and returns a pointer to the modified Slice.
func (slice *Slice[T]) AppendFunc(fn func(i int, value T) bool, values ...T) *Slice[T] {
	for i, value := range values {
		if ok := fn(i, value); ok {
			slice.Append(value)
		}
	}
	return slice
}

// AppendLength appends the specified values to the end of the Slice and returns the length of the modified Slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3}
//	length := s.AppendLength(4, 5) // s is now [1, 2, 3, 4, 5], length is 5
func (slice *Slice[T]) AppendLength(values ...T) int {
	return slice.Append(values...).Length()
}

// Bounds checks if the given integer index is within the valid range of indices for the Slice.
// It returns true if the index is greater than or equal to 0 and less than the length of the slice, indicating that the index is within bounds.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3}
//	inBounds := s.Bounds(1) // inBounds is true
//	outOfBounds := s.Bounds(5) // outOfBounds is false
func (slice *Slice[T]) Bounds(i int) bool {
	return i >= 0 && i < len(*slice)
}

// Concatenate merges the elements from the argument Slice to the tail of the receiver Slice and returns the modified Slice.
// If the provided Slice (s) is not nil, it appends its elements to the end of the receiver slice, updating the Slice in place.
//
// Example:
//
//	s1 := &Slice[int]{1, 2, 3}
//	s2 := &Slice[int]{4, 5}
//	s1.Concatenate(s2) // s1 is now [1, 2, 3, 4, 5]
func (slice *Slice[T]) Concatenate(s *Slice[T]) *Slice[T] {
	if s != nil {
		slice.Append((*s)...)
	}
	return slice
}

// ConcatenateFunc appends elements from the provided Slice (s) to the end of the receiver slice, based on the result of a filtering function.
//
// Example:
//
//	s1 := &Slice[int]{1, 2, 3}
//	s2 := &Slice[int]{4, 5, 6}
//	s1.ConcatenateFunc(func(i int, value int) bool {
//	    return value%2 == 0
//	}, s2) // s1 is now [1, 2, 3, 4, 6]
func (slice *Slice[T]) ConcatenateFunc(s *Slice[T], fn func(i int, value T) bool) *Slice[T] {
	if s != nil {
		slice.AppendFunc(fn, *s...)
	}
	return slice
}

// ConcatenateLength merges the elements from the argument Slice to the tail of the receiver Slice and returns the length of the modified Slice.
//
// Example:
//
//	s1 := &Slice[int]{1, 2, 3}
//	s2 := &Slice[int]{4, 5}
//	length := s1.ConcatenateLength(s2) // s1 is now [1, 2, 3, 4, 5], length is 5
func (slice *Slice[T]) ConcatenateLength(s *Slice[T]) int {
	return slice.Concatenate(s).Length()
}

// Contains iterates over the elements of the Slice and compares each element with the provided value.
// If a matching element is found, it returns true; otherwise, it returns false.
//
// Example:
//
//	s := &Slice[string]{"apple", "banana", "cherry"}
//	containsBanana := s.Contains("banana") // containsBanana is true
func (slice *Slice[T]) Contains(value T) bool {
	var ok bool
	slice.EachBreak(func(i int, v T) bool {
		ok = reflect.DeepEqual(v, value)
		return !ok // Each expects a "false" value to terminate the loop. We inverse the value of "ok" to make sure it exits on the condition deep equals found the right value.
	})
	return ok
}

// ContainsMany checks if multiple values exist in the Slice and returns a boolean Slice indicating their presence.
// It takes a variadic number of values and checks each of them against the elements in the Slice.
// The returned boolean Slice will have 'true' at the corresponding index if the value is found, 'false' otherwise.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	result := s.ContainsMany(2, 4, 6) // result will be [true, true, false]
func (slice *Slice[T]) ContainsMany(values ...T) *Slice[bool] {
	s := &Slice[bool]{}
	for _, value := range values {
		s.Append(slice.Contains(value))
	}
	return s
}

// Delete removes the element at the specified index from the slice, if the index is within bounds.
// It modifies the Slice in place by removing the element at the given index and returns a pointer to the modified Slice.
// If the index is out of bounds, it does not modify the Slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	s.Delete(2) // s is now [1, 2, 4, 5]
func (slice *Slice[T]) Delete(i int) *Slice[T] {
	if slice.Bounds(i) {
		*slice = append((*slice)[:i], (*slice)[i+1:]...)
	}
	return slice
}

// DeleteLength removes the element at the specified index from the slice, if the index is within bounds, and returns the new length of the modified Slice.
// It modifies the Slice in place by removing the element at the given index and returns the new length of the Slice.
// If the index is out of bounds, it does not modify the Slice and returns the current length.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	newLength := s.DeleteLength(2) // s is now [1, 2, 4, 5], newLength is 4
func (slice *Slice[T]) DeleteLength(i int) int {
	return slice.Delete(i).Length()
}

// DeleteOK removes the element at the specified index from the Slice if the index is within bounds and returns the result of the transaction.
// It modifies the Slice in place by removing the element at the given index and returns true if the deletion was successful.
// If the index is out of bounds, it does not modify the Slice and returns false.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	deleted := s.DeleteOK(2) // s is now [1, 2, 4, 5], deleted is true
func (slice *Slice[T]) DeleteOK(i int) bool {
	if slice.Bounds(i) {
		slice.Delete(i)
		return true
	}
	return false
}

// Each executes a provided function once for each element in the Slice and returns the Slice.
// It iterates over each element of the slice, invoking the provided function for each element, and returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[string]{"apple", "banana", "cherry"}
//	s.Each(func(i int, value string) {
//	    fmt.Printf("Element %d: %s\n", i, value)
//	})
func (slice *Slice[T]) Each(fn func(i int, value T)) *Slice[T] {
	slice.EachBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachBreak executes a provided function once for each element in the Slice with an optional break when the function returns false.
// It iterates over each element of the slice, invoking the provided function for each element. If the function returns false, the iteration stops.
// It returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	s.EachBreak(func(i int, value int) bool {
//	    fmt.Printf("Element %d: %d\n", i, value)
//	    return i < 3 // Stop iteration when i is less than 3
//	})
func (slice *Slice[T]) EachBreak(fn func(i int, value T) bool) *Slice[T] {
	for i, v := range *slice {
		ok := fn(i, v)
		if !ok {
			break
		}
	}
	return slice
}

// EachReverse executes a provided function once for each element in the reverse order they are stored in the Slice.
// It iterates over each element of the Slice in reverse order, invoking the provided function for each element, and returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	s.EachReverse(func(i int, value int) {
//	    fmt.Printf("Element %d: %d\n", i, value)
//	})
func (slice *Slice[T]) EachReverse(fn func(i int, value T)) *Slice[T] {
	slice.EachReverseBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachReverseBreak executes a provided function once for each element in the reverse order they are stored in the Slice with an optional break when the function returns false.
// It iterates over each element of the Slice in reverse order, invoking the provided function for each element. If the function returns false, the iteration stops.
// It returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	s.EachReverseBreak(func(i int, value int) bool {
//	    fmt.Printf("Element %d: %d\n", i, value)
//	    return i > 2 // Stop iteration when i is greater than 2
//	})
func (slice *Slice[T]) EachReverseBreak(fn func(i int, value T) bool) *Slice[T] {
	for i := len(*slice) - 1; i >= 0; i-- {
		ok := fn(i, (*slice)[i])
		if !ok {
			break
		}
	}
	return slice
}

// Fetch retrieves the element held at the specified index in the Slice.
//
// Example:
//
//	s := &Slice[string]{"apple", "banana", "cherry"}
//	fruit := s.Fetch(1)
//	// fruit will be "banana"
func (slice *Slice[T]) Fetch(i int) T {
	v, _ := slice.Get(i)
	return v
}

// FetchLength retrieves the element held at the specified index in the Slice and the length of the Slice.
// It returns the element at the specified index and the length of the Slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30, 40, 50}
//	value, length := s.FetchLength(2)
//	// value will be 30
//	// length will be 5
func (slice *Slice[T]) FetchLength(i int) (T, int) {
	return slice.Fetch(i), slice.Length()
}

// Filter creates a new Slice containing only the elements that satisfy the given predicate function.
// It iterates over the elements of the Slice and applies the predicate function to each element.
// Elements for which the predicate returns true are included in the new Slice, and others are excluded.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	filtered := s.Filter(func(x int) bool {
//		return x%2 == 0 // Keep only even numbers
//	})
//	// filtered will be &Slice[int]{2, 4}
func (slice *Slice[T]) Filter(fn func(i int, value T) bool) *Slice[T] {
	s := &Slice[T]{}
	slice.Each(func(i int, value T) {
		if fn(i, value) {
			s.Append(value)
		}
	})
	return s
}

// FindIndex finds the index of the first element that satisfies the given predicate function.
// It returns the index of the first matching element and true if found; otherwise, it returns -1 and false.
//
// Example:
//
//	s := &Slice[string]{"apple", "banana", "cherry"}
//	index, found := s.FindIndex(func(fruit string) bool {
//	    return fruit == "banana"
//	})
//	// index will be 1
//	// found will be true
func (slice *Slice[T]) FindIndex(fn func(value T) bool) (int, bool) {
	var index int
	var ok bool
	slice.EachBreak(func(i int, value T) bool {
		index = i
		ok = fn(value)
		if !ok {
			index = -1
		}
		return !ok
	})
	return index, ok
}

// Get retrieves the element held at the specified index in the Slice and a boolean indicating if it was successfully retrieved.
// It returns the element at the specified index and a boolean indicating whether the element was successfully retrieved.
//
// Example:
//
//	s := &Slice[float64]{3.14, 2.71, 1.61}
//	value, ok := s.Get(1)
//	// value will be 2.71
//	// ok will be true
func (slice *Slice[T]) Get(i int) (T, bool) {
	var (
		ok = slice.Bounds(i)
		v  T
	)
	if ok {
		v = (*slice)[i]
	}
	return v, ok
}

// GetLength retrieves the element held at the specified index in the slice, the length of the slice, and a boolean indicating if it was successfully retrieved.
// It returns the element at the specified index, the length of the slice, and a boolean indicating whether the element was successfully retrieved.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30, 40, 50}
//	value, length, ok := s.GetLength(2)
//	// value will be 30
//	// length will be 5
//	// ok will be true
func (slice *Slice[T]) GetLength(i int) (T, int, bool) {
	v, ok := slice.Get(i)
	l := slice.Length()
	return v, l, ok
}

// IsEmpty returns whether the Slice is empty.
// It returns true if the Slice is empty (length is zero), otherwise false.
//
// Example:
//
//	s := &Slice[int]{}
//	isEmpty := s.IsEmpty() // isEmpty will be true
func (slice *Slice[T]) IsEmpty() bool {
	return len(*slice) == 0
}

// IsPopulated returns whether the Slice is populated.
// It returns true if the Slice is not empty (length is greater than zero), otherwise false.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	isPopulated := s.IsPopulated() // isPopulated will be true
func (slice *Slice[T]) IsPopulated() bool {
	return !slice.IsEmpty()
}

// Length returns the number of elements in the Slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30, 40, 50}
//	length := s.Length() // length will be 5
func (slice *Slice[T]) Length() int {
	return len(*slice)
}

// Make empties the slice, sets the new Slice to the length of n, and returns the modified Slice.
// It replaces the existing Slice with a new Slice of the specified length (n) and returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	s.Make(3) // s will be an empty Slice of length 3
func (slice *Slice[T]) Make(i int) *Slice[T] {
	*slice = make(Slice[T], i)
	return slice
}

// MakeEach empties the slice, sets the new Slice to the length of n, and populates it by performing a for-each loop on the provided values.
// Finally, it returns a pointer to the modified Slice.
// It replaces the existing Slice with a new Slice of the specified length (n) and populates it by performing a for-each loop on the provided values.
//
// Example:
//
//	s := &Slice[int]{}
//	s.MakeEach(10, 20, 30) // s will be a Slice containing {10, 20, 30}
func (slice *Slice[T]) MakeEach(v ...T) *Slice[T] {
	return slice.Make(len(v)).Each(func(i int, _ T) {
		slice.Replace(i, v[i])
	})
}

// MakeEachReverse empties the slice, sets the new Slice to the length of n, and performs an inverse for-each loop on the provided values,
// inserting each entry at the appropriate index before returning the modified Slice.
// It replaces the existing Slice with a new Slice of the specified length (n) and populates it by performing an inverse for-each loop on the provided values.
//
// Example:
//
//	s := &Slice[int]{}
//	s.MakeEachReverse(10, 20, 30) // s will be a Slice containing {30, 20, 10}
func (slice *Slice[T]) MakeEachReverse(v ...T) *Slice[T] {
	return slice.Make(len(v)).EachReverse(func(i int, _ T) {
		slice.Replace(i, v[i])
	})
}

// Map executes a provided function once for each element and sets the returned value to the current index.
// It iterates over each element of the slice, applying the provided function to each element,
// and assigns the returned value to the current index in the slice. Returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	s.Map(func(i int, value int) int {
//	  return value * 2
//	}) // s will be a Slice containing {20, 40, 60}
func (slice *Slice[T]) Map(fn func(i int, value T) T) *Slice[T] {
	slice.Each(func(i int, value T) {
		slice.Replace(i, fn(i, value))
	})
	return slice
}

// MapReverse executes a provided function once for each element in reverse order and sets the returned value to the current index.
// It iterates over each element of the Slice in reverse order, applying the provided function to each element,
// and assigns the returned value to the current index in the slice. Returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	s.MapReverse(func(i int, value int) int {
//	  return value * 2
//	}) // s will be a Slice containing {60, 40, 20}
func (slice *Slice[T]) MapReverse(fn func(i int, value T) T) *Slice[T] {
	slice.EachReverse(func(i int, value T) {
		slice.Replace(i, fn(i, value))
	})
	return slice
}

// Poll removes the first element from the Slice and returns that removed element.
// It removes and returns the first element of the Slice if the Slice is not empty.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	value := s.Poll() // value will be 10, and s will be [20, 30].
func (slice *Slice[T]) Poll() T {
	var v T
	if !slice.IsEmpty() {
		v = (*slice)[0]
		*slice = (*slice)[1:]
	}
	return v
}

// PollLength removes the first element from the Slice and returns the removed element and the length of the modified Slice.
// It removes and returns the first element of the Slice along with the new length of the modified Slice if the Slice is not empty.
func (slice *Slice[T]) PollLength() (T, int) {
	return slice.Poll(), slice.Length()
}

// PollOK removes the first element from the Slice and returns a boolean on the outcome of the transaction.
// It removes the first element from the Slice and returns true if the Slice is not empty; otherwise, it returns false.
func (slice *Slice[T]) PollOK() (T, bool) {
	var (
		ok = slice.IsPopulated()
		v  T
	)
	if ok {
		v = slice.Poll()
	}
	return v, ok
}

// Pop removes the last element from the Slice and returns that element.
// It removes and returns the last element of the Slice if the Slice is not empty.
func (slice *Slice[T]) Pop() T {
	var v T
	if slice.IsPopulated() {
		l := slice.Length()
		v = (*slice)[l-1]
		*slice = (*slice)[:l-1]
	}
	return v
}

// PopLength removes the last element from the Slice and returns the removed element along with the length of the modified Slice.
// It removes and returns the last element of the Slice if the Slice is not empty, along with the new length of the modified Slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	value, length := s.PopLength() // value will be 30, length will be 2, and s will be [10, 20]
func (slice *Slice[T]) PopLength() (T, int) {
	return slice.Pop(), slice.Length()
}

// PopOK removes the last element from the Slice and returns a boolean indicating the outcome of the transaction.
// It removes the last element from the Slice and returns true if the Slice is not empty; otherwise, it returns false.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	value, ok := s.PopOK() // value will be 30, ok will be true, and s will be [10, 20]
func (slice *Slice[T]) PopOK() (T, bool) {
	var (
		ok = slice.IsPopulated()
		v  T
	)
	if ok {
		v = slice.Pop()
	}
	return v, ok
}

// Precatenate merges the elements from the argument Slice to the head of the receiver Slice and returns the modified Slice.
// If the provided Slice (s) is not nil, it prepends its elements to the receiver Slice and returns a pointer to the modified Slice.
//
// Example:
//
//	s1 := &Slice[int]{1, 2, 3}
//	s2 := &Slice[int]{4, 5}
//	s1.Precatenate(s2) // s1 will be [4, 5, 1, 2, 3]
func (slice *Slice[T]) Precatenate(s *Slice[T]) *Slice[T] {
	if s != nil {
		slice.Prepend((*s)...)
	}
	return slice
}

// PrecatenateFunc prepends elements from another Slice to the head of the receiver Slice based on a provided predicate function.
// It iterates over each element in the source slice, invoking the provided function. If the function returns 'true' for an element,
// that element is prepended to the receiver slice. If it returns 'false', the element is skipped.
//
// Example:
//
//	s1 := &Slice[int]{1, 2, 3}
//	s2 := &Slice[int]{4, 5, 6}
//	result := s1.PrecatenateFunc(s2, func(i int, value int) bool {
//	    return value%2 == 0
//	}) // s1 will be modified to [6, 4, 2, 1, 3], and 'result' will be a pointer to 's1'.
func (slice *Slice[T]) PrecatenateFunc(s *Slice[T], fn func(i int, value T) bool) *Slice[T] {
	return slice.PrependFunc(fn, (*s)...)
}

// PrecatenateLength merges the elements from the argument Slice to the head of the receiver Slice and returns the length of the modified Slice.
// If the provided Slice (s) is not nil, it prepends its elements to the receiver Slice and returns the length of the modified Slice.
//
// Example:
//
//	s1 := &Slice[int]{1, 2, 3}
//	s2 := &Slice[int]{4, 5}
//	length := s1.PrecatenateLength(s2) // length will be 5, and s1 will be [4, 5, 1, 2, 3]
func (slice *Slice[T]) PrecatenateLength(s *Slice[T]) int {
	return slice.Precatenate(s).Length()
}

// Prepend adds one element to the head of the Slice and returns the modified Slice.
// It adds the specified values to the beginning of the existing Slice and returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{2, 3}
//	s.Prepend(1) // s will be [1, 2, 3]
func (slice *Slice[T]) Prepend(values ...T) *Slice[T] {
	*slice = append(values, *slice...)
	return slice
}

// PrependFunc prepends elements to the head of the receiver Slice based on a provided predicate function.
// It iterates over each element in the 'values' argument, invoking the provided function. If the function returns 'true' for an element,
// that element is prepended to the receiver slice. If it returns 'false', the element is skipped.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3}
//	result := s.PrependFunc(func(i int, value int) bool {
//	    return value%2 == 0
//	}, 4, 5, 6) // 's' will be modified to [6, 4, 2, 1, 3], and 'result' will be a pointer to 's'.
func (slice *Slice[T]) PrependFunc(fn func(i int, value T) bool, values ...T) *Slice[T] {
	for i, value := range values {
		if fn(i, value) {
			slice.Prepend(value)
		}
	}
	return slice
}

// PrependLength adds n elements to the head of the Slice and returns the length of the modified Slice.
// It adds the specified values to the beginning of the existing Slice and returns the length of the modified Slice.
//
// Example:
//
//	s := &Slice[int]{2, 3}
//	length := s.PrependLength(1, 0) // length will be 4, and s will be [1, 0, 2, 3]
func (slice *Slice[T]) PrependLength(values ...T) int {
	return slice.Prepend(values...).Length()
}

// Reduce creates a new slice containing elements from the receiver slice that satisfy a provided predicate function.
// It iterates over each element in the receiver slice, invoking the provided function. If the function returns 'true' for an element,
// that element is appended to the new slice. If it returns 'false', the element is skipped.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	result := s.Reduce(func(i int, value int) bool {
//	    return value%2 == 0
//	}) // 'result' will be a new slice containing [2, 4].
func (slice *Slice[T]) Reduce(fn func(i int, value T) bool) *Slice[T] {
	s := &Slice[T]{}
	slice.Each(func(i int, value T) {
		if ok := fn(i, value); ok {
			s.Append(value)
		}
	})
	return s
}

// Replace changes the contents of the Slice at the argument index if it is in bounds.
// It replaces the element at the specified index with the provided value if the index is within bounds and returns true.
// Otherwise, it returns false.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3}
//	ok := s.Replace(1, 4) // ok will be true, and s will be [1, 4, 3]
func (slice *Slice[T]) Replace(i int, value T) bool {
	ok := slice.Bounds(i)
	if ok {
		(*slice)[i] = value
	}
	return ok
}

// Reverse reverses the Slice in linear time and returns the modified Slice.
// It reverses the order of elements in the Slice and returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3}
//	s.Reverse() // s will be [3, 2, 1]
func (slice *Slice[T]) Reverse() *Slice[T] {
	var (
		i = 0
		j = slice.Length() - 1
	)
	for i < j {
		slice.Swap(i, j)
		i = i + 1
		j = j - 1
	}
	return slice
}

// Set returns a unique slice, removing duplicate elements that have the same hash value.
// Returns the modified Slice at the end of the iteration.
// It removes duplicate elements from the slice, keeping only the first occurrence of each unique element.
// Returns a pointer to the modified Slice with unique elements.
//
// Example:
//
//	s := &Slice[int]{1, 2, 2, 3, 3, 3}
//	s.Set() // s will be [1, 2, 3]
func (slice *Slice[T]) Set() *Slice[T] {
	var (
		k  string
		m  = map[string]bool{}
		ok bool
		s  = &Slice[T]{}
	)
	slice.Each(func(_ int, value T) {
		k = fmt.Sprintf("%v", value) // TODO: Check if there is a better way to generate key.
		_, ok = m[k]
		if !ok {
			s.Append(value)
		}
		m[k] = true
	})
	*slice = *s
	return slice
}

// Shuffle randomly shuffles the elements of the Slice and returns the modified Slice.
// It shuffles the elements of the Slice in a random order and returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	s.Shuffle() // s will be a random permutation of [1, 2, 3, 4, 5]
func (slice *Slice[T]) Shuffle() *Slice[T] {
	rand.Shuffle(len(*slice), func(i, j int) {
		slice.Swap(i, j)
	})
	return slice
}

// Slice slices the Slice from i to j and returns the modified Slice.
// It slices the Slice from the specified start (i) index to the end (j) index (inclusive),
// and returns a pointer to the modified Slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	s.Slice(1, 3) // s will be [2, 3, 4]
func (slice *Slice[T]) Slice(i int, j int) *Slice[T] {
	if j < i {
		i, j = j, i
	}
	if slice.Bounds(i) && slice.Bounds(j) {
		*slice = (*slice)[i:j]
	}
	return slice
}

// Swap moves element i to j and j to i.
// If both indices (i and j) are within bounds, it swaps the elements at those positions in the Slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3}
//	s.Swap(0, 2) // s will be [3, 2, 1]
func (slice *Slice[T]) Swap(i int, j int) {
	if slice.Bounds(i) && slice.Bounds(j) {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

// New creates a new instance of the Slice[T] type and initializes it with the provided values.
// It allows you to create a new slice and populate it with the specified elements.
//
// Example:
//
//	s := New[int](1, 2, 3, 4, 5)
//	// 's' will be a pointer to a new slice containing [1, 2, 3, 4, 5].
func New[T any](values ...T) *Slice[T] {
	return (&Slice[T]{}).Append(values...)
}
