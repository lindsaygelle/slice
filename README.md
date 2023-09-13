# Slice
![Slice](https://repository-images.githubusercontent.com/192740394/a748b8c6-34ae-4aca-ad43-c18d5908b5e4)

Slice is a Go package that offers a versatile set of pre-built slices with extended functionality. It abstracts common list operations, such as appending, deleting, concatenating, mapping, and more, making it easier to work with slices in Go.


[![Go Reference](https://pkg.go.dev/badge/github.com/lindsaygelle/slice.svg)](https://pkg.go.dev/github.com/lindsaygelle/slice)
[![GitHub](https://img.shields.io/github/license/lindsaygelle/slice)](/LICENSE)


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

Creating a New Slice:
```Go
s := &slice.Slice[int]{}
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
