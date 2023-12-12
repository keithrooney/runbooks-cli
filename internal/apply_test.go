package internal

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {

	stdout := new(bytes.Buffer)
	rootCmd.SetOut(stdout)

	rootCmd.SetArgs([]string{"apply", "test.yml"})

	assert.NoError(t, rootCmd.Execute())
	assert.Equal(t, "Applying test.yml ...\n", stdout.String())

}
