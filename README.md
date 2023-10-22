# Slice
Slice is a [Go](https://github.com/golang/go) package that provides a generic slice with extended functionality. It abstracts common list operations, such as appending, deleting, concatenating, mapping, and more, making it easier to work with slices in Go.

![Slice](https://repository-images.githubusercontent.com/192740394/a748b8c6-34ae-4aca-ad43-c18d5908b5e4)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/lindsaygelle/slice)](https://pkg.go.dev/github.com/lindsaygelle/slice)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindsaygelle/slice)](https://goreportcard.com/report/github.com/lindsaygelle/slice)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/lindsaygelle/slice)](https://github.com/lindsaygelle/slice/releases)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/slice)](LICENSE.txt)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

## Features

### 🍕 Enhanced Functionality
Slice simplifies the complexities of slice manipulation in Go, streamlining your code by offering a rich set of operations.

### 🍰 Seamless Integration
Incorporating Slice into your Go projects is effortless. Simply import the package, specify the type, and harness its capabilities right away.

### 🧀 Type Generic
Slice embraces generics, empowering you to store and slice data flexibly and efficiently.

### 🥩 Access Safety
Slice enhances code robustness by addressing common programming pitfalls, including null pointer dereferences and out-of-bounds array accesses. Your code stays safe and reliable.

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
Appends values to the end of the slice and returns a pointer to the modified slice.
```Go
package main

import (
	"fmt"

	"github.com/lindsaygelle/slice"
)

func main() {
	newSlice := slice.New[int](1)
	newSlice.Append(1)
	fmt.Println(newSlice) // &[1]
}
```

### AppendFunc
Appends selected elements to the end of the slice based on a provided condition.
```Go
package main

import (
	"fmt"

	"github.com/lindsaygelle/slice"
)

func main() {
	newSlice := slice.New[int](1)
	newSlice.AppendFunc([]int{1, 2, 3, 4}, func(i int, value int) bool {
		return value%2 == 0
	})
	fmt.Println(newSlice) // &[2, 4]
}
```

### AppendLength
Appends values to the end of the slice and returns the length of the modified slice.
```Go
package main

import (
	"fmt"

	"github.com/lindsaygelle/slice"
)

func main() {
	newSlice := slice.New[int](0)
	length := newSlice.AppendLength(1, 2, 3, 4)
	fmt.Println(length) // 5
}
```

### Bounds
Checks if an index is within the valid range of indices for the slice.
```Go
package main

import (
	"fmt"

	"github.com/lindsaygelle/slice"
)

func main() {
	newSlice := slice.New[int]()
	inBounds := newSlice.Bounds(0)
	fmt.Println(inBounds) // false
	newSlice.Append(0)
	outOfBounds := newSlice.Bounds(0)
	fmt.Println(outOfBounds) // true
}
```

### Clone
Creates a duplicate of the current slice with a reference to a new pointer.
```Go
s1 := &slice.Slice[int]{1, 2}
s2 := s1.Clone()
s2.Replace(0, 3) // s1 is still [1, 2], s2 is [3, 2].
```

### Concatenate
Merges elements from another slice to the tail of the receiver slice.
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5}
s1.Concatenate(s2) // s1 is now [1, 2, 3, 4, 5].
```

### ConcatenateFunc
Appends elements from another slice based on a filtering function.
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5, 6}
s1.ConcatenateFunc(s2, func(i int, value int) bool {
	return value%2 == 0
}) // s1 is now [1, 2, 3, 4, 6].
```

### ConcatenateLength
Merges elements from another slice to the tail of the receiver slice and returns the length of the modified slice.
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5}
length := s1.ConcatenateLength(s2) // s1 is now [1, 2, 3, 4, 5], length is 5.
```

### Contains
Checks if a value exists in the slice.
```Go
s := &slice.Slice[string]{"apple", "banana", "cherry"}
containsBanana := s.Contains("banana") // containsBanana is true.
```

### ContainsMany
Checks if multiple values exist in the slice and returns a boolean slice indicating their presence.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
result := s.ContainsMany(2, 4, 6) // result will be [true, true, false].
```

### Delete
Safely removes an element from the slice by index.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.Delete(2) // s is now [1, 2, 4, 5].
```

### DeleteFunc
Safely removes elements from the slice based on a provided predicate function.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
even := func(i int, value int) bool { return value%2 == 0 }
s.DeleteFunc(even) // 's' will contain [1, 3, 5] after removing even elements.
```

### DeleteLength
Safely removes an element from the slice by index and returns the new length of the modified slice.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
newLength := s.DeleteLength(2) // s is now [1, 2, 4, 5], newLength is 4.
```

### DeleteOK
Safely removes an element from the slice by index and returns a boolean indicating the success of the operation.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
deleted := s.DeleteOK(2) // s is now [1, 2, 4, 5], deleted is true.
```

