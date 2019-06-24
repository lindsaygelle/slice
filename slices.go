package slice

var (
	_ slices = (*Slices)(nil)
)

// NewSlices instantiates a new empty Slices slice.
func NewSlices(slice ...*Slice) *Slices {
	return (&Slices{
		slice: &Slice{}}).Assign(slice...)
}

type slices interface {
	Append(slice *Slice) *Slices
	Assign(slices ...*Slice) *Slices
	Bounds(i int) bool
	Concatenate(slices *Slices) *Slices
	Each(f func(i int, value *Slice)) *Slices
	Empty() bool
	Fetch(i int) *Slice
	Flatten() *Slice
	Get(i int) (*Slice, bool)
	Len() int
	Map(f func(i int, slice *Slice) *Slice) *Slices
	Poll() *Slice
	Pop() *Slice
	Preassign(slices ...*Slice) *Slices
	Precatenate(slices *Slices) *Slices
	Prepend(slice *Slice) *Slices
	Push(slice *Slice) int
	Replace(i int, slice *Slice) *Slices
}

// Slices is a superset of the Slice struct whose methods manage the access, insertion and modification of Slice only values.
type Slices struct {
	slice *Slice
}

// Append method adds one Slice to the end of the Slices Slice and returns the modified Slices Slice.
func (pointer *Slices) Append(slice *Slice) *Slices {
	pointer.slice.Append(slice)
	return pointer
}

// Assign method adds zero or more Slice pointers to the end of the Slices Slice and returns the modified Slices Slice.
func (pointer *Slices) Assign(slices ...*Slice) *Slices {
	for _, slice := range slices {
		pointer.Append(slice)
	}
	return pointer
}

// Bounds checks an integer value safely sits within the range of accessible values for the Slices Slice.
func (pointer *Slices) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

// Concatenate merges two Slices Slices into a single Slices Slice.
func (pointer *Slices) Concatenate(slices *Slices) *Slices {
	slices.Each(func(_ int, slice *Slice) {
		pointer.Append(slice)
	})
	return pointer
}

// Each method executes a provided function once for each Slices Slice element.
func (pointer *Slices) Each(f func(i int, slice *Slice)) *Slices {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*Slice))
	})
	return pointer
}

// Empty returns a boolean indicating whether the Slices Slice contains zero values.
func (pointer *Slices) Empty() bool {
	return pointer.slice.Empty()
}

// Fetch retrieves the string held at the argument index. Returns nil if index exceeds Slices Slice length.
func (pointer *Slices) Fetch(i int) *Slice {
	slice, _ := pointer.Get(i)
	return slice
}

// Flatten merges all Slice's into a single Slice and returns the Slice.
func (pointer *Slices) Flatten() *Slice {
	slice := New()
	pointer.Each(func(_ int, s *Slice) {
		s.Each(func(_ int, value interface{}) {
			slice.Append(value)
		})
	})
	return slice
}

// Get returns the Slice held at the argument index and a boolean indicating if it was successfully retrieved.
func (pointer *Slices) Get(i int) (*Slice, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*Slice), ok
	}
	return nil, ok
}

// Len method returns the number of elements in the Slices Slice.
func (pointer *Slices) Len() int {
	return pointer.slice.Len()
}

// Map method executes a provided function once for each Slice element and sets the returned value to the current index.
func (pointer *Slices) Map(f func(i int, slice *Slice) *Slice) *Slices {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*Slice))
	})
	return pointer
}

// Poll method removes the first Slice from the Slices Slice and returns the removed Slice.
func (pointer *Slices) Poll() *Slice {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*Slice)
	}
	return nil
}

// Pop method removes the last Slice from the Slices Slice and returns the Slice.
func (pointer *Slices) Pop() *Slice {
	value := pointer.slice.Poll()
	if value != nil {
		return value.(*Slice)
	}
	return nil
}

// Preassign method adds zero or more Slices to the beginning of the Slices Slice and returns the modified Slices Slice.
func (pointer *Slices) Preassign(slices ...*Slice) *Slices {
	for _, slice := range slices {
		pointer.Prepend(slice)
	}
	return pointer
}

// Precatenate merges two Slices Slices, prepending the argument Slices Slice.
func (pointer *Slices) Precatenate(slices *Slices) *Slices {
	pointer.slice.Precatenate(slices.slice)
	return pointer
}

// Prepend method adds one Slice to the beginning of the Slices Slice and returns the modified Slices Slice.
func (pointer *Slices) Prepend(slice *Slice) *Slices {
	pointer.slice.Prepend(slice)
	return pointer
}

// Push method adds a new Slice to the end of the Slices Slice and returns the length of the modified Slices Slice.
func (pointer *Slices) Push(slice *Slice) int {
	return pointer.slice.Push(slice)
}

// Replace method replaces the Slice at the argument index if it is in bounds with the argument Slice.
func (pointer *Slices) Replace(i int, slice *Slice) *Slices {
	pointer.slice.Replace(i, slice)
	return pointer
}
