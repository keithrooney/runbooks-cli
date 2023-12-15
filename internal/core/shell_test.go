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
		Command: "grep",
	}

	err := shell.Execute()

	assert.Error(t, err, "Error: syntax error")

}
