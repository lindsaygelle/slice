package slice

type stringer = slicer

func (stringer *stringer) Append(s string) {
	stringer.append(s)
}

func (stringer *stringer) Each(fn func(int, string)) {
	stringer.each(func(i int, x interface{}) {
		fn(i, x.(string))
	})
}

func (stringer *stringer) Prepend(s string) {
	stringer.prepend(s)
}

func (stringer *stringer) Pop() string {
	return (stringer.pop().(string))
}

// Stringer is an interface for a collection of strings.
type Stringer interface {
	Append(string)
	Each(func(int, string))
	Len() int
	Prepend(string)
	Pop() string
}

// String initializes a Stringer interface.
func String(strings ...string) Stringer {
	return (&stringer{})
}
