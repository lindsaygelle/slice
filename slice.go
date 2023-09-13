package slice

import (
	"fmt"
)

// Slice is a list-like struct whose methods are used to perform traversal and mutation operations by numeric index.
type Slice[T any] []T

// Append adds n elements to the end of the slice and returns the modified slice.
func (slice *Slice[T]) Append(values ...T) *Slice[T] {
	(*slice) = (append(*slice, values...))
	return slice
}

// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
func (slice *Slice[T]) AppendLength(values ...T) int {
	return (slice.Append(values...).Length())
}

// Bounds checks an integer value safely sits within the range of accessible values for the slice.
func (slice *Slice[T]) Bounds(i int) bool {
	return ((i > -1) && (i < len(*slice)))
}

// Concatenate merges the elements from the argument slice to the the tail of the argument slice.
func (slice *Slice[T]) Concatenate(s *Slice[T]) *Slice[T] {
	if s != nil {
		slice.Append((*s)...)
	}
	return slice
}

// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice and returns the length of the receiver slice.
func (slice *Slice[T]) ConcatenateLength(s *Slice[T]) int {
	return (slice.Concatenate(s).Length())
}

// Delete deletes the element from the argument index and returns the modified slice.
func (slice *Slice[T]) Delete(i int) *Slice[T] {
	var (
		ok = slice.Bounds(i)
	)
	if ok {
		(*slice) = append((*slice)[:i], (*slice)[i+1:]...)
	}
	return slice
}

// DeleteLength deletes the element from the argument index and returns the new length of the slice.
func (slice *Slice[T]) DeleteLength(i int) int {
	return slice.Delete(i).Length()
}

// DeleteOK deletes the element from the argument index and returns the result of the transaction.
func (slice *Slice[T]) DeleteOK(i int) bool {
	var (
		ok = slice.Bounds(i)
	)
	if ok {
		slice.Delete(i)
	}
	return ok
}

