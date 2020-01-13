package slice

import (
	"sort"
)

// Inter is the interface that handles a int collection.
type Inter interface {
	Append(...int) Inter
	Bounds(int) bool
	Concatenate(Inter) Inter
	Each(func(int, int)) Inter
	EachBreak(func(int, int) bool) Inter
	EachReverse(func(int, int)) Inter
	EachReverseBreak(func(int, int) bool) Inter
	Fetch(int) int
	Get(int) (int, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, int) int) Inter
	Poll() int
	Pop() int
	Precatenate(Inter) Inter
	Prepend(...int) Inter
	Push(...int) int
	Replace(int, int) bool
	Set() Inter
	Slice(int, int) Inter
	Sort() Inter
	Swap(int, int)
	Unshift(...int) int
	Values() []int
}

// NewInter returns a new Inter interface.
func NewInter(i ...int) Inter {
	return (&inter{&Slice{}}).Append(i...)
}

type inter struct{ s *Slice }

func (in *inter) Append(i ...int) Inter {
	in.s.Append(intsToInterface(i...)...)
	return in
}

func (in *inter) Bounds(i int) bool {
	return in.s.Bounds(i)
}

func (in *inter) Concatenate(v Inter) Inter {
	in.s.Concatenate(v.(*inter).s)
	return in
}

func (in *inter) Each(fn func(int, int)) Inter {
	in.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return in
}

func (in *inter) EachBreak(fn func(int, int) bool) Inter {
	in.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int)))
	})
	return in
}

func (in *inter) EachReverse(fn func(int, int)) Inter {
	in.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return in
}

func (in *inter) EachReverseBreak(fn func(int, int) bool) Inter {
	in.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int)))
	})
	return in
}

func (in *inter) Fetch(i int) int {
	var n, _ = in.Get(i)
	return n
}

func (in *inter) Get(i int) (int, bool) {
	var (
		ok bool
		n  int
	)
	ok = in.Bounds(i)
	if ok {
		n = (in.s.Fetch(i)).(int)
	}
	return n, ok
}

func (in *inter) Len() int {
	return (in.s.Len())
}

func (in *inter) Less(i int, j int) bool {
	return in.Fetch(i) < in.Fetch(j)
}

func (in *inter) Map(fn func(int, int) int) Inter {
	in.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int)))
	})
	return in
}

func (in *inter) Poll() int {
	var (
		n int
		v = in.s.Poll()
	)
	if v != nil {
		n = (v.(int))
	}
	return n
}

func (in *inter) Pop() int {
	var (
		n int
		v = in.s.Pop()
	)
	if v != nil {
		n = (v.(int))
	}
	return n
}

func (in *inter) Precatenate(v Inter) Inter {
	in.s.Precatenate(v.(*inter).s)
	return in
}

func (in *inter) Prepend(i ...int) Inter {
	in.s.Prepend(intsToInterface(i...)...)
	return in
}

func (in *inter) Push(i ...int) int {
	return in.s.Push(intsToInterface(i...))
}

func (in *inter) Replace(i int, n int) bool {
	return (in.s.Replace(i, n))
}

func (in *inter) Set() Inter {
	in.s.Set()
	return in
}

func (in *inter) Slice(i int, j int) Inter {
	in.s.Slice(i, j)
	return in
}

func (in *inter) Sort() Inter {
	sort.Sort(in)
	return in
}

func (in *inter) Swap(i int, j int) {
	in.s.Swap(i, j)
}

func (in *inter) Unshift(i ...int) int {
	return (in.s.Unshift(intsToInterface(i...)))
}

func (in *inter) Values() []int {
	var v = make([]int, in.Len())
	in.Each(func(i int, n int) {
		v[i] = n
	})
	return v
}

func intsToInterface(n ...int) []interface{} {
	var (
		i int
		v int
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
