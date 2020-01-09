package slice

type slice []interface{}

func (slice *slice) append(i ...interface{}) *slice {
	(*slice) = (append(*slice, i...))
	return slice
}

func (slice *slice) concatenate(s *slice) *slice {
	slice.append((*s)...)
	return slice
}

func (slice *slice) each(fn func(int, interface{})) {
	var (
		i int
		v interface{}
	)
	for i, v = range *slice {
		fn(i, v)
	}
}

func (slice *slice) len() int { return (len(*slice)) }

func (slice *slice) precatenate(s *slice) *slice {
	slice.prepend((*s)...)
	return slice
}

func (slice *slice) prepend(i ...interface{}) *slice {
	(*slice) = (append(i, *slice...))
	return slice
}
