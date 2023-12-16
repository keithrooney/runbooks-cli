package core

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type Shell struct {
	Command   string    `yaml:"command"`
	Assertion Assertion `yaml:"assert"`
}

func (shell Shell) Execute(runner *interp.Runner) error {
	file, err := syntax.NewParser().Parse(strings.NewReader(shell.Command), "")
	if err != nil {
		return fmt.Errorf("syntax error: %w\n", err)
	}
	err = runner.Run(context.TODO(), file)
	if err != nil {
		return fmt.Errorf("execution error: %w\n", err)
	}
	successful, err := shell.Assertion.Evaluate(runner)
	if err != nil {
		return fmt.Errorf("assertion error: %w\n", err)
	}
	if !successful {
		return errors.New("assertion failure\n")
	}
	return nil
}
