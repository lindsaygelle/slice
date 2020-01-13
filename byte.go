package slice

import (
	"sort"
)

// Byter is the interface that handles a byte collection.
type Byter interface {
	Append(...byte) Byter
	Bounds(int) bool
	Concatenate(Byter) Byter
	Each(func(int, byte)) Byter
	EachBreak(func(int, byte) bool) Byter
	EachReverse(func(int, byte)) Byter
	EachReverseBreak(func(int, byte) bool) Byter
	Fetch(int) byte
	Get(int) (byte, bool)
	Len() int
	Less(int, int) bool
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

type byter struct{ s *Slice }

func (b *byter) Append(i ...byte) Byter {
	b.s.Append(byteToInterface(i...)...)
	return b
}

func (b *byter) Bounds(i int) bool {
	return b.s.Bounds(i)
}

func (b *byter) Concatenate(v Byter) Byter {
	b.s.Concatenate(v.(*byter).s)
	return b
}

func (b *byter) Each(fn func(int, byte)) Byter {
	b.s.Each(func(i int, v interface{}) {
		fn(i, (v.(byte)))
	})
	return b
}

func (b *byter) EachBreak(fn func(int, byte) bool) Byter {
	b.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(byte)))
	})
	return b
}

func (b *byter) EachReverse(fn func(int, byte)) Byter {
	b.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(byte)))
	})
	return b
}

func (b *byter) EachReverseBreak(fn func(int, byte) bool) Byter {
	b.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(byte)))
	})
	return b
}

func (b *byter) Fetch(i int) byte {
	var s, _ = b.Get(i)
	return s
}

func (b *byter) Get(i int) (byte, bool) {
	var (
		ok bool
		s  byte
	)
	ok = b.Bounds(i)
	if ok {
		s = (b.s.Fetch(i)).(byte)
	}
	return s, ok
}

func (b *byter) Len() int {
	return (b.s.Len())
}

func (b *byter) Less(i int, j int) bool {
	return b.Fetch(i) < b.Fetch(j)
}

func (b *byter) Map(fn func(int, byte) byte) Byter {
	b.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(byte)))
	})
	return b
}

func (b *byter) Poll() byte {
	var (
		s byte
		v = b.s.Poll()
	)
	if v != nil {
		s = (v.(byte))
	}
	return s
}

func (b *byter) Pop() byte {
	var (
		s byte
		v = b.s.Pop()
	)
	if v != nil {
		s = (v.(byte))
	}
	return s
}

func (b *byter) Precatenate(v Byter) Byter {
	b.s.Precatenate(v.(*byter).s)
	return b
}

func (b *byter) Prepend(i ...byte) Byter {
	b.s.Prepend(byteToInterface(i...)...)
	return b
}

func (b *byter) Push(i ...byte) int {
	return b.s.Push(byteToInterface(i...))
}

func (b *byter) Replace(i int, n byte) bool {
	return (b.s.Replace(i, n))
}

func (b *byter) Set() Byter {
	b.s.Set()
	return b
}

func (b *byter) Slice(i int, j int) Byter {
	b.s.Slice(i, j)
	return b
}

func (b *byter) Sort() Byter {
	sort.Sort(b)
	return b
}

func (b *byter) Swap(i int, j int) {
	b.s.Swap(i, j)
}

func (b *byter) Unshift(i ...byte) int {
	return (b.s.Unshift(byteToInterface(i...)))
}

func (b *byter) Values() []byte {
	var v = make([]byte, b.Len())
	b.Each(func(i int, n byte) {
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
