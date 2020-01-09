package slice

type Interfacer interface {
	Append(...interface{}) int
	Len() int
	Prepend(...interface{}) int
}

type interfacer struct{ *slice }

func (interfacer *interfacer) Append(i ...interface{}) int {
	return interfacer.append(i...)
}

func (interfacer *interfacer) Prepend(i ...interface{}) int {
	return interfacer.prepend(i...)
}
