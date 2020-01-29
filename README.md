[![Build Status](https://travis-ci.org/gellel/slice.svg?branch=master)](https://travis-ci.org/gellel/slice)
[![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/gellel/slice/blob/master/LICENSE)

# Slice

Package slice is a package of slice interfaces to handle common list-like operations on a collection of Go interfaces. Slice takes care of the allocation and deallocation of pointers and provides common list-like operations on the reference without the concern for handling the referencing and dereferencing of the pointer.

Slice contains a common slice type (exported as `slice.Slice`) that provides methods to perform traversal and mutation operations for a collection of Go interfaces. The `slice.Slice` struct can be wrapped in a generic Go struct to handle the allocation and retrieval of specific types. This has been done for all (as of writing this README) Go data types.

Package slice comes with all Go primative types as interfaces out of the box. Each interface is indicated by the _er_ suffix. 

Each slice interface comes with a unique constructor function that takes zero to n arguments of the corresponding slice interface type.

The slice interfaces do not expose the underlying `slice.Slice` to prevent a mixed collection. It is recommended to adopt this pattern when creating a custom implementation, but not required.

The package is built around the Go documentation pattern. Please consider using `godoc` when using this package. If you are using Go 1.12 or earlier, `godoc` should be included. All Go 1.13 users will need to grab this package using the `go get` flow.

## Installing

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or project's Go module dependencies.

```go get github.com/gellel/slice```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

```go get -u github.com/gellel/slice```

## Dependencies

The SDK includes a vendor folder containing the runtime dependencies of the SDK. The metadata of the SDK's dependencies can be found in the Go module file go.mod.

## Go Modules

If you are using Go modules, your go get will default to the latest tagged release version of the SDK. To get a specific release version of the SDK use `@<tag>` in your `go get` command.

```go get github.com/gelle/slice@<version>```

To get the latest SDK repository change use @latest.

## License

This SDK is distributed under the Apache License, Version 2.0, see LICENSE for more information.

## Snippets

Slice exports all primative Go types as interfaces. 

```Go
package main

import (
    "fmt"

    "github.com/gellel/slice"
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

Each interface is intended to handle a unique Go lang primative type.

A Slice interface implements all methods of slice.Slice.

```Go

import (
    "github.com/gellel/slice"
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
    "github.com/gellel/slice"
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
