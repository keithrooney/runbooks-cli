package core

import (
	"mvdan.cc/sh/v3/interp"
)

type Step struct {
	Name  string `yaml:"name"`
	Shell Shell  `yaml:"shell"`
}

func (step Step) Execute(runner *interp.Runner) error {
	return step.Shell.Execute(runner)
}
