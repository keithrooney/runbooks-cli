package runbook

import (
	"context"
	"errors"
	"strings"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type Step struct {
	Name  string `yaml:"name"`
	Shell string `yaml:"shell"`
}

func (step Step) Execute() error {
	file, err := syntax.NewParser().Parse(strings.NewReader(step.Shell), "")
	if err != nil {
		return errors.New("Error: syntax error")
	}
	runner, _ := interp.New()
	return runner.Run(context.TODO(), file)
}
