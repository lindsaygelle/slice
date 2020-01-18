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

func (pointer *runer) Append(i ...rune) Runer {
	pointer.s.Append(runeToInterface(i...)...)
	return pointer
}

func (pointer *runer) Bounds(i int) bool {
	return pointer.s.Bounds(i)
}

func (pointer *runer) Concatenate(v Runer) Runer {
	pointer.s.Concatenate(v.(*runer).s)
	return pointer
}

func (pointer *runer) Delete(i int) Runer {
	pointer.s.Delete(i)
	return pointer
}

func (pointer *runer) Each(fn func(int, rune)) Runer {
	pointer.s.Each(func(i int, v interface{}) {
		fn(i, (v.(rune)))
	})
	return pointer
}

func (pointer *runer) EachBreak(fn func(int, rune) bool) Runer {
	pointer.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(rune)))
	})
	return pointer
}

func (pointer *runer) EachReverse(fn func(int, rune)) Runer {
	pointer.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(rune)))
	})
	return pointer
}

func (pointer *runer) EachReverseBreak(fn func(int, rune) bool) Runer {
	pointer.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(rune)))
	})
	return pointer
}

func (pointer *runer) Fetch(i int) rune {
	var s, _ = pointer.Get(i)
	return s
}

func (pointer *runer) Get(i int) (rune, bool) {
	var (
		ok bool
		s  rune
	)
	ok = pointer.Bounds(i)
	if ok {
		s = (pointer.s.Fetch(i)).(rune)
	}
	return s, ok
}

func (pointer *runer) Len() int {
	return (pointer.s.Len())
}

func (pointer *runer) Less(i int, j int) bool {
	return pointer.Fetch(i) < pointer.Fetch(j)
}

func (pointer *runer) Make(i int) Runer {
	pointer.s.Make(i)
	return pointer
}

func (pointer *runer) MakeEach(i ...rune) Runer {
	pointer.s.MakeEach(runeToInterface(i...)...)
	return pointer
}

func (pointer *runer) MakeEachReverse(i ...rune) Runer {
	pointer.s.MakeEachReverse(runeToInterface(i...)...)
	return pointer
}

func (pointer *runer) Map(fn func(int, rune) rune) Runer {
	pointer.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(rune)))
	})
	return pointer
}

func (pointer *runer) Poll() rune {
	var (
		s rune
		v = pointer.s.Poll()
	)
	if v != nil {
		s = (v.(rune))
	}
	return s
}

func (pointer *runer) Pop() rune {
	var (
		s rune
		v = pointer.s.Pop()
	)
	if v != nil {
		s = (v.(rune))
	}
	return s
}

func (pointer *runer) Precatenate(v Runer) Runer {
	pointer.s.Precatenate(v.(*runer).s)
	return pointer
}

func (pointer *runer) Prepend(i ...rune) Runer {
	pointer.s.Prepend(runeToInterface(i...)...)
	return pointer
}

func (pointer *runer) Push(i ...rune) int {
	return pointer.s.Push(runeToInterface(i...))
}

func (pointer *runer) Replace(i int, n rune) bool {
	return (pointer.s.Replace(i, n))
}

func (pointer *runer) Set() Runer {
	pointer.s.Set()
	return pointer
}

func (pointer *runer) Slice(i int, j int) Runer {
	pointer.s.Slice(i, j)
	return pointer
}

func (pointer *runer) Sort() Runer {
	sort.Sort(pointer)
	return pointer
}

func (pointer *runer) Swap(i int, j int) {
	pointer.s.Swap(i, j)
}

func (pointer *runer) Unshift(i ...rune) int {
	return (pointer.s.Unshift(runeToInterface(i...)))
}

func (pointer *runer) Values() []rune {
	var v = make([]rune, pointer.Len())
	pointer.Each(func(i int, n rune) {
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
