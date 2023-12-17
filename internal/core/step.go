package core

import (
	"fmt"

	"mvdan.cc/sh/v3/interp"
)

type Step struct {
	Name    string
	Command Command
}

func (step Step) Execute(runner *interp.Runner) error {
	return step.Command.Execute(runner)
}

func (step *Step) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var fields map[string]string
	if err := unmarshal(&fields); err != nil {
		return err
	}
	step.Name = fields["name"]
	if cmd, ok := fields["shell"]; ok {
		step.Command = Shell(cmd)
	} else if cmd, ok := fields["assert"]; ok {
		step.Command = Assert(cmd)
	} else {
		return fmt.Errorf("unsupported command\n")
	}
	return nil
}