// Each executes a provided function once for each slice element and returns the slice.
func (slice *Slice[T]) Each(fn func(int, T)) *Slice[T] {
	slice.EachBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachBreak executes a provided function once for each element with an optional break when the function returns false.
// Returns the slice at the end of the iteration.
func (slice *Slice[T]) EachBreak(fn func(int, T) bool) *Slice[T] {
	var (
		i  int
		ok = true
		v  T
	)
	for i, v = range *slice {
		ok = fn(i, v)
		if !ok {
			break
		}
	}
	return slice
}

// EachReverse executes a provided function once for each element in the reverse order they are stored in the slice.
// Returns the slice at the end of the iteration.
func (slice *Slice[T]) EachReverse(fn func(int, T)) *Slice[T] {
	slice.EachReverseBreak(func(i int, value T) bool {
		fn(i, value)
		return true
	})
	return slice
}

// EachReverseBreak executes a provided function once for each element in the reverse order they are stored in the slice with an optional break when the function returns false.
// Returns the slice at the end of the iteration.
func (slice *Slice[T]) EachReverseBreak(fn func(int, T) bool) *Slice[T] {
	var (
		i  int
		ok = true
	)
	for i = len(*slice) - 1; i >= 0; i-- {
		ok = fn(i, (*slice)[i])
		if !ok {
			break
		}
	}
	return slice
}

// Fetch retrieves the element held at the argument index. Returns the default type if index exceeds slice length.
func (slice *Slice[T]) Fetch(i int) T {
	var v, _ = slice.Get(i)
	return v
}

// FetchLength retrives the element held at the argument index and the length of the slice. Returns the default type if index exceeds slice length.
func (slice *Slice[T]) FetchLength(i int) (T, int) {
	return slice.Fetch(i), slice.Length()
}

// Get returns the element held at the argument index and a boolean indicating if it was successfully retrieved.
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

// GetLength returns the element at the argument index, the length of the slice and a boolean indicating if the element was successfully retrieved.
func (slice *Slice[T]) GetLength(i int) (T, int, bool) {
	var v, ok = slice.Get(i)
	var l = slice.Length()
	return v, l, ok
}

// Length returns the number of elements in the slice.
func (slice *Slice[T]) Length() int {
	return (len(*slice))
}

// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
func (slice *Slice[T]) Make(i int) *Slice[T] {
	(*slice) = make(Slice[T], i)
	return slice
}

// MakeEach empties the slice, sets the new slice to the length of n and performs a for-each loop for the argument sequence, inserting each entry at the appropriate index before returning the modified slice.
func (slice *Slice[T]) MakeEach(v ...T) *Slice[T] {
	return slice.Make(len(v)).Each(func(i int, _ T) {
		slice.Replace(i, v[i])
	})
}

// MakeEachReverse empties the slice, sets the new slice to the length of n and performs an inverse for-each loop for the argument sequence,
// inserting each entry at the appropriate index before returning the modified slice.
func (slice *Slice[T]) MakeEachReverse(v ...T) *Slice[T] {
	return slice.Make(len(v)).EachReverse(func(i int, _ T) {
		slice.Replace(i, v[i])
	})
}

// Map executes a provided function once for each element and sets the returned value to the current index. Returns the slice at the end of the iteration.
func (slice *Slice[T]) Map(fn func(int, T) T) *Slice[T] {
	slice.Each(func(i int, value T) {
		slice.Replace(i, fn(i, value))
	})
	return slice
}

// Poll removes the first element from the slice and returns that removed element.
func (slice *Slice[T]) Poll() T {
	var (
		l  = slice.Length()
		ok = l > 0
		v  T
	)
	if ok {
		v = (*slice)[0]
		(*slice) = (*slice)[1:]
	}
	return v
}

// PollLength removes the first element from the slice and returns the removed element and the length of the modified slice.
func (slice *Slice[T]) PollLength() (T, int) {
	return slice.Poll(), slice.Length()
}

// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
func (slice *Slice[T]) PollOK() (T, bool) {
	var (
		ok = slice.Bounds(0)
		v  T
	)
	if ok {
		v = slice.Poll()
	}
	return v, ok
}

// Pop removes the last element from the slice and returns that element.
func (slice *Slice[T]) Pop() T {
	var (
		l  = slice.Length()
		ok = l > 0
		v  T
	)
	if ok {
		v = (*slice)[l-1]
		(*slice) = (*slice)[:l-1]
	}
	return v
}

// PopLength removes the last element from the slice and returns the removed element and the length of the modified slice.
func (slice *Slice[T]) PopLength() (T, int) {
	return slice.Pop(), slice.Length()
}

// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
func (slice *Slice[T]) PopOK() (T, bool) {
	var (
		ok = slice.Length() > 0
		v  T
	)
	if ok {
		v = slice.Pop()
	}
	return v, ok
}

// Precatenate merges the elements from the argument slice to the the head of the argument slice and returns the modified slice.
func (slice *Slice[T]) Precatenate(s *Slice[T]) *Slice[T] {
	if s != nil {
		slice.Prepend((*s)...)
	}
	return slice
}

// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice and returns the length of the receiver slice.
func (slice *Slice[T]) PrecatenateLength(s *Slice[T]) int {
	return (slice.Precatenate(s).Length())
}

// Prepend adds one element to the head of the slice and returns the modified slice.
func (slice *Slice[T]) Prepend(values ...T) *Slice[T] {
	(*slice) = (append(values, *slice...))
	return slice
}

// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
func (slice *Slice[T]) PrependLength(values ...T) int {
	return (slice.Prepend(values...).Length())
}

// Push adds a new element to the end of the slice and returns the length of the modified slice.
func (slice *Slice[T]) Push(values ...T) int {
	return (slice.Append(values...).Length())
}

// Replace changes the contents of the slice at the argument index if it is in bounds.
func (slice *Slice[T]) Replace(i int, value T) bool {
	var (
		ok = slice.Bounds(i)
	)
	if ok {
		(*slice)[i] = value
	}
	return ok
}

// Reverse reverses the slice in linear time. Returns the slice at the end of the iteration.
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

// Set returns a unique slice, removing duplicate elements that have the same hash value. Returns the modified at the end of the iteration.
func (slice *Slice[T]) Set() *Slice[T] {
	var (
		k  string
		m  = map[string]bool{}
		ok bool
		s  = &Slice[T]{}
	)
	slice.Each(func(_ int, value T) {
		k = fmt.Sprintf("%v", value) // Todo: Check if there is a better way to store key.
		_, ok = m[k]
		if !ok {
			s.Append(value)
		}
		m[k] = true
	})
	(*slice) = (*s)
	return slice
}

// Slice[T any] slices the slice from i to j and returns the modified slice.
func (slice *Slice[T]) Slice(i int, j int) *Slice[T] {
	if j < i {
		i, j = j, i
	}
	if slice.Bounds(i) && slice.Bounds(j) {
		(*slice) = (*slice)[i:j]
	}
	return slice
}

// Swap moves element i to j and j to i.
func (slice *Slice[T]) Swap(i int, j int) {
	if slice.Bounds(i) && slice.Bounds(j) {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

// Unshift adds one or more elements to the beginning of the slice and returns the new length of the modified slice.
func (slice *Slice[T]) Unshift(values ...T) int {
	return (slice.Prepend(values...).Length())
}
