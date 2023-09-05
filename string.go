package slice

import (
	"sort"
)

// String is the interface that handles a string collection.
type String interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...string) String
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...string) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s String) String
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s String) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) String
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, string)) String
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, string) bool) String
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, string)) String
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, string) bool) String
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) string
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (string, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (string, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (string, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) String
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...string) String
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...string) String
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, string) string) String
	// Poll removes the first element from the slice and returns that removed element.
	Poll() string
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (string, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (string, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() string
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (string, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (string, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s String) String
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s String) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...string) String
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...string) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...string) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v string) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() String
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() String
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) String
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...string) int
	// Values returns the internal values of the slice.
	Values() []string
}

// NewString returns a new String interface.
func NewString(i ...string) String {
	return (&stringContainer{&Slice{}}).Append(i...)
}

type stringContainer struct{ s *Slice }

// Append implements Append for String.
func (u *stringContainer) Append(i ...string) String {
	u.s.Append(stringToInterfaceSlice(i...)...)
	return u
}

// AppendLength implements Append for String.
func (u *stringContainer) AppendLength(i ...string) int {
	return u.Append(i...).Len()
}

// Bounds implements Bounds for String.
func (u *stringContainer) Bounds(i int) bool {
	return u.s.Bounds(i)
}

// Concatenate implements Concatenate for String.
func (u *stringContainer) Concatenate(v String) String {
	u.s.Concatenate(v.(*stringContainer).s)
	return u
}

// ConcatenateLength implements ConcatenateLength for String.
func (u *stringContainer) ConcatenateLength(v String) int {
	return u.Concatenate(u).Len()
}

// Delete implements Delete for String.
func (u *stringContainer) Delete(i int) String {
	u.s.Delete(i)
	return u
}

// DeleteLength implements DeleteLength for String.
func (u *stringContainer) DeleteLength(i int) int {
	return u.Delete(i).Len()
}

// Each implements Each for String.
func (u *stringContainer) Each(fn func(int, string)) String {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(string)))
	})
	return u
}

// EachBreak implements EachBreak for String.
func (u *stringContainer) EachBreak(fn func(int, string) bool) String {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(string)))
	})
	return u
}

// EachReverse implements EachReverse for String.
func (u *stringContainer) EachReverse(fn func(int, string)) String {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(string)))
	})
	return u
}

// EachReverseBreak implements EachReverseBreak for String.
func (u *stringContainer) EachReverseBreak(fn func(int, string) bool) String {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(string)))
	})
	return u
}

// Fetch implements Fetch for String.
func (u *stringContainer) Fetch(i int) string {
	var s, _ = u.Get(i)
	return s
}

// FetchLength implements FetchLength for String.
func (u *stringContainer) FetchLength(i int) (string, int) {
	v, i := u.s.FetchLength(i)
	return v.(string), i
}

// Get implements Get for String.
func (u *stringContainer) Get(i int) (string, bool) {
	var (
		ok bool
		s  string
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(string)
	}
	return s, ok
}

// GetLength implements GetLength for String.
func (u *stringContainer) GetLength(i int) (string, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(string), l, ok
}

// Len implements Len for String.
func (u *stringContainer) Len() int {
	return u.s.Len()
}

// Less implements Less for String.
func (u *stringContainer) Less(i int, j int) bool {
	return u.Fetch(i) < u.Fetch(j)
}

// Make implements Make for String.
func (u *stringContainer) Make(i int) String {
	u.s.Make(i)
	return u
}

// MakeEach implements MakeEach for String.
func (u *stringContainer) MakeEach(v ...string) String {
	u.s.MakeEach(stringToInterfaceSlice(v...)...)
	return u
}

// MakeEachReverse implements MakeEachReverse for String.
func (u *stringContainer) MakeEachReverse(v ...string) String {
	u.s.MakeEachReverse(stringToInterfaceSlice(v...)...)
	return u
}

// Map implements Map for String.
func (u *stringContainer) Map(fn func(int, string) string) String {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(string)))
	})
	return u
}

// Poll implements Poll for String.
func (u *stringContainer) Poll() string {
	var (
		s string
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(string))
	}
	return s
}

// PollLength implements PollLength for String.
func (u *stringContainer) PollLength() (string, int) {
	v, l := u.s.PollLength()
	return v.(string), l
}

// PollOK implements PollOK for String.
func (u *stringContainer) PollOK() (string, bool) {
	v, ok := u.s.PollOK()
	return v.(string), ok
}

// Pop implements Pop for String.
func (u *stringContainer) Pop() string {
	var (
		s string
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(string))
	}
	return s
}

// PopLength implements PopLength for String.
func (u *stringContainer) PopLength() (string, int) {
	v, l := u.s.PopLength()
	return v.(string), l
}

// PopOK implements PopOK for String.
func (u *stringContainer) PopOK() (string, bool) {
	v, ok := u.s.PopOK()
	return v.(string), ok
}

// Precatenate implements Precatenate for String.
func (u *stringContainer) Precatenate(v String) String {
	u.s.Precatenate(v.(*stringContainer).s)
	return u
}

// PrecatenateLength implements PrecatenateLength for String.
func (u *stringContainer) PrecatenateLength(v String) int {
	return u.Precatenate(v).Len()
}

// Prepend implements Prepend for String.
func (u *stringContainer) Prepend(i ...string) String {
	u.s.Prepend(stringToInterfaceSlice(i...)...)
	return u
}

// PrependLength implements PrependLength for String.
func (u *stringContainer) PrependLength(v ...string) int {
	return u.Prepend(v...).Len()
}

// Push implements Push for String.
func (u *stringContainer) Push(i ...string) int {
	return u.s.Push(stringToInterfaceSlice(i...))
}

// Replace implements Replace for String.
func (u *stringContainer) Replace(i int, n string) bool {
	return (u.s.Replace(i, n))
}

// Reverse implements Reverse for String.
func (u *stringContainer) Reverse() String {
	u.s.Reverse()
	return u
}

// Set implements Set for String.
func (u *stringContainer) Set() String {
	u.s.Set()
	return u
}

// Slice implements Slice for String.
func (u *stringContainer) Slice(i int, j int) String {
	u.s.Slice(i, j)
	return u
}

// Sort implements Sort for String.
func (u *stringContainer) Sort() String {
	sort.Sort(u)
	return u
}

// Swap implements Swap for String.
func (u *stringContainer) Swap(i int, j int) {
	u.s.Swap(i, j)
}

// Unshift implements Unshift for String.
func (u *stringContainer) Unshift(i ...string) int {
	return (u.s.Unshift(stringToInterfaceSlice(i...)))
}

// Values implements Values for String.
func (u *stringContainer) Values() []string {
	var v = make([]string, u.Len())
	u.Each(func(i int, n string) {
		v[i] = n
	})
	return v
}

func stringToInterfaceSlice(n ...string) []interface{} {
	var (
		i int
		v string
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
