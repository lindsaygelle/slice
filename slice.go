package slice

type slicer []interface{}

func (s *slicer) append(i interface{}) { (*s) = (append(*s, i)) }

func (s *slicer) each(fn func(int, interface{})) {
	for i, x := range *s {
		fn(i, x)
	}
}

func (s *slicer) poll()                 {}
func (s *slicer) pop() (x interface{})  { return x }
func (s *slicer) prepend(i interface{}) { (*s) = (append(slicer{i}, *s...)) }

func (s *slicer) Len() int { return (len(*s)) }

type Slicer interface {
	Len() int
}
