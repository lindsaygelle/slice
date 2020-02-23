package slice

import (
	"fmt"
)

var _ slicer = (&Slice{})

// slice is the private interface for a Slice.
type slicer interface {
	Append(...interface{}) *Slice
	AppendLength(...interface{}) int
	Bounds(int) bool
	Concatenate(*Slice) *Slice
	ConcatenateLength(*Slice) int
	Delete(int) *Slice
	DeleteLength(int) int
	DeleteOK(int) bool
	Each(func(int, interface{})) *Slice
	EachBreak(func(int, interface{}) bool) *Slice
	EachReverse(func(int, interface{})) *Slice
	EachReverseBreak(func(int, interface{}) bool) *Slice
	Fetch(int) interface{}
	FetchLength(int) (interface{}, int)
	Get(int) (interface{}, bool)
	GetLength(int) (interface{}, int, bool)
	Len() int
	Make(int) *Slice
	MakeEach(...interface{}) *Slice
	MakeEachReverse(...interface{}) *Slice
	Map(func(int, interface{}) interface{}) *Slice
	Poll() interface{}
	PollLength() (interface{}, int)
	PollOK() (interface{}, bool)
	Pop() interface{}
	PopLength() (interface{}, int)
	PopOK() (interface{}, bool)
	Precatenate(*Slice) *Slice
	Prepend(...interface{}) *Slice
	PrependLength(...interface{}) int
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

// Append adds n elements to the end of the slice
// and returns the modified slice.
func (slice *Slice) Append(i ...interface{}) *Slice {
	(*slice) = (append(*slice, i...))
	return slice
}

// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
func (slice *Slice) AppendLength(i ...interface{}) int {
	return (slice.Append(i...).Len())
}

// Bounds checks an integer value safely sits within the range of
// accessible values for the slice.
func (slice *Slice) Bounds(i int) bool {
	return ((i > -1) && (i < len(*slice)))
}

// Concatenate merges the elements from the argument slice
// to the the tail of the argument slice.
func (slice *Slice) Concatenate(s *Slice) *Slice {
	if s != nil {
		slice.Append((*s)...)
	}
	return slice
}

// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
// and returns the length of the receiver slice.
func (slice *Slice) ConcatenateLength(s *Slice) int {
	return (slice.Concatenate(s).Len())
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

// DeleteLength deletes the element from the argument index and returns the new length of the slice.
func (slice *Slice) DeleteLength(i int) int { return slice.Delete(i).Len() }

// DeleteOK deletes the element from the argument index and returns the result of the transaction.
func (slice *Slice) DeleteOK(i int) bool {
	var (
		ok = slice.Bounds(i)
	)
	if ok {
		(*slice) = append((*slice)[:i], (*slice)[i+1:]...)
	}
	return ok
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

// FetchLength retrives the element held at the argument index and the length of the slice.
// Returns the default type if index exceeds slice length.
func (slice *Slice) FetchLength(i int) (interface{}, int) {
	var v = slice.Fetch(i)
	var l = slice.Len()
	return v, l
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

// GetLength returns the element at the argument index, the length of the slice
// and a boolean indicating if the element was successfully retrieved.
func (slice *Slice) GetLength(i int) (interface{}, int, bool) {
	var v, ok = slice.Get(i)
	var l = slice.Len()
	return v, l, ok
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

// PollLength removes the first element from the slice and returns the removed element and the length
// of the modified slice.
func (slice *Slice) PollLength() (interface{}, int) {
	var v = slice.Poll()
	var l = slice.Len()
	return v, l
}

// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
func (slice *Slice) PollOK() (interface{}, bool) {
	var v = slice.Poll()
	return v, (v != nil)
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

// PopLength removes the last element from the slice and returns the removed element and the length
// of the modified slice.
func (slice *Slice) PopLength() (interface{}, int) {
	var v = slice.Pop()
	var l = slice.Len()
	return v, l
}

// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
func (slice *Slice) PopOK() (interface{}, bool) {
	var v = slice.Pop()
	return v, (v != nil)
}

// Precatenate merges the elements from the argument slice
// to the the head of the argument slice and returns the modified slice.
func (slice *Slice) Precatenate(s *Slice) *Slice {
	if s != nil {
		slice.Prepend((*s)...)
	}
	return slice
}

// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
// and returns the length of the receiver slice.
func (slice *Slice) PrecatenateLength(s *Slice) int {
	return (slice.Precatenate(s).Len())
}

// Prepend adds one element to the head of the slice
// and returns the modified slice.
func (slice *Slice) Prepend(i ...interface{}) *Slice {
	(*slice) = (append(i, *slice...))
	return slice
}

// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
func (slice *Slice) PrependLength(i ...interface{}) int {
	return (slice.Prepend(i...).Len())
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
