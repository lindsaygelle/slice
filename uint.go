package slice

import (
	"sort"
)

// UInt is the interface that handles a uint collection.
type UInt interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...uint) UInt
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...uint) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s UInt) UInt
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s UInt) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) UInt
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, uint)) UInt
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, uint) bool) UInt
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, uint)) UInt
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, uint) bool) UInt
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) uint
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (uint, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (uint, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (uint, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) UInt
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...uint) UInt
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...uint) UInt
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, uint) uint) UInt
	// Poll removes the first element from the slice and returns that removed element.
	Poll() uint
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (uint, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (uint, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() uint
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (uint, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (uint, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s UInt) UInt
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s UInt) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...uint) UInt
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...uint) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...uint) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v uint) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() UInt
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() UInt
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) UInt
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...uint) int
	// Values returns the internal values of the slice.
	Values() []uint
}

// NewUInt returns a new UInt interface.
func NewUInt(i ...uint) UInt {
	return (&uintContainer{&Slice{}}).Append(i...)
}

type uintContainer struct{ s *Slice }

// Append implements Append for UInt.
func (u *uintContainer) Append(i ...uint) UInt {
	u.s.Append(uintToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for UInt.
func (u *uintContainer) AppendLength(i ...uint) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for UInt.
func (u *uintContainer) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for UInt.
func (u *uintContainer) Concatenate(v UInt) UInt {
	u.s.Concatenate(v.(*uintContainer).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for UInt.
func (u *uintContainer) ConcatenateLength(v UInt) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for UInt.
func (u *uintContainer) Delete(i int) UInt {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for UInt.
func (u *uintContainer) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for UInt.
func (u *uintContainer) Each(fn func(int, uint)) UInt {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint)))
	})
	return u
}

// EachBreak implements EachBreak for UInt.
func (u *uintContainer) EachBreak(fn func(int, uint) bool) UInt {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint)))
	})
	return u
}

// EachReverse implements EachReverse for UInt.
func (u *uintContainer) EachReverse(fn func(int, uint)) UInt {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for UInt.
func (u *uintContainer) EachReverseBreak(fn func(int, uint) bool) UInt {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint)))
	})
	return u
}

// Fetch implements Fetch for UInt.
func (u *uintContainer) Fetch(i int) uint {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for UInt.
func (u *uintContainer) FetchLength(i int) (uint, int) {
	v, i := u.s.FetchLength(i)
	return v.(uint), i
}

// Get implements Get for UInt.
func (u *uintContainer) Get(i int) (uint, bool) {
	var (
		ok bool
		s  uint
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint)
	}
	return s, ok
}

// GetLength implements GetLength for UInt.
func (u *uintContainer) GetLength(i int) (uint, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(uint), l, ok
}

// Len implements Len for UInt.
func (u *uintContainer) Len() int {
	return u.s.Len()
}

// Less implements Less for UInt.
func (u *uintContainer) Less(i int, j int) bool {
	return u.Fetch(i) < u.Fetch(j)
}

// Make implements Make for UInt.
func (u *uintContainer) Make(i int) UInt {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for UInt.
func (u *uintContainer) MakeEach(v ...uint) UInt {
	u.s.MakeEach(uintToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for UInt.
func (u *uintContainer) MakeEachReverse(v ...uint) UInt {
	u.s.MakeEachReverse(uintToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for UInt.
func (u *uintContainer) Map(fn func(int, uint) uint) UInt {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint)))
	})
	return u
}

// Poll implements Poll for UInt.
func (u *uintContainer) Poll() uint {
	var (
		s uint
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint))
	}
	return s
}

// PollLength implements PollLength for UInt.
func (u *uintContainer) PollLength() (uint, int) {
	v, l := u.s.PollLength()
	return v.(uint), l
}

// PollOK implements PollOK for UInt.
func (u *uintContainer) PollOK() (uint, bool) {
	v, ok := u.s.PollOK()
	return v.(uint), ok
}

// Pop implements Pop for UInt.
func (u *uintContainer) Pop() uint {
	var (
		s uint
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint))
	}
	return s
}

// PopLength implements PopLength for UInt.
func (u *uintContainer) PopLength() (uint, int) {
	v, l := u.s.PopLength()
	return v.(uint), l
}

// PopOK implements PopOK for UInt.
func (u *uintContainer) PopOK() (uint, bool) {
	v, ok := u.s.PopOK()
	return v.(uint), ok
}

// Precatenate implements Precatenate for UInt.
func (u *uintContainer) Precatenate(v UInt) UInt {
	u.s.Precatenate(v.(*uintContainer).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for UInt.
func (u *uintContainer) PrecatenateLength(v UInt) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for UInt.
func (u *uintContainer) Prepend(i ...uint) UInt {
	u.s.Prepend(uintToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for UInt.
func (u *uintContainer) PrependLength(v ...uint) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for UInt.
func (u *uintContainer) Push(i ...uint) int {
	return u.s.Push(uintToInterfaceSlice(i...))
}

// Replace implements Replace for UInt.
func (u *uintContainer) Replace(i int, n uint) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for UInt.
func (u *uintContainer) Reverse() UInt {
	u.s.Reverse()
	return u
}

// Set implements Set for UInt.
func (u *uintContainer) Set() UInt {
	u.s.Set()
	return u
}

// Slice implements Slice for UInt.
func (u *uintContainer) Slice(i int, j int) UInt {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for UInt.
func (u *uintContainer) Sort() UInt {
	sort.Sort(u)
	return u
}

// Swap implements Swap for UInt.
func (u *uintContainer) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for UInt.
func (u *uintContainer) Unshift(i ...uint) int {
	return (u.s.Unshift(uintToInterfaceSlice(i...)))
}

// Values implements Values for UInt.
func (u *uintContainer) Values() []uint {
	var v = make([]uint, u.Len())
	u.Each(func(i int, n uint) {
		v[i] = n
	})
	return v
}

func uintToInterfaceSlice(n ...uint) []interface{} {
	var (
		i int
		v uint
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
