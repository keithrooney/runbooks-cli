package core

import (
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

	if assert.NoError(t, err) {
		assert.Equal(t, expected, *actual)
	}

}

func TestLoadFileErrorFileNotExist(t *testing.T) {

	_, err := Load("foobar.txt")

	if assert.Error(t, err) {
		assert.Equal(t, "file error: stat foobar.txt: no such file or directory\n", err.Error())
	}

}

func TestLoadFileErrorUnsupportedFormat(t *testing.T) {

	directory := t.TempDir()
	file, err := os.CreateTemp(directory, "test.txt")
	_, err = file.WriteString(`,`)

	assert.NoError(t, err)

	_, err = Load(file.Name())

	if assert.Error(t, err) {
		assert.Equal(t, "file error: yaml: did not find expected node content\n", err.Error())
	}

}
