package export

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {

	directory := t.TempDir()

	file, err := os.CreateTemp(directory, "runbook-*.yml")
	assert.NoError(t, err)

	_, err = file.WriteString("title: \"The application is crashing.\"\ndetails: \"This normally occurs when there is a resource leak of some description.\"")
	assert.NoError(t, err)

	stdout := new(bytes.Buffer)
	Command.SetOut(stdout)

	Command.SetArgs([]string{file.Name()})

	assert.NoError(t, Command.Execute())

	expected := "# The application is crashing.\n\n## Details\nThis normally occurs when there is a resource leak of some description.\n\n"
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

	assert.Error(t, err)
	assert.Equal(t, "Error: unexpected file format\n", err.Error())

}

func TestCommandTemplateFailure(t *testing.T) {

	directory := t.TempDir()

	file, err := os.CreateTemp(directory, "runbook-*.yml")
	assert.NoError(t, err)

	_, err = file.WriteString("---")
	assert.NoError(t, err)

	Command.SetArgs([]string{file.Name()})

	err = Command.Execute()

	assert.Error(t, err)
	assert.Equal(t, "Error: unexpected template error\n", err.Error())

}
