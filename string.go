package slice

import (
	"sort"
	"strings"
)

// Stringer is the interface that handles a string collection.
type Stringer interface {
	Append(...string) Stringer
	Bounds(int) bool
	Concatenate(Stringer) Stringer
	Each(func(int, string)) Stringer
	EachBreak(func(int, string) bool) Stringer
	EachReverse(func(int, string)) Stringer
	EachReverseBreak(func(int, string) bool) Stringer
	Fetch(int) string
	Get(int) (string, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, string) string) Stringer
	Precatenate(Stringer) Stringer
	Prepend(...string) Stringer
	Push(...string) int
	Replace(int, string) bool
	Set() Stringer
	Sort() Stringer
	Swap(int, int)
	Unshift(...string) int
	Values() []string
}

// NewStringer returns a new Stringer interface.
func NewStringer() Stringer {
	return (&stringer{s: &Slice{}})
}

type stringer struct{ s *Slice }

func (str *stringer) Append(s ...string) Stringer {
	str.s.Append(stringsToInterfaces(s...)...)
	return str
}

func (str *stringer) Bounds(i int) bool {
	return str.s.Bounds(i)
}

func (str *stringer) Concatenate(s Stringer) Stringer {
	str.s.Concatenate(s.(*stringer).s)
	return str
}

func (str *stringer) Each(fn func(int, string)) Stringer {
	str.s.Each(func(i int, v interface{}) {
		fn(i, (v.(string)))
	})
	return str
}

func (str *stringer) EachBreak(fn func(int, string) bool) Stringer {
	str.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(string)))
	})
	return str
}

func (str *stringer) EachReverse(fn func(int, string)) Stringer {
	str.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(string)))
	})
	return str
}

func (str *stringer) EachReverseBreak(fn func(int, string) bool) Stringer {
	str.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(string)))
	})
	return str
}

func (str *stringer) Fetch(i int) string {
	var s, _ = str.Get(i)
	return s
}

func (str *stringer) Get(i int) (string, bool) {
	var (
		ok bool
		s  string
	)
	ok = str.Bounds(i)
	if ok {
		s = (str.s.Fetch(i)).(string)
	}
	return s, ok
}

func (str *stringer) Len() int {
	return (str.s.Len())
}

func (str *stringer) Less(i int, j int) bool {
	const (
		f string = "%s"
	)
	var (
		a  = str.Fetch(i)
		b  = str.Fetch(j)
		ok bool
	)
	ok = (a == b)
	if ok {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}
	ok = a < b
	return ok
}

func (str *stringer) Map(fn func(int, string) string) Stringer {
	str.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(string)))
	})
	return str
}

func (str *stringer) Precatenate(s Stringer) Stringer {
	str.s.Precatenate(s.(*stringer).s)
	return str
}

func (str *stringer) Prepend(s ...string) Stringer {
	str.s.Prepend(stringsToInterfaces(s...)...)
	return str
}

func (str *stringer) Push(s ...string) int {
	return str.s.Push(stringsToInterfaces(s...))
}

func (str *stringer) Replace(i int, s string) bool {
	return (str.s.Replace(i, s))
}

func (str *stringer) Set() Stringer {
	str.s.Set()
	return str
}

func (str *stringer) Sort() Stringer {
	sort.Sort(str)
	return str
}

func (str *stringer) Swap(i int, j int) {
	str.s.Swap(i, j)
}

func (str *stringer) Unshift(s ...string) int {
	return (str.s.Unshift(stringsToInterfaces(s...)))
}

func (str *stringer) Values() []string {
	var strs = make([]string, str.Len())
	str.Each(func(i int, s string) {
		strs[i] = s
	})
	return strs
}

func stringsToInterfaces(s ...string) []interface{} {
	var (
		i int
		v string
		x = make([]interface{}, (len(s)))
	)
	for i, v = range s {
		x[i] = v
	}
	return x
}
