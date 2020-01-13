package slice

import (
	"fmt"
	"sort"
)

// Interfacer is the interface that handles a interface{} collection.
type Interfacer interface {
	Append(...interface{}) Interfacer
	Bounds(int) bool
	Concatenate(Interfacer) Interfacer
	Each(func(int, interface{})) Interfacer
	EachBreak(func(int, interface{}) bool) Interfacer
	EachReverse(func(int, interface{})) Interfacer
	EachReverseBreak(func(int, interface{}) bool) Interfacer
	Fetch(int) interface{}
	Get(int) (interface{}, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, interface{}) interface{}) Interfacer
	Poll() interface{}
	Pop() interface{}
	Precatenate(Interfacer) Interfacer
	Prepend(...interface{}) Interfacer
	Push(...interface{}) int
	Replace(int, interface{}) bool
	Set() Interfacer
	Slice(int, int) Interfacer
	Sort() Interfacer
	Swap(int, int)
	Unshift(...interface{}) int
	Values() []interface{}
}

// NewInterfacer returns a new Interfacer interface.
func NewInterfacer() Interfacer {
	return (&interfacer{s: &Slice{}})
}

type interfacer struct{ s *Slice }

func (in *interfacer) Append(i ...interface{}) Interfacer {
	in.s.Append(i...)
	return in
}

func (in *interfacer) Bounds(i int) bool {
	return in.s.Bounds(i)
}

func (in *interfacer) Concatenate(v Interfacer) Interfacer {
	in.s.Concatenate(v.(*interfacer).s)
	return in
}

func (in *interfacer) Each(fn func(int, interface{})) Interfacer {
	in.s.Each(func(i int, v interface{}) {
		fn(i, (v.(interface{})))
	})
	return in
}

func (in *interfacer) EachBreak(fn func(int, interface{}) bool) Interfacer {
	in.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(interface{})))
	})
	return in
}

func (in *interfacer) EachReverse(fn func(int, interface{})) Interfacer {
	in.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(interface{})))
	})
	return in
}

func (in *interfacer) EachReverseBreak(fn func(int, interface{}) bool) Interfacer {
	in.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(interface{})))
	})
	return in
}

func (in *interfacer) Fetch(i int) interface{} {
	var s, _ = in.Get(i)
	return s
}

func (in *interfacer) Get(i int) (interface{}, bool) {
	var (
		ok bool
		s  interface{}
	)
	ok = in.Bounds(i)
	if ok {
		s = (in.s.Fetch(i)).(interface{})
	}
	return s, ok
}

func (in *interfacer) Len() int {
	return (in.s.Len())
}

func (in *interfacer) Less(i int, j int) bool {
	return fmt.Sprintf("%v", in.Fetch(i)) < fmt.Sprintf("%v", in.Fetch(j))
}

func (in *interfacer) Map(fn func(int, interface{}) interface{}) Interfacer {
	in.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(interface{})))
	})
	return in
}

func (in *interfacer) Poll() interface{} {
	var (
		s interface{}
		v = in.s.Poll()
	)
	if v != nil {
		s = (v.(interface{}))
	}
	return s
}

func (in *interfacer) Pop() interface{} {
	var (
		s interface{}
		v = in.s.Pop()
	)
	if v != nil {
		s = (v.(interface{}))
	}
	return s
}

func (in *interfacer) Precatenate(v Interfacer) Interfacer {
	in.s.Precatenate(v.(*interfacer).s)
	return in
}

func (in *interfacer) Prepend(i ...interface{}) Interfacer {
	in.s.Prepend(i...)
	return in
}

func (in *interfacer) Push(i ...interface{}) int {
	return in.s.Push(i)
}

func (in *interfacer) Replace(i int, s interface{}) bool {
	return (in.s.Replace(i, s))
}

func (in *interfacer) Set() Interfacer {
	in.s.Set()
	return in
}

func (in *interfacer) Slice(i int, j int) Interfacer {
	in.s.Slice(i, j)
	return in
}

func (in *interfacer) Sort() Interfacer {
	sort.Sort(in)
	return in
}

func (in *interfacer) Swap(i int, j int) {
	in.s.Swap(i, j)
}

func (in *interfacer) Unshift(i ...interface{}) int {
	return (in.s.Unshift(i))
}

func (in *interfacer) Values() []interface{} {
	var strs = make([]interface{}, in.Len())
	in.Each(func(i int, s interface{}) {
		strs[i] = s
	})
	return strs
}
