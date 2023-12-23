package core

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApply(t *testing.T) {

	directory := t.TempDir()

	err := os.WriteFile(path.Join(directory, "state.error"), []byte("application is in a bad state"), 0644)
	require.NoError(t, err)

	require.NoError(t, Apply("testdata/apply.yml", map[string]string{"directory": directory}))

	bytes, err := os.ReadFile(path.Join(directory, "state.ok"))

	require.NoError(t, err)
	require.Equal(t, "state has been refreshed\n", string(bytes))

	_, err = os.Stat(path.Join(directory, "state.error"))

	require.Error(t, err)

}

func TestApplyLoadReturnsError(t *testing.T) {
}
