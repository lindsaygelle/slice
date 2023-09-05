package slice

import (
	"sort"
)

// Complex64 is the interface that handles a complex64 collection.
type Complex64 interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...complex64) Complex64
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...complex64) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Complex64) Complex64
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Complex64) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Complex64
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, complex64)) Complex64
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, complex64) bool) Complex64
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, complex64)) Complex64
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, complex64) bool) Complex64
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) complex64
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (complex64, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (complex64, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (complex64, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Complex64
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...complex64) Complex64
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...complex64) Complex64
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, complex64) complex64) Complex64
	// Poll removes the first element from the slice and returns that removed element.
	Poll() complex64
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (complex64, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (complex64, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() complex64
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (complex64, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (complex64, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Complex64) Complex64
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s Complex64) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...complex64) Complex64
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...complex64) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...complex64) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v complex64) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() Complex64
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Complex64
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Complex64
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...complex64) int
	// Values returns the internal values of the slice.
	Values() []complex64
}

// NewComplex64 returns a new Complex64 interface.
func NewComplex64(i ...complex64) Complex64 {
	return (&complex64Container{&Slice{}}).Append(i...)
}

type complex64Container struct{ s *Slice }

// Append implements Append for Complex64.
func (u *complex64Container) Append(i ...complex64) Complex64 {
	u.s.Append(complex64ToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for Complex64.
func (u *complex64Container) AppendLength(i ...complex64) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for Complex64.
func (u *complex64Container) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for Complex64.
func (u *complex64Container) Concatenate(v Complex64) Complex64 {
	u.s.Concatenate(v.(*complex64Container).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for Complex64.
func (u *complex64Container) ConcatenateLength(v Complex64) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for Complex64.
func (u *complex64Container) Delete(i int) Complex64 {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for Complex64.
func (u *complex64Container) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for Complex64.
func (u *complex64Container) Each(fn func(int, complex64)) Complex64 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(complex64)))
	})
	return u
}

// EachBreak implements EachBreak for Complex64.
func (u *complex64Container) EachBreak(fn func(int, complex64) bool) Complex64 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(complex64)))
	})
	return u
}

// EachReverse implements EachReverse for Complex64.
func (u *complex64Container) EachReverse(fn func(int, complex64)) Complex64 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(complex64)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for Complex64.
func (u *complex64Container) EachReverseBreak(fn func(int, complex64) bool) Complex64 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(complex64)))
	})
	return u
}

// Fetch implements Fetch for Complex64.
func (u *complex64Container) Fetch(i int) complex64 {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for Complex64.
func (u *complex64Container) FetchLength(i int) (complex64, int) {
	v, i := u.s.FetchLength(i)
	return v.(complex64), i
}

// Get implements Get for Complex64.
func (u *complex64Container) Get(i int) (complex64, bool) {
	var (
		ok bool
		s  complex64
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(complex64)
	}
	return s, ok
}

// GetLength implements GetLength for Complex64.
func (u *complex64Container) GetLength(i int) (complex64, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(complex64), l, ok
}

// Len implements Len for Complex64.
func (u *complex64Container) Len() int {
	return u.s.Len()
}

// Less implements Less for Complex64.
func (u *complex64Container) Less(i int, j int) bool {
	return real(u.Fetch(i)) < real(u.Fetch(j))
}

// Make implements Make for Complex64.
func (u *complex64Container) Make(i int) Complex64 {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for Complex64.
func (u *complex64Container) MakeEach(v ...complex64) Complex64 {
	u.s.MakeEach(complex64ToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for Complex64.
func (u *complex64Container) MakeEachReverse(v ...complex64) Complex64 {
	u.s.MakeEachReverse(complex64ToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for Complex64.
func (u *complex64Container) Map(fn func(int, complex64) complex64) Complex64 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(complex64)))
	})
	return u
}

// Poll implements Poll for Complex64.
func (u *complex64Container) Poll() complex64 {
	var (
		s complex64
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(complex64))
	}
	return s
}

// PollLength implements PollLength for Complex64.
func (u *complex64Container) PollLength() (complex64, int) {
	v, l := u.s.PollLength()
	return v.(complex64), l
}

// PollOK implements PollOK for Complex64.
func (u *complex64Container) PollOK() (complex64, bool) {
	v, ok := u.s.PollOK()
	return v.(complex64), ok
}

// Pop implements Pop for Complex64.
func (u *complex64Container) Pop() complex64 {
	var (
		s complex64
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(complex64))
	}
	return s
}

// PopLength implements PopLength for Complex64.
func (u *complex64Container) PopLength() (complex64, int) {
	v, l := u.s.PopLength()
	return v.(complex64), l
}

// PopOK implements PopOK for Complex64.
func (u *complex64Container) PopOK() (complex64, bool) {
	v, ok := u.s.PopOK()
	return v.(complex64), ok
}

// Precatenate implements Precatenate for Complex64.
func (u *complex64Container) Precatenate(v Complex64) Complex64 {
	u.s.Precatenate(v.(*complex64Container).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for Complex64.
func (u *complex64Container) PrecatenateLength(v Complex64) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for Complex64.
func (u *complex64Container) Prepend(i ...complex64) Complex64 {
	u.s.Prepend(complex64ToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for Complex64.
func (u *complex64Container) PrependLength(v ...complex64) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for Complex64.
func (u *complex64Container) Push(i ...complex64) int {
	return u.s.Push(complex64ToInterfaceSlice(i...))
}

// Replace implements Replace for Complex64.
func (u *complex64Container) Replace(i int, n complex64) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for Complex64.
func (u *complex64Container) Reverse() Complex64 {
	u.s.Reverse()
	return u
}

// Set implements Set for Complex64.
func (u *complex64Container) Set() Complex64 {
	u.s.Set()
	return u
}

// Slice implements Slice for Complex64.
func (u *complex64Container) Slice(i int, j int) Complex64 {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for Complex64.
func (u *complex64Container) Sort() Complex64 {
	sort.Sort(u)
	return u
}

// Swap implements Swap for Complex64.
func (u *complex64Container) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for Complex64.
func (u *complex64Container) Unshift(i ...complex64) int {
	return (u.s.Unshift(complex64ToInterfaceSlice(i...)))
}

// Values implements Values for Complex64.
func (u *complex64Container) Values() []complex64 {
	var v = make([]complex64, u.Len())
	u.Each(func(i int, n complex64) {
		v[i] = n
	})
	return v
}

func complex64ToInterfaceSlice(n ...complex64) []interface{} {
	var (
		i int
		v complex64
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
