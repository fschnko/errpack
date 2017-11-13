package errpack_test

import (
	"errors"
	"fmt"

	"github.com/fschnko/errpack"
)

// Add adds errors to the pack ignoring nil errors.
// It is very useful to combine functions calling with errors collecting.
func ExamplePack_Add() {
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
}
