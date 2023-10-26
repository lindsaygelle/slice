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
newSlice := &slice.Slice[int]{}
newSlice.Append(1, 2)
fmt.Println(newSlice) // &[1, 2]
```

### AppendFunc
Appends selected elements to the end of the slice based on a provided condition.
```Go
newSlice := &slice.Slice[int]{}
values := []int{1, 2, 3, 4, 5}
newSlice.AppendFunc(values, func(i int, value int) bool {
    return value%2 == 0
})
fmt.Println(newSlice) // &[2, 4]
```

### AppendLength
Appends values to the end of the slice and returns the length of the modified slice.
```Go
newSlice := &slice.Slice[int]{}
length := newSlice.AppendLength(1, 2, 3)
fmt.Println(length) // 3
```

### Bounds
Checks if an index is within the valid range of indices for the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
isWithinBounds := newSlice.Bounds(2)
fmt.Println(isWithinBounds) // true
```

### Clone
Creates a duplicate of the current slice with a reference to a new pointer.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
cloneSlice := newSlice.Clone()
fmt.Println(cloneSlice) // &[1, 2, 3]
```

### Concatenate
Merges elements from another slice to the tail of the receiver slice.
```Go
slice1 := &slice.Slice[int]{1, 2}
slice2 := &slice.Slice[int]{3, 4}
slice1.Concatenate(slice2)
fmt.Println(newSlice) // &[1, 2, 3, 4]
```

### ConcatenateFunc
Appends elements from another slice based on a filtering function.
```Go
slice1 := &slice.Slice[int]{1, 2, 3}
slice2 := &slice.Slice[int]{4, 5, 6}
slice1.ConcatenateFunc(slice2, func(i int, value int) bool {
    return value%2 == 0
})
fmt.Println(slice1) // &[2, 4, 6]
```

### ConcatenateLength
Merges elements from another slice to the tail of the receiver slice and returns the length of the modified slice.
```Go
slice1 := &slice.Slice[int]{1, 2}
slice2 := &slice.Slice[int]{3, 4}
length := slice1.ConcatenateLength(slice2)
fmt.Println(length) // 4
```

### Contains
Checks if a value exists in the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
contains := slice.Contains(2)
fmt.Println(contains) // true
```

### ContainsMany
Checks if multiple values exist in the slice and returns a boolean slice indicating their presence.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
contains := slice.ContainsMany(2, 6)
fmt.Println(contains) // &[true, false]
```

### Deduplicate
Removes values from the slice that have the same basic hash value.
```Go
newSlice := &slice.Slice[int]{1, 2, 2, 3, 4, 4, 5}
slice.Deduplicate()
fmt.Println(newSlice) // &[1, 2, 3, 4, 5]
```

### Delete
Safely removes an element from the slice by index.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice := slice.Delete(2)
fmt.Println(newSlice) // &[1, 2, 4, 5]
```

### DeleteFunc
Safely removes elements from the slice based on a provided predicate function.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice.DeleteFunc(func(i int, value int) bool {
    return value%2 == 0
})
fmt.Println(newSlice) // &[1, 3, 5]
```

### DeleteLength
Safely removes an element from the slice by index and returns the new length of the modified slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
length := slice.DeleteLength(2)
fmt.Println(length) // 4
```

### DeleteOK
Safely removes an element from the slice by index and returns a boolean indicating the success of the operation.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
success := slice.DeleteOK(2)
fmt.Println(success) // true
```

### DeleteUnsafe
Removes an element from the slice by index. Panics if index is out of bounds.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice.DeleteUnsafe(2)
fmt.Println(newSlice) // &[1, 2, 4, 5]
```

### Each
Executes a provided function for each element in the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice.Each(func(i int, value int) {
    fmt.Println(value)
})
// Output:
// 1
// 2
// 3
// 4
// 5
```

### EachBreak
Executes a provided function for each element in the slice with an optional break condition.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice.EachBreak(func(i int, value int) bool {
    fmt.Println(value)
    return value == 3
})
// Output:
// 1
// 2
// 3
```

### EachOK
Executes a provided function for each element in the slice and returns a bool breaking the loop when true.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
allPositive := slice.EachOK(func(i int, value int) bool {
    return value > 0
})
fmt.Println(allPositive) // true
```

### EachReverse
Executes a provided function for each element in reverse order.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice.EachReverse(func(i int, value int) {
    fmt.Println(value)
})
// Output:
// 5
// 4
// 3
// 2
// 1
```

### EachReverseBreak
Executes a provided function for each element in reverse order with an optional break condition.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice.EachReverseBreak(func(i int, value int) bool {
    fmt.Println(value)
    return value == 3
})
// Output:
// 5
// 4
// 3
```

### EachReverseOK
Executes a provided function for each element in the slice and returns a bool breaking the loop when true.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
allPositive := slice.EachReverseOK(func(i int, value int) bool {
    return value > 0
})
fmt.Println(allPositive) // true
```

