package core

import (
	"io"
)

func Export(file string, writer io.Writer) error {
	runbook, err := Load(file)
	if err != nil {
		return err
	}
	return runbook.Export(writer)
}