### DeleteUnsafe
Removes an element from the slice by index. Panics if index is out of bounds.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.DeleteUnsafe(2)
// The slice becomes [1, 2, 4, 5] with the element at index 2 (value 3) removed.
```

### Each
Executes a provided function for each element in the slice.
```Go
s := &slice.Slice[string]{"apple", "banana", "cherry"}
s.Each(func(i int, value string) {
	fmt.Printf("Element %d: %s\n", i, value) // i will be "1", value will be "apple", and so on.
})
```

### EachBreak
Executes a provided function for each element in the slice with an optional break condition.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.EachBreak(func(i int, value int) bool {
	fmt.Printf("Element %d: %d\n", i, value)
	return i < 3 // Stop iteration when i is less than 3.
})
```

### EachOK
Executes a provided function for each element in the slice and returns a bool breaking the loop when true.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
result := slice.EachOK(func(i int, value int) bool {
	fmt.Println(i, value)
	return value == 4 // Stop iteration at index 3.
}) // result is true.
```

### EachReverse
Executes a provided function for each element in reverse order.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.EachReverse(func(i int, value int) {
	fmt.Printf("Element %d: %d\n", i, value) // first iteration i will be 4, value will be 5.
})
```

### EachReverseBreak
Executes a provided function for each element in reverse order with an optional break condition.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.EachReverseBreak(func(i int, value int) bool {
	fmt.Printf("Element %d: %d\n", i, value)
	return i > 2 // Stop iteration when i is greater than 2.
})
```

### EachReverseOK
Executes a provided function for each element in the slice and returns a bool breaking the loop when true.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
result := slice.EachReverseOK(func(i int, value int) bool {
	fmt.Println(i, value)
	return value == 4 // Stop iteration at index 3 (runs two loops).
}) // result is true.
```

### Equal
Checks if the two slices are equal.
```Go
s1 := slice.Slice[int]{1, 2, 3}
s2 := slice.Slice[int]{1, 2, 3}
equal := s1.Equal(s2) // true.
```

### EqualFunc
Checks whether the slices are equal based on the filtering function.
```Go
s1 := slice.Slice[int]{1, 2, 3}
s2 := slice.Slice[int]{2, 4, 6}
customEqual := func(i int, value1, value2 int) bool {
    return value1*2 == value2
}
equal := s1.EqualFunc(s2, customEqual) // true.
```

### EqualLength
Checks if the two slices have the same length.
```Go
s1 := slice.Slice[int]{1, 2, 3}
s2 := slice.Slice[int]{2, 4, 6}
equal := s1.Length(s2) // true.
```

### Fetch
Retrieves the element at a specified index in the slice.
```Go
s := &slice.Slice[string]{"apple", "banana", "cherry"}
fruit := s.Fetch(1) // fruit will be "banana"
```

### FetchLength
Retrieves the element at a specified index in the slice and the length of the slice.
```Go
s := &slice.Slice[int]{10, 20, 30, 40, 50}
value, length := s.FetchLength(2)
// value will be 30.
// length will be 5.
```

