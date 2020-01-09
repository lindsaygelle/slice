package slice

type slice []interface{}

func (s *slice) each(fn func(i int, x interface{})) {
	for i, x := range *s {
		fn(i, x)
	}
}

func (s *slice) push(i interface{}) int { (*s) = (append(*s, i)); return s.Len() }

func (s *slice) Len() int      { return (len(*s)) }
func (s *slice) OK(i int) bool { return ((i > -1) && (i < (len(*s)))) }
