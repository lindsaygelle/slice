package slice

import (
	"fmt"
)

var _ slicer = (&Slice{})

type slicer interface {
	Append(...interface{}) *Slice
	Bounds(int) bool
	Concatenate(*Slice) *Slice
	Each(func(int, interface{})) *Slice
	EachBreak(func(int, interface{}) bool) *Slice
	EachReverse(func(int, interface{})) *Slice
	EachReverseBreak(func(int, interface{}) bool) *Slice
	Fetch(int) interface{}
	Get(int) (interface{}, bool)
	Len() int
	Map(func(int, interface{}) interface{}) *Slice
	Poll() interface{}
	Pop() interface{}
	Precatenate(*Slice) *Slice
	Prepend(...interface{}) *Slice
	Push(...interface{}) int
	Replace(int, interface{}) bool
	Set() *Slice
	Swap(int, int)
	Unshift(...interface{}) int
	Values() []interface{}
}

// Slice is an implementation of a *[]interface{}.
//
// Slice has methods to perform traversal and mutation operations.
// A Slice can accept any interface{} but does not implement a sort proceedure.
//
// To extend a Slice construct a struct and a supporting interface that implements the Slice methods.
type Slice []interface{}

// Append adds one element to the end of the collection
// and returns the modified collection.
func (slice *Slice) Append(i ...interface{}) *Slice {
	(*slice) = (append(*slice, i...))
	return slice
}

// Bounds checks an integer value safely sits within the range of
// accessible values for the collection.
func (slice *Slice) Bounds(i int) bool {
	return ((i > -1) && (i < len(*slice)))
}

// Concatenate merges the elements from the argument collection
// to the the tail of the argument collection.
func (slice *Slice) Concatenate(s *Slice) *Slice {
	slice.Append((*s)...)
	return slice
}

// Each executes a provided function once for each collection element.
func (slice *Slice) Each(fn func(int, interface{})) *Slice {
	var (
		i int
		v interface{}
	)
	for i, v = range *slice {
		fn(i, v)
	}
	return slice
}

// EachBreak executes a provided function once for each
// element with an optional break when the function returns false.
func (slice *Slice) EachBreak(fn func(int, interface{}) bool) *Slice {
	var (
		i  int
		ok = true
		v  interface{}
	)
	for i, v = range *slice {
		ok = fn(i, v)
		if !ok {
			break
		}
	}
	return slice
}

// EachReverse executes a provided function once for each
// element in the reverse order they are stored in the *Slice.
func (slice *Slice) EachReverse(fn func(int, interface{})) *Slice {
	var (
		i int
	)
	for i = len(*slice) - 1; i >= 0; i-- {
		fn(i, (*slice)[i])
	}
	return slice
}

// EachReverseBreak executes a provided function once for each
// element in the reverse order they are stored in the collection
// with an optional break when the function returns false.
func (slice *Slice) EachReverseBreak(fn func(int, interface{}) bool) *Slice {
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

// Fetch retrieves the element held at the argument index.
// Returns the default type if index exceeds collection length.
func (slice *Slice) Fetch(i int) interface{} {
	var v, _ = slice.Get(i)
	return v
}

// Get returns the element held at the argument index and a boolean
// indicating if it was successfully retrieved.
func (slice *Slice) Get(i int) (interface{}, bool) {
	var (
		ok = slice.Bounds(i)
	)
	if ok {
		return (*slice)[i], ok
	}
	return nil, ok
}

// Len returns the number of elements in the collection.
func (slice *Slice) Len() int { return (len(*slice)) }

// Map executes a provided function once for each element and sets
// the returned value to the current index.
func (slice *Slice) Map(fn func(int, interface{}) interface{}) *Slice {
	slice.Each(func(i int, v interface{}) {
		slice.Replace(i, fn(i, v))
	})
	return slice
}

// Precatenate merges the elements from the argument collection
// to the the head of the argument collection.
func (slice *Slice) Precatenate(s *Slice) *Slice {
	slice.Prepend((*s)...)
	return slice
}

// Poll removes the first element from the collection and returns that removed element.
func (slice *Slice) Poll() interface{} {
	var (
		l  = slice.Len()
		ok = l > 0
		v  interface{}
	)
	if ok {
		v = (*slice)[0]
		(*slice) = (*slice)[1:]
	}
	return v
}

// Pop removes the last element from the collection and returns that element.
func (slice *Slice) Pop() interface{} {
	var (
		l  = slice.Len()
		ok = l > 0
		v  interface{}
	)
	if ok {
		v = (*slice)[l-1]
		(*slice) = (*slice)[:l-1]
	}
	return v
}

// Prepend adds one element to the head of the collection
// and returns the modified collection.
func (slice *Slice) Prepend(i ...interface{}) *Slice {
	(*slice) = (append(i, *slice...))
	return slice
}

// Push adds a new element to the end of the collection and
// returns the length of the modified collection.
func (slice *Slice) Push(i ...interface{}) int {
	return (slice.Append(i...).Len())
}

// Replace changes the contents of the collection
// at the argument index if it is in bounds.
func (slice *Slice) Replace(i int, v interface{}) bool {
	var (
		ok = slice.Bounds(i)
	)
	if ok {
		(*slice)[i] = v
	}
	return ok
}

// Set returns a unique collection, removing duplicate
// elements that have the same hash value.
func (slice *Slice) Set() *Slice {
	const (
		f string = "%v"
	)
	var (
		k  string
		m  = map[string]bool{}
		ok bool
		s  = &Slice{}
	)
	slice.Each(func(_ int, v interface{}) {
		k = fmt.Sprintf(f, v)
		_, ok = m[k]
		if !ok {
			s.Append(v)
		}
		m[k] = true
	})
	(*slice) = (*s)
	return slice
}

// Slice slices the collection from i to j and returns the modified collection.
func (slice *Slice) Slice(i int, j int) *Slice {
	if j > i {
		i, j = j, i
	}
	if slice.Bounds(i) && slice.Bounds(j) {
		(*slice) = (*slice)[i:j]
	}
	return slice
}

// Swap moves element i to j and j to i.
func (slice *Slice) Swap(i int, j int) {
	(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
}

// Unshift adds one or more elements to the beginning of the collection and
// returns the new length of the modified collection.
func (slice *Slice) Unshift(i ...interface{}) int {
	return (slice.Prepend(i...).Len())
}

// Values returns the internal values of the collection.
func (slice *Slice) Values() []interface{} {
	return (*slice)
}
