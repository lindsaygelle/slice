package slice

import (
	"sort"
)

// Complexer128 is the interface that handles a complex128 collection.
type Complexer128 interface {
	Append(...complex128) Complexer128
	Bounds(int) bool
	Concatenate(Complexer128) Complexer128
	Delete(int) Complexer128
	Each(func(int, complex128)) Complexer128
	EachBreak(func(int, complex128) bool) Complexer128
	EachReverse(func(int, complex128)) Complexer128
	EachReverseBreak(func(int, complex128) bool) Complexer128
	Fetch(int) complex128
	Get(int) (complex128, bool)
	Len() int
	Less(int, int) bool
	Make(int) Complexer128
	MakeEach(...complex128) Complexer128
	MakeEachReverse(...complex128) Complexer128
	Map(func(int, complex128) complex128) Complexer128
	Poll() complex128
	Pop() complex128
	Precatenate(Complexer128) Complexer128
	Prepend(...complex128) Complexer128
	Push(...complex128) int
	Replace(int, complex128) bool
	Set() Complexer128
	Slice(int, int) Complexer128
	Sort() Complexer128
	Swap(int, int)
	Unshift(...complex128) int
	Values() []complex128
}

// NewComplexer128 returns a new Complexer128 interface.
func NewComplexer128(i ...complex128) Complexer128 {
	return (&complexer128{&Slice{}}).Append(i...)
}

type complexer128 struct{ s *Slice }

func (p *complexer128) Append(i ...complex128) Complexer128 {
	p.s.Append(complex128ToInterface(i...)...)
	return p
}

func (p *complexer128) Bounds(i int) bool {
	return p.s.Bounds(i)
}

func (p *complexer128) Concatenate(v Complexer128) Complexer128 {
	p.s.Concatenate(v.(*complexer128).s)
	return p
}

func (p *complexer128) Delete(i int) Complexer128 {
	p.s.Delete(i)
	return p
}

func (p *complexer128) Each(fn func(int, complex128)) Complexer128 {
	p.s.Each(func(i int, v interface{}) {
		fn(i, (v.(complex128)))
	})
	return p
}

func (p *complexer128) EachBreak(fn func(int, complex128) bool) Complexer128 {
	p.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(complex128)))
	})
	return p
}

func (p *complexer128) EachReverse(fn func(int, complex128)) Complexer128 {
	p.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(complex128)))
	})
	return p
}

func (p *complexer128) EachReverseBreak(fn func(int, complex128) bool) Complexer128 {
	p.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(complex128)))
	})
	return p
}

func (p *complexer128) Fetch(i int) complex128 {
	var s, _ = p.Get(i)
	return s
}

func (p *complexer128) Get(i int) (complex128, bool) {
	var (
		ok bool
		s  complex128
	)
	ok = p.Bounds(i)
	if ok {
		s = (p.s.Fetch(i)).(complex128)
	}
	return s, ok
}

func (p *complexer128) Len() int {
	return (p.s.Len())
}

func (p *complexer128) Less(i int, j int) bool {
	return ((p.s.Fetch(i).(float64)) < p.s.Fetch(j).(float64))
}

func (p *complexer128) Make(i int) Complexer128 {
	p.s.Make(i)
	return p
}

func (p *complexer128) MakeEach(i ...complex128) Complexer128 {
	p.s.MakeEach(complex128ToInterface(i...)...)
	return p
}

func (p *complexer128) MakeEachReverse(i ...complex128) Complexer128 {
	p.s.MakeEachReverse(complex128ToInterface(i...)...)
	return p
}

func (p *complexer128) Map(fn func(int, complex128) complex128) Complexer128 {
	p.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(complex128)))
	})
	return p
}

func (p *complexer128) Poll() complex128 {
	var (
		s complex128
		v = p.s.Poll()
	)
	if v != nil {
		s = (v.(complex128))
	}
	return s
}

func (p *complexer128) Pop() complex128 {
	var (
		s complex128
		v = p.s.Pop()
	)
	if v != nil {
		s = (v.(complex128))
	}
	return s
}

func (p *complexer128) Precatenate(v Complexer128) Complexer128 {
	p.s.Precatenate(v.(*complexer128).s)
	return p
}

func (p *complexer128) Prepend(i ...complex128) Complexer128 {
	p.s.Prepend(complex128ToInterface(i...)...)
	return p
}

func (p *complexer128) Push(i ...complex128) int {
	return p.s.Push(complex128ToInterface(i...))
}

func (p *complexer128) Replace(i int, n complex128) bool {
	return (p.s.Replace(i, n))
}

func (p *complexer128) Set() Complexer128 {
	p.s.Set()
	return p
}

func (p *complexer128) Slice(i int, j int) Complexer128 {
	p.s.Slice(i, j)
	return p
}

func (p *complexer128) Sort() Complexer128 {
	sort.Sort(p)
	return p
}

func (p *complexer128) Swap(i int, j int) {
	p.s.Swap(i, j)
}

func (p *complexer128) Unshift(i ...complex128) int {
	return (p.s.Unshift(complex128ToInterface(i...)))
}

func (p *complexer128) Values() []complex128 {
	var v = make([]complex128, p.Len())
	p.Each(func(i int, n complex128) {
		v[i] = n
	})
	return v
}

func complex128ToInterface(n ...complex128) []interface{} {
	var (
		i int
		v complex128
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
