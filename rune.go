package slice

import (
	"sort"
)

// Rune is the interface that handles a rune collection.
type Rune interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...rune) Rune
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...rune) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Rune) Rune
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Rune) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Rune
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, rune)) Rune
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, rune) bool) Rune
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, rune)) Rune
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, rune) bool) Rune
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) rune
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (rune, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (rune, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (rune, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Rune
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...rune) Rune
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...rune) Rune
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, rune) rune) Rune
	// Poll removes the first element from the slice and returns that removed element.
	Poll() rune
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PollLength() (rune, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (rune, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() rune
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PopLength() (rune, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	// TODO PopOK() (rune, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Rune) Rune
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	// TODO PrecatenateLength(s Rune) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...rune) Rune
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	// TODO PrependLength(i ...rune) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...rune) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v rune) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	// TODO Reverse() Rune
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Rune
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Rune
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...rune) int
	// Values returns the internal values of the slice.
	Values() []rune
}

// NewRune returns a new Rune interface.
func NewRune(i ...rune) Rune {
	return (&runeContainer{&Slice{}}).Append(i...)
}

type runeContainer struct{ s *Slice }

func (u *runeContainer) Append(i ...rune) Rune {
	u.s.Append(runeToInterfaceSlice(i...)...)
	return u
}

func (u *runeContainer) AppendLength(i ...rune) int {
	return u.Append(i...).Len()
}

func (u *runeContainer) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *runeContainer) Concatenate(v Rune) Rune {
	u.s.Concatenate(v.(*runeContainer).s)
	return u
}

func (u *runeContainer) ConcatenateLength(v Rune) int {
	return u.Concatenate(u).Len()
}

func (u *runeContainer) Delete(i int) Rune {
	u.s.Delete(i)
	return u
}

func (u *runeContainer) DeleteLength(i int) int {
	return u.s.Delete(i).Len()
}

func (u *runeContainer) Each(fn func(int, rune)) Rune {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(rune)))
	})
	return u
}

func (u *runeContainer) EachBreak(fn func(int, rune) bool) Rune {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(rune)))
	})
	return u
}

func (u *runeContainer) EachReverse(fn func(int, rune)) Rune {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(rune)))
	})
	return u
}

func (u *runeContainer) EachReverseBreak(fn func(int, rune) bool) Rune {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(rune)))
	})
	return u
}

func (u *runeContainer) Fetch(i int) rune {
	var s, _ = u.Get(i)
	return s
}

func (u *runeContainer) FetchLength(i int) (rune, int) {
	v, i := u.s.FetchLength(i)
	return v.(rune), i
}

func (u *runeContainer) Get(i int) (rune, bool) {
	var (
		ok bool
		s  rune
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(rune)
	}
	return s, ok
}

func (u *runeContainer) GetLength(i int) (rune, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(rune), l, ok
}

func (u *runeContainer) Len() int {
	return (u.s.Len())
}

func (u *runeContainer) Less(i int, j int) bool {
	return i < j
}

func (u *runeContainer) Make(i int) Rune {
	u.s.Make(i)
	return u
}

func (u *runeContainer) MakeEach(v ...rune) Rune {
	u.s.MakeEach(runeToInterfaceSlice(v...)...)
	return u
}

func (u *runeContainer) MakeEachReverse(v ...rune) Rune {
	u.s.MakeEachReverse(runeToInterfaceSlice(v...)...)
	return u
}

func (u *runeContainer) Map(fn func(int, rune) rune) Rune {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(rune)))
	})
	return u
}

func (u *runeContainer) Poll() rune {
	var (
		s rune
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(rune))
	}
	return s
}

func (u *runeContainer) PollLength() (rune, int) {
	v, l := u.s.PollLength()
	return v.(rune), l
}

func (u *runeContainer) PollOK() (rune, bool) {
	v, ok := u.s.PollOK()
	return v.(rune), ok
}

func (u *runeContainer) Pop() rune {
	var (
		s rune
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(rune))
	}
	return s
}

func (u *runeContainer) Precatenate(v Rune) Rune {
	u.s.Precatenate(v.(*runeContainer).s)
	return u
}

func (u *runeContainer) Prepend(i ...rune) Rune {
	u.s.Prepend(runeToInterfaceSlice(i...)...)
	return u
}

func (u *runeContainer) Push(i ...rune) int {
	return u.s.Push(runeToInterfaceSlice(i...))
}

func (u *runeContainer) Replace(i int, n rune) bool {
	return (u.s.Replace(i, n))
}

func (u *runeContainer) Set() Rune {
	u.s.Set()
	return u
}

func (u *runeContainer) Slice(i int, j int) Rune {
	u.s.Slice(i, j)
	return u
}

func (u *runeContainer) Sort() Rune {
	sort.Sort(u)
	return u
}

func (u *runeContainer) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *runeContainer) Unshift(i ...rune) int {
	return (u.s.Unshift(runeToInterfaceSlice(i...)))
}

func (u *runeContainer) Values() []rune {
	var v = make([]rune, u.Len())
	u.Each(func(i int, n rune) {
		v[i] = n
	})
	return v
}

func runeToInterfaceSlice(n ...rune) []interface{} {
	var (
		i int
		v rune
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
