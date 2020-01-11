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
	Sort() Inter
	Swap(int, int)
	Unshift(...int) int
	Values() []int
}

// NewInter returns a new Inter interface.
func NewInter(i ...int) Inter {
	return (&interger{&Slice{}}).Append(i...)
}

type interger struct{ s *Slice }

func (in *interger) Append(i ...int) Inter {
	in.s.Append(intsToInterface(i...)...)
	return in
}

func (in *interger) Bounds(i int) bool {
	return in.s.Bounds(i)
}

func (in *interger) Concatenate(s Inter) Inter {
	in.s.Concatenate(s.(*interger).s)
	return in
}

func (in *interger) Each(fn func(int, int)) Inter {
	in.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return in
}

func (in *interger) EachBreak(fn func(int, int) bool) Inter {
	in.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int)))
	})
	return in
}

func (in *interger) EachReverse(fn func(int, int)) Inter {
	in.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return in
}

func (in *interger) EachReverseBreak(fn func(int, int) bool) Inter {
	in.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int)))
	})
	return in
}

func (in *interger) Fetch(i int) int {
	var s, _ = in.Get(i)
	return s
}

func (in *interger) Get(i int) (int, bool) {
	var (
		ok bool
		s  int
	)
	ok = in.Bounds(i)
	if ok {
		s = (in.s.Fetch(i)).(int)
	}
	return s, ok
}

func (in *interger) Len() int {
	return (in.s.Len())
}

func (in *interger) Less(i int, j int) bool {
	return i < j
}

func (in *interger) Map(fn func(int, int) int) Inter {
	in.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int)))
	})
	return in
}

func (in *interger) Poll() int {
	var (
		s int
		v = in.s.Poll()
	)
	if v != nil {
		s = (v.(int))
	}
	return s
}

func (in *interger) Pop() int {
	var (
		s int
		v = in.s.Pop()
	)
	if v != nil {
		s = (v.(int))
	}
	return s
}

func (in *interger) Precatenate(s Inter) Inter {
	in.s.Precatenate(s.(*interger).s)
	return in
}

func (in *interger) Prepend(i ...int) Inter {
	in.s.Prepend(intsToInterface(i...)...)
	return in
}

func (in *interger) Push(i ...int) int {
	return in.s.Push(intsToInterface(i...))
}

func (in *interger) Replace(i int, s int) bool {
	return (in.s.Replace(i, s))
}

func (in *interger) Set() Inter {
	in.s.Set()
	return in
}

func (in *interger) Sort() Inter {
	sort.Sort(in)
	return in
}

func (in *interger) Swap(i int, j int) {
	in.s.Swap(i, j)
}

func (in *interger) Unshift(i ...int) int {
	return (in.s.Unshift(intsToInterface(i...)))
}

func (in *interger) Values() []int {
	var strs = make([]int, in.Len())
	in.Each(func(i int, s int) {
		strs[i] = s
	})
	return strs
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
