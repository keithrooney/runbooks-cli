package core

import (
	"fmt"

	"mvdan.cc/sh/v3/interp"
)

func Apply(file string) error {
	runbook, err := Load(file)
	if err != nil {
		return err
	}
	runner, _ := interp.New()
	applicable, err := runbook.Mitigation.Clause.Evaluate(runner)
	if err != nil {
		return err
	}
	if !applicable {
		return nil
	}
	fmt.Printf("RUNBOOK [%s]\n\n", runbook.Title)
	for _, step := range runbook.Mitigation.Steps {
		fmt.Printf("STEP [%s]\n\n", step.Name)
		if err := step.Execute(runner); err != nil {
			return err
		}
	}
	return nil
}
