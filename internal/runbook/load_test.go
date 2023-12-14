package runbook

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	directory := t.TempDir()

	file, err := os.CreateTemp(directory, "runbook-*.yml")

	assert.NoError(t, err)

	_, err = file.WriteString("title: \"This is our runbook\"\n")

	assert.NoError(t, err)

	expected := Runbook{Title: "This is our runbook"}
	actual, err := Load(file.Name())

	assert.NoError(t, err)
	assert.Equal(t, expected, *actual)

}

func TestLoadFileDoesNotExist(t *testing.T) {

	_, err := Load("foobar.txt")

	assert.Error(t, err)
	assert.Equal(t, "Error: file does not exist\n", err.Error())

}

func TestLoadInvalidFileFormat(t *testing.T) {

	directory := t.TempDir()
	file, err := os.CreateTemp(directory, "test.txt")
	_, err = file.WriteString(`,`)

	assert.NoError(t, err)

	_, err = Load(file.Name())

	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("Error: unexpected file format\n"), err.Error())

}