### Equal
Checks if the two slices are equal.
```Go
slice1 := &slice.Slice[int]{1, 2, 3}
slice2 := &slice.Slice[int]{1, 2, 3}
isEqual := slice1.Equal(slice2)
fmt.Println(isEqual) // true
```

### EqualFunc
Checks whether the slices are equal based on the filtering function.
```Go
slice1 := &slice.Slice[int]{1, 2, 3}
slice2 := &slice.Slice[int]{2, 3, 4}
isEqual := slice1.EqualFunc(slice2, func(i int, a int, b int) bool {
    return a == b-1
})
fmt.Println(isEqual) // true
```

### EqualLength
Checks if the two slices have the same length.
```Go
slice1 := &slice.Slice[int]{1, 2, 3}
slice2 := &slice.Slice[int]{4, 5, 6}
isEqualLength := slice1.EqualLength(slice2)
fmt.Println(isEqualLength) // true
```

### Fetch
Retrieves the element at a specified index in the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
value := slice.Fetch(1)
fmt.Println(value) // 2
```

### FetchLength
Retrieves the element at a specified index in the slice and the length of the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
value, length := slice.FetchLength(2)
fmt.Println(value, length) // 3, 5
```

### Filter
Creates a new slice containing elements that satisfy a given predicate function.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
filteredSlice := slice.Filter(func(i int, value int) bool {
    return value%2 == 0
})
fmt.Println(filteredSlice) // &[2, 4]
```

### FindIndex
Finds the index of the first element that satisfies a given predicate function.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
index, found := slice.FindIndex(func(value int) bool {
    return value == 3
})
fmt.Println(index, found) // 2, true
```

### Get
Retrieves the element at a specified index in the slice and returns a boolean indicating success.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
value, found := slice.Get(1)
fmt.Println(value, found) // 2, true
```

### GetLength
Retrieves the element at a specified index in the slice, a boolean indicating success and the length of the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
value, found, length := slice.GetLength(2)
fmt.Println(value, found, length) // 3, true, 5
```

### IsEmpty
Checks if the slice is empty.
```Go
newSlice := &slice.Slice[int]{}
isEmpty := newSlice.IsEmpty()
fmt.Println(isEmpty) // true
```

### IsPopulated
Checks if the slice is not empty.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
isPopulated := newSlice.IsPopulated()
fmt.Println(isPopulated) // true
```

### Length
Returns the number of elements in the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
length := slice.Length()
fmt.Println(length) // 5
```

### Make
Empties the slice and sets it to a specified length.
```Go
newSlice := (&slice.Slice[int]{}).Make(3)
fmt.Println(newSlice) // &[0, 0, 0]
```

### MakeEach
Empties the slice, sets it to a specified length, and populates it with provided values.
```Go
newSlice := (&slice.Slice[int]{}).MakeEach(1, 2, 3)
fmt.Println(newSlice) // &[1, 2, 3]
```

### MakeEachReverse
Empties the slice, sets it to a specified length, and populates it with provided values in reverse order.
```Go
newSlice := (&slice.Slice[int]{}).MakeEachReverse(1, 2, 3)
fmt.Println(newSlice) // &[3, 2, 1]
```

### Map
Executes a provided function for each element and sets the returned value to a new slice at the current index.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
mappedSlice := slice.Map(func(i int, value int) int {
    return value * 2
})
fmt.Println(mappedSlice) // &[2, 4, 6]
```

### MapReverse
Executes a provided function for each element in reverse order and sets the returned value to a new slice at the current index.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
mappedSlice := slice.MapReverse(func(i int, value int) int {
    return value * 2
})
fmt.Println(mappedSlice) // &[6, 4, 2]
```

### Modify
Applies the provided function to each element in the slice and modifies the elements in place.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
newSlice.Modify(func(i int, value int) int {
    return value + 10
})
fmt.Println(slice) // &[11, 12, 13]
```

### ModifyReverse
Applies the provided function to each element in the slice in reverse order and modifies the elements in place.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
newSlice.ModifyReverse(func(i int, value int) int {
    return value + 10
})
fmt.Println(slice) // &[13, 12, 11]
```

### Poll
Removes and returns the first element from the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
value := slice.Poll()
fmt.Println(value) // 1
```

### PollLength
Removes the first element from the slice and returns the removed element and the length of the modified slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
value, length := slice.PollLength()
fmt.Println(value, length) // 1, 4
```

### PollOK
Removes and returns the first element from the slice and returns a boolean indicating success.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
value, ok := slice.PollOK()
fmt.Println(value, ok) // 1, true
```

### Pop
Removes and returns the last element from the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
value := slice.Pop()
fmt.Println(value) // 5
```

