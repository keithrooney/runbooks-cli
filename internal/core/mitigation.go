package core

import (
	"mvdan.cc/sh/v3/interp"
)

type Mitigation struct {
	Steps     []Step    `yaml:"steps"`
	Assertion Assertion `yaml:"when"`
}

func (mitigation Mitigation) Execute(printf func(format string, i ...interface{})) error {
	runner, _ := interp.New()
	applicable, err := mitigation.Assertion.Evaluate(runner)
	if err != nil {
		return err
	}
	if !applicable {
		return nil
	}
	for _, step := range mitigation.Steps {
		printf("STEP [%s] => [%s]\n\n", step.Name, step.Shell)
		if err := step.Execute(runner); err != nil {
			return err
		}
	}
	return nil
}
