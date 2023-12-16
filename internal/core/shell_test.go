package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellExecute(t *testing.T) {

	shell := Shell{
		Command: "echo 1;",
	}

	err := shell.Execute()

	assert.NoError(t, err)

}

func TestShellExecuteSyntaxError(t *testing.T) {

	shell := Shell{
		Command: "grep `",
	}

	err := shell.Execute()

	if assert.Error(t, err) {
		assert.Equal(t, "syntax error: 1:6: reached EOF without closing quote `\n", err.Error())
	}

}
