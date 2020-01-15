package slice

import "sort"

// Floater32 is the interface that handles a float32 collection.
type Floater32 interface {
	Append(...float32) Floater32
	Bounds(int) bool
	Concatenate(Floater32) Floater32
	Delete(int) Floater32
	Each(func(int, float32)) Floater32
	EachBreak(func(int, float32) bool) Floater32
	EachReverse(func(int, float32)) Floater32
	EachReverseBreak(func(int, float32) bool) Floater32
	Fetch(int) float32
	Get(int) (float32, bool)
	Len() int
	Less(int, int) bool
	Make(int) Floater32
	MakeEach(...float32) Floater32
	MakeEachReverse(...float32) Floater32
	Map(func(int, float32) float32) Floater32
	Poll() float32
	Pop() float32
	Precatenate(Floater32) Floater32
	Prepend(...float32) Floater32
	Push(...float32) int
	Replace(int, float32) bool
	Set() Floater32
	Slice(int, int) Floater32
	Sort() Floater32
	Swap(int, int)
	Unshift(...float32) int
	Values() []float32
}

// NewFloater32 returns a new Floater32 interface.
func NewFloater32(f ...float32) Floater32 {
	return (&floater32{&Slice{}}).Append(f...)
}

type floater32 struct{ s *Slice }

func (f32 *floater32) Append(f ...float32) Floater32 {
	f32.s.Append(float32ToInterface(f...)...)
	return f32
}

func (f32 *floater32) Bounds(i int) bool {
	return f32.s.Bounds(i)
}

func (f32 *floater32) Concatenate(f Floater32) Floater32 {
	f32.s.Concatenate(f.(*floater32).s)
	return f32
}

func (f32 *floater32) Delete(i int) Floater32 {
	f32.s.Delete(i)
	return f32
}

func (f32 *floater32) Each(fn func(int, float32)) Floater32 {
	f32.s.Each(func(i int, v interface{}) {
		fn(i, (v.(float32)))
	})
	return f32
}

func (f32 *floater32) EachBreak(fn func(int, float32) bool) Floater32 {
	f32.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float32)))
	})
	return f32
}

func (f32 *floater32) EachReverse(fn func(int, float32)) Floater32 {
	f32.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(float32)))
	})
	return f32
}

func (f32 *floater32) EachReverseBreak(fn func(int, float32) bool) Floater32 {
	f32.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(float32)))
	})
	return f32
}

func (f32 *floater32) Fetch(i int) float32 {
	var s, _ = f32.Get(i)
	return s
}

func (f32 *floater32) Get(i int) (float32, bool) {
	var (
		ok bool
		s  float32
	)
	ok = f32.Bounds(i)
	if ok {
		s = (f32.s.Fetch(i)).(float32)
	}
	return s, ok
}

func (f32 *floater32) Len() int {
	return (f32.s.Len())
}

func (f32 *floater32) Less(i int, j int) bool {
	return f32.Fetch(i) < f32.Fetch(j)
}

func (f32 *floater32) Make(i int) Floater32 {
	f32.s.Make(i)
	return f32
}

func (f32 *floater32) MakeEach(i ...float32) Floater32 {
	f32.s.MakeEach(float32ToInterface(i...)...)
	return f32
}

func (f32 *floater32) MakeEachReverse(i ...float32) Floater32 {
	f32.s.MakeEachReverse(float32ToInterface(i...)...)
	return f32
}

func (f32 *floater32) Map(fn func(int, float32) float32) Floater32 {
	f32.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(float32)))
	})
	return f32
}

func (f32 *floater32) Poll() float32 {
	var (
		s float32
		v = f32.s.Poll()
	)
	if v != nil {
		s = (v.(float32))
	}
	return s
}

func (f32 *floater32) Pop() float32 {
	var (
		s float32
		v = f32.s.Pop()
	)
	if v != nil {
		s = (v.(float32))
	}
	return s
}

func (f32 *floater32) Precatenate(f Floater32) Floater32 {
	f32.s.Precatenate(f.(*floater32).s)
	return f32
}

func (f32 *floater32) Prepend(f ...float32) Floater32 {
	f32.s.Prepend(float32ToInterface(f...)...)
	return f32
}

func (f32 *floater32) Push(f ...float32) int {
	return f32.s.Push(float32ToInterface(f...))
}

func (f32 *floater32) Replace(i int, s float32) bool {
	return (f32.s.Replace(i, s))
}

func (f32 *floater32) Set() Floater32 {
	f32.s.Set()
	return f32
}

func (f32 *floater32) Slice(i int, j int) Floater32 {
	f32.s.Slice(i, j)
	return f32
}

func (f32 *floater32) Sort() Floater32 {
	sort.Sort(f32)
	return f32
}

func (f32 *floater32) Swap(i int, j int) {
	f32.s.Swap(i, j)
}

func (f32 *floater32) Unshift(f ...float32) int {
	return (f32.s.Unshift(float32ToInterface(f...)))
}

func (f32 *floater32) Values() []float32 {
	var f = make([]float32, f32.Len())
	f32.Each(func(i int, s float32) {
		f[i] = s
	})
	return f
}

func float32ToInterface(f ...float32) []interface{} {
	var (
		i int
		v float32
		x = make([]interface{}, (len(f)))
	)
	for i, v = range f {
		x[i] = v
	}
	return x
}
