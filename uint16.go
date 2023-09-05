package slice

import (
	"sort"
)

// UInt16 is the interface that handles a uint16 collection.
type UInt16 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...uint16) UInt16
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...uint16) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s UInt16) UInt16
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s UInt16) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) UInt16
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, uint16)) UInt16
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, uint16) bool) UInt16
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, uint16)) UInt16
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, uint16) bool) UInt16
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) uint16
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (uint16, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (uint16, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (uint16, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) UInt16
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...uint16) UInt16
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...uint16) UInt16
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, uint16) uint16) UInt16
	// Poll removes the first element from the slice and returns that removed element.
	Poll() uint16
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (uint16, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (uint16, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() uint16
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (uint16, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (uint16, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s UInt16) UInt16
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s UInt16) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...uint16) UInt16
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...uint16) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...uint16) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v uint16) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() UInt16
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() UInt16
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) UInt16
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...uint16) int
	// Values returns the internal values of the slice.
	Values() []uint16
}

// NewUInt16 returns a new UInt16 interface.
func NewUInt16(i ...uint16) UInt16 {
	return (&uint16Container{&Slice{}}).Append(i...)
}

type uint16Container struct{ s *Slice }

// Append implements Append for UInt16.
func (u *uint16Container) Append(i ...uint16) UInt16 {
	u.s.Append(uint16ToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for UInt16.
func (u *uint16Container) AppendLength(i ...uint16) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for UInt16.
func (u *uint16Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for UInt16.
func (u *uint16Container) Concatenate(v UInt16) UInt16 {
	u.s.Concatenate(v.(*uint16Container).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for UInt16.
func (u *uint16Container) ConcatenateLength(v UInt16) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for UInt16.
func (u *uint16Container) Delete(i int) UInt16 {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for UInt16.
func (u *uint16Container) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for UInt16.
func (u *uint16Container) Each(fn func(int, uint16)) UInt16 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint16)))
	})
	return u
}

// EachBreak implements EachBreak for UInt16.
func (u *uint16Container) EachBreak(fn func(int, uint16) bool) UInt16 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint16)))
	})
	return u
}

// EachReverse implements EachReverse for UInt16.
func (u *uint16Container) EachReverse(fn func(int, uint16)) UInt16 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint16)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for UInt16.
func (u *uint16Container) EachReverseBreak(fn func(int, uint16) bool) UInt16 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint16)))
	})
	return u
}

// Fetch implements Fetch for UInt16.
func (u *uint16Container) Fetch(i int) uint16 {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for UInt16.
func (u *uint16Container) FetchLength(i int) (uint16, int) {
	v, i := u.s.FetchLength(i)
	return v.(uint16), i
}

// Get implements Get for UInt16.
func (u *uint16Container) Get(i int) (uint16, bool) {
	var (
		ok bool
		s  uint16
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint16)
	}
	return s, ok
}

// GetLength implements GetLength for UInt16.
func (u *uint16Container) GetLength(i int) (uint16, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(uint16), l, ok
}

// Len implements Len for UInt16.
func (u *uint16Container) Len() int {
	return u.s.Len()
}

// Less implements Less for UInt16.
func (u *uint16Container) Less(i int, j int) bool {
	return u.Fetch(i) < u.Fetch(j)
}

// Make implements Make for UInt16.
func (u *uint16Container) Make(i int) UInt16 {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for UInt16.
func (u *uint16Container) MakeEach(v ...uint16) UInt16 {
	u.s.MakeEach(uint16ToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for UInt16.
func (u *uint16Container) MakeEachReverse(v ...uint16) UInt16 {
	u.s.MakeEachReverse(uint16ToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for UInt16.
func (u *uint16Container) Map(fn func(int, uint16) uint16) UInt16 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint16)))
	})
	return u
}

// Poll implements Poll for UInt16.
func (u *uint16Container) Poll() uint16 {
	var (
		s uint16
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint16))
	}
	return s
}

// PollLength implements PollLength for UInt16.
func (u *uint16Container) PollLength() (uint16, int) {
	v, l := u.s.PollLength()
	return v.(uint16), l
}

// PollOK implements PollOK for UInt16.
func (u *uint16Container) PollOK() (uint16, bool) {
	v, ok := u.s.PollOK()
	return v.(uint16), ok
}

// Pop implements Pop for UInt16.
func (u *uint16Container) Pop() uint16 {
	var (
		s uint16
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint16))
	}
	return s
}

// PopLength implements PopLength for UInt16.
func (u *uint16Container) PopLength() (uint16, int) {
	v, l := u.s.PopLength()
	return v.(uint16), l
}

// PopOK implements PopOK for UInt16.
func (u *uint16Container) PopOK() (uint16, bool) {
	v, ok := u.s.PopOK()
	return v.(uint16), ok
}

// Precatenate implements Precatenate for UInt16.
func (u *uint16Container) Precatenate(v UInt16) UInt16 {
	u.s.Precatenate(v.(*uint16Container).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for UInt16.
func (u *uint16Container) PrecatenateLength(v UInt16) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for UInt16.
func (u *uint16Container) Prepend(i ...uint16) UInt16 {
	u.s.Prepend(uint16ToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for UInt16.
func (u *uint16Container) PrependLength(v ...uint16) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for UInt16.
func (u *uint16Container) Push(i ...uint16) int {
	return u.s.Push(uint16ToInterfaceSlice(i...))
}

// Replace implements Replace for UInt16.
func (u *uint16Container) Replace(i int, n uint16) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for UInt16.
func (u *uint16Container) Reverse() UInt16 {
	u.s.Reverse()
	return u
}

// Set implements Set for UInt16.
func (u *uint16Container) Set() UInt16 {
	u.s.Set()
	return u
}

// Slice implements Slice for UInt16.
func (u *uint16Container) Slice(i int, j int) UInt16 {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for UInt16.
func (u *uint16Container) Sort() UInt16 {
	sort.Sort(u)
	return u
}

// Swap implements Swap for UInt16.
func (u *uint16Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for UInt16.
func (u *uint16Container) Unshift(i ...uint16) int {
	return (u.s.Unshift(uint16ToInterfaceSlice(i...)))
}

// Values implements Values for UInt16.
func (u *uint16Container) Values() []uint16 {
	var v = make([]uint16, u.Len())
	u.Each(func(i int, n uint16) {
		v[i] = n
	})
	return v
}

func uint16ToInterfaceSlice(n ...uint16) []interface{} {
	var (
		i int
		v uint16
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
