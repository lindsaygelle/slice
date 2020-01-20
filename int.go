package slice

import (
	"sort"
)

// Inter is the interface that handles a int collection.
type Inter interface {
	Append(...int) Inter
	Bounds(int) bool
	Concatenate(Inter) Inter
	Delete(int) Inter
	Each(func(int, int)) Inter
	EachBreak(func(int, int) bool) Inter
	EachReverse(func(int, int)) Inter
	EachReverseBreak(func(int, int) bool) Inter
	Fetch(int) int
	Get(int) (int, bool)
	Len() int
	Less(int, int) bool
	Make(int) Inter
	MakeEach(...int) Inter
	MakeEachReverse(...int) Inter
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

func (p *inter) Append(i ...int) Inter {
	p.s.Append(intToInterface(i...)...)
	return p
}

func (p *inter) Bounds(i int) bool {
	return p.s.Bounds(i)
}

func (p *inter) Concatenate(v Inter) Inter {
	p.s.Concatenate(v.(*inter).s)
	return p
}

func (p *inter) Delete(i int) Inter {
	p.s.Delete(i)
	return p
}

func (p *inter) Each(fn func(int, int)) Inter {
	p.s.Each(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return p
}

func (p *inter) EachBreak(fn func(int, int) bool) Inter {
	p.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int)))
	})
	return p
}

func (p *inter) EachReverse(fn func(int, int)) Inter {
	p.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(int)))
	})
	return p
}

func (p *inter) EachReverseBreak(fn func(int, int) bool) Inter {
	p.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(int)))
	})
	return p
}

func (p *inter) Fetch(i int) int {
	var n, _ = p.Get(i)
	return n
}

func (p *inter) Get(i int) (int, bool) {
	var (
		ok bool
		n  int
	)
	ok = p.Bounds(i)
	if ok {
		n = (p.s.Fetch(i)).(int)
	}
	return n, ok
}

func (p *inter) Len() int {
	return (p.s.Len())
}

func (p *inter) Less(i int, j int) bool {
	return p.Fetch(i) < p.Fetch(j)
}

func (p *inter) Make(i int) Inter {
	p.s.Make(i)
	return p
}

func (p *inter) MakeEach(v ...int) Inter {
	p.s.MakeEach(intToInterface(v...)...)
	return p
}

func (p *inter) MakeEachReverse(v ...int) Inter {
	p.s.MakeEachReverse(intToInterface(v...)...)
	return p
}

func (p *inter) Map(fn func(int, int) int) Inter {
	p.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(int)))
	})
	return p
}

func (p *inter) Poll() int {
	var (
		n int
		v = p.s.Poll()
	)
	if v != nil {
		n = (v.(int))
	}
	return n
}

func (p *inter) Pop() int {
	var (
		n int
		v = p.s.Pop()
	)
	if v != nil {
		n = (v.(int))
	}
	return n
}

func (p *inter) Precatenate(v Inter) Inter {
	p.s.Precatenate(v.(*inter).s)
	return p
}

func (p *inter) Prepend(i ...int) Inter {
	p.s.Prepend(intToInterface(i...)...)
	return p
}

func (p *inter) Push(i ...int) int {
	return p.s.Push(intToInterface(i...))
}

func (p *inter) Replace(i int, n int) bool {
	return (p.s.Replace(i, n))
}

func (p *inter) Set() Inter {
	p.s.Set()
	return p
}

func (p *inter) Slice(i int, j int) Inter {
	p.s.Slice(i, j)
	return p
}

func (p *inter) Sort() Inter {
	sort.Sort(p)
	return p
}

func (p *inter) Swap(i int, j int) {
	p.s.Swap(i, j)
}

func (p *inter) Unshift(i ...int) int {
	return (p.s.Unshift(intToInterface(i...)))
}

func (p *inter) Values() []int {
	var v = make([]int, p.Len())
	p.Each(func(i int, n int) {
		v[i] = n
	})
	return v
}

func intToInterface(n ...int) []interface{} {
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