### Filter
Creates a new slice containing elements that satisfy a given predicate function.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
filtered := s.Filter(func(x int) bool {
	return x%2 == 0  Keep only even numbers
}) // filtered will be &Slice[int]{2, 4}.
```

### FindIndex
Finds the index of the first element that satisfies a given predicate function.
```Go
s := &slice.Slice[string]{"apple", "banana", "cherry"}
index, found := s.FindIndex(func(fruit string) bool {
	return fruit == "banana"
})
// index will be 1.
// found will be true.
```

### Get
Retrieves the element at a specified index in the slice and returns a boolean indicating success.
```Go
s := &slice.Slice[float64]{3.14, 2.71, 1.61}
value, ok := s.Get(1)
// value will be 2.71.
// ok will be true.
```

### GetLength
Retrieves the element at a specified index in the slice, a boolean indicating success and the length of the slice.
```Go
s := &slice.Slice[int]{10, 20, 30, 40, 50}
value, ok, length := s.GetLength(2)
// value will be 30.
// ok will be true.
// length will be 5.
```

### IsEmpty
Checks if the slice is empty.
```Go
s := &slice.Slice[int]{}
isEmpty := s.IsEmpty() // isEmpty will be true.
```

### IsPopulated
Checks if the slice is not empty.
```Go
s := &slice.Slice[int]{10, 20, 30}
isPopulated := s.IsPopulated() // isPopulated will be true.
```

### Length
Returns the number of elements in the slice.
```Go
s := &slice.Slice[int]{10, 20, 30, 40, 50}
length := s.Length() // length will be 5.
```

### Make
Empties the slice and sets it to a specified length.
```Go
s := &slice.Slice[int]{10, 20, 30}
s.Make(3) // s will be an empty Slice of length 3.
```

### MakeEach
Empties the slice, sets it to a specified length, and populates it with provided values.
```Go
s := &slice.Slice[int]{}
s.MakeEach(10, 20, 30) // s will be a Slice containing {10, 20, 30}.
```

### MakeEachReverse
Empties the slice, sets it to a specified length, and populates it with provided values in reverse order.
```Go
s := &slice.Slice[int]{}
s.MakeEachReverse(10, 20, 30) //  s will be a Slice containing {30, 20, 10}.
```

### Map
Executes a provided function for each element and sets the returned value to a new slice at the current index.
```Go
s := &slice.Slice[int]{10, 20, 30}
s.Map(func(i int, value int) int {
	return value * 2
}) // s will be a Slice containing {20, 40, 60}.
```

### MapReverse
Executes a provided function for each element in reverse order and sets the returned value to a new slice at the current index.
```Go
s := &slice.Slice[int]{10, 20, 30}
s.MapReverse(func(i int, value int) int {
	return value * 2
}) // s will be a Slice containing {60, 40, 20}.
```

### Modify
Modify applies the provided function to each element in the slice and modifies the elements in place.
```Go
s := slice.Slice[int]{1, 2, 3, 4, 5}
modifiedSlice := numbers.Modify(func(i int, value int) int {
    return value * 2
})
// modifiedSlice: [2, 4, 6, 8, 10].
```

###ModifyReverseModifyReverse
Modify applies the provided function to each element in the slice in reverse order and modifies the elements in place.
```Go
s := slice.Slice[int]{1, 2, 3, 4, 5}
modifiedSlice := numbers.Modify(func(i int, value int) int {
    return value * 2
})
// modifiedSlice: [2, 4, 6, 8, 10].
```

### Poll
Removes and returns the first element from the slice.
```Go
s := &slice.Slice[int]{10, 20, 30}
value := s.Poll() // value will be 10, and s will be [20, 30].
```

### PollLength
Removes the first element from the slice and returns the removed element and the length of the modified slice.
```Go
s := &slice.Slice[int]{10, 20, 30}
value, length := s.Poll() // value will be 10, length will be 2, and s will be [20, 30].
```

### PollOK
Removes and returns the first element from the slice and returns a boolean indicating success.
```Go
s := slice.Slice[int]{1, 2, 3}
value, ok := s.PollOK() // 'value' will be 1, and 'ok' will be true as the slice is not empty.
```

### Pop
Removes and returns the last element from the slice.
```Go
s := &slice.Slice[int]{1, 2, 3}
value := s.Pop()
// 'value' will be 3, and the slice will become [1, 2].
```

### PopLength
Removes the last element from the slice and returns the removed element and the length of the modified slice.
```Go
s := &slice.Slice[int]{10, 20, 30}
value, length := s.PopLength() // value will be 30, length will be 2, and s will be [10, 20].
```

### PopOK
Removes and returns the last element from the slice and returns a boolean indicating success.
```Go
s := &slice.Slice[int]{10, 20, 30}
value, ok := s.PopOK() // value will be 30, ok will be true, and s will be [10, 20].
```

### Precatenate
Merges elements from another slice to the head of the receiver slice.
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5}
s1.Precatenate(s2) // s1 will be [4, 5, 1, 2, 3].
```

### PrecatenateFunc
Prepends elements from another slice based on a provided predicate function.
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5, 6}
result := s1.PrecatenateFunc(s2, func(i int, value int) bool {
	return value%2 == 0
})  // s1 will be modified to [6, 4, 2, 1, 3], and 'result' will be a pointer to 's1'.
```

### PrecatenateLength
Merges elements from another slice to the head of the receiver slice and returns the length of the modified slice.
```Go
s1 := &slice.Slice[int]{1, 2, 3}
s2 := &slice.Slice[int]{4, 5}
length := s1.PrecatenateLength(s2) // length will be 5, and s1 will be [4, 5, 1, 2, 3].
```

### Prepend
Adds one element to the head of the slice.
```Go
s := &slice.Slice[int]{2, 3}
s.Prepend(1) // s will be [1, 2, 3].
```

### PrependFunc
Prepends elements to the head of the slice based on a provided predicate function.
```Go
s := &slice.Slice[int]{1, 2, 3}
result := s.PrependFunc([]int{4, 5, 6}, func(i int, value int) bool {
	return value%2 == 0
}) // 's' will be modified to [6, 4, 2, 1, 3], and 'result' will be a pointer to 's'.
```

### PrependLength
Adds multiple elements to the head of the slice and returns the length of the modified slice.
```Go
s := &slice.Slice[int]{2, 3}
length := s.PrependLength(1, 0) // length will be 4, and s will be [1, 0, 2, 3].
```

### Reduce
Reduce applies the provided function to each element in the slice and reduces the elements to a single value.
```Go
s := slice.Slice[int]({1, 2, 3, 4, 5})
sum := numbers.Reduce(func(i int, currentValue int, resultValue int) int {
    return resultValue + currentValue
})
// sum: 15 (1 + 2 + 3 + 4 + 5).
```

### ReduceReverse
ReduceReverse iterates over the slice in reverse order and reduces the elements into a single value using the provided function.
```Go
s := slice.Slice[int]{1, 2, 3, 4, 5}
sum := numbers.ReduceReverse(func(i int, currentValue, resultValue int) int {
    return currentValue + resultValue
})
// sum: 15 (sum of all elements in the slice).
```

### Replace
Replaces the element at the index with the provided value.
```Go
s := &slice.Slice[int]{1, 2, 3}
ok := s.Replace(1, 4) // ok will be true, and s will be [1, 4, 3].
```

### Reverse
Reverses the order of the slice.
```Go
s := &slice.Slice[int]{1, 2, 3}
s.Reverse() // s will be [3, 2, 1].
```

### Set
Creates a new slice containing opnly unique values.
```Go
s := &slice.Slice[int]{1, 2, 2, 3, 3, 3}
s.Set() // s will be [1, 2, 3].
```

### Shuffle
Randomly shuffles elements in the slice.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
s.Shuffle() // s will be a random permutation of [1, 2, 3, 4, 5].
```

