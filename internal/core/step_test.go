package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestStepUnmarshallShellCommand(t *testing.T) {

	bytes, err := os.ReadFile("testdata/shell.yml")

	assert.NoError(t, err)

	step := Step{}
	err = yaml.Unmarshal(bytes, &step)

	if assert.NoError(t, err) {
		_, ok := step.Command.(Shell)
		assert.True(t, ok)
	}

}

func TestStepUnmarshallAssertCommand(t *testing.T) {

	bytes, err := os.ReadFile("testdata/assert.yml")

	assert.NoError(t, err)

	step := Step{}
	err = yaml.Unmarshal(bytes, &step)

	if assert.NoError(t, err) {
		_, ok := step.Command.(Assert)
		assert.True(t, ok)
	}

}

func TestStepUnmarshallIncludeCommand(t *testing.T) {

	bytes, err := os.ReadFile("testdata/include.yml")

	assert.NoError(t, err)

	step := Step{}
	err = yaml.Unmarshal(bytes, &step)

	if assert.NoError(t, err) {
		_, ok := step.Command.(Include)
		assert.True(t, ok)
	}

}

func TestStepUnmarshallUnsupportedCommand(t *testing.T) {

	bytes, err := os.ReadFile("testdata/unsupported.yml")

	assert.NoError(t, err)

	step := Step{}
	err = yaml.Unmarshal(bytes, &step)

	if assert.Error(t, err) {
		assert.Equal(t, "unsupported command\n", err.Error())
	}

}
