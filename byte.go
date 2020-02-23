package slice

import (
	"sort"
)

// Byter is the interface that handles a byte collection.
type Byter interface {
	Append(...byte) Byter
	Bounds(int) bool
	Concatenate(Byter) Byter
	Delete(int) Byter
	DeleteOK(int) bool
	DeleteLength(int) int
	Each(func(int, byte)) Byter
	EachBreak(func(int, byte) bool) Byter
	EachReverse(func(int, byte)) Byter
	EachReverseBreak(func(int, byte) bool) Byter
	Fetch(int) byte
	FetchLength(int) (byte, int)
	Get(int) (byte, bool)
	GetLength(int) (byte, int, bool)
	Len() int
	Less(int, int) bool
	Make(int) Byter
	MakeEach(...byte) Byter
	MakeEachReverse(...byte) Byter
	Map(func(int, byte) byte) Byter
	Poll() byte
	Pop() byte
	Precatenate(Byter) Byter
	Prepend(...byte) Byter
	Push(...byte) int
	Replace(int, byte) bool
	Set() Byter
	Slice(int, int) Byter
	Sort() Byter
	Swap(int, int)
	Unshift(...byte) int
	Values() []byte
}

// NewByter returns a new Byter interface.
func NewByter(i ...byte) Byter {
	return (&byter{&Slice{}}).Append(i...)
}

// byter is the private struct that implements the Byter interface.
type byter struct{ s *Slice }

func (p *byter) Append(i ...byte) Byter {
	p.s.Append(byteToInterface(i...)...)
	return p
}

func (p *byter) Bounds(i int) bool {
	return p.s.Bounds(i)
}

func (p *byter) Concatenate(v Byter) Byter {
	p.s.Concatenate(v.(*byter).s)
	return p
}

func (p *byter) Delete(i int) Byter {
	p.s.Delete(i)
	return p
}

func (p *byter) Each(fn func(int, byte)) Byter {
	p.s.Each(func(i int, v interface{}) {
		fn(i, (v.(byte)))
	})
	return p
}

func (p *byter) EachBreak(fn func(int, byte) bool) Byter {
	p.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(byte)))
	})
	return p
}

func (p *byter) EachReverse(fn func(int, byte)) Byter {
	p.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(byte)))
	})
	return p
}

func (p *byter) EachReverseBreak(fn func(int, byte) bool) Byter {
	p.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(byte)))
	})
	return p
}

func (p *byter) Fetch(i int) byte {
	var s, _ = p.Get(i)
	return s
}

func (p, *byter) FetchLength(i int) (byte, int) {
	var s, _ = p.Fetch(i)
	var l = p.Len()
	return s, l
}

func (p *byter) Get(i int) (byte, bool) {
	var (
		ok bool
		s  byte
	)
	ok = p.Bounds(i)
	if ok {
		s = (p.s.Fetch(i)).(byte)
	}
	return s, ok
}

func (p *byter) GetLength(i int) (byte, int, bool) {
	var s, ok = p.Get(i)
	var l = p.Len()
	return s, l, ok
}

func (p *byter) Len() int {
	return (p.s.Len())
}

func (p *byter) Less(i int, j int) bool {
	return p.Fetch(i) < p.Fetch(j)
}

func (p *byter) Make(i int) Byter {
	p.s.Make(i)
	return p
}

func (p *byter) MakeEach(i ...byte) Byter {
	p.s.MakeEach(byteToInterface(i...)...)
	return p
}

func (p *byter) MakeEachReverse(i ...byte) Byter {
	p.s.MakeEachReverse(byteToInterface(i...)...)
	return p
}

func (p *byter) Map(fn func(int, byte) byte) Byter {
	p.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(byte)))
	})
	return p
}

func (p *byter) Poll() byte {
	var (
		s byte
		v = p.s.Poll()
	)
	if v != nil {
		s = (v.(byte))
	}
	return s
}

func (p *byter) Pop() byte {
	var (
		s byte
		v = p.s.Pop()
	)
	if v != nil {
		s = (v.(byte))
	}
	return s
}

func (p *byter) Precatenate(v Byter) Byter {
	p.s.Precatenate(v.(*byter).s)
	return p
}

func (p *byter) Prepend(i ...byte) Byter {
	p.s.Prepend(byteToInterface(i...)...)
	return p
}

func (p *byter) Push(i ...byte) int {
	return p.s.Push(byteToInterface(i...))
}

func (p *byter) Replace(i int, n byte) bool {
	return (p.s.Replace(i, n))
}

func (p *byter) Set() Byter {
	p.s.Set()
	return p
}

func (p *byter) Slice(i int, j int) Byter {
	p.s.Slice(i, j)
	return p
}

func (p *byter) Sort() Byter {
	sort.Sort(p)
	return p
}

func (p *byter) Swap(i int, j int) {
	p.s.Swap(i, j)
}

func (p *byter) Unshift(i ...byte) int {
	return (p.s.Unshift(byteToInterface(i...)))
}

func (p *byter) Values() []byte {
	var v = make([]byte, p.Len())
	p.Each(func(i int, n byte) {
		v[i] = n
	})
	return v
}

func byteToInterface(n ...byte) []interface{} {
	var (
		i int
		v byte
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
