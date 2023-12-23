package core

import (
	"fmt"

	"mvdan.cc/sh/v3/interp"
)

func Apply(file string, inputs map[string]string) error {
	runbook, err := Load(file)
	if err != nil {
		return err
	}
	if err := runbook.Parameters.Valid(inputs); err != nil {
		return err
	}
	runner, _ := interp.New(interp.Env(Environment(inputs)))
	if applicable, err := runbook.Mitigation.Clause.Evaluate(runner); err != nil {
		return err
	} else if !applicable {
		fmt.Printf("RUNBOOK [%s] => [not applicable]\n", runbook.Title)
		return nil
	}
	fmt.Printf("RUNBOOK [%s] => [applicable]\n\n", runbook.Title)
	for _, step := range runbook.Mitigation.Steps {
		if err := step.Execute(runner); err != nil {
			fmt.Printf("STEP [%s] => [error]\n\n", step.Name)
			return err
		}
		fmt.Printf("STEP [%s] => [success]\n\n", step.Name)
	}
	return nil
}
