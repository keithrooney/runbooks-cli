package core

import (
	"context"
	"fmt"
	"strings"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type Shell string

func (shell Shell) Execute(runner *interp.Runner) error {
	file, err := syntax.NewParser().Parse(strings.NewReader(string(shell)), "")
	if err != nil {
		return fmt.Errorf("syntax error: %w\n", err)
	}
	err = runner.Run(context.TODO(), file)
	if err != nil {
		return fmt.Errorf("execution error: %w\n", err)
	}
	return nil
}

func (shell Shell) Format(options FormatOptions) string {
	return Format(string(shell), options)
}
