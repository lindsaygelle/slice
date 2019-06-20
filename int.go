package slice

var (
	_ i = (*Int)(nil)
)

// NewInt instantiates a new empty Int slice.
func NewInt() *Int {
	return &Int{
		slice: &Slice{}}
}

// NewIntSlice instantiates a new populated or empty Int slice.
func NewIntSlice(i ...int) *Int {
	return NewInt().Assign(i...)
}

type i interface {
	Append(number int) *Int
	Assign(numbers ...int) *Int
	Bounds(i int) bool
	Concatenate(s *Int) *Int
	Each(f func(i int, number int)) *Int
	Empty() bool
	Fetch(i int) int
	Get(i int) (int, bool)
	Len() int
	Map(func(i int, number int) int) *Int
	Max() int
	Min() int
	Poll() int
	Pop() int
	Preassign(numbers ...int) *Int
	Precatenate(s *Int) *Int
	Prepend(number int) *Int
	Replace(i int, number int) bool
	Set() *Int
	Sort() *Int
	Sum() int
}

// Int is a superset of the Slice struct whose methods manage the access, insertion and modification of int only values.
type Int struct {
	slice *Slice
}

// Append method adds one int to the end of the Int Slice and returns the modified Int Slice.
func (pointer *Int) Append(number int) *Int {
	pointer.slice.Append(number)
	return pointer
}

// Assign method adds zero or more ints to the end of the Int Slice and returns the modified Int Slice.
func (pointer *Int) Assign(numbers ...int) *Int {
	for i := range numbers {
		pointer.Append(numbers[i])
	}
	return pointer
}

// Bounds checks an integer value safely sits within the range of accessible values for the Int Slice.
func (pointer *Int) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

// Concatenate merges two Int Slices into a single Int Slice.
func (pointer *Int) Concatenate(i *Int) *Int {
	pointer.slice.Concatenate(i.slice)
	return pointer
}

// Each method executes a provided function once for each Int Slice element.
func (pointer *Int) Each(f func(i int, number int)) *Int {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(int))
	})
	return pointer
}

// Empty returns a boolean indicating whether the Int Slice contains zero values.
func (pointer *Int) Empty() bool {
	return pointer.slice.Empty()
}

// Fetch retrieves the int held at the argument index. Returns nil int if index exceeds Int Slice length.
func (pointer *Int) Fetch(i int) int {
	return pointer.slice.Fetch(i).(int)
}

// Get returns the int held at the argument index and a boolean indicating if it was successfully retrieved.
func (pointer *Int) Get(i int) (int, bool) {
	n, ok := pointer.slice.Get(i)
	return n.(int), ok
}

// Len method returns the number of elements in the Int Slice.
func (pointer *Int) Len() int {
	return pointer.slice.Len()
}

// Map method executes a provided function once for each Int Slice element and sets the returned value to the current index.
func (pointer *Int) Map(f func(i int, number int) int) *Int {
	for i, value := range *pointer.slice {
		pointer.slice.Replace(i, f(i, value.(int)))
	}
	return pointer
}

// Max returns the largest int in the Int Slice.
func (pointer *Int) Max() int {
	if pointer.Empty() {
		return 0
	}
	i := 0
	j := pointer.Len() - 1
	m := 0
	for {
		a := pointer.Fetch(i)
		b := pointer.Fetch(j)
		if a > m {
			m = a
		}
		if b > m {
			m = b
		}
		i = i + 1
		j = j - 1
		if i >= j {
			break
		}
	}
	return m
}

// Min returns the smallest int in the Int Slice.
func (pointer *Int) Min() int {
	if pointer.Empty() {
		return 0
	}
	i := 1
	j := pointer.Len() - 1
	m := pointer.Fetch(0)
	for {
		a := pointer.Fetch(i)
		b := pointer.Fetch(j)
		if a < m {
			m = a
		}
		if b < m {
			m = b
		}
		i = i + 1
		j = j - 1
		if i >= j {
			break
		}
	}
	return m
}

// Poll method removes the first int from the Int Slice and returns that removed int.
func (pointer *Int) Poll() int {
	return pointer.slice.Poll().(int)
}

// Pop method removes the last int from the Int Slice and returns that int.
func (pointer *Int) Pop() int {
	return pointer.slice.Pop().(int)
}

// Preassign method adds zero or more ints to the beginning of the Int Slice and returns the modified Int Slice.
func (pointer *Int) Preassign(numbers ...int) *Int {
	for _, number := range numbers {
		pointer.slice.Prepend(number)
	}
	return pointer
}

// Precatenate merges two Int Slices, prepending the argument Int Slice.
func (pointer *Int) Precatenate(s *Int) *Int {
	pointer.slice.Precatenate(s.slice)
	return pointer
}

// Prepend method adds one int to the beginning of the Int Slice and returns the modified Int Slice.
func (pointer *Int) Prepend(number int) *Int {
	pointer.slice.Prepend(number)
	return pointer
}

// Replace method changes the contents of the Int Slice at the argument index if it is in bounds.
func (pointer *Int) Replace(i int, number int) bool {
	return pointer.slice.Replace(i, number)
}

// Set method returns a unique Int Slice, removing duplicate elements that have the same int value.
func (pointer *Int) Set() *Int {
	pointer.slice.Set()
	return pointer
}

// Sort alphabetically organises each element in the Int Slice.
func (pointer *Int) Sort() *Int {
	pointer.slice.Sort()
	return pointer
}

// Sum returns the total of all integers in the Int Slice.
func (pointer *Int) Sum() int {
	i := 0
	pointer.Each(func(_, n int) {
		i = i + n
	})
	return i
}
