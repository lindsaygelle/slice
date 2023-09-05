package slice

import (
	"sort"
)

// Byte is the interface that handles a byte collection.
type Byte interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...byte) Byte
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...byte) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Byte) Byte
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Byte) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Byte
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, byte)) Byte
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, byte) bool) Byte
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, byte)) Byte
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, byte) bool) Byte
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) byte
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (byte, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (byte, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (byte, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Byte
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...byte) Byte
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...byte) Byte
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, byte) byte) Byte
	// Poll removes the first element from the slice and returns that removed element.
	Poll() byte
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (byte, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (byte, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() byte
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (byte, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (byte, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Byte) Byte
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s Byte) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...byte) Byte
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...byte) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...byte) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v byte) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() Byte
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Byte
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Byte
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...byte) int
	// Values returns the internal values of the slice.
	Values() []byte
}

// NewByte returns a new Byte interface.
func NewByte(i ...byte) Byte {
	return (&byteContainer{&Slice{}}).Append(i...)
}

type byteContainer struct{ s *Slice }

// Append implements Append for Byte.
func (u *byteContainer) Append(i ...byte) Byte {
	u.s.Append(byteToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for Byte.
func (u *byteContainer) AppendLength(i ...byte) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for Byte.
func (u *byteContainer) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for Byte.
func (u *byteContainer) Concatenate(v Byte) Byte {
	u.s.Concatenate(v.(*byteContainer).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for Byte.
func (u *byteContainer) ConcatenateLength(v Byte) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for Byte.
func (u *byteContainer) Delete(i int) Byte {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for Byte.
func (u *byteContainer) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for Byte.
func (u *byteContainer) Each(fn func(int, byte)) Byte {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(byte)))
	})
	return u
}

// EachBreak implements EachBreak for Byte.
func (u *byteContainer) EachBreak(fn func(int, byte) bool) Byte {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(byte)))
	})
	return u
}

// EachReverse implements EachReverse for Byte.
func (u *byteContainer) EachReverse(fn func(int, byte)) Byte {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(byte)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for Byte.
func (u *byteContainer) EachReverseBreak(fn func(int, byte) bool) Byte {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(byte)))
	})
	return u
}

// Fetch implements Fetch for Byte.
func (u *byteContainer) Fetch(i int) byte {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for Byte.
func (u *byteContainer) FetchLength(i int) (byte, int) {
	v, i := u.s.FetchLength(i)
	return v.(byte), i
}

// Get implements Get for Byte.
func (u *byteContainer) Get(i int) (byte, bool) {
	var (
		ok bool
		s  byte
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(byte)
	}
	return s, ok
}

// GetLength implements GetLength for Byte.
func (u *byteContainer) GetLength(i int) (byte, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(byte), l, ok
}

// Len implements Len for Byte.
func (u *byteContainer) Len() int {
	return u.s.Len()
}

// Less implements Less for Byte.
func (u *byteContainer) Less(i int, j int) bool {
	return u.Fetch(i) < u.Fetch(j)
}

// Make implements Make for Byte.
func (u *byteContainer) Make(i int) Byte {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for Byte.
func (u *byteContainer) MakeEach(v ...byte) Byte {
	u.s.MakeEach(byteToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for Byte.
func (u *byteContainer) MakeEachReverse(v ...byte) Byte {
	u.s.MakeEachReverse(byteToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for Byte.
func (u *byteContainer) Map(fn func(int, byte) byte) Byte {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(byte)))
	})
	return u
}

// Poll implements Poll for Byte.
func (u *byteContainer) Poll() byte {
	var (
		s byte
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(byte))
	}
	return s
}

// PollLength implements PollLength for Byte.
func (u *byteContainer) PollLength() (byte, int) {
	v, l := u.s.PollLength()
	return v.(byte), l
}

// PollOK implements PollOK for Byte.
func (u *byteContainer) PollOK() (byte, bool) {
	v, ok := u.s.PollOK()
	return v.(byte), ok
}

// Pop implements Pop for Byte.
func (u *byteContainer) Pop() byte {
	var (
		s byte
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(byte))
	}
	return s
}

// PopLength implements PopLength for Byte.
func (u *byteContainer) PopLength() (byte, int) {
	v, l := u.s.PopLength()
	return v.(byte), l
}

// PopOK implements PopOK for Byte.
func (u *byteContainer) PopOK() (byte, bool) {
	v, ok := u.s.PopOK()
	return v.(byte), ok
}

// Precatenate implements Precatenate for Byte.
func (u *byteContainer) Precatenate(v Byte) Byte {
	u.s.Precatenate(v.(*byteContainer).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for Byte.
func (u *byteContainer) PrecatenateLength(v Byte) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for Byte.
func (u *byteContainer) Prepend(i ...byte) Byte {
	u.s.Prepend(byteToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for Byte.
func (u *byteContainer) PrependLength(v ...byte) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for Byte.
func (u *byteContainer) Push(i ...byte) int {
	return u.s.Push(byteToInterfaceSlice(i...))
}

// Replace implements Replace for Byte.
func (u *byteContainer) Replace(i int, n byte) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for Byte.
func (u *byteContainer) Reverse() Byte {
	u.s.Reverse()
	return u
}

// Set implements Set for Byte.
func (u *byteContainer) Set() Byte {
	u.s.Set()
	return u
}

// Slice implements Slice for Byte.
func (u *byteContainer) Slice(i int, j int) Byte {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for Byte.
func (u *byteContainer) Sort() Byte {
	sort.Sort(u)
	return u
}

// Swap implements Swap for Byte.
func (u *byteContainer) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for Byte.
func (u *byteContainer) Unshift(i ...byte) int {
	return (u.s.Unshift(byteToInterfaceSlice(i...)))
}

// Values implements Values for Byte.
func (u *byteContainer) Values() []byte {
	var v = make([]byte, u.Len())
	u.Each(func(i int, n byte) {
		v[i] = n
	})
	return v
}

func byteToInterfaceSlice(n ...byte) []interface{} {
	var (
		i int
		v byte
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
