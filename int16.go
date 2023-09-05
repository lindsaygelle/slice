package slice

import (
	"sort"
)

// Int16 is the interface that handles a int16 collection.
type Int16 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...int16) Int16
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...int16) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Int16) Int16
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Int16) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Int16
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, int16)) Int16
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, int16) bool) Int16
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, int16)) Int16
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, int16) bool) Int16
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) int16
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (int16, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (int16, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (int16, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Int16
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...int16) Int16
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...int16) Int16
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, int16) int16) Int16
	// Poll removes the first element from the slice and returns that removed element.
	Poll() int16
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (int16, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (int16, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() int16
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (int16, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (int16, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Int16) Int16
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s Int16) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...int16) Int16
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...int16) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...int16) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v int16) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() Int16
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Int16
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Int16
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...int16) int
	// Values returns the internal values of the slice.
	Values() []int16
}

// NewInt16 returns a new Int16 interface.
func NewInt16(i ...int16) Int16 {
	return (&int16Container{&Slice{}}).Append(i...)
}

type int16Container struct{ s *Slice }

// Append implements Append for Int16.
func (u *int16Container) Append(i ...int16) Int16 {
	u.s.Append(int16ToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for Int16.
func (u *int16Container) AppendLength(i ...int16) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for Int16.
func (u *int16Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for Int16.
func (u *int16Container) Concatenate(v Int16) Int16 {
	u.s.Concatenate(v.(*int16Container).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for Int16.
func (u *int16Container) ConcatenateLength(v Int16) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for Int16.
func (u *int16Container) Delete(i int) Int16 {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for Int16.
func (u *int16Container) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for Int16.
func (u *int16Container) Each(fn func(int, int16)) Int16 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int16)))
	})
	return u
}

// EachBreak implements EachBreak for Int16.
func (u *int16Container) EachBreak(fn func(int, int16) bool) Int16 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int16)))
	})
	return u
}

// EachReverse implements EachReverse for Int16.
func (u *int16Container) EachReverse(fn func(int, int16)) Int16 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int16)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for Int16.
func (u *int16Container) EachReverseBreak(fn func(int, int16) bool) Int16 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int16)))
	})
	return u
}

// Fetch implements Fetch for Int16.
func (u *int16Container) Fetch(i int) int16 {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for Int16.
func (u *int16Container) FetchLength(i int) (int16, int) {
	v, i := u.s.FetchLength(i)
	return v.(int16), i
}

// Get implements Get for Int16.
func (u *int16Container) Get(i int) (int16, bool) {
	var (
		ok bool
		s  int16
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(int16)
	}
	return s, ok
}

// GetLength implements GetLength for Int16.
func (u *int16Container) GetLength(i int) (int16, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(int16), l, ok
}

// Len implements Len for Int16.
func (u *int16Container) Len() int {
	return u.s.Len()
}

// Less implements Less for Int16.
func (u *int16Container) Less(i int, j int) bool {
	return u.Fetch(i) < u.Fetch(j)
}

// Make implements Make for Int16.
func (u *int16Container) Make(i int) Int16 {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for Int16.
func (u *int16Container) MakeEach(v ...int16) Int16 {
	u.s.MakeEach(int16ToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for Int16.
func (u *int16Container) MakeEachReverse(v ...int16) Int16 {
	u.s.MakeEachReverse(int16ToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for Int16.
func (u *int16Container) Map(fn func(int, int16) int16) Int16 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int16)))
	})
	return u
}

// Poll implements Poll for Int16.
func (u *int16Container) Poll() int16 {
	var (
		s int16
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(int16))
	}
	return s
}

// PollLength implements PollLength for Int16.
func (u *int16Container) PollLength() (int16, int) {
	v, l := u.s.PollLength()
	return v.(int16), l
}

// PollOK implements PollOK for Int16.
func (u *int16Container) PollOK() (int16, bool) {
	v, ok := u.s.PollOK()
	return v.(int16), ok
}

// Pop implements Pop for Int16.
func (u *int16Container) Pop() int16 {
	var (
		s int16
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(int16))
	}
	return s
}

// PopLength implements PopLength for Int16.
func (u *int16Container) PopLength() (int16, int) {
	v, l := u.s.PopLength()
	return v.(int16), l
}

// PopOK implements PopOK for Int16.
func (u *int16Container) PopOK() (int16, bool) {
	v, ok := u.s.PopOK()
	return v.(int16), ok
}

// Precatenate implements Precatenate for Int16.
func (u *int16Container) Precatenate(v Int16) Int16 {
	u.s.Precatenate(v.(*int16Container).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for Int16.
func (u *int16Container) PrecatenateLength(v Int16) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for Int16.
func (u *int16Container) Prepend(i ...int16) Int16 {
	u.s.Prepend(int16ToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for Int16.
func (u *int16Container) PrependLength(v ...int16) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for Int16.
func (u *int16Container) Push(i ...int16) int {
	return u.s.Push(int16ToInterfaceSlice(i...))
}

// Replace implements Replace for Int16.
func (u *int16Container) Replace(i int, n int16) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for Int16.
func (u *int16Container) Reverse() Int16 {
	u.s.Reverse()
	return u
}

// Set implements Set for Int16.
func (u *int16Container) Set() Int16 {
	u.s.Set()
	return u
}

// Slice implements Slice for Int16.
func (u *int16Container) Slice(i int, j int) Int16 {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for Int16.
func (u *int16Container) Sort() Int16 {
	sort.Sort(u)
	return u
}

// Swap implements Swap for Int16.
func (u *int16Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for Int16.
func (u *int16Container) Unshift(i ...int16) int {
	return (u.s.Unshift(int16ToInterfaceSlice(i...)))
}

// Values implements Values for Int16.
func (u *int16Container) Values() []int16 {
	var v = make([]int16, u.Len())
	u.Each(func(i int, n int16) {
		v[i] = n
	})
	return v
}

func int16ToInterfaceSlice(n ...int16) []interface{} {
	var (
		i int
		v int16
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
