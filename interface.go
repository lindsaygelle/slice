package slice

import (
	"sort"
)

// Interface is the interface that handles a interface{} collection.
type Interface interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...interface{}) Interface
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...interface{}) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s Interface) Interface
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s Interface) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) Interface
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, interface{})) Interface
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, interface{}) bool) Interface
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, interface{})) Interface
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, interface{}) bool) Interface
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) interface{}
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (interface{}, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (interface{}, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (interface{}, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) Interface
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...interface{}) Interface
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...interface{}) Interface
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, interface{}) interface{}) Interface
	// Poll removes the first element from the slice and returns that removed element.
	Poll() interface{}
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PollLength() (interface{}, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (interface{}, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() interface{}
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	// TODO PopLength() (interface{}, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	// TODO PopOK() (interface{}, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s Interface) Interface
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	// TODO PrecatenateLength(s Interface) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...interface{}) Interface
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	// TODO PrependLength(i ...interface{}) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...interface{}) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v interface{}) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	// TODO Reverse() Interface
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() Interface
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) Interface
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...interface{}) int
	// Values returns the internal values of the slice.
	Values() []interface{}
}

// NewInterface returns a new Interface interface.
func NewInterface(i ...interface{}) Interface {
	return (&interfaceContainer{&Slice{}}).Append(i...)
}

type interfaceContainer struct{ s *Slice }

func (u *interfaceContainer) Append(i ...interface{}) Interface {
	u.s.Append(interfaceToInterfaceSlice(i...)...)
	return u
}

func (u *interfaceContainer) AppendLength(i ...interface{}) int {
	return u.Append(i...).Len()
}

func (u *interfaceContainer) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *interfaceContainer) Concatenate(v Interface) Interface {
	u.s.Concatenate(v.(*interfaceContainer).s)
	return u
}

func (u *interfaceContainer) ConcatenateLength(v Interface) int {
	return u.Concatenate(u).Len()
}

func (u *interfaceContainer) Delete(i int) Interface {
	u.s.Delete(i)
	return u
}

func (u *interfaceContainer) DeleteLength(i int) int {
	return u.s.Delete(i).Len()
}

func (u *interfaceContainer) Each(fn func(int, interface{})) Interface {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(interface{})))
	})
	return u
}

func (u *interfaceContainer) EachBreak(fn func(int, interface{}) bool) Interface {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(interface{})))
	})
	return u
}

func (u *interfaceContainer) EachReverse(fn func(int, interface{})) Interface {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(interface{})))
	})
	return u
}

func (u *interfaceContainer) EachReverseBreak(fn func(int, interface{}) bool) Interface {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(interface{})))
	})
	return u
}

func (u *interfaceContainer) Fetch(i int) interface{} {
	var s, _ = u.Get(i)
	return s
}

func (u *interfaceContainer) FetchLength(i int) (interface{}, int) {
	v, i := u.s.FetchLength(i)
	return v.(interface{}), i
}

func (u *interfaceContainer) Get(i int) (interface{}, bool) {
	var (
		ok bool
		s  interface{}
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(interface{})
	}
	return s, ok
}

func (u *interfaceContainer) GetLength(i int) (interface{}, int, bool) {
	v, l, ok := u.s.GetLength(i)
	return v.(interface{}), l, ok
}

func (u *interfaceContainer) Len() int {
	return (u.s.Len())
}

func (u *interfaceContainer) Less(i int, j int) bool {
	return i < j
}

func (u *interfaceContainer) Make(i int) Interface {
	u.s.Make(i)
	return u
}

func (u *interfaceContainer) MakeEach(v ...interface{}) Interface {
	u.s.MakeEach(interfaceToInterfaceSlice(v...)...)
	return u
}

func (u *interfaceContainer) MakeEachReverse(v ...interface{}) Interface {
	u.s.MakeEachReverse(interfaceToInterfaceSlice(v...)...)
	return u
}

func (u *interfaceContainer) Map(fn func(int, interface{}) interface{}) Interface {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(interface{})))
	})
	return u
}

func (u *interfaceContainer) Poll() interface{} {
	var (
		s interface{}
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(interface{}))
	}
	return s
}

func (u *interfaceContainer) PollLength() (interface{}, int) {
	v, l := u.s.PollLength()
	return v.(interface{}), l
}

func (u *interfaceContainer) PollOK() (interface{}, bool) {
	v, ok := u.s.PollOK()
	return v.(interface{}), ok
}

func (u *interfaceContainer) Pop() interface{} {
	var (
		s interface{}
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(interface{}))
	}
	return s
}

func (u *interfaceContainer) Precatenate(v Interface) Interface {
	u.s.Precatenate(v.(*interfaceContainer).s)
	return u
}

func (u *interfaceContainer) Prepend(i ...interface{}) Interface {
	u.s.Prepend(interfaceToInterfaceSlice(i...)...)
	return u
}

func (u *interfaceContainer) Push(i ...interface{}) int {
	return u.s.Push(interfaceToInterfaceSlice(i...))
}

func (u *interfaceContainer) Replace(i int, n interface{}) bool {
	return (u.s.Replace(i, n))
}

func (u *interfaceContainer) Set() Interface {
	u.s.Set()
	return u
}

func (u *interfaceContainer) Slice(i int, j int) Interface {
	u.s.Slice(i, j)
	return u
}

func (u *interfaceContainer) Sort() Interface {
	sort.Sort(u)
	return u
}

func (u *interfaceContainer) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *interfaceContainer) Unshift(i ...interface{}) int {
	return (u.s.Unshift(interfaceToInterfaceSlice(i...)))
}

func (u *interfaceContainer) Values() []interface{} {
	var v = make([]interface{}, u.Len())
	u.Each(func(i int, n interface{}) {
		v[i] = n
	})
	return v
}

func interfaceToInterfaceSlice(n ...interface{}) []interface{} {
	var (
		i int
		v interface{}
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
