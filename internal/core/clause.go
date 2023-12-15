package core

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"strings"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type Clause struct {
	Shell Shell `yaml:"shell"`
}

func (clause Clause) Evaluate() (bool, error) {
	if clause.Shell.Command == "" {
		return true, nil
	}
	src, err := syntax.NewParser().Parse(strings.NewReader(clause.Shell.Command), "")
	if err != nil {
		return false, errors.New("Error: syntax error")
	}
	var stdio bytes.Buffer
	var stderr bytes.Buffer
	runner, _ := interp.New(
		// The returns no error, hence why the error is ignored.
		interp.StdIO(nil, &stdio, &stderr),
	)
	err = runner.Run(context.TODO(), src)
	if err != nil {
		return false, errors.New("Error: unsupported syntax")
	}
	output := stdio.String()
	boolean, err := strconv.ParseBool(strings.TrimSpace(output))
	if err != nil {
		return false, errors.New("Error: unsupported boolean value")
	}
	return boolean, nil
}
