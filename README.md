# Slice

The Slice package is a versatile Go library designed to simplify common list-like operations on collections of Go interfaces. It abstracts away the complexities of memory allocation and deallocation of pointers while providing a rich set of list operations. Whether you're dealing with a homogeneous or heterogeneous collection, Slice has you covered.

Key Features:

* Common List Operations: Slice provides a comprehensive set of list operations like appending, deleting, concatenating, mapping, reversing, and more, all seamlessly integrated into a simple API.
* Dynamic Memory Management: Forget about manual memory management; Slice takes care of allocating and deallocating memory as needed, making your code cleaner and safer.
* Type-Agnostic: Slice works effortlessly with a mix of data types, making it a valuable tool for dealing with collections of heterogeneous data.
* Generic Go Struct Wrapping: You can easily wrap the slice.Slice type in your custom Go struct to handle type-specific operations, allowing for robust, type-safe code.
* Intuitive Constructor Functions: Each Slice interface comes with a dedicated constructor function for creating slices conveniently.
* No Underlying Slice Exposure: To prevent mixing incompatible data types, Slice interfaces do not expose the underlying slice.Slice. While recommended for custom implementations, it's not mandatory.

# Usage:

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

## Docker-compose
A [`docker-compose.yml`](./docker-compose.yml) file has also been added for convenience.

# Usage
To get started with the slice package, follow these steps:

Install the slice package using the go get command:

```sh
go get github.com/lindsaygelle/slice
```

Import the package in your Go code:

```Go
import "github.com/lindsaygelle/slice"
```

Create Slice interfaces for the specific data types you want to work with. Use the provided constructor functions to initialize them.

```Go
// Example for string slice
s := slice.NewStringer("apple", "banana", "cherry")
```

Use the methods provided by the Slice interface to perform various list-like operations. Here are some common operations:

```Go
// Append elements to the slice
s.Append("date", "fig")

// Check if an index is within bounds
inBounds := s.Bounds(2)

// Pop the last element from the slice
poppedElement := s.Pop()
```

You can sort the slice and iterate over its elements easily:

```Go
// Example for sorting and iteration
sortedSlice := s.Sort()
sortedSlice.Each(func(index int, value string) {
    fmt.Println(index, value)
})
```

That's it! You're ready to use the slice package to simplify your slice operations in Go. Explore the package documentation for a complete list of available methods and advanced usage patterns.


## Snippets

Each Slice interface is designed to handle a specific Go primitive type. You can easily create, manipulate, and check bounds on these slices.

```Go
package main

import (
    "fmt"

    "github.com/lindsaygelle/slice"
)

var (
    b    slice.Byter        // []byte
    c64  slice.Complexer64  // []complex64
    c128 slice.Complexer128 // []complex128
    f32  slice.Floater32    // []float32
    f64  slice.Floater64    // []float64
    i    slice.Inter        // []interface{}
    i8   slice.Inter8       // []int8
    i16  slice.Inter16      // []int16
    i32  slice.Inter32      // []int32
    i64  slice.Inter64      // []int64
    r    slice.Runer        // []rune
    s    *slice.Slice       // []interface{}
    u    slice.UInter       // []uint
    u8   slice.UInter8      // []uint8
    u16  slice.UInter16     // []uint16
    u32  slice.UInter32     // []uint32
    u64  slice.UInter64     // []uint64
    v    slice.Interfacer   // []interface{}
)

func main() {
    var (
        s = slice.NewStringer("a", "b", "c", "go!")
    )
    s.Bounds(0)          // true
    fmt.Println(s.Pop()) // "go!"
}
```

Each interface is intended to handle a unique Go lang primitive type.

A Slice interface implements all methods of slice.Slice.

```Go

import (
    "github.com/lindsaygelle/slice"
)

func main() {

    var (
        numbers = slice.NewInter(6, 1, 2, 3)
    )
    numbers.Sort().Each(func(i int, n int) {
        fmt.Println(i, n) // (0, 1), (1, 2), (2, 3), (3, 6)
    })
}
```

## Extending

Slice supports interface extension by wrapping the Slice in an struct and exposing a corresponding interface.

This is the pattern implemented by this package and is used for the provided interface types.

```Go
package food

import (
    "github.com/lindsaygelle/slice"
)

// Food is a struct that describes food.
type Food struct {
    Name string
}

// Fooder is an interface that contains a collection of Food.
type Fooder interface {
    Append(Food) Fooder
    Prepend(Food) Fooder
}

// fooder is a struct that interfaces with slice.Slice.
type fooder struct { s *slice.Slice }

// Append adds Food structs to the tail of the Fooder.
func (f *fooder) Append(food ...Food) Fooder {
    f.s.Append(food...)
    return f
}

// Prepend adds Food structs to the head of the Fooder.
func (f *fooder) Prepend(food ...Food) Fooder {
    f.s.Prepend(food...)
    return f
}
```

## Contributing
Contributions to Slice are welcome! If you have any ideas, bug reports, or enhancements, please submit them as GitHub issues or create a pull request with your changes. For major contributions, it is recommended to discuss your ideas first by creating an issue to ensure alignment with the project's goals and direction. Please see the [CONTRIBUTION](./CONTRIBUTING.md) file fore more details.

## License
Slice is licensed under the MIT License. Feel free to use, modify, and distribute the code within this repository as per the terms of the license. Please see the [LICENSE](./LICENSE) file for more details.
