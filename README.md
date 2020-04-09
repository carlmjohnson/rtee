# Rtee [![GoDoc](https://godoc.org/github.com/carlmjohnson/rtee?status.svg)](https://godoc.org/github.com/carlmjohnson/rtee) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/rtee)](https://goreportcard.com/report/github.com/carlmjohnson/rtee)

Like tee but with automatic process substitution.

## Installation

First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN="$(pwd)" GOPATH="$(mktemp -d)" go get github.com/carlmjohnson/rtee
```

## TODO

- [ ] Option to run output through shell
