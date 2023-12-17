package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"mvdan.cc/sh/v3/interp"
)

func TestShellExecute(t *testing.T) {

	shell := Shell("echo 1;")

	runner, _ := interp.New()
	err := shell.Execute(runner)

	assert.NoError(t, err)

}

func TestShellExecuteSyntaxError(t *testing.T) {

	shell := Shell("grep `")

	runner, _ := interp.New()
	err := shell.Execute(runner)

	if assert.Error(t, err) {
		assert.Equal(t, "syntax error: 1:6: reached EOF without closing quote `\n", err.Error())
	}

}
