package slice

// Stringer is the interface that accesses the string collection.
type Stringer interface {
	Append(...string) Stringer
	Bounds(int) bool
	Concatenate(Stringer) Stringer
	Each(func(int, string)) Stringer
	EachBreak(func(int, string) bool) Stringer
	EachReverse(func(int, string)) Stringer
	EachReverseBreak(func(int, string) bool) Stringer
	Len() int
	Map(func(int, string) string) Stringer
	Precatenate(Stringer) Stringer
	Prepend(...string) Stringer
	Push(...string) int
	Replace(int, string) bool
	Swap(int, int) Stringer
	Unshift(...string) int
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

func (str *stringer) Len() int {
	return (str.s.Len())
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

func (str *stringer) Swap(i int, j int) Stringer {
	str.s.Swap(i, j)
	return str
}

func (str *stringer) Unshift(s ...string) int {
	return (str.s.Unshift(stringsToInterfaces(s...)))
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
