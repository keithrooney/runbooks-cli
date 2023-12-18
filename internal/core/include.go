package core

import "mvdan.cc/sh/v3/interp"

type Include string

func (include Include) Execute(runner *interp.Runner) error {
	return Apply(string(include))
}
