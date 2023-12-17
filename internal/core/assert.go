package core

import (
	"errors"
	"fmt"

	"mvdan.cc/sh/v3/interp"
)

type Assert string

func (assert Assert) Execute(runner *interp.Runner) error {
	clause := Clause(assert)
	successful, err := clause.Evaluate(runner)
	if err != nil {
		return fmt.Errorf("assertion error: %w\n", err)
	}
	if !successful {
		return errors.New("assertion failure\n")
	}
	return nil
}
