package slice

import (
	"sort"
)

// Int32 is the interface that handles a int32 collection.
type Int32 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...int32) Int32
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...int32) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Int32) Int32
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Int32) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Int32
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, int32)) Int32
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, int32) bool) Int32
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, int32)) Int32
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, int32) bool) Int32
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) int32
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (int32, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (int32, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (int32, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Int32
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...int32) Int32
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...int32) Int32
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, int32) int32) Int32
	// Poll removes the first element from the slice and returns that removed element.
	Poll() int32
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PollLength() (int32, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (int32, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() int32
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PopLength() (int32, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	// TODO PopOK() (int32, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Int32) Int32
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	// TODO PrecatenateLength(s Int32) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...int32) Int32
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	// TODO PrependLength(i ...int32) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...int32) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v int32) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	// TODO Reverse() Int32
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Int32
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Int32
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...int32) int
	// Values returns the internal values of the slice.
	Values() []int32
}

// NewInt32 returns a new Int32 interface.
func NewInt32(i ...int32) Int32 {
	return (&int32Container{&Slice{}}).Append(i...)
}

type int32Container struct{ s *Slice }

func (u *int32Container) Append(i ...int32) Int32 {
	u.s.Append(int32ToInterfaceSlice(i...)...)
	return u
}

func (u *int32Container) AppendLength(i ...int32) int {
	return u.Append(i...).Len()
}

func (u *int32Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *int32Container) Concatenate(v Int32) Int32 {
	u.s.Concatenate(v.(*int32Container).s)
	return u
}

func (u *int32Container) ConcatenateLength(v Int32) int {
	return u.Concatenate(u).Len()
}

func (u *int32Container) Delete(i int) Int32 {
	u.s.Delete(i)
	return u
}

func (u *int32Container) DeleteLength(i int) int {
	return u.s.Delete(i).Len()
}

func (u *int32Container) Each(fn func(int, int32)) Int32 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int32)))
	})
	return u
}

func (u *int32Container) EachBreak(fn func(int, int32) bool) Int32 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int32)))
	})
	return u
}

func (u *int32Container) EachReverse(fn func(int, int32)) Int32 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int32)))
	})
	return u
}

func (u *int32Container) EachReverseBreak(fn func(int, int32) bool) Int32 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int32)))
	})
	return u
}

func (u *int32Container) Fetch(i int) int32 {
	var s, _ = u.Get(i)
	return s
}

func (u *int32Container) FetchLength(i int) (int32, int) {
	v, i := u.s.FetchLength(i)
	return v.(int32), i
}

func (u *int32Container) Get(i int) (int32, bool) {
	var (
		ok bool
		s  int32
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(int32)
	}
	return s, ok
}

func (u *int32Container) GetLength(i int) (int32, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(int32), l, ok
}

func (u *int32Container) Len() int {
	return (u.s.Len())
}

func (u *int32Container) Less(i int, j int) bool {
	return i < j
}

func (u *int32Container) Make(i int) Int32 {
	u.s.Make(i)
	return u
}

func (u *int32Container) MakeEach(v ...int32) Int32 {
	u.s.MakeEach(int32ToInterfaceSlice(v...)...)
	return u
}

func (u *int32Container) MakeEachReverse(v ...int32) Int32 {
	u.s.MakeEachReverse(int32ToInterfaceSlice(v...)...)
	return u
}

func (u *int32Container) Map(fn func(int, int32) int32) Int32 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int32)))
	})
	return u
}

func (u *int32Container) Poll() int32 {
	var (
		s int32
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(int32))
	}
	return s
}

func (u *int32Container) PollLength() (int32, int) {
	v, l := u.s.PollLength()
	return v.(int32), l
}

func (u *int32Container) PollOK() (int32, bool) {
	v, ok := u.s.PollOK()
	return v.(int32), ok
}

func (u *int32Container) Pop() int32 {
	var (
		s int32
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(int32))
	}
	return s
}

func (u *int32Container) Precatenate(v Int32) Int32 {
	u.s.Precatenate(v.(*int32Container).s)
	return u
}

func (u *int32Container) Prepend(i ...int32) Int32 {
	u.s.Prepend(int32ToInterfaceSlice(i...)...)
	return u
}

func (u *int32Container) Push(i ...int32) int {
	return u.s.Push(int32ToInterfaceSlice(i...))
}

func (u *int32Container) Replace(i int, n int32) bool {
	return (u.s.Replace(i, n))
}

func (u *int32Container) Set() Int32 {
	u.s.Set()
	return u
}

func (u *int32Container) Slice(i int, j int) Int32 {
	u.s.Slice(i, j)
	return u
}

func (u *int32Container) Sort() Int32 {
	sort.Sort(u)
	return u
}

func (u *int32Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *int32Container) Unshift(i ...int32) int {
	return (u.s.Unshift(int32ToInterfaceSlice(i...)))
}

func (u *int32Container) Values() []int32 {
	var v = make([]int32, u.Len())
	u.Each(func(i int, n int32) {
		v[i] = n
	})
	return v
}

func int32ToInterfaceSlice(n ...int32) []interface{} {
	var (
		i int
		v int32
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
