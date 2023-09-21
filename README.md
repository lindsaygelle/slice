# Slice
Slice is a [Go](https://github.com/golang/go) package that provides a generic slice with extended functionality. It abstracts common list operations, such as appending, deleting, concatenating, mapping, and more, making it easier to work with slices in Go.

![Slice](https://repository-images.githubusercontent.com/192740394/a748b8c6-34ae-4aca-ad43-c18d5908b5e4)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/lindsaygelle/slice)](https://pkg.go.dev/github.com/lindsaygelle/slice)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindsaygelle/slice)](https://goreportcard.com/report/github.com/lindsaygelle/slice)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/lindsaygelle/slice)](https://github.com/lindsaygelle/slice/releases)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/slice)](LICENSE.txt)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

## Features

## ✨ Enhanced Functionality
Slice abstracts the intricacies of working with slices in Go, providing a comprehensive set of operations to simplify your code and reduce redundancy.

## 🚀 Seamless Integration
Getting started with Slice is straightforward. Import the package, provide a type, and you're ready to leverage its capabilities within your Go projects.


## Installation
You can install it in your Go project using `go get`:

```sh
go get github.com/lindsaygelle/slice
```

## Usage
Import the package into your Go code:

```Go
import (
	"github.com/lindsaygelle/slice"
)
```

## Methods
Provided methods for `&slice.Slice[T]`.

### Append
```Go
s := &slice.Slice[int]{1, 2, 3}
s.Append(4, 5) // s is now [1, 2, 3, 4, 5]
```

### AppendFunc
```Go
s := &slice.slice.Slice[int]{}
s.AppendFunc(func(i int, value int) bool {
	return value%2 == 0 // Append even numbers to the Slice.
}, 1, 2, 3, 4, 5)
```

### AppendLength
```Go
s := &slice.Slice[int]{1, 2, 3}
length := s.AppendLength(4, 5) // s is now [1, 2, 3, 4, 5], length is 5
```

### Bounds
```Go
s := &slice.Slice[int]{1, 2, 3}
inBounds := s.Bounds(1) // inBounds is true
outOfBounds := s.Bounds(5) // outOfBounds is false
```

### Concatenate
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5}
s1.Concatenate(s2) // s1 is now [1, 2, 3, 4, 5]
```

### ConcatenateFunc
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5, 6}
s1.ConcatenateFunc(s2, func(i int, value int) bool {
	return value%2 == 0
}) // s1 is now [1, 2, 3, 4, 6]
```

### ConcatenateLength
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5}
length := s1.ConcatenateLength(s2) // s1 is now [1, 2, 3, 4, 5], length is 5
```

### Contains
```Go
s := &slice.Slice[string]{"apple", "banana", "cherry"}
containsBanana := s.Contains("banana") // containsBanana is true
```

### ContainsMany
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
result := s.ContainsMany(2, 4, 6) // result will be [true, true, false]
```

### Delete
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.Delete(2) // s is now [1, 2, 4, 5]
```

### DeleteFunc
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
even := func(i int, value int) bool { return value%2 == 0 }
s.DeleteFunc(even) // 's' will contain [1, 3, 5] after removing even elements.
```

### DeleteLength
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
newLength := s.DeleteLength(2) // s is now [1, 2, 4, 5], newLength is 4
```

### DeleteOK
```
s := &slice.Slice[int]{1, 2, 3, 4, 5}
deleted := s.DeleteOK(2) // s is now [1, 2, 4, 5], deleted is true
```

### Each
```Go
s := &slice.Slice[string]{"apple", "banana", "cherry"}
s.Each(func(i int, value string) {
	fmt.Printf("Element %d: %s\n", i, value)
})
```

### EachBreak
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.EachBreak(func(i int, value int) bool {
	fmt.Printf("Element %d: %d\n", i, value)
	return i < 3 // Stop iteration when i is less than 3
})
```

### EachReverse
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.EachReverse(func(i int, value int) {
	fmt.Printf("Element %d: %d\n", i, value)
})
```

### EachReverseBreak
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.EachReverseBreak(func(i int, value int) bool {
	fmt.Printf("Element %d: %d\n", i, value)
	return i > 2 // Stop iteration when i is greater than 2
})
```

### Fetch
```Go
s := &slice.Slice[string]{"apple", "banana", "cherry"}
fruit := s.Fetch(1) // fruit will be "banana"
```

### FetchLength
```Go
s := &slice.Slice[int]{10, 20, 30, 40, 50}
value, length := s.FetchLength(2)
// value will be 30
// length will be 5
```

