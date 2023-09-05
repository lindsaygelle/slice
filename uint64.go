package slice

import (
	"sort"
)

// UInt64 is the interface that handles a uint64 collection.
type UInt64 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...uint64) UInt64
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...uint64) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s UInt64) UInt64
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s UInt64) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) UInt64
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, uint64)) UInt64
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, uint64) bool) UInt64
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, uint64)) UInt64
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, uint64) bool) UInt64
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) uint64
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (uint64, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (uint64, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (uint64, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) UInt64
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...uint64) UInt64
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...uint64) UInt64
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, uint64) uint64) UInt64
	// Poll removes the first element from the slice and returns that removed element.
	Poll() uint64
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PollLength() (uint64, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (uint64, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() uint64
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PopLength() (uint64, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	// TODO PopOK() (uint64, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s UInt64) UInt64
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	// TODO PrecatenateLength(s UInt64) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...uint64) UInt64
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	// TODO PrependLength(i ...uint64) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...uint64) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v uint64) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	// TODO Reverse() UInt64
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() UInt64
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) UInt64
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...uint64) int
	// Values returns the internal values of the slice.
	Values() []uint64
}

// NewUInt64 returns a new UInt64 interface.
func NewUInt64(i ...uint64) UInt64 {
	return (&uint64Container{&Slice{}}).Append(i...)
}

type uint64Container struct{ s *Slice }

func (u *uint64Container) Append(i ...uint64) UInt64 {
	u.s.Append(uint64ToInterfaceSlice(i...)...)
	return u
}

func (u *uint64Container) AppendLength(i ...uint64) int {
	return u.Append(i...).Len()
}

func (u *uint64Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *uint64Container) Concatenate(v UInt64) UInt64 {
	u.s.Concatenate(v.(*uint64Container).s)
	return u
}

func (u *uint64Container) ConcatenateLength(v UInt64) int {
	return u.Concatenate(u).Len()
}

func (u *uint64Container) Delete(i int) UInt64 {
	u.s.Delete(i)
	return u
}

func (u *uint64Container) DeleteLength(i int) int {
	return u.s.Delete(i).Len()
}

func (u *uint64Container) Each(fn func(int, uint64)) UInt64 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint64)))
	})
	return u
}

func (u *uint64Container) EachBreak(fn func(int, uint64) bool) UInt64 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint64)))
	})
	return u
}

func (u *uint64Container) EachReverse(fn func(int, uint64)) UInt64 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint64)))
	})
	return u
}

func (u *uint64Container) EachReverseBreak(fn func(int, uint64) bool) UInt64 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint64)))
	})
	return u
}

func (u *uint64Container) Fetch(i int) uint64 {
	var s, _ = u.Get(i)
	return s
}

func (u *uint64Container) FetchLength(i int) (uint64, int) {
	v, i := u.s.FetchLength(i)
	return v.(uint64), i
}

func (u *uint64Container) Get(i int) (uint64, bool) {
	var (
		ok bool
		s  uint64
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint64)
	}
	return s, ok
}

func (u *uint64Container) GetLength(i int) (uint64, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(uint64), l, ok
}

func (u *uint64Container) Len() int {
	return (u.s.Len())
}

func (u *uint64Container) Less(i int, j int) bool {
	return i < j
}

func (u *uint64Container) Make(i int) UInt64 {
	u.s.Make(i)
	return u
}

func (u *uint64Container) MakeEach(v ...uint64) UInt64 {
	u.s.MakeEach(uint64ToInterfaceSlice(v...)...)
	return u
}

func (u *uint64Container) MakeEachReverse(v ...uint64) UInt64 {
	u.s.MakeEachReverse(uint64ToInterfaceSlice(v...)...)
	return u
}

func (u *uint64Container) Map(fn func(int, uint64) uint64) UInt64 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint64)))
	})
	return u
}

func (u *uint64Container) Poll() uint64 {
	var (
		s uint64
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint64))
	}
	return s
}

func (u *uint64Container) PollLength() (uint64, int) {
	v, l := u.s.PollLength()
	return v.(uint64), l
}

func (u *uint64Container) PollOK() (uint64, bool) {
	v, ok := u.s.PollOK()
	return v.(uint64), ok
}

func (u *uint64Container) Pop() uint64 {
	var (
		s uint64
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint64))
	}
	return s
}

func (u *uint64Container) Precatenate(v UInt64) UInt64 {
	u.s.Precatenate(v.(*uint64Container).s)
	return u
}

func (u *uint64Container) Prepend(i ...uint64) UInt64 {
	u.s.Prepend(uint64ToInterfaceSlice(i...)...)
	return u
}

func (u *uint64Container) Push(i ...uint64) int {
	return u.s.Push(uint64ToInterfaceSlice(i...))
}

func (u *uint64Container) Replace(i int, n uint64) bool {
	return (u.s.Replace(i, n))
}

func (u *uint64Container) Set() UInt64 {
	u.s.Set()
	return u
}

func (u *uint64Container) Slice(i int, j int) UInt64 {
	u.s.Slice(i, j)
	return u
}

func (u *uint64Container) Sort() UInt64 {
	sort.Sort(u)
	return u
}

func (u *uint64Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *uint64Container) Unshift(i ...uint64) int {
	return (u.s.Unshift(uint64ToInterfaceSlice(i...)))
}

func (u *uint64Container) Values() []uint64 {
	var v = make([]uint64, u.Len())
	u.Each(func(i int, n uint64) {
		v[i] = n
	})
	return v
}

func uint64ToInterfaceSlice(n ...uint64) []interface{} {
	var (
		i int
		v uint64
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
