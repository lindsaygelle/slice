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
	Precatenate(Stringer) Stringer
	Prepend(...string) Stringer
	Push(...string) int
	Replace(int, string) bool
	Unshift(...string) int
}

// NewStringer returns a new Stringer interface.
func NewStringer() Stringer {
	return (&stringer{&slice{}})
}

type stringer struct{ *slice }

func (str *stringer) Append(s ...string) Stringer {
	str.append(stringsToInterfaces(s...)...)
	return str
}

func (str *stringer) Bounds(i int) bool {
	return str.bounds(i)
}

func (str *stringer) Concatenate(s Stringer) Stringer {
	str.concatenate(s.(*stringer).slice)
	return str
}

func (str *stringer) Each(fn func(int, string)) Stringer {
	str.each(func(i int, v interface{}) {
		fn(i, (v.(string)))
	})
	return str
}

func (str *stringer) EachBreak(fn func(int, string) bool) Stringer {
	str.eachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(string)))
	})
	return str
}

func (str *stringer) EachReverse(fn func(int, string)) Stringer {
	str.eachReverse(func(i int, v interface{}) {
		fn(i, (v.(string)))
	})
	return str
}

func (str *stringer) EachReverseBreak(fn func(int, string) bool) Stringer {
	str.eachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(string)))
	})
	return str
}

func (str *stringer) Len() int {
	return (str.len())
}

func (str *stringer) Precatenate(s Stringer) Stringer {
	str.precatenate(s.(*stringer).slice)
	return str
}

func (str *stringer) Prepend(s ...string) Stringer {
	str.prepend(stringsToInterfaces(s...)...)
	return str
}

func (str *stringer) Push(s ...string) int {
	return (str.Append(s...).Len())
}

func (str *stringer) Replace(i int, s string) bool {
	return (str.replace(i, s))
}

func (str *stringer) Unshift(s ...string) int {
	return (str.Prepend(s...).Len())
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
