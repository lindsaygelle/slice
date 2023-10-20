package slice_test

import (
	"fmt"
	"net/http"

	"github.com/lindsaygelle/slice"
)

func ExampleNew() {
	// Creating a new integer slice
	intSlice := slice.New[int](1, 2, 3, 4, 5)

	// Print the created slice
	fmt.Println("Integer Slice:", intSlice)

	// Creating a new string slice
	stringSlice := slice.New[string]("apple", "orange", "banana", "grape")

	// Print the created slice
	fmt.Println("String Slice:", stringSlice)

	// Creating a new http.Request slice
	httpRequestSlice := slice.New[http.Request](http.Request{})

	// Print the created slice
	fmt.Println(httpRequestSlice)

	// Define a custom struct type
	type Animal struct{ Name string }

	// Creating a custom slice
	customSlice := slice.New[Animal](Animal{"Cat"}, Animal{"Dog"})

	// Print the created slice
	fmt.Println(customSlice)
}
