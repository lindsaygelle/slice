package slice

import (
	"sort"
)

// UInt8 is the interface that handles a uint8 collection.
type UInt8 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...uint8) UInt8
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...uint8) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s UInt8) UInt8
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s UInt8) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) UInt8
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, uint8)) UInt8
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, uint8) bool) UInt8
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, uint8)) UInt8
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, uint8) bool) UInt8
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) uint8
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (uint8, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (uint8, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (uint8, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) UInt8
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...uint8) UInt8
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...uint8) UInt8
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, uint8) uint8) UInt8
	// Poll removes the first element from the slice and returns that removed element.
	Poll() uint8
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PollLength() (uint8, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (uint8, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() uint8
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PopLength() (uint8, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	// TODO PopOK() (uint8, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s UInt8) UInt8
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	// TODO PrecatenateLength(s UInt8) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...uint8) UInt8
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	// TODO PrependLength(i ...uint8) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...uint8) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v uint8) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	// TODO Reverse() UInt8
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() UInt8
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) UInt8
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...uint8) int
	// Values returns the internal values of the slice.
	Values() []uint8
}

// NewUInt8 returns a new UInt8 interface.
func NewUInt8(i ...uint8) UInt8 {
	return (&uint8Container{&Slice{}}).Append(i...)
}

type uint8Container struct{ s *Slice }

func (u *uint8Container) Append(i ...uint8) UInt8 {
	u.s.Append(uint8ToInterfaceSlice(i...)...)
	return u
}

func (u *uint8Container) AppendLength(i ...uint8) int {
	return u.Append(i...).Len()
}

func (u *uint8Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *uint8Container) Concatenate(v UInt8) UInt8 {
	u.s.Concatenate(v.(*uint8Container).s)
	return u
}

func (u *uint8Container) ConcatenateLength(v UInt8) int {
	return u.Concatenate(u).Len()
}

func (u *uint8Container) Delete(i int) UInt8 {
	u.s.Delete(i)
	return u
}

func (u *uint8Container) DeleteLength(i int) int {
	return u.s.Delete(i).Len()
}

func (u *uint8Container) Each(fn func(int, uint8)) UInt8 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint8)))
	})
	return u
}

func (u *uint8Container) EachBreak(fn func(int, uint8) bool) UInt8 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint8)))
	})
	return u
}

func (u *uint8Container) EachReverse(fn func(int, uint8)) UInt8 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint8)))
	})
	return u
}

func (u *uint8Container) EachReverseBreak(fn func(int, uint8) bool) UInt8 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint8)))
	})
	return u
}

func (u *uint8Container) Fetch(i int) uint8 {
	var s, _ = u.Get(i)
	return s
}

func (u *uint8Container) FetchLength(i int) (uint8, int) {
	v, i := u.s.FetchLength(i)
	return v.(uint8), i
}

func (u *uint8Container) Get(i int) (uint8, bool) {
	var (
		ok bool
		s  uint8
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint8)
	}
	return s, ok
}

func (u *uint8Container) GetLength(i int) (uint8, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(uint8), l, ok
}

func (u *uint8Container) Len() int {
	return (u.s.Len())
}

func (u *uint8Container) Less(i int, j int) bool {
	return i < j
}

func (u *uint8Container) Make(i int) UInt8 {
	u.s.Make(i)
	return u
}

func (u *uint8Container) MakeEach(v ...uint8) UInt8 {
	u.s.MakeEach(uint8ToInterfaceSlice(v...)...)
	return u
}

func (u *uint8Container) MakeEachReverse(v ...uint8) UInt8 {
	u.s.MakeEachReverse(uint8ToInterfaceSlice(v...)...)
	return u
}

func (u *uint8Container) Map(fn func(int, uint8) uint8) UInt8 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint8)))
	})
	return u
}

func (u *uint8Container) Poll() uint8 {
	var (
		s uint8
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint8))
	}
	return s
}

func (u *uint8Container) PollLength() (uint8, int) {
	v, l := u.s.PollLength()
	return v.(uint8), l
}

func (u *uint8Container) PollOK() (uint8, bool) {
	v, ok := u.s.PollOK()
	return v.(uint8), ok
}

func (u *uint8Container) Pop() uint8 {
	var (
		s uint8
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint8))
	}
	return s
}

func (u *uint8Container) Precatenate(v UInt8) UInt8 {
	u.s.Precatenate(v.(*uint8Container).s)
	return u
}

func (u *uint8Container) Prepend(i ...uint8) UInt8 {
	u.s.Prepend(uint8ToInterfaceSlice(i...)...)
	return u
}

func (u *uint8Container) Push(i ...uint8) int {
	return u.s.Push(uint8ToInterfaceSlice(i...))
}

func (u *uint8Container) Replace(i int, n uint8) bool {
	return (u.s.Replace(i, n))
}

func (u *uint8Container) Set() UInt8 {
	u.s.Set()
	return u
}

func (u *uint8Container) Slice(i int, j int) UInt8 {
	u.s.Slice(i, j)
	return u
}

func (u *uint8Container) Sort() UInt8 {
	sort.Sort(u)
	return u
}

func (u *uint8Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *uint8Container) Unshift(i ...uint8) int {
	return (u.s.Unshift(uint8ToInterfaceSlice(i...)))
}

func (u *uint8Container) Values() []uint8 {
	var v = make([]uint8, u.Len())
	u.Each(func(i int, n uint8) {
		v[i] = n
	})
	return v
}

func uint8ToInterfaceSlice(n ...uint8) []interface{} {
	var (
		i int
		v uint8
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
