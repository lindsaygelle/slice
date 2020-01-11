package slice

import "sort"

// Floater64 is the interface that handles a float64 collection.
type Floater64 interface {
	Append(...float64) Floater64
	Bounds(int) bool
	Concatenate(Floater64) Floater64
	Each(func(int, float64)) Floater64
	EachBreak(func(int, float64) bool) Floater64
	EachReverse(func(int, float64)) Floater64
	EachReverseBreak(func(int, float64) bool) Floater64
	Fetch(int) float64
	Get(int) (float64, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, float64) float64) Floater64
	Poll() float64
	Pop() float64
	Precatenate(Floater64) Floater64
	Prepend(...float64) Floater64
	Push(...float64) int
	Replace(int, float64) bool
	Set() Floater64
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

func (f64 *floater64) Append(f ...float64) Floater64 {
	f64.s.Append(float64ToInterface(f...)...)
	return f64
}

func (f64 *floater64) Bounds(i int) bool {
	return f64.s.Bounds(i)
}

func (f64 *floater64) Concatenate(f Floater64) Floater64 {
	f64.s.Concatenate(f.(*floater64).s)
	return f64
}

func (f64 *floater64) Each(fn func(int, float64)) Floater64 {
	f64.s.Each(func(i int, v interface{}) {
		fn(i, (v.(float64)))
	})
	return f64
}

func (f64 *floater64) EachBreak(fn func(int, float64) bool) Floater64 {
	f64.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float64)))
	})
	return f64
}

func (f64 *floater64) EachReverse(fn func(int, float64)) Floater64 {
	f64.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(float64)))
	})
	return f64
}

func (f64 *floater64) EachReverseBreak(fn func(int, float64) bool) Floater64 {
	f64.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float64)))
	})
	return f64
}

func (f64 *floater64) Fetch(i int) float64 {
	var s, _ = f64.Get(i)
	return s
}

func (f64 *floater64) Get(i int) (float64, bool) {
	var (
		ok bool
		s  float64
	)
	ok = f64.Bounds(i)
	if ok {
		s = (f64.s.Fetch(i)).(float64)
	}
	return s, ok
}

func (f64 *floater64) Len() int {
	return (f64.s.Len())
}

func (f64 *floater64) Less(i int, j int) bool {
	return f64.Fetch(i) < f64.Fetch(j)
}

func (f64 *floater64) Map(fn func(int, float64) float64) Floater64 {
	f64.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(float64)))
	})
	return f64
}

func (f64 *floater64) Poll() float64 {
	var (
		s float64
		v = f64.s.Poll()
	)
	if v != nil {
		s = (v.(float64))
	}
	return s
}

func (f64 *floater64) Pop() float64 {
	var (
		s float64
		v = f64.s.Pop()
	)
	if v != nil {
		s = (v.(float64))
	}
	return s
}

func (f64 *floater64) Precatenate(f Floater64) Floater64 {
	f64.s.Precatenate(f.(*floater64).s)
	return f64
}

func (f64 *floater64) Prepend(f ...float64) Floater64 {
	f64.s.Prepend(float64ToInterface(f...)...)
	return f64
}

func (f64 *floater64) Push(f ...float64) int {
	return f64.s.Push(float64ToInterface(f...))
}

func (f64 *floater64) Replace(i int, s float64) bool {
	return (f64.s.Replace(i, s))
}

func (f64 *floater64) Set() Floater64 {
	f64.s.Set()
	return f64
}

func (f64 *floater64) Sort() Floater64 {
	sort.Sort(f64)
	return f64
}

func (f64 *floater64) Swap(i int, j int) {
	f64.s.Swap(i, j)
}

func (f64 *floater64) Unshift(f ...float64) int {
	return (f64.s.Unshift(float64ToInterface(f...)))
}

func (f64 *floater64) Values() []float64 {
	var f = make([]float64, f64.Len())
	f64.Each(func(i int, s float64) {
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
