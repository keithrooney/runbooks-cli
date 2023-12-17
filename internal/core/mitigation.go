package core

import (
	"fmt"
	"io"

	"mvdan.cc/sh/v3/interp"
)

type Mitigation struct {
	Steps  []Step `yaml:"steps"`
	Clause Clause `yaml:"when"`
}

func (mitigation Mitigation) Execute(writer io.Writer) error {
	runner, _ := interp.New()
	applicable, err := mitigation.Clause.Evaluate(runner)
	if err != nil {
		return err
	}
	if !applicable {
		return nil
	}
	for _, step := range mitigation.Steps {
		fmt.Fprintf(writer, "STEP [%s] => [%s]\n\n", step.Name, step.Command)
		if err := step.Execute(runner); err != nil {
			return err
		}
	}
	return nil
}
