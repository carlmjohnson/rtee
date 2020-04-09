# Rtee [![GoDoc](https://godoc.org/github.com/carlmjohnson/rtee?status.svg)](https://godoc.org/github.com/carlmjohnson/rtee) [![Go Report Card](https://goreportcard.com/badge/github.com/carlmjohnson/rtee)](https://goreportcard.com/report/github.com/carlmjohnson/rtee)

Like tee but with automatic process substitution.

## Screenshots
```
$ rtee -h
rtee - Like tee but with automatic process substitution.

        Options:
  -dst value
        secondary output file or URL (default stdout)

$ ls | rtee pbcopy
LICENSE
README.md
app
go.mod
go.sum
main.go

$ pbpaste
LICENSE
README.md
app
go.mod
go.sum
main.go

$ ls | rtee -dst output.txt pbcopy

$ cat output.txt
LICENSE
README.md
app
go.mod
go.sum
main.go
```

## Installation

First install [Go](http://golang.org).

If you just want to install the binary to your current directory and don't care about the source code, run

```bash
GOBIN="$(pwd)" GOPATH="$(mktemp -d)" go get github.com/carlmjohnson/rtee
```

## TODO

- [ ] Option to run arguments through shell
