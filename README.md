# slice
Golang slice structure. Accepts any interface as a value and offers common methods to access, modify and traverse the interface.

Get it:

```
go get github.com/gellel/slice
```

Import it:

```
import (
	"github.com/gellel/slice"
)
```

## Usage

Creating a basic slice pointer.

```go
package main

import (
	"fmt"

	"github.com/gellel/slice"
)

func main() {

    slice := slice.New("i", 1, map[string]string{})

    fmt.Println(slice.Len())

    fmt.Println(slice.Remove(1))
}
```

Creating a slice wrapper to accept specific data.

```go
package main

import (
    "github.com/gellel/slice"
)

type T struct{}

type Types struct {
    slice *slice.Slice
}

func (pointer *Types) Add(t T) {
    pointer.slice.Append(t)
}
```

Using a built-in string slice

```go
package main

import (
    "github.com/gellel/slice"
)

func main() {

    a := slice.NewString()

    b := slice.NewStringSlice("a","b","c")

    c := slice.NewInt()

    d := slice.NewIntSlice(4, 2, 0)
}
```

## License

[MIT](https://github.com/gellel/slice/blob/master/LICENSE)
