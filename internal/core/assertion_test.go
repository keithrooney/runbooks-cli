package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndefinedAssertionReturnsTrue(t *testing.T) {

	var assertion Assertion = ""

	executable, err := assertion.Evaluate()

	if assert.NoError(t, err) {
		assert.True(t, executable)
	}

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
		if assert.NoError(t, err) {
			assert.True(t, executable)
		}
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
		if assert.NoError(t, err) {
			assert.False(t, executable)
		}
	}

}

func TestAssertionSyntaxError(t *testing.T) {

	var assertion Assertion = "grep `"

	executable, err := assertion.Evaluate()

	if assert.Error(t, err) {
		assert.False(t, executable)
		assert.Equal(t, "syntax error: 1:6: reached EOF without closing quote `\n", err.Error())
	}

}

func TestAssertionTypeError(t *testing.T) {

	var assertion Assertion = "echo \"Truthy\""

	executable, err := assertion.Evaluate()

	if assert.Error(t, err) {
		assert.False(t, executable)
		assert.Equal(t, "type error: strconv.ParseBool: parsing \"Truthy\": invalid syntax\n", err.Error())
	}

}
