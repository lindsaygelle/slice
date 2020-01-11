[![Build Status](https://travis-ci.org/gellel/slice.svg?branch=master)](https://travis-ci.org/gellel/slice)
[![Apache V2 License](https://img.shields.io/badge/license-Apache%20V2-blue.svg)](https://github.com/gellel/slice/blob/master/LICENSE)

# Slice

Slice is a package of interfaces to add functionality to slice-like structs.

The package is built around the Go API reference documentation. Please consider using `godoc`
to build custom integrations. If you are using Go 1.12 or earlier, godoc should be included. All
Go 1.13 users will need to grab this package using the `go get` flow.

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

This SDK is distributed under the Apache License, Version 2.0, see LICENSE.txt and NOTICE.txt for more information.

## Exports

Slice exports all base Go types as interfaces.

```Go
package main

var (
    f32 slice.Floater32
    f64 slice.Floater64
    i   slice.Inter
    i8  slice.Inter8
    i16 slice.Inter16
    i32 slice.Inter32
    i64 slice.Inter64
    v   slice.Interfacer
)

func main() {}
```

## Example

```Go
package main 

import (
    "github.com/gellel/slice"
)

var (
    // integers is a slice of Go int.
    integers = slice.NewInteger(1, 2, 3, 9)
)

func main() {}
```