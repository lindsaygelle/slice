package slice

import "sort"

// Inter32 is the interface that handles a int32 collection.
type Inter32 interface {
	Append(...int32) Inter32
	Bounds(int) bool
	Concatenate(Inter32) Inter32
	Each(func(int, int32)) Inter32
	EachBreak(func(int, int32) bool) Inter32
	EachReverse(func(int, int32)) Inter32
	EachReverseBreak(func(int, int32) bool) Inter32
	Fetch(int) int32
	Get(int) (int32, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, int32) int32) Inter32
	Poll() int32
	Pop() int32
	Precatenate(Inter32) Inter32
	Prepend(...int32) Inter32
	Push(...int32) int
	Replace(int, int32) bool
	Set() Inter32
	Slice(int, int) Inter32
	Sort() Inter32
	Swap(int, int)
	Unshift(...int32) int
	Values() []int32
}

// NewInter32 returns a new Inter32 interface.
func NewInter32(i ...int32) Inter32 {
	return (&inter32{&Slice{}}).Append(i...)
}

type inter32 struct{ s *Slice }

func (i32 *inter32) Append(i ...int32) Inter32 {
	i32.s.Append(int32ToInterface(i...)...)
	return i32
}

func (i32 *inter32) Bounds(i int) bool {
	return i32.s.Bounds(i)
}

func (i32 *inter32) Concatenate(f Inter32) Inter32 {
	i32.s.Concatenate(f.(*inter32).s)
	return i32
}

func (i32 *inter32) Each(fn func(int, int32)) Inter32 {
	i32.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int32)))
	})
	return i32
}

func (i32 *inter32) EachBreak(fn func(int, int32) bool) Inter32 {
	i32.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int32)))
	})
	return i32
}

func (i32 *inter32) EachReverse(fn func(int, int32)) Inter32 {
	i32.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int32)))
	})
	return i32
}

func (i32 *inter32) EachReverseBreak(fn func(int, int32) bool) Inter32 {
	i32.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int32)))
	})
	return i32
}

func (i32 *inter32) Fetch(i int) int32 {
	var s, _ = i32.Get(i)
	return s
}

func (i32 *inter32) Get(i int) (int32, bool) {
	var (
		ok bool
		s  int32
	)
	ok = i32.Bounds(i)
	if ok {
		s = (i32.s.Fetch(i)).(int32)
	}
	return s, ok
}

func (i32 *inter32) Len() int {
	return (i32.s.Len())
}

func (i32 *inter32) Less(i int, j int) bool {
	return i32.Fetch(i) < i32.Fetch(j)
}

func (i32 *inter32) Map(fn func(int, int32) int32) Inter32 {
	i32.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int32)))
	})
	return i32
}

func (i32 *inter32) Poll() int32 {
	var (
		n int32
		v = i32.s.Poll()
	)
	if v != nil {
		n = (v.(int32))
	}
	return n
}

func (i32 *inter32) Pop() int32 {
	var (
		n int32
		v = i32.s.Pop()
	)
	if v != nil {
		n = (v.(int32))
	}
	return n
}

func (i32 *inter32) Precatenate(f Inter32) Inter32 {
	i32.s.Precatenate(f.(*inter32).s)
	return i32
}

func (i32 *inter32) Prepend(i ...int32) Inter32 {
	i32.s.Prepend(int32ToInterface(i...)...)
	return i32
}

func (i32 *inter32) Push(i ...int32) int {
	return i32.s.Push(int32ToInterface(i...))
}

func (i32 *inter32) Replace(i int, n int32) bool {
	return (i32.s.Replace(i, n))
}

func (i32 *inter32) Set() Inter32 {
	i32.s.Set()
	return i32
}

func (i32 *inter32) Slice(i int, j int) Inter32 {
	i32.s.Slice(i, j)
	return i32
}

func (i32 *inter32) Sort() Inter32 {
	sort.Sort(i32)
	return i32
}

func (i32 *inter32) Swap(i int, j int) {
	i32.s.Swap(i, j)
}

func (i32 *inter32) Unshift(i ...int32) int {
	return (i32.s.Unshift(int32ToInterface(i...)))
}

func (i32 *inter32) Values() []int32 {
	var v = make([]int32, i32.Len())
	i32.Each(func(i int, n int32) {
		v[i] = n
	})
	return v
}

func int32ToInterface(n ...int32) []interface{} {
	var (
		i int
		v int32
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
