package slice

import (
	"sort"
)

// Complexer64 is the interface that handles a complex64 collection.
type Complexer64 interface {
	Append(...complex64) Complexer64
	Bounds(int) bool
	Concatenate(Complexer64) Complexer64
	Delete(int) Complexer64
	Each(func(int, complex64)) Complexer64
	EachBreak(func(int, complex64) bool) Complexer64
	EachReverse(func(int, complex64)) Complexer64
	EachReverseBreak(func(int, complex64) bool) Complexer64
	Fetch(int) complex64
	Get(int) (complex64, bool)
	Len() int
	Less(int, int) bool
	Make(int) Complexer64
	MakeEach(...complex64) Complexer64
	MakeEachReverse(...complex64) Complexer64
	Map(func(int, complex64) complex64) Complexer64
	Poll() complex64
	Pop() complex64
	Precatenate(Complexer64) Complexer64
	Prepend(...complex64) Complexer64
	Push(...complex64) int
	Replace(int, complex64) bool
	Set() Complexer64
	Slice(int, int) Complexer64
	Sort() Complexer64
	Swap(int, int)
	Unshift(...complex64) int
	Values() []complex64
}

// NewComplexer returns a new Complexer64 interface.
func NewComplexer(i ...complex64) Complexer64 {
	return (&complexer64{&Slice{}}).Append(i...)
}

type complexer64 struct{ s *Slice }

func (p *complexer64) Append(i ...complex64) Complexer64 {
	p.s.Append(complex64ToInterface(i...)...)
	return p
}

func (p *complexer64) Bounds(i int) bool {
	return p.s.Bounds(i)
}

func (p *complexer64) Concatenate(v Complexer64) Complexer64 {
	p.s.Concatenate(v.(*complexer64).s)
	return p
}

func (p *complexer64) Delete(i int) Complexer64 {
	p.s.Delete(i)
	return p
}

func (p *complexer64) Each(fn func(int, complex64)) Complexer64 {
	p.s.Each(func(i int, v interface{}) {
		fn(i, (v.(complex64)))
	})
	return p
}

func (p *complexer64) EachBreak(fn func(int, complex64) bool) Complexer64 {
	p.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(complex64)))
	})
	return p
}

func (p *complexer64) EachReverse(fn func(int, complex64)) Complexer64 {
	p.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(complex64)))
	})
	return p
}

func (p *complexer64) EachReverseBreak(fn func(int, complex64) bool) Complexer64 {
	p.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(complex64)))
	})
	return p
}

func (p *complexer64) Fetch(i int) complex64 {
	var s, _ = p.Get(i)
	return s
}

func (p *complexer64) Get(i int) (complex64, bool) {
	var (
		ok bool
		s  complex64
	)
	ok = p.Bounds(i)
	if ok {
		s = (p.s.Fetch(i)).(complex64)
	}
	return s, ok
}

func (p *complexer64) Len() int {
	return (p.s.Len())
}

func (p *complexer64) Less(i int, j int) bool {
	return ((p.s.Fetch(i).(float64)) < p.s.Fetch(j).(float64))
}

func (p *complexer64) Make(i int) Complexer64 {
	p.s.Make(i)
	return p
}

func (p *complexer64) MakeEach(i ...complex64) Complexer64 {
	p.s.MakeEach(complex64ToInterface(i...)...)
	return p
}

func (p *complexer64) MakeEachReverse(i ...complex64) Complexer64 {
	p.s.MakeEachReverse(complex64ToInterface(i...)...)
	return p
}

func (p *complexer64) Map(fn func(int, complex64) complex64) Complexer64 {
	p.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(complex64)))
	})
	return p
}

func (p *complexer64) Poll() complex64 {
	var (
		s complex64
		v = p.s.Poll()
	)
	if v != nil {
		s = (v.(complex64))
	}
	return s
}

func (p *complexer64) Pop() complex64 {
	var (
		s complex64
		v = p.s.Pop()
	)
	if v != nil {
		s = (v.(complex64))
	}
	return s
}

func (p *complexer64) Precatenate(v Complexer64) Complexer64 {
	p.s.Precatenate(v.(*complexer64).s)
	return p
}

func (p *complexer64) Prepend(i ...complex64) Complexer64 {
	p.s.Prepend(complex64ToInterface(i...)...)
	return p
}

func (p *complexer64) Push(i ...complex64) int {
	return p.s.Push(complex64ToInterface(i...))
}

func (p *complexer64) Replace(i int, n complex64) bool {
	return (p.s.Replace(i, n))
}

func (p *complexer64) Set() Complexer64 {
	p.s.Set()
	return p
}

func (p *complexer64) Slice(i int, j int) Complexer64 {
	p.s.Slice(i, j)
	return p
}

func (p *complexer64) Sort() Complexer64 {
	sort.Sort(p)
	return p
}

func (p *complexer64) Swap(i int, j int) {
	p.s.Swap(i, j)
}

func (p *complexer64) Unshift(i ...complex64) int {
	return (p.s.Unshift(complex64ToInterface(i...)))
}

func (p *complexer64) Values() []complex64 {
	var v = make([]complex64, p.Len())
	p.Each(func(i int, n complex64) {
		v[i] = n
	})
	return v
}

func complex64ToInterface(n ...complex64) []interface{} {
	var (
		i int
		v complex64
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
