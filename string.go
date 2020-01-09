package slice

type stringer = slicer

func (stringer *stringer) Each(fn func(int, string)) {
	stringer.eachString(fn)
}

func (stringer *stringer) Push(s string) {
	stringer.pushString(s)
}

type Stringer interface {
	Each(func(int, string))
	Push(string)
}

func String(strings ...string) Stringer {
	return (&stringer{})
}
