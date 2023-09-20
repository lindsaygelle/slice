# Slice

## ![Slice](https://repository-images.githubusercontent.com/192740394/a748b8c6-34ae-4aca-ad43-c18d5908b5e4)

Slice is a Go package that offers a versatile set of pre-built slices with extended functionality. It abstracts common list operations, such as appending, deleting, concatenating, mapping, and more, making it easier to work with slices in Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/lindsaygelle/slice)](https://goreportcard.com/report/github.com/lindsaygelle/slice)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/lindsaygelle/slice)](https://github.com/lindsaygelle/slice/releases)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/lindsaygelle/slice)](https://pkg.go.dev/github.com/lindsaygelle/slice)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/slice)](LICENSE.txt)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)


## Installation
Getting started with Slice is a breeze. You can install it in your Go project using `go get`:

```sh
go get github.com/lindsaygelle/slice
```

## Usage
To begin using Slice, simply import the package into your Go code:

```Go
import (
	"github.com/lindsaygelle/slice"
)
```

Creating and Initializing a Slice:
```Go
// Create an empty slice of integers
s := &slice.Slice[int]{}
```

Appending Elements to a Slice:
```Go
// Append values to the slice
s.Append(1, 2, 3)
```

Getting the Length of a Slice:
```Go
// Get the length of the slice
length := s.Length()
```

Deleting an Element from a Slice:
```Go
// Delete an element at index 2
s.Delete(2)
```

Iterating Over a Slice:
```Go
// Iterate over the slice and print each element
s.Each(func(index int, value int) {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
})
```

Reversing a Slice:
```Go
// Reverse the order of elements in the slice
s.Reverse()
```

Slicing a Slice:
```Go
// Slice the slice from index 1 to 3
s.Slice(1, 3)
```

Swapping Elements in a Slice:
```Go
// Swap elements at indices 1 and 2
s.Swap(1, 2)
```

More complicated examples:
```Go
// Create a slice of strings
strSlice := &slice.Slice[string]{"apple", "banana", "cherry"}

// Append multiple values to the slice
strSlice.Append("date", "elderberry")

// Check if the slice contains a specific value
containsCherry := strSlice.Contains("cherry") // Should return true

// Replace the element at index 2 with "grape"
strSlice.Replace(2, "grape")

// Get the length of the slice
strLength := strSlice.Length()

// Iterate over the slice and print each element
strSlice.Each(func(index int, value string) {
    fmt.Printf("Index %d: %s\n", index, value)
})
```

Using a complex type:
```Go
// Define a custom struct
type Person struct {
    Name  string
    Age   int
    Email string
}

// Create a slice of Person structs
people := &slice.Slice[Person]{
    {Name: "Alice", Age: 30, Email: "alice@example.com"},
    {Name: "Bob", Age: 25, Email: "bob@example.com"},
}

// Append a new person to the slice
newPerson := Person{Name: "Charlie", Age: 35, Email: "charlie@example.com"}
people.Append(newPerson)

// Find the index of a person with a specific email address
index := people.FindIndex(func(p Person) bool {
    return p.Email == "bob@example.com"
})

// Slice the slice to include only people aged 30 or older
people.Slice(1, people.Length())

// Reverse the order of people in the slice
people.Reverse()

// Iterate over the slice and print each person's details
people.Each(func(index int, person Person) {
    fmt.Printf("Index %d: Name: %s, Age: %d, Email: %s\n", index, person.Name, person.Age, person.Email)
})
```

## Docker
Slice is Docker-friendly! You can easily incorporate Slice into a Docker container using the provided Dockerfile. Here are the steps to build and run the container:

Building the Docker container:
```sh
docker build . -t slice
```

Developing and running Go within the Docker container:
```sh
docker run -it --rm --name slice slice
```

A docker-compose file has also been included for convenience:
```sh
docker-compose up -d
```

## Contributing
We warmly welcome contributions to Slice. Whether you have innovative ideas, bug reports, or enhancements in mind, please share them with us by submitting GitHub issues or creating pull requests. For substantial contributions, it's a good practice to start a discussion by creating an issue to ensure alignment with the project's goals and direction. Refer to the [CONTRIBUTING](./CONTRIBUTING.md) file for comprehensive details.

## Branching
For a smooth collaboration experience, we have established branch naming conventions and guidelines. Please consult the [BRANCH_NAMING_CONVENTION](./BRANCH_NAMING_CONVENTION.md) document for comprehensive information and best practices.

## License
Slice is released under the MIT License, granting you the freedom to use, modify, and distribute the code within this repository in accordance with the terms of the license. For additional information, please review the [LICENSE](./LICENSE) file.

## Acknowledgements
We express our gratitude to [egonelbre/gophers](https://github.com/egonelbre/gophers) for providing the delightful Gopher artwork used in our social preview. Don't hesitate to pay them a visit!
