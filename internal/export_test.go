package internal

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {

	directory := t.TempDir()

	file, err := os.CreateTemp(directory, "runbook-*.yml")
	assert.NoError(t, err)

	_, err = file.WriteString("title: \"This is our runbook\"")
	assert.NoError(t, err)

	stdout := new(bytes.Buffer)
	rootCmd.SetOut(stdout)

	rootCmd.SetArgs([]string{"export", file.Name()})

	assert.NoError(t, rootCmd.Execute())
	assert.Equal(t, fmt.Sprintf("Exporting %s ...\n", file.Name()), stdout.String())

}
