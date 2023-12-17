package core

import (
	"io"
)

func Apply(file string, writer io.Writer) error {
	runbook, err := Load(file)
	if err != nil {
		return err
	}
	return runbook.Apply(writer)
}
