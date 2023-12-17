package core

import "mvdan.cc/sh/v3/interp"

type Command interface {
	Execute(runner *interp.Runner) error
}
