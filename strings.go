package slice

func NewStringsSlice() *Strings {
	return &Strings{
		slices: NewSlices()}
}

type Strings struct {
	slices *Slices
}

func (pointer *Strings) Append(s *String) *Strings {
	pointer.slices.Append(s.slice)
	return pointer
}

func (pointer *Strings) Prepend(s *String) *Strings {
	pointer.slices.Prepend(s.slice)
	return pointer
}
