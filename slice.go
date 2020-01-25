package slice

import (
	"fmt"
)

var _ slicer = (&Slice{})

// slice is the private interface for a Slice.
type slicer interface {
	Append(...interface{}) *Slice
	Bounds(int) bool
	Concatenate(*Slice) *Slice
	Delete(int) *Slice
	Each(func(int, interface{})) *Slice
	EachBreak(func(int, interface{}) bool) *Slice
	EachReverse(func(int, interface{})) *Slice
	EachReverseBreak(func(int, interface{}) bool) *Slice
	Fetch(int) interface{}
	Get(int) (interface{}, bool)
	Len() int
	Make(int) *Slice
	MakeEach(...interface{}) *Slice
	MakeEachReverse(...interface{}) *Slice
	Map(func(int, interface{}) interface{}) *Slice
	Poll() interface{}
	Pop() interface{}
	Precatenate(*Slice) *Slice
	Prepend(...interface{}) *Slice
	Push(...interface{}) int
	Replace(int, interface{}) bool
	Reverse() *Slice
	Set() *Slice
	Slice(int, int) *Slice
	Swap(int, int)
	Unshift(...interface{}) int
	Values() []interface{}
}

// NewSlice returns a new Slice.
func NewSlice(v ...interface{}) *Slice {
	return (&Slice{}).MakeEach(v...)
}

// Slice is a list-like struct whose methods are used to perform traversal and mutation operations by numeric index.
//
// The Slice does not use a fixed address size and will dynamically allocate and deallocate space as new entries are pushed into the sequence.
//
// Slice is written to handle a mix content type. By default the Slice assumes that the data returned is something or nil.
// To handle returning a element from the Slice as a non-interface type it is best to create a custom
// interface or struct to handle the type-casting.
//
// To implement the Slice as single type, create a struct that contains a Slice pointer as a hidden field and
// compose an interface that exports the Slice's methods, using the wrapping struct to
// handle the transaction between the struct and the Slice.
type Slice []interface{}

// Append adds one element to the end of the slice
// and returns the modified slice.
func (slice *Slice) Append(i ...interface{}) *Slice {
	(*slice) = (append(*slice, i...))
	return slice
}

// Bounds checks an integer value safely sits within the range of
// accessible values for the slice.
func (slice *Slice) Bounds(i int) bool {
	return ((i > -1) && (i < len(*slice)))
}

// Concatenate merges the elements from the argument slice
// to the the tail of the argument slice.
func (slice *Slice) Concatenate(s *Slice) *Slice {
	slice.Append((*s)...)
	return slice
}

// Delete deletes the element from the argument index and returns the modified slice.
func (slice *Slice) Delete(i int) *Slice {
	var (
		ok = slice.Bounds(i)
	)
	if ok {
		(*slice) = append((*slice)[:i], (*slice)[i+1:]...)
	}
	return slice
}

// Each executes a provided function once for each slice element and returns the slice.
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
// Returns the slice at the end of the iteration.
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
// element in the reverse order they are stored in the slice.
// Returns the slice at the end of the iteration.
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
// element in the reverse order they are stored in the slice
// with an optional break when the function returns false.
// Returns the slice at the end of the iteration.
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
// Returns the default type if index exceeds slice length.
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

// Len returns the number of elements in the slice.
func (slice *Slice) Len() int { return (len(*slice)) }

// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
func (slice *Slice) Make(i int) *Slice {
	(*slice) = make(Slice, i)
	return slice
}

// MakeEach empties the slice, sets the new slice to the length of n and performs
// a for-each loop for the argument sequence, inserting each entry at the
// appropriate index before returning the modified slice.
func (slice *Slice) MakeEach(v ...interface{}) *Slice {
	return slice.Make(len(v)).Each(func(i int, _ interface{}) {
		slice.Replace(i, v[i])
	})
}

// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
// an inverse for-each loop for the argument sequence, inserting each entry at the
// appropriate index before returning the modified slice.
func (slice *Slice) MakeEachReverse(v ...interface{}) *Slice {
	return slice.Make(len(v)).EachReverse(func(i int, _ interface{}) {
		slice.Replace(i, v[i])
	})
}

// Map executes a provided function once for each element and sets
// the returned value to the current index.
// Returns the slice at the end of the iteration.
func (slice *Slice) Map(fn func(int, interface{}) interface{}) *Slice {
	slice.Each(func(i int, v interface{}) {
		slice.Replace(i, fn(i, v))
	})
	return slice
}

// Precatenate merges the elements from the argument slice
// to the the head of the argument slice and returns the modified slice.
func (slice *Slice) Precatenate(s *Slice) *Slice {
	slice.Prepend((*s)...)
	return slice
}

// Poll removes the first element from the slice and returns that removed element.
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

// Pop removes the last element from the slice and returns that element.
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

// Prepend adds one element to the head of the slice
// and returns the modified slice.
func (slice *Slice) Prepend(i ...interface{}) *Slice {
	(*slice) = (append(i, *slice...))
	return slice
}

// Push adds a new element to the end of the slice and
// returns the length of the modified slice.
func (slice *Slice) Push(i ...interface{}) int {
	return (slice.Append(i...).Len())
}

// Replace changes the contents of the slice
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

// Reverse reverses the slice in linear time.
// Returns the slice at the end of the iteration.
func (slice *Slice) Reverse() *Slice {
	var (
		i = 0
		j = slice.Len()
	)
	for i < j {
		slice.Swap(i, j)
		i = i + 1
		j = j - 1
	}
	return slice
}

// Set returns a unique slice, removing duplicate
// elements that have the same hash value.
// Returns the modified at the end of the iteration.
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

// Slice slices the slice from i to j and returns the modified slice.
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

// Unshift adds one or more elements to the beginning of the slice and
// returns the new length of the modified slice.
func (slice *Slice) Unshift(i ...interface{}) int {
	return (slice.Prepend(i...).Len())
}

// Values returns the internal values of the slice.
func (slice *Slice) Values() []interface{} {
	return (*slice)
}
