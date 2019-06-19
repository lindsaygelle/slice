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

type Strings struct {
    slice *slice.Slice
}

func (pointer *Strings) Add(s string) {
    pointer.slice.Append(s)
}
```

## License

[MIT](https://github.com/gellel/slice/blob/master/LICENSE)