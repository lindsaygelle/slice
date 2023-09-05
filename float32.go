package slice

import (
	"sort"
)

// Float32 is the interface that handles a float32 collection.
type Float32 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...float32) Float32
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...float32) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Float32) Float32
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Float32) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Float32
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, float32)) Float32
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, float32) bool) Float32
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, float32)) Float32
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, float32) bool) Float32
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) float32
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (float32, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (float32, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (float32, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Float32
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...float32) Float32
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...float32) Float32
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, float32) float32) Float32
	// Poll removes the first element from the slice and returns that removed element.
	Poll() float32
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (float32, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (float32, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() float32
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (float32, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (float32, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Float32) Float32
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s Float32) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...float32) Float32
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...float32) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...float32) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v float32) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() Float32
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Float32
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Float32
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...float32) int
	// Values returns the internal values of the slice.
	Values() []float32
}

// NewFloat32 returns a new Float32 interface.
func NewFloat32(i ...float32) Float32 {
	return (&float32Container{&Slice{}}).Append(i...)
}

type float32Container struct{ s *Slice }

// Append implements Append for Float32.
func (u *float32Container) Append(i ...float32) Float32 {
	u.s.Append(float32ToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for Float32.
func (u *float32Container) AppendLength(i ...float32) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for Float32.
func (u *float32Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for Float32.
func (u *float32Container) Concatenate(v Float32) Float32 {
	u.s.Concatenate(v.(*float32Container).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for Float32.
func (u *float32Container) ConcatenateLength(v Float32) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for Float32.
func (u *float32Container) Delete(i int) Float32 {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for Float32.
func (u *float32Container) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for Float32.
func (u *float32Container) Each(fn func(int, float32)) Float32 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(float32)))
	})
	return u
}

// EachBreak implements EachBreak for Float32.
func (u *float32Container) EachBreak(fn func(int, float32) bool) Float32 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float32)))
	})
	return u
}

// EachReverse implements EachReverse for Float32.
func (u *float32Container) EachReverse(fn func(int, float32)) Float32 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(float32)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for Float32.
func (u *float32Container) EachReverseBreak(fn func(int, float32) bool) Float32 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float32)))
	})
	return u
}

// Fetch implements Fetch for Float32.
func (u *float32Container) Fetch(i int) float32 {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for Float32.
func (u *float32Container) FetchLength(i int) (float32, int) {
	v, i := u.s.FetchLength(i)
	return v.(float32), i
}

// Get implements Get for Float32.
func (u *float32Container) Get(i int) (float32, bool) {
	var (
		ok bool
		s  float32
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(float32)
	}
	return s, ok
}

// GetLength implements GetLength for Float32.
func (u *float32Container) GetLength(i int) (float32, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(float32), l, ok
}

// Len implements Len for Float32.
func (u *float32Container) Len() int {
	return u.s.Len()
}

// Less implements Less for Float32.
func (u *float32Container) Less(i int, j int) bool {
	return u.Fetch(i) < u.Fetch(j)
}

// Make implements Make for Float32.
func (u *float32Container) Make(i int) Float32 {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for Float32.
func (u *float32Container) MakeEach(v ...float32) Float32 {
	u.s.MakeEach(float32ToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for Float32.
func (u *float32Container) MakeEachReverse(v ...float32) Float32 {
	u.s.MakeEachReverse(float32ToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for Float32.
func (u *float32Container) Map(fn func(int, float32) float32) Float32 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(float32)))
	})
	return u
}

// Poll implements Poll for Float32.
func (u *float32Container) Poll() float32 {
	var (
		s float32
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(float32))
	}
	return s
}

// PollLength implements PollLength for Float32.
func (u *float32Container) PollLength() (float32, int) {
	v, l := u.s.PollLength()
	return v.(float32), l
}

// PollOK implements PollOK for Float32.
func (u *float32Container) PollOK() (float32, bool) {
	v, ok := u.s.PollOK()
	return v.(float32), ok
}

// Pop implements Pop for Float32.
func (u *float32Container) Pop() float32 {
	var (
		s float32
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(float32))
	}
	return s
}

// PopLength implements PopLength for Float32.
func (u *float32Container) PopLength() (float32, int) {
	v, l := u.s.PopLength()
	return v.(float32), l
}

// PopOK implements PopOK for Float32.
func (u *float32Container) PopOK() (float32, bool) {
	v, ok := u.s.PopOK()
	return v.(float32), ok
}

// Precatenate implements Precatenate for Float32.
func (u *float32Container) Precatenate(v Float32) Float32 {
	u.s.Precatenate(v.(*float32Container).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for Float32.
func (u *float32Container) PrecatenateLength(v Float32) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for Float32.
func (u *float32Container) Prepend(i ...float32) Float32 {
	u.s.Prepend(float32ToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for Float32.
func (u *float32Container) PrependLength(v ...float32) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for Float32.
func (u *float32Container) Push(i ...float32) int {
	return u.s.Push(float32ToInterfaceSlice(i...))
}

// Replace implements Replace for Float32.
func (u *float32Container) Replace(i int, n float32) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for Float32.
func (u *float32Container) Reverse() Float32 {
	u.s.Reverse()
	return u
}

// Set implements Set for Float32.
func (u *float32Container) Set() Float32 {
	u.s.Set()
	return u
}

// Slice implements Slice for Float32.
func (u *float32Container) Slice(i int, j int) Float32 {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for Float32.
func (u *float32Container) Sort() Float32 {
	sort.Sort(u)
	return u
}

// Swap implements Swap for Float32.
func (u *float32Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for Float32.
func (u *float32Container) Unshift(i ...float32) int {
	return (u.s.Unshift(float32ToInterfaceSlice(i...)))
}

// Values implements Values for Float32.
func (u *float32Container) Values() []float32 {
	var v = make([]float32, u.Len())
	u.Each(func(i int, n float32) {
		v[i] = n
	})
	return v
}

func float32ToInterfaceSlice(n ...float32) []interface{} {
	var (
		i int
		v float32
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
