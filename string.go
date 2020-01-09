package slice

// Stringer is the interface that accesses the string collection.
type Stringer interface {
	Append(...string) Stringer
	Concatenate(Stringer) Stringer
	Len() int
	Precatenate(Stringer) Stringer
	Prepend(...string) Stringer
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

func (str *stringer) Concatenate(s Stringer) Stringer {
	str.concatenate(s.(*stringer).slice)
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
