package slice

import (
	"sort"
)

// UInter32 is the interface that handles a uint32 collection.
type UInter32 interface {
	Append(...uint32) UInter32
	Bounds(int) bool
	Concatenate(UInter32) UInter32
	Each(func(int, uint32)) UInter32
	EachBreak(func(int, uint32) bool) UInter32
	EachReverse(func(int, uint32)) UInter32
	EachReverseBreak(func(int, uint32) bool) UInter32
	Fetch(int) uint32
	Get(int) (uint32, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, uint32) uint32) UInter32
	Poll() uint32
	Pop() uint32
	Precatenate(UInter32) UInter32
	Prepend(...uint32) UInter32
	Push(...uint32) int
	Replace(int, uint32) bool
	Set() UInter32
	Sort() UInter32
	Swap(int, int)
	Unshift(...uint32) int
	Values() []uint32
}

// NewUInter32 returns a new UInter32 interface.
func NewUInter32(i ...uint32) UInter32 {
	return (&uinter32{&Slice{}}).Append(i...)
}

type uinter32 struct{ s *Slice }

func (u *uinter32) Append(i ...uint32) UInter32 {
	u.s.Append(uint32ToInterface(i...)...)
	return u
}

func (u *uinter32) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *uinter32) Concatenate(v UInter32) UInter32 {
	u.s.Concatenate(v.(*uinter32).s)
	return u
}

func (u *uinter32) Each(fn func(int, uint32)) UInter32 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint32)))
	})
	return u
}

func (u *uinter32) EachBreak(fn func(int, uint32) bool) UInter32 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint32)))
	})
	return u
}

func (u *uinter32) EachReverse(fn func(int, uint32)) UInter32 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint32)))
	})
	return u
}

func (u *uinter32) EachReverseBreak(fn func(int, uint32) bool) UInter32 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint32)))
	})
	return u
}

func (u *uinter32) Fetch(i int) uint32 {
	var s, _ = u.Get(i)
	return s
}

func (u *uinter32) Get(i int) (uint32, bool) {
	var (
		ok bool
		s  uint32
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint32)
	}
	return s, ok
}

func (u *uinter32) Len() int {
	return (u.s.Len())
}

func (u *uinter32) Less(i int, j int) bool {
	return i < j
}

func (u *uinter32) Map(fn func(int, uint32) uint32) UInter32 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint32)))
	})
	return u
}

func (u *uinter32) Poll() uint32 {
	var (
		s uint32
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint32))
	}
	return s
}

func (u *uinter32) Pop() uint32 {
	var (
		s uint32
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint32))
	}
	return s
}

func (u *uinter32) Precatenate(v UInter32) UInter32 {
	u.s.Precatenate(v.(*uinter32).s)
	return u
}

func (u *uinter32) Prepend(i ...uint32) UInter32 {
	u.s.Prepend(uint32ToInterface(i...)...)
	return u
}

func (u *uinter32) Push(i ...uint32) int {
	return u.s.Push(uint32ToInterface(i...))
}

func (u *uinter32) Replace(i int, n uint32) bool {
	return (u.s.Replace(i, n))
}

func (u *uinter32) Set() UInter32 {
	u.s.Set()
	return u
}

func (u *uinter32) Sort() UInter32 {
	sort.Sort(u)
	return u
}

func (u *uinter32) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *uinter32) Unshift(i ...uint32) int {
	return (u.s.Unshift(uint32ToInterface(i...)))
}

func (u *uinter32) Values() []uint32 {
	var v = make([]uint32, u.Len())
	u.Each(func(i int, n uint32) {
		v[i] = n
	})
	return v
}

func uint32ToInterface(n ...uint32) []interface{} {
	var (
		i int
		v uint32
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
