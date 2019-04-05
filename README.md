## Requirements

Make sure the following software is installed:

- `Golang` (examples were tested with `go`, version `1.12.1`)
- `make` - help out to build and run all project code from a single place

Make sure you set `GOPATH` and `GOBIN` environment variables.

The first one is usally set to `~/go` and its folder structure should follow official Golang guidelines.

The `GOBIN` should point at the destination folder of the installed artifacts, it usually points at `~/go/bin` folder.

## Getting Started

1. To see a list of available `targets`, type:

        make

    or alternatively

        make help

2. To build a selected example:

        make hello

3. To run it later on:

        make run_hello

4. To build all examples:

        make all

## Tips

1. To learn more about `GOPATH`, type: `go help gopath`

## Documentation

- Official Golang documentation, read it up [here](https://golang.org/doc/)
- Example `Makefile` structure for `Golang` [project](https://kodfabrik.com/journal/a-good-makefile-for-go/)
