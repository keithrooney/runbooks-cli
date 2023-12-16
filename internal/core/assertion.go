package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type Assertion string

func (assertion Assertion) Evaluate(runner *interp.Runner) (bool, error) {
	if assertion == "" {
		return true, nil
	}
	src, err := syntax.NewParser().Parse(strings.NewReader(string(assertion)), "")
	if err != nil {
		return false, fmt.Errorf("syntax error: %w\n", err)
	}
	var stdio bytes.Buffer
	var stderr bytes.Buffer
	if err := reconfigure(runner, nil, &stdio, &stderr); err != nil {
		return false, err
	}
	err = runner.Run(context.TODO(), src)
	if err := reconfigure(runner, nil, nil, nil); err != nil {
		return false, err
	}
	if err != nil {
		return false, fmt.Errorf("execution error: %w\n", err)
	}
	output := stdio.String()
	boolean, err := strconv.ParseBool(strings.TrimSpace(output))
	if err != nil {
		return false, fmt.Errorf("type error: %w\n", err)
	}
	return boolean, nil
}

func reconfigure(runner *interp.Runner, in io.Reader, out, err io.Writer) error {
	fn := interp.StdIO(in, out, err)
	if err := fn(runner); err != nil {
		return err
	}
	return nil
}
