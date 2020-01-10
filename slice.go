package slice

type slice []interface{}

func (slice *slice) append(i ...interface{}) *slice {
	(*slice) = (append(*slice, i...))
	return slice
}

func (slice *slice) bounds(i int) bool {
	return ((i > -1) && (i < len(*slice)))
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

func (slice *slice) eachBreak(fn func(int, interface{}) bool) {
	var (
		i  int
		ok = true
		v  interface{}
	)
	for i, v = range *slice {
		ok = fn(i, v)
		if !ok {
			break
		}
	}
}

func (slice *slice) eachReverse(fn func(int, interface{})) {
	var (
		i int
	)
	for i = len(*slice) - 1; i >= 0; i-- {
		fn(i, (*slice)[i])
	}
}

func (slice *slice) eachReverseBreak(fn func(int, interface{}) bool) {
	var (
		i  int
		ok = true
	)
	for i = len(*slice) - 1; i >= 0; i-- {
		ok = fn(i, (*slice)[i])
		if !ok {
			break
		}
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

func (slice *slice) replace(i int, v interface{}) bool {
	var (
		ok = slice.bounds(i)
	)
	if ok {
		(*slice)[i] = v
	}
	return ok
}
