package apply

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {

	Command.SetArgs([]string{"testdata/runbook.yml"})

	assert.NoError(t, Command.Execute())

	bytes, err := os.ReadFile("/tmp/application/state")

	assert.NoError(t, err)

	assert.Equal(t, "state has been refreshed\n", string(bytes))

}
