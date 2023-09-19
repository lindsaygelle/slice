package slice

import (
	"fmt"
	"reflect"
)

// Slice is a generic list-like struct that allows for the manipulation and traversal of elements by numeric index.
// It is parameterized with the type `T`, allowing it to work with slices of any data type.
//
// Example usage:
//
//	// Create a slice of integers.
//	s := &Slice[int]{1, 2, 3, 4, 5}
//
//	// Append values to the slice.
//	s.Append(6, 7, 8)
//
//	// Check if a value exists in the slice.
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
// This type provides a set of methods for common slice operations and is designed to be generic for flexibility.
type Slice[T any] []T

// Append appends the specified values to the end of the slice and returns a pointer to the modified slice.
// It extends the slice by adding the provided values to the end, updating the slice in place.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3}
//	s.Append(4, 5) // s is now [1, 2, 3, 4, 5]
func (slice *Slice[T]) Append(values ...T) *Slice[T] {
	*slice = append(*slice, values...)
	return slice
}

// AppendLength appends the specified values to the end of the slice and returns the length of the modified slice.
// It extends the slice by adding the provided values to the end, updating the slice in place, and returning the new length.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3}
//	length := s.AppendLength(4, 5) // s is now [1, 2, 3, 4, 5], length is 5
func (slice *Slice[T]) AppendLength(values ...T) int {
	return slice.Append(values...).Length()
}

// AppendSome appends selected elements to the end of the slice based on a provided condition.
// It iterates over the specified values, invoking the provided function for each element.
// If the function returns true for an element, that element is added to the end of the slice.
//
// The `fn` function is a callback function that takes two arguments: the index `i` and the current value of the element.
// It should return `true` to include the element in the slice or `false` to exclude it.
//
// Example:
//
//	s := &slice.Slice[int]{}
//	s.AppendSome(func(i int, value int) bool {
//	  return value%2 == 0 // Append even numbers to the slice.
//	}, 1, 2, 3, 4, 5)
//
// After this operation, `s` will contain [2, 4].
//
// This method modifies the original slice and returns a pointer to the modified slice.
func (slice *Slice[T]) AppendSome(fn func(int, T) bool, values ...T) *Slice[T] {
	for i, value := range values {
		if ok := fn(i, value); ok {
			slice.Append(value)
		}
	}
	return slice
}

// Bounds checks if the given integer index is within the valid range of indices for the slice.
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

// Concatenate merges the elements from the argument slice to the tail of the receiver slice and returns the modified slice.
// If the provided slice (s) is not nil, it appends its elements to the end of the receiver slice, updating the slice in place.
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

// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice and returns the length of the modified slice.
// If the provided slice (s) is not nil, it appends its elements to the end of the receiver slice, updating the slice in place.
//
// Example:
//
//	s1 := &Slice[int]{1, 2, 3}
//	s2 := &Slice[int]{4, 5}
//	length := s1.ConcatenateLength(s2) // s1 is now [1, 2, 3, 4, 5], length is 5
func (slice *Slice[T]) ConcatenateLength(s *Slice[T]) int {
	return slice.Concatenate(s).Length()
}

// ConcatenateSome appends elements from the provided slice (s) to the end of the receiver slice, based on the result of a filtering function.
// It iterates over the elements of the provided slice and appends those elements for which the provided function returns true.
//
// Example:
//
//	s1 := &Slice[int]{1, 2, 3}
//	s2 := &Slice[int]{4, 5, 6}
//	s1.ConcatenateSome(func(i int, value int) bool {
//	    return value%2 == 0
//	}, s2) // s1 is now [1, 2, 3, 4, 6]
func (slice *Slice[T]) ConcatenateSome(fn func(i int, value T) bool, s *Slice[T]) *Slice[T] {
	if s != nil {
		slice.AppendSome(fn, *s...)
	}
	return slice
}

