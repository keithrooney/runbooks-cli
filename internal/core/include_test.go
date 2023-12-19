package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncludeFormat(t *testing.T) {
	testCases := []struct {
		filename string
		expected string
	}{
		{
			filename: "foobar.yml",
			expected: "```text\nfoobar.md\n```",
		},
		{
			filename: "path/to/file/foobar.yml",
			expected: "```text\npath/to/file/foobar.md\n```",
		},
	}
	for _, testCase := range testCases {
		include := Include(testCase.filename)
		assert.Equal(t, testCase.expected, include.Format(FormatOptions{}))
	}
}
