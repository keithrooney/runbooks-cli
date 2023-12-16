package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"mvdan.cc/sh/v3/interp"
)

func TestEmptyClauseReturnsTrue(t *testing.T) {

	var clause Clause = ""

	runner, _ := interp.New()
	executable, err := clause.Evaluate(runner)

	if assert.NoError(t, err) {
		assert.True(t, executable)
	}

}

func TestClauseReturnsTrue(t *testing.T) {

	clauses := []Clause{
		"echo 1",
		"echo \"t\"",
		"echo \"T\"",
		"echo \"true\"",
		"echo \"TRUE\"",
		"echo \"True\"",
	}

	for _, clause := range clauses {
		runner, _ := interp.New()
		executable, err := clause.Evaluate(runner)
		if assert.NoError(t, err) {
			assert.True(t, executable)
		}
	}

}

func TestClauseReturnsFalse(t *testing.T) {

	clauses := []Clause{
		"echo 0",
		"echo \"f\"",
		"echo \"F\"",
		"echo \"false\"",
		"echo \"FALSE\"",
		"echo \"False\"",
	}

	for _, clause := range clauses {
		runner, _ := interp.New()
		executable, err := clause.Evaluate(runner)
		if assert.NoError(t, err) {
			assert.False(t, executable)
		}
	}

}

func TestClauseSyntaxError(t *testing.T) {

	var clause Clause = "grep `"

	runner, _ := interp.New()
	executable, err := clause.Evaluate(runner)

	if assert.Error(t, err) {
		assert.False(t, executable)
		assert.Equal(t, "syntax error: 1:6: reached EOF without closing quote `\n", err.Error())
	}

}

func TestClauseTypeError(t *testing.T) {

	var clause Clause = "echo \"Truthy\""

	runner, _ := interp.New()
	executable, err := clause.Evaluate(runner)

	if assert.Error(t, err) {
		assert.False(t, executable)
		assert.Equal(t, "type error: strconv.ParseBool: parsing \"Truthy\": invalid syntax\n", err.Error())
	}

}
