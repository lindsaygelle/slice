# Slice
![Slice](https://repository-images.githubusercontent.com/192740394/a748b8c6-34ae-4aca-ad43-c18d5908b5e4)

Slice is a Go package that offers a versatile set of pre-built slices with extended functionality. It abstracts common list operations, such as appending, deleting, concatenating, mapping, and more, making it easier to work with slices in Go.

Gopher artwork was sourced from [egonelbre/gophers](https://github.com/egonelbre/gophers).

[![Go Reference](https://pkg.go.dev/badge/github.com/lindsaygelle/slice.svg)](https://pkg.go.dev/github.com/lindsaygelle/slice)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/slice)](/LICENSE)


## Installation
To use the slice package in your Go project, you can install it via `go get`:

```sh
go get github.com/lindsaygelle/slice
```

## Usage
Import the slice package into your Go code:

```Go
import (
	"github.com/lindsaygelle/slice"
)
```

Creating a New Slice:
```Go
s := &slice.Slice[int]{}
```

## Docker
You can also use Slice within a Docker container with the provided Dockerfile. Here are the steps to build and run the container:

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
Contributions to Slice are highly encouraged! Whether you have ideas, bug reports, or enhancements, please submit them as GitHub issues or create a pull request with your changes. For significant contributions, it's advisable to initiate a discussion by creating an issue to ensure alignment with the project's goals and direction. Refer to the [CONTRIBUTING](./CONTRIBUTING.md) file for more details.

## License
Slice is released under the MIT License. You are welcome to use, modify, and distribute the code within this repository in accordance with the terms of the license. For additional information, please review the [LICENSE](./LICENSE) file.
