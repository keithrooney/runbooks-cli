package export

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {

	stdout := new(bytes.Buffer)
	Command.SetOut(stdout)

	Command.SetArgs([]string{"testdata/runbook.yml"})

	assert.NoError(t, Command.Execute())

	bytes, err := os.ReadFile("testdata/runbook.md")

	assert.NoError(t, err)

	expected := string(bytes)
	actual := stdout.String()

	assert.Equal(t, expected, actual)

}

func TestCommandLoadFailure(t *testing.T) {

	directory := t.TempDir()

	file, err := os.CreateTemp(directory, "runbook-*.yml")
	assert.NoError(t, err)

	_, err = file.WriteString("title: \"")
	assert.NoError(t, err)

	Command.SetArgs([]string{file.Name()})

	err = Command.Execute()

	if assert.Error(t, err) {
		assert.Equal(t, "file error: yaml: found unexpected end of stream\n", err.Error())
	}

}

func TestCommandTemplateFailure(t *testing.T) {

	directory := t.TempDir()

	file, err := os.CreateTemp(directory, "runbook-*.yml")
	assert.NoError(t, err)

	_, err = file.WriteString("---")
	assert.NoError(t, err)

	Command.SetArgs([]string{file.Name()})

	err = Command.Execute()

	if assert.Error(t, err) {
		assert.Equal(t, "template error: template: export:1:5: executing \"export\" at <.Title>: nil pointer evaluating *core.Runbook.Title\n", err.Error())
	}

}
