package slice

import (
	"sort"
)

// Int64 is the interface that handles a int64 collection.
type Int64 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...int64) Int64
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...int64) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Int64) Int64
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Int64) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Int64
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, int64)) Int64
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, int64) bool) Int64
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, int64)) Int64
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, int64) bool) Int64
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) int64
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (int64, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (int64, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (int64, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Int64
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...int64) Int64
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...int64) Int64
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, int64) int64) Int64
	// Poll removes the first element from the slice and returns that removed element.
	Poll() int64
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (int64, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (int64, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() int64
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (int64, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (int64, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Int64) Int64
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s Int64) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...int64) Int64
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...int64) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...int64) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v int64) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() Int64
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Int64
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Int64
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...int64) int
	// Values returns the internal values of the slice.
	Values() []int64
}

// NewInt64 returns a new Int64 interface.
func NewInt64(i ...int64) Int64 {
	return (&int64Container{&Slice{}}).Append(i...)
}

type int64Container struct{ s *Slice }

// Append implements Append for Int64.
func (u *int64Container) Append(i ...int64) Int64 {
	u.s.Append(int64ToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for Int64.
func (u *int64Container) AppendLength(i ...int64) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for Int64.
func (u *int64Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for Int64.
func (u *int64Container) Concatenate(v Int64) Int64 {
	u.s.Concatenate(v.(*int64Container).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for Int64.
func (u *int64Container) ConcatenateLength(v Int64) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for Int64.
func (u *int64Container) Delete(i int) Int64 {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for Int64.
func (u *int64Container) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for Int64.
func (u *int64Container) Each(fn func(int, int64)) Int64 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int64)))
	})
	return u
}

// EachBreak implements EachBreak for Int64.
func (u *int64Container) EachBreak(fn func(int, int64) bool) Int64 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int64)))
	})
	return u
}

// EachReverse implements EachReverse for Int64.
func (u *int64Container) EachReverse(fn func(int, int64)) Int64 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int64)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for Int64.
func (u *int64Container) EachReverseBreak(fn func(int, int64) bool) Int64 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int64)))
	})
	return u
}

// Fetch implements Fetch for Int64.
func (u *int64Container) Fetch(i int) int64 {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for Int64.
func (u *int64Container) FetchLength(i int) (int64, int) {
	v, i := u.s.FetchLength(i)
	return v.(int64), i
}

// Get implements Get for Int64.
func (u *int64Container) Get(i int) (int64, bool) {
	var (
		ok bool
		s  int64
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(int64)
	}
	return s, ok
}

// GetLength implements GetLength for Int64.
func (u *int64Container) GetLength(i int) (int64, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(int64), l, ok
}

// Len implements Len for Int64.
func (u *int64Container) Len() int {
	return u.s.Len()
}

// Less implements Less for Int64.
func (u *int64Container) Less(i int, j int) bool {
	return u.Fetch(i) < u.Fetch(j)
}

// Make implements Make for Int64.
func (u *int64Container) Make(i int) Int64 {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for Int64.
func (u *int64Container) MakeEach(v ...int64) Int64 {
	u.s.MakeEach(int64ToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for Int64.
func (u *int64Container) MakeEachReverse(v ...int64) Int64 {
	u.s.MakeEachReverse(int64ToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for Int64.
func (u *int64Container) Map(fn func(int, int64) int64) Int64 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int64)))
	})
	return u
}

// Poll implements Poll for Int64.
func (u *int64Container) Poll() int64 {
	var (
		s int64
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(int64))
	}
	return s
}

// PollLength implements PollLength for Int64.
func (u *int64Container) PollLength() (int64, int) {
	v, l := u.s.PollLength()
	return v.(int64), l
}

// PollOK implements PollOK for Int64.
func (u *int64Container) PollOK() (int64, bool) {
	v, ok := u.s.PollOK()
	return v.(int64), ok
}

// Pop implements Pop for Int64.
func (u *int64Container) Pop() int64 {
	var (
		s int64
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(int64))
	}
	return s
}

// PopLength implements PopLength for Int64.
func (u *int64Container) PopLength() (int64, int) {
	v, l := u.s.PopLength()
	return v.(int64), l
}

// PopOK implements PopOK for Int64.
func (u *int64Container) PopOK() (int64, bool) {
	v, ok := u.s.PopOK()
	return v.(int64), ok
}

// Precatenate implements Precatenate for Int64.
func (u *int64Container) Precatenate(v Int64) Int64 {
	u.s.Precatenate(v.(*int64Container).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for Int64.
func (u *int64Container) PrecatenateLength(v Int64) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for Int64.
func (u *int64Container) Prepend(i ...int64) Int64 {
	u.s.Prepend(int64ToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for Int64.
func (u *int64Container) PrependLength(v ...int64) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for Int64.
func (u *int64Container) Push(i ...int64) int {
	return u.s.Push(int64ToInterfaceSlice(i...))
}

// Replace implements Replace for Int64.
func (u *int64Container) Replace(i int, n int64) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for Int64.
func (u *int64Container) Reverse() Int64 {
	u.s.Reverse()
	return u
}

// Set implements Set for Int64.
func (u *int64Container) Set() Int64 {
	u.s.Set()
	return u
}

// Slice implements Slice for Int64.
func (u *int64Container) Slice(i int, j int) Int64 {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for Int64.
func (u *int64Container) Sort() Int64 {
	sort.Sort(u)
	return u
}

// Swap implements Swap for Int64.
func (u *int64Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for Int64.
func (u *int64Container) Unshift(i ...int64) int {
	return (u.s.Unshift(int64ToInterfaceSlice(i...)))
}

// Values implements Values for Int64.
func (u *int64Container) Values() []int64 {
	var v = make([]int64, u.Len())
	u.Each(func(i int, n int64) {
		v[i] = n
	})
	return v
}

func int64ToInterfaceSlice(n ...int64) []interface{} {
	var (
		i int
		v int64
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
