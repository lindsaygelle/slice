# Slice
![Slice](https://repository-images.githubusercontent.com/192740394/a748b8c6-34ae-4aca-ad43-c18d5908b5e4)

Slice is a Go package that offers a versatile set of pre-built slices with extended functionality. It abstracts common list operations, such as appending, deleting, concatenating, mapping, and more, making it easier to work with slices in Go.

Gopher artwork was sourced from [egonelbre/gophers](https://github.com/egonelbre/gophers).

[![Go Reference](https://pkg.go.dev/badge/github.com/lindsaygelle/slice.svg)](https://pkg.go.dev/github.com/lindsaygelle/slice)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/slice)](/LICENSE)

## Features
* Common List Operations: Slice provides a comprehensive set of list operations like appending, deleting, concatenating, mapping, reversing, and more, all seamlessly integrated into single package.
* Memory Management: Slice takes care of allocating and deallocating memory as needed, making your code cleaner and safer.
* Type-Agnostic: Slice works effortlessly with a mix of data types, making it a valuable tool for dealing with collections of data.
* Generic Go Struct Wrapping: You can easily wrap the slice.Slice type in your custom Go struct to handle type-specific operations, allowing for robust, type-safe code.
* Intuitive Constructor Functions: Each Slice interface comes with a dedicated constructor function for creating slices conveniently.
* No Underlying Slice Exposure: To prevent mixing incompatible data types, Slice interfaces do not expose the underlying slice.Slice. While recommended for custom implementations, it's not mandatory.


## Installation
Add slice as a dependency to your Go project using the following command:

```sh
go get -u github.com/lindsaygelle/slice
```

## Docker
You can utilize slice within a Docker container with the provided Dockerfile. Here's how to build and run the container:

Building the Docker container.
```sh
docker build . -t slice
```

Developing and running Go from within the Docker container.
```sh
docker run -it --rm --name slice slice
```

A docker-compose file has also been added for convenience.
```sh
docker-compose up -d
```

That's it! You're ready to use the slice package to simplify your slice operations in Go. Explore the package documentation for a complete list of available methods and advanced usage patterns.

## Usage

### Creating a Slice
Slice implements all of the basic Golang types as slices. To create a slice of the required type, you can use a provided constructor function.

Creating a `string` slice.
```Go
// Example for string slice
stringSlice := slice.NewString("apple", "banana", "cherry")
```

Creating an `int` slice.
```Go
intSlice := slice.NewInt(1, 2, 10, 999)
```

Creating an `int64` slice.
```Go
int64Slice := slice.NewInt64(242252525823428, 987843616165656643890)
```

### Methods & Functionality
Each slice implements the same suite of methods. Here are some common operations:

```Go
type <SliceType> interface {
	// Append adds n elements to the end of the slice
	// and returns the modified slice.
	Append(i ...T) <SliceType>
	// AppendLength adds n elements to the end of the slice and returns the length of the modified slice.
	AppendLength(i ...T) int
	// Bounds checks an integer value safely sits within the range of
	// accessible values for the slice.
	Bounds(i int) bool
	// Concatenate merges the elements from the argument slice
	// to the the tail of the argument slice.
	Concatenate(s <SliceType>) <SliceType>
	// ConcatenateLength merges the elements from the argument slice to the tail of the receiver slice
	// and returns the length of the receiver slice.
	ConcatenateLength(s <SliceType>) int
	// Delete deletes the element from the argument index and returns the modified slice.
	Delete(i int) <SliceType>
	// DeleteLength deletes the element from the argument index and returns the new length of the slice.
	DeleteLength(i int) int
	// Each executes a provided function once for each slice element and returns the slice.
	Each(fn func(int, T)) <SliceType>
	// EachBreak executes a provided function once for each
	// element with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachBreak(fn func(int, T) bool) <SliceType>
	// EachReverse executes a provided function once for each
	// element in the reverse order they are stored in the slice.
	// Returns the slice at the end of the iteration.
	EachReverse(fn func(int, T)) <SliceType>
	// EachReverseBreak executes a provided function once for each
	// element in the reverse order they are stored in the slice
	// with an optional break when the function returns false.
	// Returns the slice at the end of the iteration.
	EachReverseBreak(fn func(int, T) bool) <SliceType>
	// Fetch retrieves the element held at the argument index.
	// Returns the default type if index exceeds slice length.
	Fetch(i int) T
	// FetchLength retrives the element held at the argument index and the length of the slice.
	// Returns the default type if index exceeds slice length.
	FetchLength(i int) (T, int)
	// Get returns the element held at the argument index and a boolean
	// indicating if it was successfully retrieved.
	Get(i int) (T, bool)
	// GetLength returns the element at the argument index, the length of the slice
	// and a boolean indicating if the element was successfully retrieved.
	GetLength(i int) (T, int, bool)
	// Len returns the number of elements in the slice.
	Len() int
	// Make empties the slice, sets the new slice to the length of n and returns the modified slice.
	Make(i int) <SliceType>
	// MakeEach empties the slice, sets the new slice to the length of n and performs
	// a for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEach(v ...T) <SliceType>
	// MakeEachReverse empties the slice, sets the new slice to the length of n and performs
	// an inverse for-each loop for the argument sequence, inserting each entry at the
	// appropriate index before returning the modified slice.
	MakeEachReverse(v ...T) <SliceType>
	// Map executes a provided function once for each element and sets
	// the returned value to the current index.
	// Returns the slice at the end of the iteration.
	Map(fn func(int, T) T) <SliceType>
	// Poll removes the first element from the slice and returns that removed element.
	Poll() T
	// PollLength removes the first element from the slice and returns the removed element and the length
	// of the modified slice.
	PollLength() (T, int)
	// PollOK removes the first element from the slice and returns a boolean on the outcome of the transaction.
	PollOK() (T, bool)
	// Pop removes the last element from the slice and returns that element.
	Pop() T
	// PopLength removes the last element from the slice and returns the removed element and the length
	// of the modified slice.
	PopLength() (T, int)
	// PopOK removes the last element from the slice and returns a boolean on the outcome of the transaction.
	PopOK() (T, bool)
	// Precatenate merges the elements from the argument slice
	// to the the head of the argument slice and returns the modified slice.
	Precatenate(s <SliceType>) <SliceType>
	// PrecatenateLength merges the elements from the argument slice to the head of the receiver slice
	// and returns the length of the receiver slice.
	PrecatenateLength(s <SliceType>) int
	// Prepend adds one element to the head of the slice
	// and returns the modified slice.
	Prepend(i ...T) <SliceType>
	// PrependLength adds n elements to the head of the slice and returns the length of the modified slice.
	PrependLength(i ...T) int
	// Push adds a new element to the end of the slice and
	// returns the length of the modified slice.
	Push(i ...T) int
	// Replace changes the contents of the slice
	// at the argument index if it is in bounds.
	Replace(i int, v T) bool
	// Reverse reverses the slice in linear time.
	// Returns the slice at the end of the iteration.
	Reverse() <SliceType>
	// Set returns a unique slice, removing duplicate
	// elements that have the same hash value.
	// Returns the modified at the end of the iteration.
	Set() <SliceType>
	// Slice slices the slice from i to j and returns the modified slice.
	Slice(i int, j int) <SliceType>
	// Swap moves element i to j and j to i.
	Swap(i int, j int)
	// Unshift adds one or more elements to the beginning of the slice and
	// returns the new length of the modified slice.
	Unshift(i ...T) int
	// Values returns the internal values of the slice.
	Values() []T
}
```

## Slices
Each slice is designed to handle a specific Go type. You can easily create, manipulate, and check bounds on these slices through the provided methods.

```Go
package main

import (
    "fmt"

    "github.com/lindsaygelle/slice"
)

var (
    // []byte
    b slice.Byte
    // []complex64
    c64 slice.Complex64
    // []complex128
    c128 slice.Complex128
    // []float32
    f32 slice.Float32
    // []float64
    f64 slice.Float64
    // []interface{}
    i slice.Int
    // []int8
    i8 slice.Int8
    // []int16
    i16 slice.Int16
    // []int32
    i32 slice.Int32
    // []int64
    i64 slice.Int64
    // []rune
    r slice.Rune
    // []interface{}
    s *slice.Slice
    // []uint
    u slice.UInt
    // []uint8
    u8 slice.UInt8
    // []uint16
    u16 slice.UInt16
    // []uint32
    u32 slice.UInt32
    // []uint64
    u64 slice.UInt64
)
```

```Go
import (
    "github.com/lindsaygelle/slice"
)

func main() {
    var (
        numbers = slice.NewInt(6, 1, 2, 3)
    )
    numbers.Sort().Each(func(i int, n int) {
        fmt.Println(i, n) // (0, 1), (1, 2), (2, 3), (3, 6)
    })
}
```

## Extending

Slice supports interface extension by wrapping the Slice in an struct and exposing a corresponding interface. This is the same pattern implemented by this package and is used for the provided interfaces.

```Go
package food

import (
    "github.com/lindsaygelle/slice"
)

// Food is a struct that describes food.
type Food struct {
    Name string
}

// FoodSlice is an interface that contains a collection of Food.
type FoodSlice interface {
    Append(Food) FoodSlice
    Prepend(Food) FoodSlice
}

// food is a struct that interfaces with slice.Slice.
type food struct { s *slice.Slice }

// Append adds Food structs to the tail of the FoodSlice.
func (f *food) Append(food ...Food) FoodSlice {
    f.s.Append(food...)
    return f
}

// Prepend adds Food structs to the head of the FoodSlice.
func (f *food) Prepend(food ...Food) FoodSlice {
    f.s.Prepend(food...)
    return f
}
```

## Contributing
Contributions to Slice are welcome! If you have any ideas, bug reports, or enhancements, please submit them as GitHub issues or create a pull request with your changes. For major contributions, it is recommended to discuss your ideas first by creating an issue to ensure alignment with the project's goals and direction. Please see the [CONTRIBUTION](./CONTRIBUTING.md) file fore more details.

## Branching
For branch conventions and guidelines, please review the [BRANCH_NAMING_CONVENTION](./BRANCH_NAMING_CONVENTION.md) for more details.

## Roadmap
Planned roadmap for [slice](https://github.com/lindsaygelle/slice).

- [ ] Create GitHub action to automatically generated changelog edits.
- [ ] Create tests for each slice implementation.
- [ ] Create an examples folder.
- [ ] Could all be replaced with generics?

## License
Slice is licensed under the MIT License. Feel free to use, modify, and distribute the code within this repository as per the terms of the license. Please see the [LICENSE](./LICENSE) file for more details.
