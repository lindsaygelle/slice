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

## License

[MIT](https://github.com/gellel/slice/blob/master/LICENSE)