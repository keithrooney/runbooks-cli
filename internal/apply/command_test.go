package apply

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {

	directory := t.TempDir()

	file, err := os.CreateTemp(directory, "runbook-*.yml")
	assert.NoError(t, err)

	_, err = file.WriteString("title: \"This is our runbook\"")
	assert.NoError(t, err)

	stdout := new(bytes.Buffer)
	Command.SetOut(stdout)

	Command.SetArgs([]string{file.Name()})

	assert.NoError(t, Command.Execute())
	assert.Equal(t, fmt.Sprintf("Applying %s ...\n", file.Name()), stdout.String())

}