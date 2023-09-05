package slice

import (
	"sort"
)

// Int is the interface that handles a int collection.
type Int interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...int) Int
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...int) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Int) Int
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Int) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Int
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, int)) Int
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, int) bool) Int
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, int)) Int
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, int) bool) Int
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) int
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (int, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (int, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (int, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Int
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...int) Int
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...int) Int
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, int) int) Int
	// Poll removes the first element from the slice and returns that removed element.
	Poll() int
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (int, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (int, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() int
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (int, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (int, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Int) Int
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s Int) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...int) Int
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...int) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...int) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v int) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() Int
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Int
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Int
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...int) int
	// Values returns the internal values of the slice.
	Values() []int
}

// NewInt returns a new Int interface.
func NewInt(i ...int) Int {
	return (&intContainer{&Slice{}}).Append(i...)
}

type intContainer struct{ s *Slice }

// Append implements Append for Int.
func (u *intContainer) Append(i ...int) Int {
	u.s.Append(intToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for Int.
func (u *intContainer) AppendLength(i ...int) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for Int.
func (u *intContainer) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for Int.
func (u *intContainer) Concatenate(v Int) Int {
	u.s.Concatenate(v.(*intContainer).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for Int.
func (u *intContainer) ConcatenateLength(v Int) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for Int.
func (u *intContainer) Delete(i int) Int {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for Int.
func (u *intContainer) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for Int.
func (u *intContainer) Each(fn func(int, int)) Int {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return u
}

// EachBreak implements EachBreak for Int.
func (u *intContainer) EachBreak(fn func(int, int) bool) Int {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int)))
	})
	return u
}

// EachReverse implements EachReverse for Int.
func (u *intContainer) EachReverse(fn func(int, int)) Int {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for Int.
func (u *intContainer) EachReverseBreak(fn func(int, int) bool) Int {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int)))
	})
	return u
}

// Fetch implements Fetch for Int.
func (u *intContainer) Fetch(i int) int {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for Int.
func (u *intContainer) FetchLength(i int) (int, int) {
	v, i := u.s.FetchLength(i)
	return v.(int), i
}

// Get implements Get for Int.
func (u *intContainer) Get(i int) (int, bool) {
	var (
		ok bool
		s  int
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(int)
	}
	return s, ok
}

// GetLength implements GetLength for Int.
func (u *intContainer) GetLength(i int) (int, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(int), l, ok
}

// Len implements Len for Int.
func (u *intContainer) Len() int {
	return u.s.Len()
}

// Less implements Less for Int.
func (u *intContainer) Less(i int, j int) bool {
	return u.Fetch(i) < u.Fetch(j)
}

// Make implements Make for Int.
func (u *intContainer) Make(i int) Int {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for Int.
func (u *intContainer) MakeEach(v ...int) Int {
	u.s.MakeEach(intToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for Int.
func (u *intContainer) MakeEachReverse(v ...int) Int {
	u.s.MakeEachReverse(intToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for Int.
func (u *intContainer) Map(fn func(int, int) int) Int {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int)))
	})
	return u
}

// Poll implements Poll for Int.
func (u *intContainer) Poll() int {
	var (
		s int
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(int))
	}
	return s
}

// PollLength implements PollLength for Int.
func (u *intContainer) PollLength() (int, int) {
	v, l := u.s.PollLength()
	return v.(int), l
}

// PollOK implements PollOK for Int.
func (u *intContainer) PollOK() (int, bool) {
	v, ok := u.s.PollOK()
	return v.(int), ok
}

// Pop implements Pop for Int.
func (u *intContainer) Pop() int {
	var (
		s int
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(int))
	}
	return s
}

// PopLength implements PopLength for Int.
func (u *intContainer) PopLength() (int, int) {
	v, l := u.s.PopLength()
	return v.(int), l
}

// PopOK implements PopOK for Int.
func (u *intContainer) PopOK() (int, bool) {
	v, ok := u.s.PopOK()
	return v.(int), ok
}

// Precatenate implements Precatenate for Int.
func (u *intContainer) Precatenate(v Int) Int {
	u.s.Precatenate(v.(*intContainer).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for Int.
func (u *intContainer) PrecatenateLength(v Int) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for Int.
func (u *intContainer) Prepend(i ...int) Int {
	u.s.Prepend(intToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for Int.
func (u *intContainer) PrependLength(v ...int) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for Int.
func (u *intContainer) Push(i ...int) int {
	return u.s.Push(intToInterfaceSlice(i...))
}

// Replace implements Replace for Int.
func (u *intContainer) Replace(i int, n int) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for Int.
func (u *intContainer) Reverse() Int {
	u.s.Reverse()
	return u
}

// Set implements Set for Int.
func (u *intContainer) Set() Int {
	u.s.Set()
	return u
}

// Slice implements Slice for Int.
func (u *intContainer) Slice(i int, j int) Int {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for Int.
func (u *intContainer) Sort() Int {
	sort.Sort(u)
	return u
}

// Swap implements Swap for Int.
func (u *intContainer) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for Int.
func (u *intContainer) Unshift(i ...int) int {
	return (u.s.Unshift(intToInterfaceSlice(i...)))
}

// Values implements Values for Int.
func (u *intContainer) Values() []int {
	var v = make([]int, u.Len())
	u.Each(func(i int, n int) {
		v[i] = n
	})
	return v
}

func intToInterfaceSlice(n ...int) []interface{} {
	var (
		i int
		v int
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
