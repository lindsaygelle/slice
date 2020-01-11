package slice

import (
	"sort"
)

// UInter8 is the interface that handles a uint8 collection.
type UInter8 interface {
	Append(...uint8) UInter8
	Bounds(int) bool
	Concatenate(UInter8) UInter8
	Each(func(int, uint8)) UInter8
	EachBreak(func(int, uint8) bool) UInter8
	EachReverse(func(int, uint8)) UInter8
	EachReverseBreak(func(int, uint8) bool) UInter8
	Fetch(int) uint8
	Get(int) (uint8, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, uint8) uint8) UInter8
	Poll() uint8
	Pop() uint8
	Precatenate(UInter8) UInter8
	Prepend(...uint8) UInter8
	Push(...uint8) int
	Replace(int, uint8) bool
	Set() UInter8
	Sort() UInter8
	Swap(int, int)
	Unshift(...uint8) int
	Values() []uint8
}

// NewUInter8 returns a new UInter8 interface.
func NewUInter8(i ...uint8) UInter8 {
	return (&uinter8{&Slice{}}).Append(i...)
}

type uinter8 struct{ s *Slice }

func (u *uinter8) Append(i ...uint8) UInter8 {
	u.s.Append(uint8ToInterface(i...)...)
	return u
}

func (u *uinter8) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *uinter8) Concatenate(v UInter8) UInter8 {
	u.s.Concatenate(v.(*uinter8).s)
	return u
}

func (u *uinter8) Each(fn func(int, uint8)) UInter8 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint8)))
	})
	return u
}

func (u *uinter8) EachBreak(fn func(int, uint8) bool) UInter8 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint8)))
	})
	return u
}

func (u *uinter8) EachReverse(fn func(int, uint8)) UInter8 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint8)))
	})
	return u
}

func (u *uinter8) EachReverseBreak(fn func(int, uint8) bool) UInter8 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint8)))
	})
	return u
}

func (u *uinter8) Fetch(i int) uint8 {
	var s, _ = u.Get(i)
	return s
}

func (u *uinter8) Get(i int) (uint8, bool) {
	var (
		ok bool
		s  uint8
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint8)
	}
	return s, ok
}

func (u *uinter8) Len() int {
	return (u.s.Len())
}

func (u *uinter8) Less(i int, j int) bool {
	return i < j
}

func (u *uinter8) Map(fn func(int, uint8) uint8) UInter8 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint8)))
	})
	return u
}

func (u *uinter8) Poll() uint8 {
	var (
		s uint8
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint8))
	}
	return s
}

func (u *uinter8) Pop() uint8 {
	var (
		s uint8
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint8))
	}
	return s
}

func (u *uinter8) Precatenate(v UInter8) UInter8 {
	u.s.Precatenate(v.(*uinter8).s)
	return u
}

func (u *uinter8) Prepend(i ...uint8) UInter8 {
	u.s.Prepend(uint8ToInterface(i...)...)
	return u
}

func (u *uinter8) Push(i ...uint8) int {
	return u.s.Push(uint8ToInterface(i...))
}

func (u *uinter8) Replace(i int, n uint8) bool {
	return (u.s.Replace(i, n))
}

func (u *uinter8) Set() UInter8 {
	u.s.Set()
	return u
}

func (u *uinter8) Sort() UInter8 {
	sort.Sort(u)
	return u
}

func (u *uinter8) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *uinter8) Unshift(i ...uint8) int {
	return (u.s.Unshift(uint8ToInterface(i...)))
}

func (u *uinter8) Values() []uint8 {
	var v = make([]uint8, u.Len())
	u.Each(func(i int, n uint8) {
		v[i] = n
	})
	return v
}

func uint8ToInterface(n ...uint8) []interface{} {
	var (
		i int
		v uint8
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
