package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClauseExecuteUndefinedWhenReturnsTrue(t *testing.T) {

	clause := Clause{}

	executable, err := clause.Evaluate()

	assert.NoError(t, err)
	assert.True(t, executable)

}

func TestClauseExecuteReturnsTrue(t *testing.T) {

	commands := []string{
		"echo 1",
		"echo \"t\"",
		"echo \"T\"",
		"echo \"true\"",
		"echo \"TRUE\"",
		"echo \"True\"",
	}

	for _, command := range commands {
		clause := Clause{
			Shell: Shell{
				Command: command,
			},
		}

		executable, err := clause.Evaluate()

		assert.NoError(t, err)
		assert.True(t, executable)
	}

}

func TestClauseExecuteReturnsFalse(t *testing.T) {

	commands := []string{
		"echo 0",
		"echo \"f\"",
		"echo \"F\"",
		"echo \"false\"",
		"echo \"FALSE\"",
		"echo \"False\"",
	}

	for _, command := range commands {
		clause := Clause{
			Shell: Shell{
				Command: command,
			},
		}

		executable, err := clause.Evaluate()

		assert.NoError(t, err)
		assert.False(t, executable)
	}

}

func TestClauseExecuteSyntaxError(t *testing.T) {

	clause := Clause{
		Shell: Shell{
			Command: "grep",
		},
	}

	executable, err := clause.Evaluate()

	assert.Error(t, err, "Error: syntax error")
	assert.False(t, executable)

}

func TestClauseExecuteUnsupportedBooleanValue(t *testing.T) {

	clause := Clause{
		Shell: Shell{
			Command: "echo \"Truthy\"",
		},
	}

	executable, err := clause.Evaluate()

	assert.Error(t, err, "Error: unsupported boolean value")
	assert.False(t, executable)

}
