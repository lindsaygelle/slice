package slice

import (
	"sort"
)

// Complex128 is the interface that handles a complex128 collection.
type Complex128 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...complex128) Complex128
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...complex128) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Complex128) Complex128
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Complex128) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Complex128
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, complex128)) Complex128
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, complex128) bool) Complex128
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, complex128)) Complex128
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, complex128) bool) Complex128
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) complex128
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (complex128, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (complex128, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (complex128, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Complex128
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...complex128) Complex128
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...complex128) Complex128
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, complex128) complex128) Complex128
	// Poll removes the first element from the slice and returns that removed element.
	Poll() complex128
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PollLength() (complex128, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (complex128, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() complex128
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PopLength() (complex128, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	// TODO PopOK() (complex128, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Complex128) Complex128
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	// TODO PrecatenateLength(s Complex128) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...complex128) Complex128
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	// TODO PrependLength(i ...complex128) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...complex128) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v complex128) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	// TODO Reverse() Complex128
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Complex128
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Complex128
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...complex128) int
	// Values returns the internal values of the slice.
	Values() []complex128
}

// NewComplex128 returns a new Complex128 interface.
func NewComplex128(i ...complex128) Complex128 {
	return (&complex128Container{&Slice{}}).Append(i...)
}

type complex128Container struct{ s *Slice }

func (u *complex128Container) Append(i ...complex128) Complex128 {
	u.s.Append(complex128ToInterfaceSlice(i...)...)
	return u
}

func (u *complex128Container) AppendLength(i ...complex128) int {
	return u.Append(i...).Len()
}

func (u *complex128Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *complex128Container) Concatenate(v Complex128) Complex128 {
	u.s.Concatenate(v.(*complex128Container).s)
	return u
}

func (u *complex128Container) ConcatenateLength(v Complex128) int {
	return u.Concatenate(u).Len()
}

func (u *complex128Container) Delete(i int) Complex128 {
	u.s.Delete(i)
	return u
}

func (u *complex128Container) DeleteLength(i int) int {
	return u.s.Delete(i).Len()
}

func (u *complex128Container) Each(fn func(int, complex128)) Complex128 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(complex128)))
	})
	return u
}

func (u *complex128Container) EachBreak(fn func(int, complex128) bool) Complex128 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(complex128)))
	})
	return u
}

func (u *complex128Container) EachReverse(fn func(int, complex128)) Complex128 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(complex128)))
	})
	return u
}

func (u *complex128Container) EachReverseBreak(fn func(int, complex128) bool) Complex128 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(complex128)))
	})
	return u
}

func (u *complex128Container) Fetch(i int) complex128 {
	var s, _ = u.Get(i)
	return s
}

func (u *complex128Container) FetchLength(i int) (complex128, int) {
	v, i := u.s.FetchLength(i)
	return v.(complex128), i
}

func (u *complex128Container) Get(i int) (complex128, bool) {
	var (
		ok bool
		s  complex128
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(complex128)
	}
	return s, ok
}

func (u *complex128Container) GetLength(i int) (complex128, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(complex128), l, ok
}

func (u *complex128Container) Len() int {
	return (u.s.Len())
}

func (u *complex128Container) Less(i int, j int) bool {
	return i < j
}

func (u *complex128Container) Make(i int) Complex128 {
	u.s.Make(i)
	return u
}

func (u *complex128Container) MakeEach(v ...complex128) Complex128 {
	u.s.MakeEach(complex128ToInterfaceSlice(v...)...)
	return u
}

func (u *complex128Container) MakeEachReverse(v ...complex128) Complex128 {
	u.s.MakeEachReverse(complex128ToInterfaceSlice(v...)...)
	return u
}

func (u *complex128Container) Map(fn func(int, complex128) complex128) Complex128 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(complex128)))
	})
	return u
}

func (u *complex128Container) Poll() complex128 {
	var (
		s complex128
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(complex128))
	}
	return s
}

func (u *complex128Container) PollLength() (complex128, int) {
	v, l := u.s.PollLength()
	return v.(complex128), l
}

func (u *complex128Container) PollOK() (complex128, bool) {
	v, ok := u.s.PollOK()
	return v.(complex128), ok
}

func (u *complex128Container) Pop() complex128 {
	var (
		s complex128
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(complex128))
	}
	return s
}

func (u *complex128Container) Precatenate(v Complex128) Complex128 {
	u.s.Precatenate(v.(*complex128Container).s)
	return u
}

func (u *complex128Container) Prepend(i ...complex128) Complex128 {
	u.s.Prepend(complex128ToInterfaceSlice(i...)...)
	return u
}

func (u *complex128Container) Push(i ...complex128) int {
	return u.s.Push(complex128ToInterfaceSlice(i...))
}

func (u *complex128Container) Replace(i int, n complex128) bool {
	return (u.s.Replace(i, n))
}

func (u *complex128Container) Set() Complex128 {
	u.s.Set()
	return u
}

func (u *complex128Container) Slice(i int, j int) Complex128 {
	u.s.Slice(i, j)
	return u
}

func (u *complex128Container) Sort() Complex128 {
	sort.Sort(u)
	return u
}

func (u *complex128Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *complex128Container) Unshift(i ...complex128) int {
	return (u.s.Unshift(complex128ToInterfaceSlice(i...)))
}

func (u *complex128Container) Values() []complex128 {
	var v = make([]complex128, u.Len())
	u.Each(func(i int, n complex128) {
		v[i] = n
	})
	return v
}

func complex128ToInterfaceSlice(n ...complex128) []interface{} {
	var (
		i int
		v complex128
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
