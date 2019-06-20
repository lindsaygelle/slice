package slice

import "fmt"

var (
	_ s = (*String)(nil)
)

func NewString() *String {
	return &String{
		slice: &Slice{}}
}

func NewStringSlice(s ...string) *String {
	return NewString().Assign(s...)
}

type s interface {
	Append(value string) *String
	Assign(values ...string) *String
	Bounds(i int) bool
	Concatenate(s *String) *String
	Each(f func(i int, value string)) *String
	Fetch(i int) string
	Get(i int) (string, bool)
	Map(func(i int, value string) string) *String
	Poll() string
	Pop() string
	Preassign(values ...string) *String
	Precatenate(s *String) *String
	Prepend(value string) *String
	Replace(i int, value string) bool
}

// String is a superset of the Slice struct whose methods manage the access, insertion and modification of string only values.
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

func (pointer *String) Each(f func(i int, s string)) *String {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(string))
	})
	return pointer
}

func (pointer *String) Fetch(i int) string {
	return pointer.slice.Fetch(i).(string)
}

func (pointer *String) Get(i int) (string, bool) {
	s, ok := pointer.slice.Get(i)
	return fmt.Sprintf("%v", s), ok
}

func (pointer *String) Len() int {
	return pointer.slice.Len()
}
func (pointer *String) Map(f func(i int, value string) string) *String {
	for i, value := range *pointer.slice {
		pointer.slice.Replace(i, f(i, value.(string)))
	}
	return pointer
}

func (pointer *String) Poll() string {
	return fmt.Sprintf("%v", pointer.slice.Poll())
}

func (pointer *String) Pop() string {
	return fmt.Sprintf("%v", pointer.slice.Pop())
}

func (pointer *String) Preassign(s ...string) *String {
	for _, s := range s {
		pointer.slice.Prepend(s)
	}
	return pointer
}

func (pointer *String) Precatenate(s *String) *String {
	pointer.slice.Precatenate(s.slice)
	return pointer
}

func (pointer *String) Prepend(s string) *String {
	pointer.slice.Prepend(s)
	return pointer
}

func (pointer *String) Replace(i int, s string) bool {
	return pointer.slice.Replace(i, s)
}
