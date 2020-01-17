package slice

import (
	"sort"
)

// UInter is the interface that handles a uint collection.
type UInter interface {
	Append(...uint) UInter
	Bounds(int) bool
	Concatenate(UInter) UInter
	Delete(int) UInter
	Each(func(int, uint)) UInter
	EachBreak(func(int, uint) bool) UInter
	EachReverse(func(int, uint)) UInter
	EachReverseBreak(func(int, uint) bool) UInter
	Fetch(int) uint
	Get(int) (uint, bool)
	Len() int
	Less(int, int) bool
	Make(int) UInter
	MakeEach(...uint) UInter
	MakeEachReverse(...uint) UInter
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
	return (&uinter{&Slice{}}).Append(i...)
}

type uinter struct{ s *Slice }

func (u *uinter) Append(i ...uint) UInter {
	u.s.Append(uintToInterface(i...)...)
	return u
}

func (u *uinter) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *uinter) Concatenate(v UInter) UInter {
	u.s.Concatenate(v.(*uinter).s)
	return u
}

func (u *uinter) Delete(i int) UInter {
	u.s.Delete(i)
	return u
}

func (u *uinter) Each(fn func(int, uint)) UInter {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint)))
	})
	return u
}

func (u *uinter) EachBreak(fn func(int, uint) bool) UInter {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint)))
	})
	return u
}

func (u *uinter) EachReverse(fn func(int, uint)) UInter {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint)))
	})
	return u
}

func (u *uinter) EachReverseBreak(fn func(int, uint) bool) UInter {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint)))
	})
	return u
}

func (u *uinter) Fetch(i int) uint {
	var s, _ = u.Get(i)
	return s
}

func (u *uinter) Get(i int) (uint, bool) {
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

func (u *uinter) Len() int {
	return (u.s.Len())
}

func (u *uinter) Less(i int, j int) bool {
	return i < j
}

func (u *uinter) Make(i int) UInter {
	u.s.Make(i)
	return u
}

func (u *uinter) MakeEach(v ...uint) UInter {
	u.s.MakeEach(uintToInterface(v...)...)
	return u
}

func (u *uinter) MakeEachReverse(v ...uint) UInter {
	u.s.MakeEachReverse(uintToInterface(v...)...)
	return u
}

func (u *uinter) Map(fn func(int, uint) uint) UInter {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint)))
	})
	return u
}

func (u *uinter) Poll() uint {
	var (
		s uint
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint))
	}
	return s
}

func (u *uinter) Pop() uint {
	var (
		s uint
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint))
	}
	return s
}

func (u *uinter) Precatenate(v UInter) UInter {
	u.s.Precatenate(v.(*uinter).s)
	return u
}

func (u *uinter) Prepend(i ...uint) UInter {
	u.s.Prepend(uintToInterface(i...)...)
	return u
}

func (u *uinter) Push(i ...uint) int {
	return u.s.Push(uintToInterface(i...))
}

func (u *uinter) Replace(i int, n uint) bool {
	return (u.s.Replace(i, n))
}

func (u *uinter) Set() UInter {
	u.s.Set()
	return u
}

func (u *uinter) Slice(i int, j int) UInter {
	u.s.Slice(i, j)
	return u
}

func (u *uinter) Sort() UInter {
	sort.Sort(u)
	return u
}

func (u *uinter) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *uinter) Unshift(i ...uint) int {
	return (u.s.Unshift(uintToInterface(i...)))
}

func (u *uinter) Values() []uint {
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
