package slice

// Interfacer is the interface that handles a interface{} collection.
type Interfacer interface {
	Append(...interface{}) Interfacer
	Len() int
	Prepend(...interface{}) Interfacer
}

type interfacer struct{ s *Slice }
