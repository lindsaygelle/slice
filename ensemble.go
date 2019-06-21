package slice

var (
	_ ensemble = (*Ensemble)(nil)
)

// NewEnsemble instantiates a new empty Ensemble slice.
func NewEnsemble() *Ensemble {
	return &Ensemble{
		slice: &Slice{}}
}

// NewEnsembleSlice instantiates a new populated or empty Ensemble slice.
func NewEnsembleSlice(s ...*Slice) *Ensemble {
	return NewEnsemble()
}

type ensemble interface {
	Append(slice *Slice) *Ensemble
	Assign(slices ...*Slice) *Ensemble
	Bounds(i int) bool
	Concatenate(ensemble *Ensemble) *Ensemble
	Each(f func(i int, value *Slice)) *Ensemble
	Empty() bool
	Fetch(i int) *Slice
	Get(i int) (*Slice, bool)
	Len() int
	Map(f func(i int, slice *Slice) *Slice) *Ensemble
	Poll() *Slice
	Pop() *Slice
	Preassign(slices ...*Slice) *Ensemble
	Precatenate(ensemble *Ensemble) *Ensemble
	Prepend(slice *Slice) *Ensemble
	Push(slice *Slice) int
	Replace(i int, slice *Slice) *Ensemble
}

type Ensemble struct {
	slice *Slice
}

func (pointer *Ensemble) Append(slice *Slice) *Ensemble {
	pointer.slice.Append(slice)
	return pointer
}

func (pointer *Ensemble) Assign(slices ...*Slice) *Ensemble {
	for _, slice := range slices {
		pointer.Append(slice)
	}
	return pointer
}

func (pointer *Ensemble) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

func (pointer *Ensemble) Concatenate(ensemble *Ensemble) *Ensemble {
	ensemble.Each(func(_ int, slice *Slice) {
		pointer.Append(slice)
	})
	return pointer
}

func (pointer *Ensemble) Each(f func(i int, slice *Slice)) *Ensemble {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*Slice))
	})
	return pointer
}

func (pointer *Ensemble) Empty() bool {
	return pointer.slice.Empty()
}

func (pointer *Ensemble) Fetch(i int) *Slice {
	slice, _ := pointer.Get(i)
	return slice
}

func (pointer *Ensemble) Get(i int) (*Slice, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*Slice), ok
	}
	return nil, ok
}

func (pointer *Ensemble) Len() int {
	return pointer.slice.Len()
}

func (pointer *Ensemble) Map(f func(i int, slice *Slice) *Slice) *Ensemble {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*Slice))
	})
	return pointer
}

func (pointer *Ensemble) Poll() *Slice {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*Slice)
	}
	return nil
}

func (pointer *Ensemble) Pop() *Slice {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*Slice)
	}
	return nil
}

func (pointer *Ensemble) Precatenate(ensemble *Ensemble) *Ensemble {
	pointer.slice.Precatenate(ensemble.slice)
	return pointer
}

func (pointer *Ensemble) Preassign(slices ...*Slice) *Ensemble {
	for _, slice := range slices {
		pointer.Prepend(slice)
	}
	return pointer
}

func (pointer *Ensemble) Prepend(slice *Slice) *Ensemble {
	pointer.slice.Prepend(slice)
	return pointer
}

func (pointer *Ensemble) Push(slice *Slice) int {
	return pointer.slice.Push(slice)
}

func (pointer *Ensemble) Replace(i int, slice *Slice) *Ensemble {
	pointer.slice.Replace(i, slice)
	return pointer
}
