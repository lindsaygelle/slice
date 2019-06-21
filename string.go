package slice

import "fmt"

var (
	_ s = (*String)(nil)
)

// NewString instantiates a new empty String slice.
func NewString() *String {
	return &String{
		slice: &Slice{}}
}

// NewStringSlice instantiates a new populated or empty String slice.
func NewStringSlice(s ...string) *String {
	return NewString().Assign(s...)
}

type s interface {
	Append(value string) *String
	Assign(values ...string) *String
	Bounds(i int) bool
	Concatenate(s *String) *String
	Each(f func(i int, value string)) *String
	Empty() bool
	Fetch(i int) string
	Get(i int) (string, bool)
	Join(character string) string
	Len() int
	Map(func(i int, value string) string) *String
	Poll() string
	Pop() string
	Preassign(values ...string) *String
	Precatenate(s *String) *String
	Prepend(value string) *String
	Push(value string) int
	Replace(i int, value string) bool
	Set() *String
	Sort() *String
}

// String is a superset of the Slice struct whose methods manage the access, insertion and modification of string only values.
type String struct {
	slice *Slice
}

// Append method adds one string to the end of the String Slice and returns the modified String Slice.
func (pointer *String) Append(s string) *String {
	pointer.slice.Append(s)
	return pointer
}

// Assign method adds zero or more strings to the end of the String Slice and returns the modified String Slice.
func (pointer *String) Assign(s ...string) *String {
	for i := range s {
		pointer.Append(s[i])
	}
	return pointer
}

// Bounds checks an integer value safely sits within the range of accessible values for the String Slice.
func (pointer *String) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

// Concatenate merges two String Slices into a single String Slice.
func (pointer *String) Concatenate(s *String) *String {
	pointer.slice.Concatenate(s.slice)
	return pointer
}

// Each method executes a provided function once for each String Slice element.
func (pointer *String) Each(f func(i int, s string)) *String {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(string))
	})
	return pointer
}

// Empty returns a boolean indicating whether the String Slice contains zero values.
func (pointer *String) Empty() bool {
	return pointer.slice.Empty()
}

// Fetch retrieves the string held at the argument index. Returns nil string if index exceeds String Slice length.
func (pointer *String) Fetch(i int) string {
	s, _ := pointer.Get(i)
	return s
}

// Get returns the string held at the argument index and a boolean indicating if it was successfully retrieved.
func (pointer *String) Get(i int) (string, bool) {
	s, ok := pointer.slice.Get(i)
	return fmt.Sprintf("%v", s), ok
}

// Join merges all elements in the String Slice into a single string.
func (pointer *String) Join(character string) string {
	return pointer.slice.Join(character)
}

// Len method returns the number of elements in the String Slice.
func (pointer *String) Len() int {
	return pointer.slice.Len()
}

// Map method executes a provided function once for each String Slice element and sets the returned value to the current index.
func (pointer *String) Map(f func(i int, value string) string) *String {
	for i, value := range *pointer.slice {
		pointer.slice.Replace(i, f(i, value.(string)))
	}
	return pointer
}

// Poll method removes the first string from the String Slice and returns that removed string.
func (pointer *String) Poll() string {
	return fmt.Sprintf("%v", pointer.slice.Poll())
}

// Pop method removes the last string from the String Slice and returns that string.
func (pointer *String) Pop() string {
	return fmt.Sprintf("%v", pointer.slice.Pop())
}

// Preassign method adds zero or more strings to the beginning of the String Slice and returns the modified String Slice.
func (pointer *String) Preassign(s ...string) *String {
	for _, s := range s {
		pointer.slice.Prepend(s)
	}
	return pointer
}

// Precatenate merges two String Slices, prepending the argument String Slice.
func (pointer *String) Precatenate(s *String) *String {
	pointer.slice.Precatenate(s.slice)
	return pointer
}

// Prepend method adds one string to the beginning of the String Slice and returns the modified String Slice.
func (pointer *String) Prepend(s string) *String {
	pointer.slice.Prepend(s)
	return pointer
}

// Push method adds a new string to the end of the String Slice and returns the length of the modified String Slice.
func (pointer *String) Push(s string) int {
	return pointer.slice.Push(s)
}

// Replace method changes the contents of the String Slice at the argument index if it is in bounds.
func (pointer *String) Replace(i int, s string) bool {
	return pointer.slice.Replace(i, s)
}

// Set method returns a unique String Slice, removing duplicate elements that have the same string value.
func (pointer *String) Set() *String {
	pointer.slice.Set()
	return pointer
}

// Sort alphabetically organises each element in the String Slice.
func (pointer *String) Sort() *String {
	pointer.slice.Sort()
	return pointer
}
