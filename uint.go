package slice

import (
	"sort"
)

// UInter is the interface that handles a uint collection.
type UInter interface {
	Append(...uint) UInter
	Bounds(int) bool
	Concatenate(UInter) UInter
	Each(func(int, uint)) UInter
	EachBreak(func(int, uint) bool) UInter
	EachReverse(func(int, uint)) UInter
	EachReverseBreak(func(int, uint) bool) UInter
	Fetch(int) uint
	Get(int) (uint, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, uint) uint) UInter
	Poll() uint
	Pop() uint
	Precatenate(UInter) UInter
	Prepend(...uint) UInter
	Push(...uint) int
	Replace(int, uint) bool
	Set() UInter
	Slice(int, int) UInter
	Sort() UInter
	Swap(int, int)
	Unshift(...uint) int
	Values() []uint
}

// NewUInter returns a new UInter interface.
func NewUInter(i ...uint) UInter {
	return (&uinterger{&Slice{}}).Append(i...)
}

type uinterger struct{ s *Slice }

func (u *uinterger) Append(i ...uint) UInter {
	u.s.Append(uintToInterface(i...)...)
	return u
}

func (u *uinterger) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *uinterger) Concatenate(v UInter) UInter {
	u.s.Concatenate(v.(*uinterger).s)
	return u
}

func (u *uinterger) Each(fn func(int, uint)) UInter {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint)))
	})
	return u
}

func (u *uinterger) EachBreak(fn func(int, uint) bool) UInter {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint)))
	})
	return u
}

func (u *uinterger) EachReverse(fn func(int, uint)) UInter {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint)))
	})
	return u
}

func (u *uinterger) EachReverseBreak(fn func(int, uint) bool) UInter {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint)))
	})
	return u
}

func (u *uinterger) Fetch(i int) uint {
	var s, _ = u.Get(i)
	return s
}

func (u *uinterger) Get(i int) (uint, bool) {
	var (
		ok bool
		s  uint
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint)
	}
	return s, ok
}

func (u *uinterger) Len() int {
	return (u.s.Len())
}

func (u *uinterger) Less(i int, j int) bool {
	return i < j
}

func (u *uinterger) Map(fn func(int, uint) uint) UInter {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint)))
	})
	return u
}

func (u *uinterger) Poll() uint {
	var (
		s uint
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint))
	}
	return s
}

func (u *uinterger) Pop() uint {
	var (
		s uint
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint))
	}
	return s
}

func (u *uinterger) Precatenate(v UInter) UInter {
	u.s.Precatenate(v.(*uinterger).s)
	return u
}

func (u *uinterger) Prepend(i ...uint) UInter {
	u.s.Prepend(uintToInterface(i...)...)
	return u
}

func (u *uinterger) Push(i ...uint) int {
	return u.s.Push(uintToInterface(i...))
}

func (u *uinterger) Replace(i int, n uint) bool {
	return (u.s.Replace(i, n))
}

func (u *uinterger) Set() UInter {
	u.s.Set()
	return u
}

func (u *uinteger) Slice(i int, j int) UInter {
	u.s.Slice(i, j)
	return u
}

func (u *uinterger) Sort() UInter {
	sort.Sort(u)
	return u
}

func (u *uinterger) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *uinterger) Unshift(i ...uint) int {
	return (u.s.Unshift(uintToInterface(i...)))
}

func (u *uinterger) Values() []uint {
	var v = make([]uint, u.Len())
	u.Each(func(i int, n uint) {
		v[i] = n
	})
	return v
}

func uintToInterface(n ...uint) []interface{} {
	var (
		i int
		v uint
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