### Slice
Creates a subset of the values based on the beginning and end index.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice := s.Slice(1, 3) // newSlice will be [2, 3, 4].
```

### SortFunc
Sorts elements in the slice that satisfy a provided predicate function.
```Go
s := &slice.Slice[int]{2, 1, 5, 3}
s.SortFunc(func(i int, j int, a int, b int) bool {
	return a < b
}) // s will be [1, 2, 3, 5].
```

### Splice
Modifies the slice to include only the values based on the beginning and end index.
```Go
s := &slice.Slice[int]{1, 2, 3, 4, 5}
slice.Split(1, 3) // s will be [2, 3, 4].
```

### Split
Divides the slice into two slices at the specified index and returns the two new slices.
```Go
s := &slice.Slice[int]{1, 2, 3, 4}
a, b := s.Split(2) // a will be [1, 2], be will be [3, 4].
```

### SplitFunc
Divides the slice into two slices based on the provided function and returns the two new slices.
```Go
s := &slice.Slice[int]{1, 2, 3, 4}
a, b := s.SplitFunc(func(i int, value int) bool {
	return value%2 == 0
}) // a will be [1, 3], b will be [2, 4].
```

### SplitOK
Divides the slice into two slices at the specified index and returns the two new slices, or false if the index is invalid.
```Go
s := &slice.Slice[int]{1, 2, 3, 4}
a, b, ok := s.Split(2) // a will be [1, 2], be will be [3, 4], ok will be true.
```

### Swap
Swaps values at indexes i and j.
```Go
s := &slice.Slice[int]{1, 2, 3}
s.Swap(0, 2) // s will be [3, 2, 1].
```

## Examples
### Struct
```Go
// Define a custom struct.
type Person struct {
    Name  string
    Age   int
    Email string
}

// Create a slice of Person structs.
people := &slice.Slice[Person]{
    {Name: "Alice", Age: 30, Email: "alice@example.com"},
    {Name: "Bob", Age: 25, Email: "bob@example.com"},
}

// Append a new person to the slice.
newPerson := Person{Name: "Charlie", Age: 35, Email: "charlie@example.com"}
people.Append(newPerson)

// Find the index of a person with a specific email address.
index := people.FindIndex(func(p Person) bool {
    return p.Email == "bob@example.com"
})

// Slice the slice to include only people aged 30 or older.
people.Slice(1, people.Length())

// Reverse the order of people in the slice.
people.Reverse()

// Iterate over the slice and print each person's details.
people.Each(func(index int, person Person) {
    fmt.Printf("Index %d: Name: %s, Age: %d, Email: %s\n", index, person.Name, person.Age, person.Email)
})
```

### Chaining
```Go
s := (&slice.Slice[int64]{1, 2, 3}).Append(4, 5, 6).Filter(func(_ int, value int64) bool {
  return value%2 == 0
})
fmt.Println(s) // 2, 4, 6.
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

## Security
If you discover a security vulnerability within this project, please consult the [SECURITY](./SECURITY.md) document for information and next steps.

## Code Of Conduct
This project has adopted the [Amazon Open Source Code of Conduct](https://aws.github.io/code-of-conduct). For additional information, please review the [CODE_OF_CONDUCT](./CODE_OF_CONDUCT.md) file.

## Acknowledgements
Big thanks to [egonelbre/gophers](https://github.com/egonelbre/gophers) for providing the delightful Gopher artwork used in the social preview. Don't hesitate to pay them a visit!
