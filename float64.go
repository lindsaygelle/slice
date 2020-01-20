package slice

import "sort"

// Floater64 is the interface that handles a float64 collection.
type Floater64 interface {
	Append(...float64) Floater64
	Bounds(int) bool
	Concatenate(Floater64) Floater64
	Delete(int) Floater64
	Each(func(int, float64)) Floater64
	EachBreak(func(int, float64) bool) Floater64
	EachReverse(func(int, float64)) Floater64
	EachReverseBreak(func(int, float64) bool) Floater64
	Fetch(int) float64
	Get(int) (float64, bool)
	Len() int
	Less(int, int) bool
	Make(int) Floater64
	MakeEach(...float64) Floater64
	MakeEachReverse(...float64) Floater64
	Map(func(int, float64) float64) Floater64
	Poll() float64
	Pop() float64
	Precatenate(Floater64) Floater64
	Prepend(...float64) Floater64
	Push(...float64) int
	Replace(int, float64) bool
	Set() Floater64
	Slice(int, int) Floater64
	Sort() Floater64
	Swap(int, int)
	Unshift(...float64) int
	Values() []float64
}

// NewFloater64 returns a new Floater64 interface.
func NewFloater64(f ...float64) Floater64 {
	return (&floater64{&Slice{}}).Append(f...)
}

type floater64 struct{ s *Slice }

func (p *floater64) Append(f ...float64) Floater64 {
	p.s.Append(float64ToInterface(f...)...)
	return p
}

func (p *floater64) Bounds(i int) bool {
	return p.s.Bounds(i)
}

func (p *floater64) Concatenate(f Floater64) Floater64 {
	p.s.Concatenate(f.(*floater64).s)
	return p
}

func (p *floater64) Delete(i int) Floater64 {
	p.s.Delete(i)
	return p
}

func (p *floater64) Each(fn func(int, float64)) Floater64 {
	p.s.Each(func(i int, v interface{}) {
		fn(i, (v.(float64)))
	})
	return p
}

func (p *floater64) EachBreak(fn func(int, float64) bool) Floater64 {
	p.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float64)))
	})
	return p
}

func (p *floater64) EachReverse(fn func(int, float64)) Floater64 {
	p.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(float64)))
	})
	return p
}

func (p *floater64) EachReverseBreak(fn func(int, float64) bool) Floater64 {
	p.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float64)))
	})
	return p
}

func (p *floater64) Fetch(i int) float64 {
	var s, _ = p.Get(i)
	return s
}

func (p *floater64) Get(i int) (float64, bool) {
	var (
		ok bool
		s  float64
	)
	ok = p.Bounds(i)
	if ok {
		s = (p.s.Fetch(i)).(float64)
	}
	return s, ok
}

func (p *floater64) Len() int {
	return (p.s.Len())
}

func (p *floater64) Less(i int, j int) bool {
	return p.Fetch(i) < p.Fetch(j)
}

func (p *floater64) Make(i int) Floater64 {
	p.s.Make(i)
	return p
}

func (p *floater64) MakeEach(i ...float64) Floater64 {
	p.s.MakeEach(float64ToInterface(i...)...)
	return p
}

func (p *floater64) MakeEachReverse(i ...float64) Floater64 {
	p.s.MakeEachReverse(float64ToInterface(i...)...)
	return p
}

func (p *floater64) Map(fn func(int, float64) float64) Floater64 {
	p.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(float64)))
	})
	return p
}

func (p *floater64) Poll() float64 {
	var (
		s float64
		v = p.s.Poll()
	)
	if v != nil {
		s = (v.(float64))
	}
	return s
}

func (p *floater64) Pop() float64 {
	var (
		s float64
		v = p.s.Pop()
	)
	if v != nil {
		s = (v.(float64))
	}
	return s
}

func (p *floater64) Precatenate(f Floater64) Floater64 {
	p.s.Precatenate(f.(*floater64).s)
	return p
}

func (p *floater64) Prepend(f ...float64) Floater64 {
	p.s.Prepend(float64ToInterface(f...)...)
	return p
}

func (p *floater64) Push(f ...float64) int {
	return p.s.Push(float64ToInterface(f...))
}

func (p *floater64) Replace(i int, s float64) bool {
	return (p.s.Replace(i, s))
}

func (p *floater64) Set() Floater64 {
	p.s.Set()
	return p
}

func (p *floater64) Slice(i int, j int) Floater64 {
	p.s.Slice(i, j)
	return p
}

func (p *floater64) Sort() Floater64 {
	sort.Sort(p)
	return p
}

func (p *floater64) Swap(i int, j int) {
	p.s.Swap(i, j)
}

func (p *floater64) Unshift(f ...float64) int {
	return (p.s.Unshift(float64ToInterface(f...)))
}

func (p *floater64) Values() []float64 {
	var f = make([]float64, p.Len())
	p.Each(func(i int, s float64) {
		f[i] = s
	})
	return f
}

func float64ToInterface(f ...float64) []interface{} {
	var (
		i int
		v float64
		x = make([]interface{}, (len(f)))
	)
	for i, v = range f {
		x[i] = v
	}
	return x
}