// Contains checks if a value exists in the slice.
// It iterates over the elements of the slice and compares each element with the provided value using reflect.DeepEqual.
// If a matching element is found, it returns true; otherwise, it returns false.
//
// Example:
//
//	s := &Slice[string]{"apple", "banana", "cherry"}
//	containsBanana := s.Contains("banana") // containsBanana is true
func (slice *Slice[T]) Contains(value T) bool {
	for _, v := range *slice {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Delete removes the element at the specified index from the slice, if the index is within bounds.
// It modifies the slice in place by removing the element at the given index and returns a pointer to the modified slice.
// If the index is out of bounds, it does not modify the slice.
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

// DeleteLength removes the element at the specified index from the slice, if the index is within bounds, and returns the new length of the modified slice.
// It modifies the slice in place by removing the element at the given index and returns the new length of the slice.
// If the index is out of bounds, it does not modify the slice and returns the current length.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	newLength := s.DeleteLength(2) // s is now [1, 2, 4, 5], newLength is 4
func (slice *Slice[T]) DeleteLength(i int) int {
	return slice.Delete(i).Length()
}

// DeleteOK removes the element at the specified index from the slice if the index is within bounds and returns the result of the transaction.
// It modifies the slice in place by removing the element at the given index and returns true if the deletion was successful.
// If the index is out of bounds, it does not modify the slice and returns false.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	deleted := s.DeleteOK(2) // s is now [1, 2, 4, 5], deleted is true
func (slice *Slice[T]) DeleteOK(i int) bool {
	ok := slice.Bounds(i)
	if ok {
		slice.Delete(i)
	}
	return ok
}

// Each executes a provided function once for each element in the slice and returns the slice.
// It iterates over each element of the slice, invoking the provided function for each element, and returns a pointer to the modified slice.
//
// Example:
//
//	s := &Slice[string]{"apple", "banana", "cherry"}
//	s.Each(func(i int, value string) {
//	    fmt.Printf("Element %d: %s\n", i, value)
//	})
func (slice *Slice[T]) Each(fn func(int, T)) *Slice[T] {
	slice.EachBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachBreak executes a provided function once for each element in the slice with an optional break when the function returns false.
// It iterates over each element of the slice, invoking the provided function for each element. If the function returns false, the iteration stops.
// It returns a pointer to the modified slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	s.EachBreak(func(i int, value int) bool {
//	    fmt.Printf("Element %d: %d\n", i, value)
//	    return i < 3 // Stop iteration when i is less than 3
//	})
func (slice *Slice[T]) EachBreak(fn func(int, T) bool) *Slice[T] {
	for i, v := range *slice {
		ok := fn(i, v)
		if !ok {
			break
		}
	}
	return slice
}

// EachReverse executes a provided function once for each element in the reverse order they are stored in the slice.
// It iterates over each element of the slice in reverse order, invoking the provided function for each element, and returns a pointer to the modified slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	s.EachReverse(func(i int, value int) {
//	    fmt.Printf("Element %d: %d\n", i, value)
//	})
func (slice *Slice[T]) EachReverse(fn func(int, T)) *Slice[T] {
	slice.EachReverseBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachReverseBreak executes a provided function once for each element in the reverse order they are stored in the slice with an optional break when the function returns false.
// It iterates over each element of the slice in reverse order, invoking the provided function for each element. If the function returns false, the iteration stops.
// It returns a pointer to the modified slice.
//
// Example:
//
//	s := &Slice[int]{1, 2, 3, 4, 5}
//	s.EachReverseBreak(func(i int, value int) bool {
//	    fmt.Printf("Element %d: %d\n", i, value)
//	    return i > 2 // Stop iteration when i is greater than 2
//	})
func (slice *Slice[T]) EachReverseBreak(fn func(int, T) bool) *Slice[T] {
	for i := len(*slice) - 1; i >= 0; i-- {
		ok := fn(i, (*slice)[i])
		if !ok {
			break
		}
	}
	return slice
}

// Fetch retrieves the element held at the specified index in the slice.
// It returns the element at the specified index.
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

// FetchLength retrieves the element held at the specified index in the slice and the length of the slice.
// It returns the element at the specified index and the length of the slice.
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
func (slice *Slice[T]) FindIndex(fn func(T) bool) (int, bool) {
	for i, v := range *slice {
		if fn(v) {
			return i, true
		}
	}
	return -1, false
}

// Get retrieves the element held at the specified index in the slice and a boolean indicating if it was successfully retrieved.
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

// IsEmpty returns whether the slice is empty.
// It returns true if the slice is empty (length is zero), otherwise false.
//
// Example:
//
//	s := &Slice[int]{}
//	isEmpty := s.IsEmpty() // isEmpty will be true
func (slice *Slice[T]) IsEmpty() bool {
	return len(*slice) == 0
}

// IsPopulated returns whether the slice is populated.
// It returns true if the slice is not empty (length is greater than zero), otherwise false.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	isPopulated := s.IsPopulated() // isPopulated will be true
func (slice *Slice[T]) IsPopulated() bool {
	return !slice.IsEmpty()
}

// Length returns the number of elements in the slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30, 40, 50}
//	length := s.Length() // length will be 5
func (slice *Slice[T]) Length() int {
	return len(*slice)
}

// Make empties the slice, sets the new slice to the length of n, and returns the modified slice.
// It replaces the existing slice with a new slice of the specified length (n) and returns a pointer to the modified slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	s.Make(3) // s will be an empty slice of length 3
func (slice *Slice[T]) Make(i int) *Slice[T] {
	*slice = make(Slice[T], i)
	return slice
}

// MakeEach empties the slice, sets the new slice to the length of n, and populates it by performing a for-each loop on the provided values.
// Finally, it returns a pointer to the modified slice.
// It replaces the existing slice with a new slice of the specified length (n) and populates it by performing a for-each loop on the provided values.
//
// Example:
//
//	s := &Slice[int]{}
//	s.MakeEach(10, 20, 30) // s will be a slice containing {10, 20, 30}
func (slice *Slice[T]) MakeEach(v ...T) *Slice[T] {
	return slice.Make(len(v)).Each(func(i int, _ T) {
		slice.Replace(i, v[i])
	})
}

// MakeEachReverse empties the slice, sets the new slice to the length of n, and performs an inverse for-each loop on the provided values,
// inserting each entry at the appropriate index before returning the modified slice.
// It replaces the existing slice with a new slice of the specified length (n) and populates it by performing an inverse for-each loop on the provided values.
//
// Example:
//
//	s := &Slice[int]{}
//	s.MakeEachReverse(10, 20, 30) // s will be a slice containing {30, 20, 10}
func (slice *Slice[T]) MakeEachReverse(v ...T) *Slice[T] {
	return slice.Make(len(v)).EachReverse(func(i int, _ T) {
		slice.Replace(i, v[i])
	})
}

// Map executes a provided function once for each element and sets the returned value to the current index.
// It iterates over each element of the slice, applying the provided function to each element,
// and assigns the returned value to the current index in the slice. Returns a pointer to the modified slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	s.Map(func(i int, value int) int {
//	  return value * 2
//	}) // s will be a slice containing {20, 40, 60}
func (slice *Slice[T]) Map(fn func(int, T) T) *Slice[T] {
	slice.Each(func(i int, value T) {
		slice.Replace(i, fn(i, value))
	})
	return slice
}

// MapReverse executes a provided function once for each element in reverse order and sets the returned value to the current index.
// It iterates over each element of the slice in reverse order, applying the provided function to each element,
// and assigns the returned value to the current index in the slice. Returns a pointer to the modified slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	s.MapReverse(func(i int, value int) int {
//	  return value * 2
//	}) // s will be a slice containing {60, 40, 20}
func (slice *Slice[T]) MapReverse(fn func(int, T) T) *Slice[T] {
	slice.EachReverse(func(i int, value T) {
		slice.Replace(i, fn(i, value))
	})
	return slice
}

// Poll removes the first element from the slice and returns that removed element.
// It removes and returns the first element of the slice if the slice is not empty.
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

// PollLength removes the first element from the slice and returns the removed element and the length of the modified slice.
// It removes and returns the first element of the slice along with the new length of the modified slice if the slice is not empty.
func (slice *Slice[T]) PollLength() (T, int) {
	return slice.Poll(), slice.Length()
}

// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
// It removes the first element from the slice and returns true if the slice is not empty; otherwise, it returns false.
func (slice *Slice[T]) PollOK() (T, bool) {
	var (
		ok = !slice.IsEmpty()
		v  T
	)
	if ok {
		v = slice.Poll()
	}
	return v, ok
}

// Pop removes the last element from the slice and returns that element.
// It removes and returns the last element of the slice if the slice is not empty.
func (slice *Slice[T]) Pop() T {
	var v T
	if !slice.IsEmpty() {
		l := slice.Length()
		v = (*slice)[l-1]
		*slice = (*slice)[:l-1]
	}
	return v
}

// PopLength removes the last element from the slice and returns the removed element along with the length of the modified slice.
// It removes and returns the last element of the slice if the slice is not empty, along with the new length of the modified slice.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	value, length := s.PopLength() // value will be 30, length will be 2, and s will be [10, 20]
func (slice *Slice[T]) PopLength() (T, int) {
	return slice.Pop(), slice.Length()
}

// PopOK removes the last element from the slice and returns a boolean indicating the outcome of the transaction.
// It removes the last element from the slice and returns true if the slice is not empty; otherwise, it returns false.
//
// Example:
//
//	s := &Slice[int]{10, 20, 30}
//	value, ok := s.PopOK() // value will be 30, ok will be true, and s will be [10, 20]
func (slice *Slice[T]) PopOK() (T, bool) {
	var (
		ok = !slice.IsEmpty()
		v  T
	)
	if ok {
		v = slice.Pop()
	}
	return v, ok
}

// Precatenate merges the elements from the argument slice to the head of the receiver slice and returns the modified slice.
// If the provided slice (s) is not nil, it prepends its elements to the receiver slice and returns a pointer to the modified slice.
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

// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice and returns the length of the modified slice.
// If the provided slice (s) is not nil, it prepends its elements to the receiver slice and returns the length of the modified slice.
//
// Example:
//
//	s1 := &Slice[int]{1, 2, 3}
//	s2 := &Slice[int]{4, 5}
//	length := s1.PrecatenateLength(s2) // length will be 5, and s1 will be [4, 5, 1, 2, 3]
func (slice *Slice[T]) PrecatenateLength(s *Slice[T]) int {
	return slice.Precatenate(s).Length()
}

// Prepend adds one element to the head of the slice and returns the modified slice.
// It adds the specified values to the beginning of the existing slice and returns a pointer to the modified slice.
//
// Example:
//
//	s := &Slice[int]{2, 3}
//	s.Prepend(1) // s will be [1, 2, 3]
func (slice *Slice[T]) Prepend(values ...T) *Slice[T] {
	*slice = append(values, *slice...)
	return slice
}

// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
// It adds the specified values to the beginning of the existing slice and returns the length of the modified slice.
//
// Example:
//
//	s := &Slice[int]{2, 3}
//	length := s.PrependLength(1, 0) // length will be 4, and s will be [1, 0, 2, 3]
func (slice *Slice[T]) PrependLength(values ...T) int {
	return slice.Prepend(values...).Length()
}

// Replace changes the contents of the slice at the argument index if it is in bounds.
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

// Reverse reverses the slice in linear time and returns the modified slice.
// It reverses the order of elements in the slice and returns a pointer to the modified slice.
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
// Returns the modified slice at the end of the iteration.
// It removes duplicate elements from the slice, keeping only the first occurrence of each unique element.
// Returns a pointer to the modified slice with unique elements.
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
		k = fmt.Sprintf("%v", value) // Todo: Check if there is a better way to generate key.
		_, ok = m[k]
		if !ok {
			s.Append(value)
		}
		m[k] = true
	})
	*slice = *s
	return slice
}

// Slice slices the slice from i to j and returns the modified slice.
// It slices the slice from the specified start (i) index to the end (j) index (inclusive),
// and returns a pointer to the modified slice.
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
// If both indices (i and j) are within bounds, it swaps the elements at those positions in the slice.
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
