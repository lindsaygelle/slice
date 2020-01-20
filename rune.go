package slice

import (
	"sort"
)

// Runer is the interface that handles a rune collection.
type Runer interface {
	Append(...rune) Runer
	Bounds(int) bool
	Concatenate(Runer) Runer
	Delete(int) Runer
	Each(func(int, rune)) Runer
	EachBreak(func(int, rune) bool) Runer
	EachReverse(func(int, rune)) Runer
	EachReverseBreak(func(int, rune) bool) Runer
	Fetch(int) rune
	Get(int) (rune, bool)
	Len() int
	Less(int, int) bool
	Make(int) Runer
	MakeEach(...rune) Runer
	MakeEachReverse(...rune) Runer
	Map(func(int, rune) rune) Runer
	Poll() rune
	Pop() rune
	Precatenate(Runer) Runer
	Prepend(...rune) Runer
	Push(...rune) int
	Replace(int, rune) bool
	Set() Runer
	Slice(int, int) Runer
	Sort() Runer
	Swap(int, int)
	Unshift(...rune) int
	Values() []rune
}

// NewRuner returns a new Runer interface.
func NewRuner(i ...rune) Runer {
	return (&runer{&Slice{}}).Append(i...)
}

type runer struct{ s *Slice }

func (p *runer) Append(i ...rune) Runer {
	p.s.Append(runeToInterface(i...)...)
	return p
}

func (p *runer) Bounds(i int) bool {
	return p.s.Bounds(i)
}

func (p *runer) Concatenate(v Runer) Runer {
	p.s.Concatenate(v.(*runer).s)
	return p
}

func (p *runer) Delete(i int) Runer {
	p.s.Delete(i)
	return p
}

func (p *runer) Each(fn func(int, rune)) Runer {
	p.s.Each(func(i int, v interface{}) {
		fn(i, (v.(rune)))
	})
	return p
}

func (p *runer) EachBreak(fn func(int, rune) bool) Runer {
	p.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(rune)))
	})
	return p
}

func (p *runer) EachReverse(fn func(int, rune)) Runer {
	p.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(rune)))
	})
	return p
}

func (p *runer) EachReverseBreak(fn func(int, rune) bool) Runer {
	p.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(rune)))
	})
	return p
}

func (p *runer) Fetch(i int) rune {
	var s, _ = p.Get(i)
	return s
}

func (p *runer) Get(i int) (rune, bool) {
	var (
		ok bool
		s  rune
	)
	ok = p.Bounds(i)
	if ok {
		s = (p.s.Fetch(i)).(rune)
	}
	return s, ok
}

func (p *runer) Len() int {
	return (p.s.Len())
}

func (p *runer) Less(i int, j int) bool {
	return p.Fetch(i) < p.Fetch(j)
}

func (p *runer) Make(i int) Runer {
	p.s.Make(i)
	return p
}

func (p *runer) MakeEach(i ...rune) Runer {
	p.s.MakeEach(runeToInterface(i...)...)
	return p
}

func (p *runer) MakeEachReverse(i ...rune) Runer {
	p.s.MakeEachReverse(runeToInterface(i...)...)
	return p
}

func (p *runer) Map(fn func(int, rune) rune) Runer {
	p.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(rune)))
	})
	return p
}

func (p *runer) Poll() rune {
	var (
		s rune
		v = p.s.Poll()
	)
	if v != nil {
		s = (v.(rune))
	}
	return s
}

func (p *runer) Pop() rune {
	var (
		s rune
		v = p.s.Pop()
	)
	if v != nil {
		s = (v.(rune))
	}
	return s
}

func (p *runer) Precatenate(v Runer) Runer {
	p.s.Precatenate(v.(*runer).s)
	return p
}

func (p *runer) Prepend(i ...rune) Runer {
	p.s.Prepend(runeToInterface(i...)...)
	return p
}

func (p *runer) Push(i ...rune) int {
	return p.s.Push(runeToInterface(i...))
}

func (p *runer) Replace(i int, n rune) bool {
	return (p.s.Replace(i, n))
}

func (p *runer) Set() Runer {
	p.s.Set()
	return p
}

func (p *runer) Slice(i int, j int) Runer {
	p.s.Slice(i, j)
	return p
}

func (p *runer) Sort() Runer {
	sort.Sort(p)
	return p
}

func (p *runer) Swap(i int, j int) {
	p.s.Swap(i, j)
}

func (p *runer) Unshift(i ...rune) int {
	return (p.s.Unshift(runeToInterface(i...)))
}

func (p *runer) Values() []rune {
	var v = make([]rune, p.Len())
	p.Each(func(i int, n rune) {
		v[i] = n
	})
	return v
}

func runeToInterface(n ...rune) []interface{} {
	var (
		i int
		v rune
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
