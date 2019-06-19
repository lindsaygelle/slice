package slice

type String struct {
	slice *Slice
}

func (pointer *String) Append(s string) *String {
	pointer.slice.Append(s)
	return pointer
}

func (pointer *String) Assign(s ...string) *String {
	for i := range s {
		pointer.Append(s[i])
	}
	return pointer
}

func (pointer *String) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

func (pointer *String) Concatenate(s *String) *String {
	pointer.slice.Concatenate(s.slice)
	return pointer
}