### Filter
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
filtered := s.Filter(func(x int) bool {
	return x%2 == 0  Keep only even numbers
}) // filtered will be &Slice[int]{2, 4}
```

### FindIndex
```Go
s := &slice.Slice[string]{"apple", "banana", "cherry"}
index, found := s.FindIndex(func(fruit string) bool {
	return fruit == "banana"
})
// index will be 1
// found will be true
```

### Get
```Go
s := &slice.Slice[float64]{3.14, 2.71, 1.61}
value, ok := s.Get(1)
// value will be 2.71
// ok will be true
```

### GetLength
```Go
s := &slice.Slice[int]{10, 20, 30, 40, 50}
value, length, ok := s.GetLength(2)
// value will be 30
// length will be 5
// ok will be true
```

### IsEmpty
```Go
s := &slice.Slice[int]{}
isEmpty := s.IsEmpty() // isEmpty will be true
```

### IsPopulated
```Go
s := &slice.Slice[int]{10, 20, 30}
isPopulated := s.IsPopulated() // isPopulated will be true
```

### Length
```Go
s := &slice.Slice[int]{10, 20, 30, 40, 50}
length := s.Length() // length will be 5
```

### Make
```Go
s := &slice.Slice[int]{10, 20, 30}
s.Make(3) // s will be an empty Slice of length 3
```

### MakeEach
```Go
s := &slice.Slice[int]{}
s.MakeEach(10, 20, 30) // s will be a Slice containing {10, 20, 30}
```

### MakeEachReverse
```Go
s := &slice.Slice[int]{}
s.MakeEachReverse(10, 20, 30) //  s will be a Slice containing {30, 20, 10}
```

### Map
```Go
s := &slice.Slice[int]{10, 20, 30}
s.Map(func(i int, value int) int {
	return value * 2
}) // s will be a Slice containing {20, 40, 60}
```

### MapReverse
```Go
s := &slice.Slice[int]{10, 20, 30}
s.MapReverse(func(i int, value int) int {
	return value * 2
}) // s will be a Slice containing {60, 40, 20}
```

### Poll
```Go
s := &slice.Slice[int]{10, 20, 30}
value := s.Poll() // value will be 10, and s will be [20, 30].
```

### PopLength
```Go
s := &slice.Slice[int]{10, 20, 30}
value, length := s.PopLength() // value will be 30, length will be 2, and s will be [10, 20]
```

### PollOK
```Go
s := slice.New[int](1, 2, 3)
value, ok := s.PollOK() // 'value' will be 1, and 'ok' will be true as the slice is not empty.
```

### PopLength
```Go
s := &slice.Slice[int]{10, 20, 30}
value, length := s.PopLength() // value will be 30, length will be 2, and s will be [10, 20]
```

### PopOK
```Go
s := &slice.Slice[int]{10, 20, 30}
value, ok := s.PopOK() // value will be 30, ok will be true, and s will be [10, 20]
```

### Precatenate
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5}
s1.Precatenate(s2) // s1 will be [4, 5, 1, 2, 3]
```

### PrecatenateFunc
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5, 6}
result := s1.PrecatenateFunc(s2, func(i int, value int) bool {
	return value%2 == 0
})  // s1 will be modified to [6, 4, 2, 1, 3], and 'result' will be a pointer to 's1'.
```

### PrecatenateLength
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5}
length := s1.PrecatenateLength(s2) // length will be 5, and s1 will be [4, 5, 1, 2, 3]
```

### Prepend
```Go
s := &slice.Slice[int]{2, 3}
s.Prepend(1) // s will be [1, 2, 3]
```

### PrependFunc
```Go
s := &slice.Slice[int]{1, 2, 3}
result := s.PrependFunc(func(i int, value int) bool {
	return value%2 == 0
}, 4, 5, 6) // 's' will be modified to [6, 4, 2, 1, 3], and 'result' will be a pointer to 's'.
```

### PrependLength
```Go
s := &slice.Slice[int]{2, 3}
length := s.PrependLength(1, 0) // length will be 4, and s will be [1, 0, 2, 3]
```

### Reduce
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
result := s.Reduce(func(i int, value int) bool {
	return value%2 == 0
}) // 'result' will be a new slice containing [2, 4].
```

### Replace
```Go
s := &slice.Slice[int]{1, 2, 3}
ok := s.Replace(1, 4) // ok will be true, and s will be [1, 4, 3]
```

### Reverse
```Go
s := &slice.Slice[int]{1, 2, 3}
s.Reverse() // s will be [3, 2, 1]
```

### Set
```Go
s := &slice.Slice[int]{1, 2, 2, 3, 3, 3}
s.Set() // s will be [1, 2, 3]
```

### Shuffle
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.Shuffle() // s will be a random permutation of [1, 2, 3, 4, 5]
```

### Slice
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.Slice(1, 3) // s will be [2, 3, 4]
```

### Swap
```Go
s := &slice.Slice[int]{1, 2, 3}
s.Swap(0, 2) // s will be [3, 2, 1]
```

## Examples
### Struct
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

### Chaining
```Go
s := (&slice.Slice[int64]{1, 2, 3}).Append(4, 5, 6).Filter(func(_ int, value int64) bool {
  return value%2 == 0
})
fmt.Println(s) // 2, 4, 6
```

## Docker
A [Dockerfile](./Dockerfile) is provided for individuals that prefer containerized development.

### Building
Building the Docker container:
```sh
docker build . -t slice
```

### Running
Developing and running Go within the Docker container:
```sh
docker run -it --rm --name slice slice
```

## Docker Compose
A [docker-compose](./docker-compose.yml) file has also been included for convenience:
### Running
Running the compose file.
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
