package slice

import "sort"

// Inter64 is the interface that handles a int64 collection.
type Inter64 interface {
	Append(...int64) Inter64
	Bounds(int) bool
	Concatenate(Inter64) Inter64
	Each(func(int, int64)) Inter64
	EachBreak(func(int, int64) bool) Inter64
	EachReverse(func(int, int64)) Inter64
	EachReverseBreak(func(int, int64) bool) Inter64
	Fetch(int) int64
	Get(int) (int64, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, int64) int64) Inter64
	Poll() int64
	Pop() int64
	Precatenate(Inter64) Inter64
	Prepend(...int64) Inter64
	Push(...int64) int
	Replace(int, int64) bool
	Set() Inter64
	Slice(int, int) Inter64
	Sort() Inter64
	Swap(int, int)
	Unshift(...int64) int
	Values() []int64
}

// NewInter64 returns a new Inter64 interface.
func NewInter64(i ...int64) Inter64 {
	return (&inter64{&Slice{}}).Append(i...)
}

type inter64 struct{ s *Slice }

func (i64 *inter64) Append(i ...int64) Inter64 {
	i64.s.Append(int64ToInterface(i...)...)
	return i64
}

func (i64 *inter64) Bounds(i int) bool {
	return i64.s.Bounds(i)
}

func (i64 *inter64) Concatenate(f Inter64) Inter64 {
	i64.s.Concatenate(f.(*inter64).s)
	return i64
}

func (i64 *inter64) Each(fn func(int, int64)) Inter64 {
	i64.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int64)))
	})
	return i64
}

func (i64 *inter64) EachBreak(fn func(int, int64) bool) Inter64 {
	i64.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int64)))
	})
	return i64
}

func (i64 *inter64) EachReverse(fn func(int, int64)) Inter64 {
	i64.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int64)))
	})
	return i64
}

func (i64 *inter64) EachReverseBreak(fn func(int, int64) bool) Inter64 {
	i64.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int64)))
	})
	return i64
}

func (i64 *inter64) Fetch(i int) int64 {
	var s, _ = i64.Get(i)
	return s
}

func (i64 *inter64) Get(i int) (int64, bool) {
	var (
		ok bool
		s  int64
	)
	ok = i64.Bounds(i)
	if ok {
		s = (i64.s.Fetch(i)).(int64)
	}
	return s, ok
}

func (i64 *inter64) Len() int {
	return (i64.s.Len())
}

func (i64 *inter64) Less(i int, j int) bool {
	return i64.Fetch(i) < i64.Fetch(j)
}

func (i64 *inter64) Map(fn func(int, int64) int64) Inter64 {
	i64.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int64)))
	})
	return i64
}

func (i64 *inter64) Poll() int64 {
	var (
		n int64
		v = i64.s.Poll()
	)
	if v != nil {
		n = (v.(int64))
	}
	return n
}

func (i64 *inter64) Pop() int64 {
	var (
		n int64
		v = i64.s.Pop()
	)
	if v != nil {
		n = (v.(int64))
	}
	return n
}

func (i64 *inter64) Precatenate(f Inter64) Inter64 {
	i64.s.Precatenate(f.(*inter64).s)
	return i64
}

func (i64 *inter64) Prepend(i ...int64) Inter64 {
	i64.s.Prepend(int64ToInterface(i...)...)
	return i64
}

func (i64 *inter64) Push(i ...int64) int {
	return i64.s.Push(int64ToInterface(i...))
}

func (i64 *inter64) Replace(i int, n int64) bool {
	return (i64.s.Replace(i, n))
}

func (i64 *inter64) Set() Inter64 {
	i64.s.Set()
	return i64
}

func (i64 *inter64) Slice(i int, j int) Inter64 {
	i64.s.Slice(i, j)
	return i64
}

func (i64 *inter64) Sort() Inter64 {
	sort.Sort(i64)
	return i64
}

func (i64 *inter64) Swap(i int, j int) {
	i64.s.Swap(i, j)
}

func (i64 *inter64) Unshift(i ...int64) int {
	return (i64.s.Unshift(int64ToInterface(i...)))
}

func (i64 *inter64) Values() []int64 {
	var v = make([]int64, i64.Len())
	i64.Each(func(i int, n int64) {
		v[i] = n
	})
	return v
}

func int64ToInterface(n ...int64) []interface{} {
	var (
		i int
		v int64
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
