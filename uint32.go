package slice

import (
	"sort"
)

// UInt32 is the interface that handles a uint32 collection.
type UInt32 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...uint32) UInt32
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...uint32) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s UInt32) UInt32
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s UInt32) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) UInt32
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, uint32)) UInt32
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, uint32) bool) UInt32
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, uint32)) UInt32
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, uint32) bool) UInt32
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) uint32
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (uint32, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (uint32, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (uint32, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) UInt32
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...uint32) UInt32
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...uint32) UInt32
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, uint32) uint32) UInt32
	// Poll removes the first element from the slice and returns that removed element.
	Poll() uint32
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (uint32, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (uint32, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() uint32
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (uint32, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (uint32, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s UInt32) UInt32
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s UInt32) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...uint32) UInt32
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...uint32) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...uint32) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v uint32) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() UInt32
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() UInt32
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) UInt32
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...uint32) int
	// Values returns the internal values of the slice.
	Values() []uint32
}

// NewUInt32 returns a new UInt32 interface.
func NewUInt32(i ...uint32) UInt32 {
	return (&uint32Container{&Slice{}}).Append(i...)
}

type uint32Container struct{ s *Slice }

// Append implements Append for UInt32.
func (u *uint32Container) Append(i ...uint32) UInt32 {
	u.s.Append(uint32ToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for UInt32.
func (u *uint32Container) AppendLength(i ...uint32) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for UInt32.
func (u *uint32Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for UInt32.
func (u *uint32Container) Concatenate(v UInt32) UInt32 {
	u.s.Concatenate(v.(*uint32Container).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for UInt32.
func (u *uint32Container) ConcatenateLength(v UInt32) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for UInt32.
func (u *uint32Container) Delete(i int) UInt32 {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for UInt32.
func (u *uint32Container) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for UInt32.
func (u *uint32Container) Each(fn func(int, uint32)) UInt32 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint32)))
	})
	return u
}

// EachBreak implements EachBreak for UInt32.
func (u *uint32Container) EachBreak(fn func(int, uint32) bool) UInt32 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint32)))
	})
	return u
}

// EachReverse implements EachReverse for UInt32.
func (u *uint32Container) EachReverse(fn func(int, uint32)) UInt32 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint32)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for UInt32.
func (u *uint32Container) EachReverseBreak(fn func(int, uint32) bool) UInt32 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint32)))
	})
	return u
}

// Fetch implements Fetch for UInt32.
func (u *uint32Container) Fetch(i int) uint32 {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for UInt32.
func (u *uint32Container) FetchLength(i int) (uint32, int) {
	v, i := u.s.FetchLength(i)
	return v.(uint32), i
}

// Get implements Get for UInt32.
func (u *uint32Container) Get(i int) (uint32, bool) {
	var (
		ok bool
		s  uint32
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint32)
	}
	return s, ok
}

// GetLength implements GetLength for UInt32.
func (u *uint32Container) GetLength(i int) (uint32, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(uint32), l, ok
}

// Len implements Len for UInt32.
func (u *uint32Container) Len() int {
	return u.s.Len()
}

// Less implements Less for UInt32.
func (u *uint32Container) Less(i int, j int) bool {
	return u.Fetch(i) < u.Fetch(j)
}

// Make implements Make for UInt32.
func (u *uint32Container) Make(i int) UInt32 {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for UInt32.
func (u *uint32Container) MakeEach(v ...uint32) UInt32 {
	u.s.MakeEach(uint32ToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for UInt32.
func (u *uint32Container) MakeEachReverse(v ...uint32) UInt32 {
	u.s.MakeEachReverse(uint32ToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for UInt32.
func (u *uint32Container) Map(fn func(int, uint32) uint32) UInt32 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint32)))
	})
	return u
}

// Poll implements Poll for UInt32.
func (u *uint32Container) Poll() uint32 {
	var (
		s uint32
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint32))
	}
	return s
}

// PollLength implements PollLength for UInt32.
func (u *uint32Container) PollLength() (uint32, int) {
	v, l := u.s.PollLength()
	return v.(uint32), l
}

// PollOK implements PollOK for UInt32.
func (u *uint32Container) PollOK() (uint32, bool) {
	v, ok := u.s.PollOK()
	return v.(uint32), ok
}

// Pop implements Pop for UInt32.
func (u *uint32Container) Pop() uint32 {
	var (
		s uint32
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint32))
	}
	return s
}

// PopLength implements PopLength for UInt32.
func (u *uint32Container) PopLength() (uint32, int) {
	v, l := u.s.PopLength()
	return v.(uint32), l
}

// PopOK implements PopOK for UInt32.
func (u *uint32Container) PopOK() (uint32, bool) {
	v, ok := u.s.PopOK()
	return v.(uint32), ok
}

// Precatenate implements Precatenate for UInt32.
func (u *uint32Container) Precatenate(v UInt32) UInt32 {
	u.s.Precatenate(v.(*uint32Container).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for UInt32.
func (u *uint32Container) PrecatenateLength(v UInt32) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for UInt32.
func (u *uint32Container) Prepend(i ...uint32) UInt32 {
	u.s.Prepend(uint32ToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for UInt32.
func (u *uint32Container) PrependLength(v ...uint32) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for UInt32.
func (u *uint32Container) Push(i ...uint32) int {
	return u.s.Push(uint32ToInterfaceSlice(i...))
}

// Replace implements Replace for UInt32.
func (u *uint32Container) Replace(i int, n uint32) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for UInt32.
func (u *uint32Container) Reverse() UInt32 {
	u.s.Reverse()
	return u
}

// Set implements Set for UInt32.
func (u *uint32Container) Set() UInt32 {
	u.s.Set()
	return u
}

// Slice implements Slice for UInt32.
func (u *uint32Container) Slice(i int, j int) UInt32 {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for UInt32.
func (u *uint32Container) Sort() UInt32 {
	sort.Sort(u)
	return u
}

// Swap implements Swap for UInt32.
func (u *uint32Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for UInt32.
func (u *uint32Container) Unshift(i ...uint32) int {
	return (u.s.Unshift(uint32ToInterfaceSlice(i...)))
}

// Values implements Values for UInt32.
func (u *uint32Container) Values() []uint32 {
	var v = make([]uint32, u.Len())
	u.Each(func(i int, n uint32) {
		v[i] = n
	})
	return v
}

func uint32ToInterfaceSlice(n ...uint32) []interface{} {
	var (
		i int
		v uint32
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
