package slice

import (
	"sort"
)

// UInter64 is the interface that handles a uint64 collection.
type UInter64 interface {
	Append(...uint64) UInter64
	Bounds(int) bool
	Concatenate(UInter64) UInter64
	Delete(int) UInter64
	Each(func(int, uint64)) UInter64
	EachBreak(func(int, uint64) bool) UInter64
	EachReverse(func(int, uint64)) UInter64
	EachReverseBreak(func(int, uint64) bool) UInter64
	Fetch(int) uint64
	Get(int) (uint64, bool)
	Len() int
	Less(int, int) bool
	Make(int) UInter64
	MakeEach(...uint64) UInter64
	MakeEachReverse(...uint64) UInter64
	Map(func(int, uint64) uint64) UInter64
	Poll() uint64
	Pop() uint64
	Precatenate(UInter64) UInter64
	Prepend(...uint64) UInter64
	Push(...uint64) int
	Replace(int, uint64) bool
	Set() UInter64
	Slice(int, int) UInter64
	Sort() UInter64
	Swap(int, int)
	Unshift(...uint64) int
	Values() []uint64
}

// NewUInter64 returns a new UInter64 interface.
func NewUInter64(i ...uint64) UInter64 {
	return (&uinter64{&Slice{}}).Append(i...)
}

type uinter64 struct{ s *Slice }

func (u *uinter64) Append(i ...uint64) UInter64 {
	u.s.Append(uint64ToInterface(i...)...)
	return u
}

func (u *uinter64) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *uinter64) Concatenate(v UInter64) UInter64 {
	u.s.Concatenate(v.(*uinter64).s)
	return u
}

func (u *uinter64) Delete(i int) UInter64 {
	u.s.Delete(i)
	return u
}

func (u *uinter64) Each(fn func(int, uint64)) UInter64 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint64)))
	})
	return u
}

func (u *uinter64) EachBreak(fn func(int, uint64) bool) UInter64 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint64)))
	})
	return u
}

func (u *uinter64) EachReverse(fn func(int, uint64)) UInter64 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint64)))
	})
	return u
}

func (u *uinter64) EachReverseBreak(fn func(int, uint64) bool) UInter64 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint64)))
	})
	return u
}

func (u *uinter64) Fetch(i int) uint64 {
	var s, _ = u.Get(i)
	return s
}

func (u *uinter64) Get(i int) (uint64, bool) {
	var (
		ok bool
		s  uint64
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint64)
	}
	return s, ok
}

func (u *uinter64) Len() int {
	return (u.s.Len())
}

func (u *uinter64) Less(i int, j int) bool {
	return i < j
}

func (u *uinter64) Make(i int) UInter64 {
	u.s.Make(i)
	return u
}

func (u *uinter64) MakeEach(v ...uint64) UInter64 {
	u.s.MakeEach(uint64ToInterface(v...)...)
	return u
}

func (u *uinter64) MakeEachReverse(v ...uint64) UInter64 {
	u.s.MakeEachReverse(uint64ToInterface(v...)...)
	return u
}

func (u *uinter64) Map(fn func(int, uint64) uint64) UInter64 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint64)))
	})
	return u
}

func (u *uinter64) Poll() uint64 {
	var (
		s uint64
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint64))
	}
	return s
}

func (u *uinter64) Pop() uint64 {
	var (
		s uint64
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint64))
	}
	return s
}

func (u *uinter64) Precatenate(v UInter64) UInter64 {
	u.s.Precatenate(v.(*uinter64).s)
	return u
}

func (u *uinter64) Prepend(i ...uint64) UInter64 {
	u.s.Prepend(uint64ToInterface(i...)...)
	return u
}

func (u *uinter64) Push(i ...uint64) int {
	return u.s.Push(uint64ToInterface(i...))
}

func (u *uinter64) Replace(i int, n uint64) bool {
	return (u.s.Replace(i, n))
}

func (u *uinter64) Set() UInter64 {
	u.s.Set()
	return u
}

func (u *uinter64) Slice(i int, j int) UInter64 {
	u.s.Slice(i, j)
	return u
}

func (u *uinter64) Sort() UInter64 {
	sort.Sort(u)
	return u
}

func (u *uinter64) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *uinter64) Unshift(i ...uint64) int {
	return (u.s.Unshift(uint64ToInterface(i...)))
}

func (u *uinter64) Values() []uint64 {
	var v = make([]uint64, u.Len())
	u.Each(func(i int, n uint64) {
		v[i] = n
	})
	return v
}

func uint64ToInterface(n ...uint64) []interface{} {
	var (
		i int
		v uint64
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
