package slice

type Interfacer interface {
	Append(...interface{}) Interfacer
	Len() int
	Prepend(...interface{}) Interfacer
}

type interfacer struct{ s *Slice }
