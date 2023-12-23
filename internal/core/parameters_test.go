package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParametersValid(t *testing.T) {

	testCases := []struct {
		inputs     map[string]string
		parameters Parameters
	}{
		{
			inputs:     map[string]string{},
			parameters: Parameters([]Parameter{}),
		},
		{
			inputs: map[string]string{
				"directory": "/tmp",
			},
			parameters: Parameters([]Parameter{
				{
					Name: "directory",
				},
			}),
		},
	}

	for _, testCase := range testCases {
		require.Nil(t, testCase.parameters.Valid(testCase.inputs))
	}

}

func TestParametersValidReturnsErrors(t *testing.T) {

	testCases := []struct {
		inputs map[string]string
		err    error
	}{
		{
			err: fmt.Errorf("argument errors: insufficient number of arguments provided\n"),
		},
		{
			inputs: map[string]string{
				"foobar": "foobar",
			},
			err: fmt.Errorf("argument errors: directory was not provided\n"),
		},
	}

	parameters := Parameters([]Parameter{
		{
			Name: "directory",
		},
	})

	for _, testCase := range testCases {
		if err := parameters.Valid(testCase.inputs); assert.Error(t, err) {
			require.Equal(t, testCase.err, err)
		}
	}

}
