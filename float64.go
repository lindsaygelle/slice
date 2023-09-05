package slice

import (
	"sort"
)

// Float64 is the interface that handles a float64 collection.
type Float64 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...float64) Float64
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...float64) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Float64) Float64
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Float64) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Float64
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, float64)) Float64
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, float64) bool) Float64
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, float64)) Float64
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, float64) bool) Float64
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) float64
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (float64, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (float64, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (float64, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Float64
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...float64) Float64
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...float64) Float64
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, float64) float64) Float64
	// Poll removes the first element from the slice and returns that removed element.
	Poll() float64
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PollLength() (float64, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (float64, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() float64
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PopLength() (float64, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	// TODO PopOK() (float64, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Float64) Float64
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	// TODO PrecatenateLength(s Float64) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...float64) Float64
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	// TODO PrependLength(i ...float64) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...float64) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v float64) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	// TODO Reverse() Float64
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Float64
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Float64
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...float64) int
	// Values returns the internal values of the slice.
	Values() []float64
}

// NewFloat64 returns a new Float64 interface.
func NewFloat64(i ...float64) Float64 {
	return (&float64Container{&Slice{}}).Append(i...)
}

type float64Container struct{ s *Slice }

func (u *float64Container) Append(i ...float64) Float64 {
	u.s.Append(float64ToInterfaceSlice(i...)...)
	return u
}

func (u *float64Container) AppendLength(i ...float64) int {
	return u.Append(i...).Len()
}

func (u *float64Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *float64Container) Concatenate(v Float64) Float64 {
	u.s.Concatenate(v.(*float64Container).s)
	return u
}

func (u *float64Container) ConcatenateLength(v Float64) int {
	return u.Concatenate(u).Len()
}

func (u *float64Container) Delete(i int) Float64 {
	u.s.Delete(i)
	return u
}

func (u *float64Container) DeleteLength(i int) int {
	return u.s.Delete(i).Len()
}

func (u *float64Container) Each(fn func(int, float64)) Float64 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(float64)))
	})
	return u
}

func (u *float64Container) EachBreak(fn func(int, float64) bool) Float64 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float64)))
	})
	return u
}

func (u *float64Container) EachReverse(fn func(int, float64)) Float64 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(float64)))
	})
	return u
}

func (u *float64Container) EachReverseBreak(fn func(int, float64) bool) Float64 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float64)))
	})
	return u
}

func (u *float64Container) Fetch(i int) float64 {
	var s, _ = u.Get(i)
	return s
}

func (u *float64Container) FetchLength(i int) (float64, int) {
	v, i := u.s.FetchLength(i)
	return v.(float64), i
}

func (u *float64Container) Get(i int) (float64, bool) {
	var (
		ok bool
		s  float64
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(float64)
	}
	return s, ok
}

func (u *float64Container) GetLength(i int) (float64, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(float64), l, ok
}

func (u *float64Container) Len() int {
	return (u.s.Len())
}

func (u *float64Container) Less(i int, j int) bool {
	return i < j
}

func (u *float64Container) Make(i int) Float64 {
	u.s.Make(i)
	return u
}

func (u *float64Container) MakeEach(v ...float64) Float64 {
	u.s.MakeEach(float64ToInterfaceSlice(v...)...)
	return u
}

func (u *float64Container) MakeEachReverse(v ...float64) Float64 {
	u.s.MakeEachReverse(float64ToInterfaceSlice(v...)...)
	return u
}

func (u *float64Container) Map(fn func(int, float64) float64) Float64 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(float64)))
	})
	return u
}

func (u *float64Container) Poll() float64 {
	var (
		s float64
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(float64))
	}
	return s
}

func (u *float64Container) PollLength() (float64, int) {
	v, l := u.s.PollLength()
	return v.(float64), l
}

func (u *float64Container) PollOK() (float64, bool) {
	v, ok := u.s.PollOK()
	return v.(float64), ok
}

func (u *float64Container) Pop() float64 {
	var (
		s float64
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(float64))
	}
	return s
}

func (u *float64Container) Precatenate(v Float64) Float64 {
	u.s.Precatenate(v.(*float64Container).s)
	return u
}

func (u *float64Container) Prepend(i ...float64) Float64 {
	u.s.Prepend(float64ToInterfaceSlice(i...)...)
	return u
}

func (u *float64Container) Push(i ...float64) int {
	return u.s.Push(float64ToInterfaceSlice(i...))
}

func (u *float64Container) Replace(i int, n float64) bool {
	return (u.s.Replace(i, n))
}

func (u *float64Container) Set() Float64 {
	u.s.Set()
	return u
}

func (u *float64Container) Slice(i int, j int) Float64 {
	u.s.Slice(i, j)
	return u
}

func (u *float64Container) Sort() Float64 {
	sort.Sort(u)
	return u
}

func (u *float64Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *float64Container) Unshift(i ...float64) int {
	return (u.s.Unshift(float64ToInterfaceSlice(i...)))
}

func (u *float64Container) Values() []float64 {
	var v = make([]float64, u.Len())
	u.Each(func(i int, n float64) {
		v[i] = n
	})
	return v
}

func float64ToInterfaceSlice(n ...float64) []interface{} {
	var (
		i int
		v float64
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
