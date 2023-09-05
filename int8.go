package slice

import (
	"sort"
)

// Int8 is the interface that handles a int8 collection.
type Int8 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...int8) Int8
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...int8) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Int8) Int8
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Int8) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Int8
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, int8)) Int8
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, int8) bool) Int8
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, int8)) Int8
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, int8) bool) Int8
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) int8
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (int8, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (int8, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (int8, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Int8
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...int8) Int8
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...int8) Int8
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, int8) int8) Int8
	// Poll removes the first element from the slice and returns that removed element.
	Poll() int8
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PollLength() (int8, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (int8, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() int8
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PopLength() (int8, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	// TODO PopOK() (int8, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Int8) Int8
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	// TODO PrecatenateLength(s Int8) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...int8) Int8
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	// TODO PrependLength(i ...int8) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...int8) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v int8) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	// TODO Reverse() Int8
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Int8
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Int8
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...int8) int
	// Values returns the internal values of the slice.
	Values() []int8
}

// NewInt8 returns a new Int8 interface.
func NewInt8(i ...int8) Int8 {
	return (&int8Container{&Slice{}}).Append(i...)
}

type int8Container struct{ s *Slice }

func (u *int8Container) Append(i ...int8) Int8 {
	u.s.Append(int8ToInterfaceSlice(i...)...)
	return u
}

func (u *int8Container) AppendLength(i ...int8) int {
	return u.Append(i...).Len()
}

func (u *int8Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *int8Container) Concatenate(v Int8) Int8 {
	u.s.Concatenate(v.(*int8Container).s)
	return u
}

func (u *int8Container) ConcatenateLength(v Int8) int {
	return u.Concatenate(u).Len()
}

func (u *int8Container) Delete(i int) Int8 {
	u.s.Delete(i)
	return u
}

func (u *int8Container) DeleteLength(i int) int {
	return u.s.Delete(i).Len()
}

func (u *int8Container) Each(fn func(int, int8)) Int8 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int8)))
	})
	return u
}

func (u *int8Container) EachBreak(fn func(int, int8) bool) Int8 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int8)))
	})
	return u
}

func (u *int8Container) EachReverse(fn func(int, int8)) Int8 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int8)))
	})
	return u
}

func (u *int8Container) EachReverseBreak(fn func(int, int8) bool) Int8 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int8)))
	})
	return u
}

func (u *int8Container) Fetch(i int) int8 {
	var s, _ = u.Get(i)
	return s
}

func (u *int8Container) FetchLength(i int) (int8, int) {
	v, i := u.s.FetchLength(i)
	return v.(int8), i
}

func (u *int8Container) Get(i int) (int8, bool) {
	var (
		ok bool
		s  int8
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(int8)
	}
	return s, ok
}

func (u *int8Container) GetLength(i int) (int8, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(int8), l, ok
}

func (u *int8Container) Len() int {
	return (u.s.Len())
}

func (u *int8Container) Less(i int, j int) bool {
	return i < j
}

func (u *int8Container) Make(i int) Int8 {
	u.s.Make(i)
	return u
}

func (u *int8Container) MakeEach(v ...int8) Int8 {
	u.s.MakeEach(int8ToInterfaceSlice(v...)...)
	return u
}

func (u *int8Container) MakeEachReverse(v ...int8) Int8 {
	u.s.MakeEachReverse(int8ToInterfaceSlice(v...)...)
	return u
}

func (u *int8Container) Map(fn func(int, int8) int8) Int8 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int8)))
	})
	return u
}

func (u *int8Container) Poll() int8 {
	var (
		s int8
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(int8))
	}
	return s
}

func (u *int8Container) PollLength() (int8, int) {
	v, l := u.s.PollLength()
	return v.(int8), l
}

func (u *int8Container) PollOK() (int8, bool) {
	v, ok := u.s.PollOK()
	return v.(int8), ok
}

func (u *int8Container) Pop() int8 {
	var (
		s int8
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(int8))
	}
	return s
}

func (u *int8Container) Precatenate(v Int8) Int8 {
	u.s.Precatenate(v.(*int8Container).s)
	return u
}

func (u *int8Container) Prepend(i ...int8) Int8 {
	u.s.Prepend(int8ToInterfaceSlice(i...)...)
	return u
}

func (u *int8Container) Push(i ...int8) int {
	return u.s.Push(int8ToInterfaceSlice(i...))
}

func (u *int8Container) Replace(i int, n int8) bool {
	return (u.s.Replace(i, n))
}

func (u *int8Container) Set() Int8 {
	u.s.Set()
	return u
}

func (u *int8Container) Slice(i int, j int) Int8 {
	u.s.Slice(i, j)
	return u
}

func (u *int8Container) Sort() Int8 {
	sort.Sort(u)
	return u
}

func (u *int8Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *int8Container) Unshift(i ...int8) int {
	return (u.s.Unshift(int8ToInterfaceSlice(i...)))
}

func (u *int8Container) Values() []int8 {
	var v = make([]int8, u.Len())
	u.Each(func(i int, n int8) {
		v[i] = n
	})
	return v
}

func int8ToInterfaceSlice(n ...int8) []interface{} {
	var (
		i int
		v int8
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
