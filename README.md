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

## Usage
To create a slice for a provided data type, you can use a provided constructor function.

```Go
// Example for string slice
s := slice.NewString("apple", "banana", "cherry")
```

Use the methods provided by the initialized slice interface to perform various list-like operations. Here are some common operations:

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

## License
Slice is licensed under the MIT License. Feel free to use, modify, and distribute the code within this repository as per the terms of the license. Please see the [LICENSE](./LICENSE) file for more details.
