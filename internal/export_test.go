package internal

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {

	stdout := new(bytes.Buffer)
	rootCmd.SetOut(stdout)

	rootCmd.SetArgs([]string{"export", "test.yml"})

	assert.NoError(t, rootCmd.Execute())
	assert.Equal(t, "Exporting test.yml ...\n", stdout.String())

}
