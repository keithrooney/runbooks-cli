package core

import (
	"mvdan.cc/sh/v3/interp"
)

type Command interface {
	Formatter
	Execute(runner *interp.Runner) error
}
