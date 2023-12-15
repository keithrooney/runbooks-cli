package core

import (
	"context"
	"errors"
	"strings"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type Shell struct {
	Command   string    `yaml:"command"`
	Assertion Assertion `yaml:"assert"`
}

func (shell Shell) Execute() error {
	file, err := syntax.NewParser().Parse(strings.NewReader(shell.Command), "")
	if err != nil {
		return errors.New("Error: syntax error")
	}
	runner, _ := interp.New() // No options are provided, so no errors are returned.
	err = runner.Run(context.TODO(), file)
	if err != nil {
		return errors.New("Error: ")
	}
	successful, err := shell.Assertion.Evaluate()
	if err != nil || !successful {
		return errors.New("Error: assert failed")
	}
	return nil
}
