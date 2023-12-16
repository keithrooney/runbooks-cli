package core

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"

	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
)

type Assertion string

func (assertion Assertion) Evaluate() (bool, error) {
	if assertion == "" {
		return true, nil
	}
	src, err := syntax.NewParser().Parse(strings.NewReader(string(assertion)), "")
	if err != nil {
		return false, fmt.Errorf("syntax error: %w\n", err)
	}
	var stdio bytes.Buffer
	var stderr bytes.Buffer
	runner, _ := interp.New(
		// The method below returns nil by default, hence why the error is ignored.
		interp.StdIO(nil, &stdio, &stderr),
	)
	err = runner.Run(context.TODO(), src)
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
