package slice

func NewStrings() *Strings {
	return &Strings{
		slice: &Slice{}}
}

type Strings struct {
	slice *Slice
}

func (pointer *Strings) Append(s *String) *Strings {
	pointer.slice.Append(s.slice)
	return pointer
}

func (pointer *Strings) Assign(s ...*String) *Strings {
	for _, s := range s {
		pointer.Append(s)
	}
	return pointer
}

func (pointer *Strings) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

func (pointer *Strings) Concatenate(strings *Strings) *Strings {
	pointer.slice.Concatenate(strings.slice)
	return pointer
}

func (pointer *Strings) Each(f func(i int, s *String)) *Strings {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*String))
	})
	return pointer
}

func (pointer *Strings) Empty() bool {
	return pointer.slice.Empty()
}

func (pointer *Strings) Fetch(i int) *String {
	s, _ := pointer.Get(i)
	return s
}

func (pointer *Strings) Flatten() *String {
	s := NewString()
	pointer.slice.Each(func(i int, value interface{}) {
		s.Append(value.(string))
	})
	return s
}

func (pointer *Strings) Get(i int) (*String, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*String), ok
	}
	return nil, ok
}

func (pointer *Strings) Len() int {
	return pointer.slice.Len()
}

func (pointer *Strings) Map(f func(i int, s *String) *String) *Strings {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*String))
	})
	return pointer
}

func (pointer *Strings) Poll() *String {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*String)
	}
	return nil
}

func (pointer *Strings) Pop() *String {
	value := pointer.slice.Pop()
	if value != nil {
		return value.(*String)
	}
	return nil
}

func (pointer *Strings) Preassign(strings ...*String) *Strings {
	for _, s := range strings {
		pointer.Prepend(s)
	}
	return pointer
}

func (pointer *Strings) Precatenate(s *String) *Strings {
	pointer.slice.Precatenate(s.slice)
	return pointer
}

func (pointer *Strings) Prepend(s *String) *Strings {
	pointer.slice.Prepend(s)
	return pointer
}

func (pointer *Strings) Push(s *String) int {
	return pointer.slice.Push(s)
}

func (pointer *Strings) Replace(i int, s *String) *Strings {
	pointer.slice.Replace(i, s)
	return pointer
}

func (pointer *Strings) Sort() *Strings {
	pointer.Each(func(_ int, s *String) {
		s.Sort()
	})
	return pointer
}