### PopLength
Removes the last element from the slice and returns the removed element and the length of the modified slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
value, length := slice.PopLength()
fmt.Println(value, length) // 5, 4
```

### PopOK
Removes and returns the last element from the slice and returns a boolean indicating success.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
value, ok := slice.PopOK()
fmt.Println(value, ok) // 3, true
```

### Precatenate
Merges elements from another slice to the head of the receiver slice.
```Go
slice1 := &slice.Slice[int]{1, 2}
slice2 := &slice.Slice[int]{3, 4}
slice1.Precatenate(slice2)
fmt.Println(slice1) // &[1, 2, 3, 4]
```

### PrecatenateFunc
Prepends elements from another slice based on a provided predicate function.
```Go
slice1 := &slice.Slice[int]{1, 2, 3}
slice2 := &slice.Slice[int]{3, 4, 5}
slice1.PrecatenateFunc(slice2, func(i int, value int) bool {
    return value%2 == 0
})
fmt.Println(slice1) // &[2, 4]
```

### PrecatenateLength
Merges elements from another slice to the head of the receiver slice and returns the length of the modified slice.
```Go
slice1 := &slice.Slice[int]{1, 2}
slice2 := &slice.Slice[int]{3, 4}
length := slice1.PrecatenateLength(slice2)
fmt.Println(length) // 4
```

### Prepend
Adds one element to the head of the slice.
```Go
newSlice := &slice.Slice[int]{3, 4, 5}
newSlice.Prepend(1, 2)
fmt.Println(newSlice) // &[1, 2, 3, 4, 5]
```

### PrependFunc
Prepends elements to the head of the slice based on a provided predicate function.
```Go
newSlice := &slice.Slice[int]{3, 4, 5}
values := []int{1, 2}
slice.PrependFunc(values, func(i int, value int) bool {
    return value%2 == 0
})
fmt.Println(newSlice) // &[2, 4, 3, 4, 5]
```

### PrependLength
Adds multiple elements to the head of the slice and returns the length of the modified slice.
```Go
newSlice := &slice.Slice[int]{3, 4, 5}
length := newSlice.PrependLength(1, 2)
fmt.Println(length) // 5
```

### Reduce
Reduce applies the provided function to each element in the slice and reduces the elements to a single value.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
result := newSlice.Reduce(func(i int, currentValue int, resultValue int) int {
    return resultValue + currentValue
})
fmt.Println(result) // 15
```

### ReduceReverse
ReduceReverse iterates over the slice in reverse order and reduces the elements into a single value using the provided function.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
result := newSlice.ReduceReverse(func(i int, currentValue int, resultValue int) int {
    return resultValue + currentValue
})
fmt.Println(result) // 15
```

### Replace
Replaces the element at the index with the provided value.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
success := slice.Replace(2, 10)
fmt.Println(slice) // &[1, 2, 10, 4, 5]
fmt.Println(success) // true
```

### Reverse
Reverses the order of the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3}
newSlice.Reverse()
fmt.Println(newSlice) // &[3, 2, 1]
```

### Shuffle
Randomly shuffles elements in the slice.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice.Shuffle()
fmt.Println(newSlice) // Randomly shuffled slice
```

### Slice
Creates a subset of the values based on the beginning and end index.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
subSlice := newSlice.Slice(1, 4)
fmt.Println(subSlice) // &[2, 3, 4]
```

### SortFunc
Sorts elements in the slice that satisfy a provided predicate function.
```Go
newSlice := &slice.Slice[int]{5, 2, 1, 4, 3}
newSlice.SortFunc(func(i int, j int, a int, b int) bool {
    return a < b
})
fmt.Println(newSlice) // &[1, 2, 3, 4, 5]
```

### Splice
Modifies the slice to include only the values based on the beginning and end index.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
splicedSlice := newSlice.Splice(1, 3)
fmt.Println(splicedSlice) // &[2, 3]
```

### Split
Divides the slice into two slices at the specified index and returns the two new slices.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
left, right := newSlice.Split(2)
fmt.Println(left, right) // &[1, 2], &[3, 4, 5]
```

### SplitFunc
Divides the slice into two slices based on the provided function and returns the two new slices.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
left, right := newSlice.SplitFunc(func(i int, value int) bool {
    return value > 2
})
fmt.Println(left, right) // &[1, 2], &[3, 4, 5]
```

### SplitOK
Divides the slice into two slices at the specified index and returns the two new slices, or false if the index is invalid.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
left, right, exists := newSlice.SplitOK(2)
fmt.Println(left, right, exists) // &[1, 2], &[3, 4, 5], true
```

### Swap
Swaps values at indexes i and j.
```Go
newSlice := &slice.Slice[int]{1, 2, 3, 4, 5}
newSlice.Swap(0, 4)
fmt.Println(newSlice) // &[5, 2, 3, 4, 1]
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
