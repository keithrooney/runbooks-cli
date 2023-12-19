package core

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExport(t *testing.T) {

	writer := new(bytes.Buffer)

	err := Export("testdata/runbook.yml", writer)

	require.NoError(t, err)

	bytes, err := os.ReadFile("testdata/runbook.md")

	require.NoError(t, err)

	expected := string(bytes)
	actual := writer.String()

	require.Equal(t, expected, actual)

}

func TestExportLoadFailure(t *testing.T) {

	directory := t.TempDir()
	file, err := os.CreateTemp(directory, "runbook-*.yml")

	require.NoError(t, err)

	_, err = file.WriteString("title: \"")
	require.NoError(t, err)

	writer := new(bytes.Buffer)
	err = Export(file.Name(), writer)

	require.Error(t, err)
	require.Equal(t, "file error: yaml: found unexpected end of stream\n", err.Error())

}
