package runbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteUndefinedWhenReturnsTrue(t *testing.T) {

	clause := Clause{}

	executable, err := clause.Evaluate()

	assert.NoError(t, err)
	assert.True(t, executable)

}

func TestExecuteReturnsTrue(t *testing.T) {

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
			Shell: command,
		}

		executable, err := clause.Evaluate()

		assert.NoError(t, err)
		assert.True(t, executable)
	}

}

func TestExecuteReturnsFalse(t *testing.T) {

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
			Shell: command,
		}

		executable, err := clause.Evaluate()

		assert.NoError(t, err)
		assert.False(t, executable)
	}

}

func TestExecuteSyntaxError(t *testing.T) {

	clause := Clause{
		Shell: "grep",
	}

	executable, err := clause.Evaluate()

	assert.Error(t, err, "Error: syntax error")
	assert.False(t, executable)

}

func TestExecuteUnsupportedBooleanValue(t *testing.T) {

	clause := Clause{
		Shell: "echo \"Truthy\"",
	}

	executable, err := clause.Evaluate()

	assert.Error(t, err, "Error: unsupported boolean value")
	assert.False(t, executable)

}
