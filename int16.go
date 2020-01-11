package slice

import "sort"

// Inter16 is the interface that handles a int16 collection.
type Inter16 interface {
	Append(...int16) Inter16
	Bounds(int) bool
	Concatenate(Inter16) Inter16
	Each(func(int, int16)) Inter16
	EachBreak(func(int, int16) bool) Inter16
	EachReverse(func(int, int16)) Inter16
	EachReverseBreak(func(int, int16) bool) Inter16
	Fetch(int) int16
	Get(int) (int16, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, int16) int16) Inter16
	Poll() int16
	Pop() int16
	Precatenate(Inter16) Inter16
	Prepend(...int16) Inter16
	Push(...int16) int
	Replace(int, int16) bool
	Set() Inter16
	Sort() Inter16
	Swap(int, int)
	Unshift(...int16) int
	Values() []int16
}

// NewInter16 returns a new Inter16 interface.
func NewInter16(i ...int16) Inter16 {
	return (&inter16{&Slice{}}).Append(i...)
}

type inter16 struct{ s *Slice }

func (i16 *inter16) Append(i ...int16) Inter16 {
	i16.s.Append(int16ToInterface(i...)...)
	return i16
}

func (i16 *inter16) Bounds(i int) bool {
	return i16.s.Bounds(i)
}

func (i16 *inter16) Concatenate(f Inter16) Inter16 {
	i16.s.Concatenate(f.(*inter16).s)
	return i16
}

func (i16 *inter16) Each(fn func(int, int16)) Inter16 {
	i16.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int16)))
	})
	return i16
}

func (i16 *inter16) EachBreak(fn func(int, int16) bool) Inter16 {
	i16.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int16)))
	})
	return i16
}

func (i16 *inter16) EachReverse(fn func(int, int16)) Inter16 {
	i16.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int16)))
	})
	return i16
}

func (i16 *inter16) EachReverseBreak(fn func(int, int16) bool) Inter16 {
	i16.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int16)))
	})
	return i16
}

func (i16 *inter16) Fetch(i int) int16 {
	var s, _ = i16.Get(i)
	return s
}

func (i16 *inter16) Get(i int) (int16, bool) {
	var (
		ok bool
		s  int16
	)
	ok = i16.Bounds(i)
	if ok {
		s = (i16.s.Fetch(i)).(int16)
	}
	return s, ok
}

func (i16 *inter16) Len() int {
	return (i16.s.Len())
}

func (i16 *inter16) Less(i int, j int) bool {
	return i16.Fetch(i) < i16.Fetch(j)
}

func (i16 *inter16) Map(fn func(int, int16) int16) Inter16 {
	i16.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int16)))
	})
	return i16
}

func (i16 *inter16) Poll() int16 {
	var (
		n int16
		v = i16.s.Poll()
	)
	if v != nil {
		n = (v.(int16))
	}
	return n
}

func (i16 *inter16) Pop() int16 {
	var (
		n int16
		v = i16.s.Pop()
	)
	if v != nil {
		n = (v.(int16))
	}
	return n
}

func (i16 *inter16) Precatenate(f Inter16) Inter16 {
	i16.s.Precatenate(f.(*inter16).s)
	return i16
}

func (i16 *inter16) Prepend(i ...int16) Inter16 {
	i16.s.Prepend(int16ToInterface(i...)...)
	return i16
}

func (i16 *inter16) Push(i ...int16) int {
	return i16.s.Push(int16ToInterface(i...))
}

func (i16 *inter16) Replace(i int, n int16) bool {
	return (i16.s.Replace(i, n))
}

func (i16 *inter16) Set() Inter16 {
	i16.s.Set()
	return i16
}

func (i16 *inter16) Sort() Inter16 {
	sort.Sort(i16)
	return i16
}

func (i16 *inter16) Swap(i int, j int) {
	i16.s.Swap(i, j)
}

func (i16 *inter16) Unshift(i ...int16) int {
	return (i16.s.Unshift(int16ToInterface(i...)))
}

func (i16 *inter16) Values() []int16 {
	var v = make([]int16, i16.Len())
	i16.Each(func(i int, n int16) {
		v[i] = n
	})
	return v
}

func int16ToInterface(n ...int16) []interface{} {
	var (
		i int
		v int16
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
