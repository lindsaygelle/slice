package slice

import (
	"sort"
)

// Integer is the interface that handles a int collection.
type Integer interface {
	Append(...int) Integer
	Bounds(int) bool
	Concatenate(Integer) Integer
	Each(func(int, int)) Integer
	EachBreak(func(int, int) bool) Integer
	EachReverse(func(int, int)) Integer
	EachReverseBreak(func(int, int) bool) Integer
	Fetch(int) int
	Get(int) (int, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, int) int) Integer
	Poll() int
	Pop() int
	Precatenate(Integer) Integer
	Prepend(...int) Integer
	Push(...int) int
	Replace(int, int) bool
	Set() Integer
	Sort() Integer
	Swap(int, int)
	Unshift(...int) int
	Values() []int
}

// NewInteger returns a new Integer interface.
func NewInteger() Integer {
	return (&interger{s: &Slice{}})
}

type interger struct{ s *Slice }

func (in *interger) Append(i ...int) Integer {
	in.s.Append(intsToInterface(i...)...)
	return in
}

func (in *interger) Bounds(i int) bool {
	return in.s.Bounds(i)
}

func (in *interger) Concatenate(s Integer) Integer {
	in.s.Concatenate(s.(*interger).s)
	return in
}

func (in *interger) Each(fn func(int, int)) Integer {
	in.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return in
}

func (in *interger) EachBreak(fn func(int, int) bool) Integer {
	in.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int)))
	})
	return in
}

func (in *interger) EachReverse(fn func(int, int)) Integer {
	in.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return in
}

func (in *interger) EachReverseBreak(fn func(int, int) bool) Integer {
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

func (in *interger) Map(fn func(int, int) int) Integer {
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

func (in *interger) Precatenate(s Integer) Integer {
	in.s.Precatenate(s.(*interger).s)
	return in
}

func (in *interger) Prepend(i ...int) Integer {
	in.s.Prepend(intsToInterface(i...)...)
	return in
}

func (in *interger) Push(i ...int) int {
	return in.s.Push(intsToInterface(i...))
}

func (in *interger) Replace(i int, s int) bool {
	return (in.s.Replace(i, s))
}

func (in *interger) Set() Integer {
	in.s.Set()
	return in
}

func (in *interger) Sort() Integer {
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
