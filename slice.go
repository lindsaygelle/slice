package slice

import (
	"fmt"
	"reflect"
)

// Slice is a list-like struct whose methods are used to perform traversal and mutation operations by numeric index.
type Slice[T any] []T

// Append adds n elements to the end of the slice and returns the modified slice.
// It appends the specified values to the existing slice and returns a pointer to the modified slice.
func (slice *Slice[T]) Append(values ...T) *Slice[T] {
	*slice = append(*slice, values...)
	return slice
}

// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
// It appends the specified values to the existing slice and returns the length of the modified slice.
func (slice *Slice[T]) AppendLength(values ...T) int {
	return slice.Append(values...).Length()
}

// Bounds checks if an integer value safely sits within the range of accessible values for the slice.
// It returns true if the index is within bounds, otherwise false.
func (slice *Slice[T]) Bounds(i int) bool {
	return i >= 0 && i < len(*slice)
}

// Concatenate merges the elements from the argument slice to the tail of the argument slice.
// If the provided slice (s) is not nil, it appends its elements to the receiver slice and returns a pointer to the modified slice.
func (slice *Slice[T]) Concatenate(s *Slice[T]) *Slice[T] {
	if s != nil {
		slice.Append((*s)...)
	}
	return slice
}

// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice and returns the length of the receiver slice.
// If the provided slice (s) is not nil, it appends its elements to the receiver slice and returns the length of the modified slice.
func (slice *Slice[T]) ConcatenateLength(s *Slice[T]) int {
	return slice.Concatenate(s).Length()
}

