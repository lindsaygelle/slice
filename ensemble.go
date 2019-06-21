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
	Flatten() *Slice
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

// Ensemble is a superset of the Slice struct whose methods manage the access, insertion and modification of Slice only values.
type Ensemble struct {
	slice *Slice
}

// Append method adds one Slice to the end of the Ensemble Slice and returns the modified Ensemble Slice.
func (pointer *Ensemble) Append(slice *Slice) *Ensemble {
	pointer.slice.Append(slice)
	return pointer
}

// Assign method adds zero or more Slice pointers to the end of the Ensemble Slice and returns the modified Ensemble Slice.
func (pointer *Ensemble) Assign(slices ...*Slice) *Ensemble {
	for _, slice := range slices {
		pointer.Append(slice)
	}
	return pointer
}

// Bounds checks an integer value safely sits within the range of accessible values for the Ensemble Slice.
func (pointer *Ensemble) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

// Concatenate merges two Ensemble Slices into a single Ensemble Slice.
func (pointer *Ensemble) Concatenate(ensemble *Ensemble) *Ensemble {
	ensemble.Each(func(_ int, slice *Slice) {
		pointer.Append(slice)
	})
	return pointer
}

// Each method executes a provided function once for each Ensemble Slice element.
func (pointer *Ensemble) Each(f func(i int, slice *Slice)) *Ensemble {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*Slice))
	})
	return pointer
}

// Empty returns a boolean indicating whether the Ensemble Slice contains zero values.
func (pointer *Ensemble) Empty() bool {
	return pointer.slice.Empty()
}

// Fetch retrieves the string held at the argument index. Returns nil if index exceeds Ensemble Slice length.
func (pointer *Ensemble) Fetch(i int) *Slice {
	slice, _ := pointer.Get(i)
	return slice
}

func (pointer *Ensemble) Flatten() *Slice {
	slice := New()
	pointer.Each(func(_ int, s *Slice) {
		s.Each(func(_ int, value interface{}) {
			slice.Append(value)
		})
	})
	return slice
}

// Get returns the Slice held at the argument index and a boolean indicating if it was successfully retrieved.
func (pointer *Ensemble) Get(i int) (*Slice, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*Slice), ok
	}
	return nil, ok
}

// Len method returns the number of elements in the Ensemble Slice.
func (pointer *Ensemble) Len() int {
	return pointer.slice.Len()
}

// Map method executes a provided function once for each Slice element and sets the returned value to the current index.
func (pointer *Ensemble) Map(f func(i int, slice *Slice) *Slice) *Ensemble {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*Slice))
	})
	return pointer
}

// Poll method removes the first Slice from the Ensemble Slice and returns the removed Slice.
func (pointer *Ensemble) Poll() *Slice {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*Slice)
	}
	return nil
}

// Pop method removes the last Slice from the Ensemble Slice and returns the Slice.
func (pointer *Ensemble) Pop() *Slice {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*Slice)
	}
	return nil
}

// Preassign method adds zero or more Slices to the beginning of the Ensemble Slice and returns the modified Ensemble Slice.
func (pointer *Ensemble) Preassign(slices ...*Slice) *Ensemble {
	for _, slice := range slices {
		pointer.Prepend(slice)
	}
	return pointer
}

// Precatenate merges two Ensemble Slices, prepending the argument Ensemble Slice.
func (pointer *Ensemble) Precatenate(ensemble *Ensemble) *Ensemble {
	pointer.slice.Precatenate(ensemble.slice)
	return pointer
}

// Prepend method adds one Slice to the beginning of the Ensemble Slice and returns the modified Ensemble Slice.
func (pointer *Ensemble) Prepend(slice *Slice) *Ensemble {
	pointer.slice.Prepend(slice)
	return pointer
}

// Push method adds a new Slice to the end of the Ensemble Slice and returns the length of the modified Ensemble Slice.
func (pointer *Ensemble) Push(slice *Slice) int {
	return pointer.slice.Push(slice)
}

// Replace method replaces the Slice at the argument index if it is in bounds with the argument Slice.
func (pointer *Ensemble) Replace(i int, slice *Slice) *Ensemble {
	pointer.slice.Replace(i, slice)
	return pointer
}
