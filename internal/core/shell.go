package core

import (
	"context"
	"errors"
	"strings"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type Shell struct {
	Command string `yaml:"command"`
}

func (shell Shell) Execute() error {
	file, err := syntax.NewParser().Parse(strings.NewReader(shell.Command), "")
	if err != nil {
		return errors.New("Error: syntax error")
	}
	runner, _ := interp.New() // No options are provided, so no errors are returned.
	return runner.Run(context.TODO(), file)
}
