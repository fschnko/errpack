// Copyright © 2017-2019 Artem Feshchenko. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package errpack

import (
	"bytes"
	"fmt"
)

// The Pack is a package of errors.
type Pack struct {
	name string
	errs []error
}

// New creates new errors pack.
func New(name string) *Pack {
	return &Pack{name: name}
}

// Add adds errors to the pack.
// Nil errors are not be added.
func (p *Pack) Add(errs ...error) {
	for _, err := range errs {
		if err != nil {
			p.errs = append(p.errs, err)
		}
	}
}

// Name returns the pack name.
func (p *Pack) Name() string {
	return p.name
}

// ErrorOrNil returns an error if it has any errors or nil.
func (p *Pack) ErrorOrNil() error {
	if p == nil {
		return nil
	}
	if !p.IsEmpty() {
		return p
	}
	return nil
}

// IsEmpty returns true if the pack does not contain errors.
func (p *Pack) IsEmpty() bool {
	return len(p.errs) == 0
}

func (p *Pack) Error() string {
	buf := &bytes.Buffer{}
	for _, err := range p.errs {
		buf.WriteString(fmt.Sprintf("%s :%s\n", p.name, err))
	}
	return buf.String()
}

// Errors returns a list of raw errors.
func (p *Pack) Errors() []error {
	return p.errs
}
