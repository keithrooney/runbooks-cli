package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndefinedAssertionReturnsTrue(t *testing.T) {

	var assertion Assertion = ""

	executable, err := assertion.Evaluate()

	assert.NoError(t, err)
	assert.True(t, executable)

}

func TestAssertionReturnsTrue(t *testing.T) {

	assertions := []Assertion{
		"echo 1",
		"echo \"t\"",
		"echo \"T\"",
		"echo \"true\"",
		"echo \"TRUE\"",
		"echo \"True\"",
	}

	for _, assertion := range assertions {

		executable, err := assertion.Evaluate()

		assert.NoError(t, err)
		assert.True(t, executable)
	}

}

func TestAssertionReturnsFalse(t *testing.T) {

	assertions := []Assertion{
		"echo 0",
		"echo \"f\"",
		"echo \"F\"",
		"echo \"false\"",
		"echo \"FALSE\"",
		"echo \"False\"",
	}

	for _, assertion := range assertions {

		executable, err := assertion.Evaluate()

		assert.NoError(t, err)
		assert.False(t, executable)
	}

}

func TestAssertionSyntaxError(t *testing.T) {

	var assertion Assertion = "grep"

	executable, err := assertion.Evaluate()

	assert.Error(t, err, "Error: syntax error")
	assert.False(t, executable)

}

func TestAssertionUnsupportedBooleanValue(t *testing.T) {

	var assertion Assertion = "echo \"Truthy\""

	executable, err := assertion.Evaluate()

	assert.Error(t, err, "Error: unsupported boolean value")
	assert.False(t, executable)

}
