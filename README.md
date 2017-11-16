[![GoDoc](https://godoc.org/github.com/fschnko/errpack?status.svg)](https://godoc.org/github.com/fschnko/errpack)       [![Build Status](https://travis-ci.org/fschnko/errpack.svg?branch=master)](https://travis-ci.org/fschnko/errpack)       [![Coverage Status](https://coveralls.io/repos/github/fschnko/errpack/badge.svg?branch=master)](https://coveralls.io/github/fschnko/errpack?branch=master)      [![Go Report Card](https://goreportcard.com/badge/github.com/fschnko/errpack)](https://goreportcard.com/report/github.com/fschnko/errpack)
# errpack
Package errpack declares the Pack type for errors collecting.
This is useful for cyclic or unrelated operations when journaling is not allowed (library packages, etc.).

 ---
**Usage:**
```shell
go get github.com/fschnko/repeat
```

**Example:**
```golang
	errp := errpack.New("test error pack")
	errp.Add(
		func() error { return errors.New("func fail") }(),              // always fails.
		func() error { return nil }(),                                  // always succeed.
		func(msg string) error { return errors.New(msg) }("new error"), // returns new error with message.
	)

	for _, err := range errp.Errors() {
		fmt.Printf("%s : %s\n", errp.Name(), err)
	}
	// Output:
	// test error pack : func fail
	// test error pack : new error
```