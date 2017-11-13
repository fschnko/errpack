package main

import (
	"errors"
	"log"

	"github.com/fschnko/errpack"
)

func main() {
	errp := errpack.New("test error pack")

	errp.Add(
		FailFunc(),
		NewError("new error"),
	)

	for _, err := range errp.Errors() {
		log.Printf("%s : %s\n", errp.Name(), err)
	}
}

func FailFunc() error {
	return errors.New("func fail")
}

func NewError(msg string) error {
	return errors.New(msg)
}
