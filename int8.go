package slice

import "sort"

// Inter8 is the interface that handles a int8 collection.
type Inter8 interface {
	Append(...int8) Inter8
	Bounds(int) bool
	Concatenate(Inter8) Inter8
	Each(func(int, int8)) Inter8
	EachBreak(func(int, int8) bool) Inter8
	EachReverse(func(int, int8)) Inter8
	EachReverseBreak(func(int, int8) bool) Inter8
	Fetch(int) int8
	Get(int) (int8, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, int8) int8) Inter8
	Poll() int8
	Pop() int8
	Precatenate(Inter8) Inter8
	Prepend(...int8) Inter8
	Push(...int8) int
	Replace(int, int8) bool
	Set() Inter8
	Sort() Inter8
	Swap(int, int)
	Unshift(...int8) int
	Values() []int8
}

// NewInter8 returns a new Inter8 interface.
func NewInter8(i ...int8) Inter8 {
	return (&inter8{&Slice{}}).Append(i...)
}

type inter8 struct{ s *Slice }

func (i8 *inter8) Append(i ...int8) Inter8 {
	i8.s.Append(int8ToInterface(i...)...)
	return i8
}

func (i8 *inter8) Bounds(i int) bool {
	return i8.s.Bounds(i)
}

func (i8 *inter8) Concatenate(f Inter8) Inter8 {
	i8.s.Concatenate(f.(*inter8).s)
	return i8
}

func (i8 *inter8) Each(fn func(int, int8)) Inter8 {
	i8.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int8)))
	})
	return i8
}

func (i8 *inter8) EachBreak(fn func(int, int8) bool) Inter8 {
	i8.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int8)))
	})
	return i8
}

func (i8 *inter8) EachReverse(fn func(int, int8)) Inter8 {
	i8.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int8)))
	})
	return i8
}

func (i8 *inter8) EachReverseBreak(fn func(int, int8) bool) Inter8 {
	i8.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int8)))
	})
	return i8
}

func (i8 *inter8) Fetch(i int) int8 {
	var s, _ = i8.Get(i)
	return s
}

func (i8 *inter8) Get(i int) (int8, bool) {
	var (
		ok bool
		s  int8
	)
	ok = i8.Bounds(i)
	if ok {
		s = (i8.s.Fetch(i)).(int8)
	}
	return s, ok
}

func (i8 *inter8) Len() int {
	return (i8.s.Len())
}

func (i8 *inter8) Less(i int, j int) bool {
	return i8.Fetch(i) < i8.Fetch(j)
}

func (i8 *inter8) Map(fn func(int, int8) int8) Inter8 {
	i8.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int8)))
	})
	return i8
}

func (i8 *inter8) Poll() int8 {
	var (
		n int8
		v = i8.s.Poll()
	)
	if v != nil {
		n = (v.(int8))
	}
	return n
}

func (i8 *inter8) Pop() int8 {
	var (
		n int8
		v = i8.s.Pop()
	)
	if v != nil {
		n = (v.(int8))
	}
	return n
}

func (i8 *inter8) Precatenate(f Inter8) Inter8 {
	i8.s.Precatenate(f.(*inter8).s)
	return i8
}

func (i8 *inter8) Prepend(i ...int8) Inter8 {
	i8.s.Prepend(int8ToInterface(i...)...)
	return i8
}

func (i8 *inter8) Push(i ...int8) int {
	return i8.s.Push(int8ToInterface(i...))
}

func (i8 *inter8) Replace(i int, n int8) bool {
	return (i8.s.Replace(i, n))
}

func (i8 *inter8) Set() Inter8 {
	i8.s.Set()
	return i8
}

func (i8 *inter8) Sort() Inter8 {
	sort.Sort(i8)
	return i8
}

func (i8 *inter8) Swap(i int, j int) {
	i8.s.Swap(i, j)
}

func (i8 *inter8) Unshift(i ...int8) int {
	return (i8.s.Unshift(int8ToInterface(i...)))
}

func (i8 *inter8) Values() []int8 {
	var v = make([]int8, i8.Len())
	i8.Each(func(i int, n int8) {
		v[i] = n
	})
	return v
}

func int8ToInterface(n ...int8) []interface{} {
	var (
		i int
		v int8
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