// Contains checks if a value exists in the slice.
// It returns true if the value is found in the slice, otherwise false.
func (slice *Slice[T]) Contains(value T) bool {
	for _, v := range *slice {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Delete deletes the element from the argument index and returns the modified slice.
// If the index is within bounds, it removes the element at that index and returns a pointer to the modified slice.
func (slice *Slice[T]) Delete(i int) *Slice[T] {
	if slice.Bounds(i) {
		*slice = append((*slice)[:i], (*slice)[i+1:]...)
	}
	return slice
}

// DeleteLength deletes the element from the argument index and returns the new length of the slice.
// If the index is within bounds, it removes the element at that index and returns the new length of the modified slice.
func (slice *Slice[T]) DeleteLength(i int) int {
	return slice.Delete(i).Length()
}

// DeleteOK deletes the element from the argument index and returns the result of the transaction.
// If the index is within bounds, it removes the element at that index and returns true; otherwise, it returns false.
func (slice *Slice[T]) DeleteOK(i int) bool {
	ok := slice.Bounds(i)
	if ok {
		slice.Delete(i)
	}
	return ok
}

// Each executes a provided function once for each slice element and returns the slice.
// It iterates over each element of the slice, invoking the provided function, and returns a pointer to the modified slice.
func (slice *Slice[T]) Each(fn func(int, T)) *Slice[T] {
	slice.EachBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachBreak executes a provided function once for each element with an optional break when the function returns false.
// It iterates over each element of the slice, invoking the provided function. If the function returns false, the iteration stops.
// Returns the slice at the end of the iteration.
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
// It iterates over each element of the slice in reverse order, invoking the provided function, and returns a pointer to the modified slice.
func (slice *Slice[T]) EachReverse(fn func(int, T)) *Slice[T] {
	slice.EachReverseBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachReverseBreak executes a provided function once for each element in the reverse order they are stored in the slice with an optional break when the function returns false.
// It iterates over each element of the slice in reverse order, invoking the provided function. If the function returns false, the iteration stops.
// Returns the slice at the end of the iteration.
func (slice *Slice[T]) EachReverseBreak(fn func(int, T) bool) *Slice[T] {
	for i := len(*slice) - 1; i >= 0; i-- {
		ok := fn(i, (*slice)[i])
		if !ok {
			break
		}
	}
	return slice
}

// Fetch retrieves the element held at the argument index. Returns the default type if index exceeds slice length.
// It returns the element at the specified index and a boolean indicating whether the element was successfully retrieved.
func (slice *Slice[T]) Fetch(i int) T {
	v, _ := slice.Get(i)
	return v
}

// FetchLength retrieves the element held at the argument index and the length of the slice. Returns the default type if index exceeds slice length.
// It returns the element at the specified index, the length of the slice, and a boolean indicating whether the element was successfully retrieved.
func (slice *Slice[T]) FetchLength(i int) (T, int) {
	return slice.Fetch(i), slice.Length()
}

// FindIndex finds the index of the first element that satisfies the given predicate function.
// It returns the index of the first matching element and true if found; otherwise, it returns -1 and false.
func (slice *Slice[T]) FindIndex(fn func(T) bool) (int, bool) {
	for i, v := range *slice {
		if fn(v) {
			return i, true
		}
	}
	return -1, false
}

// Get returns the element held at the argument index and a boolean indicating if it was successfully retrieved.
// It returns the element at the specified index and a boolean indicating whether the element was successfully retrieved.
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

// GetLength returns the element at the argument index, the length of the slice, and a boolean indicating if the element was successfully retrieved.
// It returns the element at the specified index, the length of the slice, and a boolean indicating whether the element was successfully retrieved.
func (slice *Slice[T]) GetLength(i int) (T, int, bool) {
	v, ok := slice.Get(i)
	l := slice.Length()
	return v, l, ok
}

// IsEmpty returns whether the slice is empty.
// It returns true if the slice is empty (length is zero), otherwise false.
func (slice Slice[T]) IsEmpty() bool {
	return len(slice) == 0
}

// Length returns the number of elements in the slice.
// It returns the length of the slice.
func (slice *Slice[T]) Length() int {
	return len(*slice)
}

// Make empties the slice, sets the new slice to the length of n, and returns the modified slice.
// It replaces the existing slice with a new slice of the specified length (n) and returns a pointer to the modified slice.
func (slice *Slice[T]) Make(i int) *Slice[T] {
	*slice = make(Slice[T], i)
	return slice
}

// MakeEach empties the slice, sets the new slice to the length of n, and performs a for-each loop for the argument sequence,
// inserting each entry at the appropriate index before returning the modified slice.
// It replaces the existing slice with a new slice of the specified length (n) and populates it by performing a for-each loop on the provided values.
// Finally, it returns a pointer to the modified slice.
func (slice *Slice[T]) MakeEach(v ...T) *Slice[T] {
	return slice.Make(len(v)).Each(func(i int, _ T) {
		slice.Replace(i, v[i])
	})
}

// MakeEachReverse empties the slice, sets the new slice to the length of n, and performs an inverse for-each loop for the argument sequence,
// inserting each entry at the appropriate index before returning the modified slice.
// It replaces the existing slice with a new slice of the specified length (n) and populates it by performing an inverse for-each loop on the provided values.
// Finally, it returns a pointer to the modified slice.
func (slice *Slice[T]) MakeEachReverse(v ...T) *Slice[T] {
	return slice.Make(len(v)).EachReverse(func(i int, _ T) {
		slice.Replace(i, v[i])
	})
}

// Map executes a provided function once for each element and sets the returned value to the current index.
// Returns the slice at the end of the iteration.
// It iterates over each element of the slice, applying the provided function to each element,
// and assigns the returned value to the current index in the slice. Returns a pointer to the modified slice.
func (slice *Slice[T]) Map(fn func(int, T) T) *Slice[T] {
	slice.Each(func(i int, value T) {
		slice.Replace(i, fn(i, value))
	})
	return slice
}

// Poll removes the first element from the slice and returns that removed element.
// It removes and returns the first element of the slice if the slice is not empty.
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

// PopLength removes the last element from the slice and returns the removed element and the length of the modified slice.
// It removes and returns the last element of the slice along with the new length of the modified slice if the slice is not empty.
func (slice *Slice[T]) PopLength() (T, int) {
	return slice.Pop(), slice.Length()
}

// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
// It removes the last element from the slice and returns true if the slice is not empty; otherwise, it returns false.
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

// Precatenate merges the elements from the argument slice to the head of the argument slice and returns the modified slice.
// If the provided slice (s) is not nil, it prepends its elements to the receiver slice and returns a pointer to the modified slice.
func (slice *Slice[T]) Precatenate(s *Slice[T]) *Slice[T] {
	if s != nil {
		slice.Prepend((*s)...)
	}
	return slice
}

// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice and returns the length of the receiver slice.
// If the provided slice (s) is not nil, it prepends its elements to the receiver slice and returns the length of the modified slice.
func (slice *Slice[T]) PrecatenateLength(s *Slice[T]) int {
	return slice.Precatenate(s).Length()
}

// Prepend adds one element to the head of the slice and returns the modified slice.
// It adds the specified values to the beginning of the existing slice and returns a pointer to the modified slice.
func (slice *Slice[T]) Prepend(values ...T) *Slice[T] {
	*slice = append(values, *slice...)
	return slice
}

// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
// It adds the specified values to the beginning of the existing slice and returns the length of the modified slice.
func (slice *Slice[T]) PrependLength(values ...T) int {
	return slice.Prepend(values...).Length()
}

// Replace changes the contents of the slice at the argument index if it is in bounds.
// It replaces the element at the specified index with the provided value if the index is within bounds and returns true.
// Otherwise, it returns false.
func (slice *Slice[T]) Replace(i int, value T) bool {
	ok := slice.Bounds(i)
	if ok {
		(*slice)[i] = value
	}
	return ok
}

// Reverse reverses the slice in linear time and returns the modified slice.
// It reverses the order of elements in the slice and returns a pointer to the modified slice.
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

// Slice[T any] slices the slice from i to j and returns the modified slice.
// It slices the slice from the specified start (i) index to the end (j) index (inclusive),
// and returns a pointer to the modified slice.
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
func (slice *Slice[T]) Swap(i int, j int) {
	if slice.Bounds(i) && slice.Bounds(j) {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}
