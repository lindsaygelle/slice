// Package slice provides interface for managing collections of abstract data in an Array-like structure.
package slice

import (
	"fmt"
	"sort"
	"strings"
)

var (
	_ slice = (*Slice)(nil)
)

// new instantiates a new empty or populated Slice value.
func new(values ...interface{}) Slice {
	slice := Slice{}
	for _, value := range values {
		slice = append(slice, value)
	}
	return slice
}

// New instantiates a new empty or populated, Slice pointer.
// Slice pointers are mutable and hold and arguments as an interface.
// Any dataset retrieved from the Slice is intended to be cast to a required type.
// Unlike basic list-like objects, the Slice provides safe getters and setters,
// aiming to reduce the likelyhood of an exception being thrown during an operation.
func New(values ...interface{}) *Slice {
	return (&Slice{}).Assign(values...)
}

type slice interface {
	Append(value interface{}) *Slice
	Assign(values ...interface{}) *Slice
	Bounds(i int) bool
	Concatenate(slice *Slice) *Slice
	Each(f func(i int, value interface{})) *Slice
	Fetch(i int) interface{}
	Get(i int) (interface{}, bool)
	Join(character string) string
	Len() int
	Less(i, j int) bool
	Map(func(i int, value interface{}) interface{}) *Slice
	Poll() interface{}
	Pop() interface{}
	Preassign(values ...interface{}) *Slice
	Precatenate(slice *Slice) *Slice
	Prepend(value interface{}) *Slice
	Replace(i int, value interface{}) bool
	Set() *Slice
	Slice(start, end int) *Slice
	Splice(start, end int) *Slice
	Sort() *Slice
	Swap(i, j int)
}

// Slice is a list-like object whose methods are used to perform traversal and mutation operations.
type Slice []interface{}

// Append method adds one element to the end of the Slice and returns the modified Slice.
func (pointer *Slice) Append(value interface{}) *Slice {
	(*pointer) = append(*pointer, value)
	return pointer
}

// Assign method adds zero or more elements to the end of the Slice and returns the modified Slice.
func (pointer *Slice) Assign(values ...interface{}) *Slice {
	(*pointer) = append(*pointer, values...)
	return pointer
}

// Bounds checks an integer value safely sits within the range of accessible values for the Slice.
func (pointer *Slice) Bounds(i int) bool {
	return ((i > -1) && (i < len(*pointer)))
}

// Concatenate merges two Slices into a single Slice.
func (pointer *Slice) Concatenate(slice *Slice) *Slice {
	entries := []interface{}{}
	entries = append(entries, *pointer...)
	entries = append(entries, *slice...)
	(*pointer) = new(entries...)
	return pointer
}

// Each method executes a provided function once for each Slice element.
func (pointer *Slice) Each(f func(i int, value interface{})) *Slice {
	for i, value := range *pointer {
		f(i, value)
	}
	return pointer
}

// Fetch retrieves the interface held at the argument index. Returns nil if index exceeds Slice length.
func (pointer *Slice) Fetch(i int) interface{} {
	value, _ := pointer.Get(i)
	return value
}

// Get returns the interface held at the argument index and a boolean indicating if it was successfully retrieved.
func (pointer *Slice) Get(i int) (interface{}, bool) {
	ok := pointer.Bounds(i)
	if ok == true {
		return (*pointer)[i], ok
	}
	return nil, ok
}

// Join merges all elements in the Slice into a single string.
func (pointer *Slice) Join(character string) string {
	s := []string{}
	pointer.Each(func(_ int, i interface{}) {
		s = append(s, fmt.Sprintf("%v", i))
	})
	return strings.Join(s, character)
}

// Len method returns the number of elements in the Slice.
func (pointer *Slice) Len() int {
	return len(*pointer)
}

// Less checks the string value of two elements in the slice and checks which element has the lower value.
func (pointer *Slice) Less(i, j int) bool {
	s := *pointer
	a, b := fmt.Sprintf("%s", s[i]), fmt.Sprintf("%s", s[j])
	if ok := (a == b); ok {
		a, b = strings.ToLower(a), strings.ToLower(b)
	}
	ok := a < b
	return ok
}

// Map method executes a provided function once for each Slice element and sets the returned value to the current index.
func (pointer *Slice) Map(f func(i int, value interface{}) interface{}) *Slice {
	for i, value := range *pointer {
		response := f(i, value)
		if response != nil {
			pointer.Replace(i, response)
		}
	}
	return pointer
}

// Poll method removes the first element from the Slice and returns that removed element.
func (pointer *Slice) Poll() interface{} {
	length := pointer.Len()
	ok := length > 0
	if ok == true {
		value := (*pointer)[0]
		(*pointer) = (*pointer)[0:]
		return value
	}
	return nil
}

// Pop method removes the last element from the Slice and returns that element.
func (pointer *Slice) Pop() interface{} {
	length := pointer.Len()
	ok := length > 0
	if ok == true {
		value := (*pointer)[length-1]
		(*pointer) = (*pointer)[:length-1]
		return value
	}
	return nil
}

// Preassign method adds zero or more elements to the beginning of the Slice and returns the modified Slice.
func (pointer *Slice) Preassign(values ...interface{}) *Slice {
	(*pointer) = append(new(values...), (*pointer)...)
	return pointer
}

// Precatenate merges two Slices, prepending the argument Slice.
func (pointer *Slice) Precatenate(slice *Slice) *Slice {
	entries := []interface{}{}
	entries = append(entries, *slice...)
	entries = append(entries, *pointer...)
	(*pointer) = new(entries...)
	return pointer
}

// Prepend method adds one element to the beginning of the Slice and returns the modified Slice.
func (pointer *Slice) Prepend(value interface{}) *Slice {
	(*pointer) = append(Slice{value}, (*pointer)...)
	return pointer
}

// Replace method changes the contents of the Slice at the argument index if it is in bounds.
func (pointer *Slice) Replace(i int, value interface{}) bool {
	ok := pointer.Bounds(i)
	if ok == true {
		(*pointer)[i] = value
	}
	return ok
}

// Set method returns a unique Slice, removing duplicate elements that have the same hash value.
func (pointer *Slice) Set() *Slice {
	slice := Slice{}
	m := map[string]bool{}
	for _, value := range *pointer {
		key := fmt.Sprintf("%v", value)
		if _, ok := m[key]; ok != true {
			slice = append(slice, value)
		}
		m[key] = true
	}
	(*pointer) = slice
	return pointer
}

// Slice method returns a shallow copy of a portion of the Slice into a new Slice type selected from begin to end (end not included).
// The original Slice will not be modified.
func (pointer *Slice) Slice(start, end int) *Slice {
	if ok := start < end; ok != true {
		return pointer.Slice(end, start)
	}
	if ok := pointer.Bounds(start) && pointer.Bounds(end); ok != true {
		return &Slice{}
	}
	return (&Slice{}).Assign((*pointer)[start:end]...)
}

// Splice method changes the contents of the Slice by removing existing elements.
func (pointer *Slice) Splice(start, end int) *Slice {
	if ok := start < end; ok != true {
		return pointer.Splice(end, start)
	}
	if ok := pointer.Bounds(start) && pointer.Bounds(end); ok != true {
		return pointer
	}
	(*pointer) = (*pointer)[start:end]
	return pointer
}

// Sort alphabetically organises each element in the Slice.
func (pointer *Slice) Sort() *Slice {
	sort.Sort(pointer)
	return pointer
}

// Swap moves element i to j and j to i.
func (pointer *Slice) Swap(i int, j int) {
	s := *pointer
	s[i], s[j] = s[j], s[i]
	*pointer = s
}
