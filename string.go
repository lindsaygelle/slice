package slice

type stringer struct{ *slice }

func (s *stringer) Each(fn func(int, string)) {
	s.slice.each(func(i int, x interface{}) {
		fn(i, x.(string))
	})
}

type Stringer interface {
	Each(func(int, string))
	Len() int
	OK(int) bool
}

func String(strings ...string) Stringer {
	var (
		s = &slice{}
	)
	for _, i := range strings {
		s.push(i)
	}
	return &stringer{s}
}
