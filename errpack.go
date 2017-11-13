package errpack

import (
	"bytes"
	"fmt"
)

type Pack struct {
	name string
	errs []error
}

func New(name string) *Pack {
	return &Pack{name: name}
}

func (p *Pack) Add(errs ...error) {
	for _, err := range errs {
		if err != nil {
			p.errs = append(p.errs, err)
		}
	}
}

func (p *Pack) Name() string {
	return p.name
}

func (p *Pack) ErrorOrNil() error {
	if len(p.errs) == 0 {
		return nil
	}
	return p
}

func (p *Pack) Error() string {
	buf := &bytes.Buffer{}
	for _, err := range p.errs {
		buf.WriteString(fmt.Sprintf("%s : %s\n", p.name, err))
	}
	return buf.String()
}

func (p *Pack) Errors() []error {
	return p.errs
}
