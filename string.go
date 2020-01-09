package slice

// Stringer is the interface that accesses the string collection.
type Stringer interface {
	Append(...string) int
	Len() int
	Prepend(...string) int
}

// NewStringer returns a new Stringer interface.
func NewStringer() Stringer {
	return (&stringer{&slice{}})
}

type stringer struct{ *slice }

func (stringer *stringer) to(s ...string) []interface{} {
	var (
		i int
		v string
		x = make([]interface{}, len(s))
	)
	for i, v = range s {
		x[i] = v
	}
	return x
}

func (stringer *stringer) Append(s ...string) int  { return (stringer.append(stringer.to(s...)...)) }
func (stringer *stringer) Len() int                { return (stringer.len()) }
func (stringer *stringer) Prepend(s ...string) int { return (stringer.prepend(stringer.to(s...)...)) }
