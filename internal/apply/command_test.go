package apply

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {

	os.RemoveAll("/tmp/runbooks-cli")
	os.Mkdir("/tmp/runbooks-cli", os.ModePerm)

	err := os.WriteFile("/tmp/runbooks-cli/state.error", []byte("state has to be refreshed"), 0644)

	assert.NoError(t, err)

	Command.SetArgs([]string{"testdata/runbook.yml"})

	assert.NoError(t, Command.Execute())

	bytes, err := os.ReadFile("/tmp/runbooks-cli/state.ok")

	if assert.NoError(t, err) {
		assert.Equal(t, "state has been refreshed\n", string(bytes))
	}

	_, err = os.Stat("/tmp/runbooks-cli/state.error")

	assert.Error(t, err)

}
