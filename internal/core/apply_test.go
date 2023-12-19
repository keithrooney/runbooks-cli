package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApply(t *testing.T) {

	os.RemoveAll("/tmp/runbooks-cli")
	os.Mkdir("/tmp/runbooks-cli", os.ModePerm)

	err := os.WriteFile("/tmp/runbooks-cli/state.error", []byte("state has to be refreshed"), 0644)

	require.NoError(t, err)
	require.NoError(t, Apply("testdata/apply.yml"))

	bytes, err := os.ReadFile("/tmp/runbooks-cli/state.ok")

	require.NoError(t, err)
	require.Equal(t, "state has been refreshed\n", string(bytes))

	_, err = os.Stat("/tmp/runbooks-cli/state.error")

	require.Error(t, err)

}
