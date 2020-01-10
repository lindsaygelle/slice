package slice

type Slice []interface{}

// Append method adds one element to the end of the collection
// and returns the modified collection.
func (slice *Slice) Append(i ...interface{}) *Slice {
	(*slice) = (append(*slice, i...))
	return slice
}

// Bounds checks an integer value safely sits within the range of
// accessible values for the collection.
func (slice *Slice) Bounds(i int) bool {
	return ((i > -1) && (i < len(*slice)))
}

// Concatenate merges the elements from the argument collection
// to the the tail of the argument collection.
func (slice *Slice) Concatenate(s *Slice) *Slice {
	slice.Append((*s)...)
	return slice
}

// Each method executes a provided function once for each collection element.
func (slice *Slice) Each(fn func(int, interface{})) {
	var (
		i int
		v interface{}
	)
	for i, v = range *slice {
		fn(i, v)
	}
}

// EachBreak method executes a provided function once for each
// element with an optional break when the function returns false.
func (slice *Slice) EachBreak(fn func(int, interface{}) bool) {
	var (
		i  int
		ok = true
		v  interface{}
	)
	for i, v = range *slice {
		ok = fn(i, v)
		if !ok {
			break
		}
	}
}

// EachReverse method executes a provided function once for each
// element in the reverse order they are stored in the *Slice.
func (slice *Slice) EachReverse(fn func(int, interface{})) {
	var (
		i int
	)
	for i = len(*slice) - 1; i >= 0; i-- {
		fn(i, (*slice)[i])
	}
}

// EachReverseBreak method executes a provided function once for each
// element in the reverse order they are stored in the collection
// with an optional break when the function returns false.
func (slice *Slice) EachReverseBreak(fn func(int, interface{}) bool) {
	var (
		i  int
		ok = true
	)
	for i = len(*slice) - 1; i >= 0; i-- {
		ok = fn(i, (*slice)[i])
		if !ok {
			break
		}
	}
}

// Fetch retrieves the element held at the argument index.
// Returns the default type if index exceeds collection length.
func (slice *Slice) Fetch(i int) interface{} {
	var v, _ = slice.Get(i)
	return v
}

// Get returns the element held at the argument index and a boolean
// indicating if it was successfully retrieved.
func (slice *Slice) Get(i int) (interface{}, bool) {
	var (
		ok = slice.Bounds(i)
	)
	if ok {
		return (*slice)[i], ok
	}
	return nil, ok
}

// Len method returns the number of elements in the collection.
func (slice *Slice) Len() int { return (len(*slice)) }

// Map method executes a provided function once for each element and sets
// the returned value to the current index.
func (slice *Slice) Map(fn func(int, interface{}) interface{}) {
	slice.Each(func(i int, v interface{}) {
		slice.Replace(i, fn(i, v))
	})
}

// Precatenate merges the elements from the argument collection
// to the the head of the argument collection.
func (slice *Slice) Precatenate(s *Slice) *Slice {
	slice.Prepend((*s)...)
	return slice
}

// Prepend method adds one element to the head of the collection
// and returns the modified collection.
func (slice *Slice) Prepend(i ...interface{}) *Slice {
	(*slice) = (append(i, *slice...))
	return slice
}

// Replace method changes the contents of the collection
// at the argument index if it is in bounds.
func (slice *Slice) Replace(i int, v interface{}) bool {
	var (
		ok = slice.Bounds(i)
	)
	if ok {
		(*slice)[i] = v
	}
	return ok
}

// Set method returns a unique Slice, removing duplicate
// elements that have the same hash value.
func (slice *Slice) Set() *Slice {
	return slice
}

// Swap moves element i to j and j to i.
func (slice *Slice) Swap(i int, j int) {
	(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
}
