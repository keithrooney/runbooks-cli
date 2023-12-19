package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
		options  FormatOptions
	}{
		{
			text:     "ls -aux",
			expected: "```text\nls -aux\n```",
		},
		{
			text:     "\n   ls -aux   \n",
			expected: "```text\nls -aux\n```",
		},
		{
			text:     "while True;\ndo\n\techo 'Hello, World';\ndone\n",
			expected: "```text\nwhile True;\ndo\n\techo 'Hello, World';\ndone\n```",
		},
		{
			text: "while True;\ndo\n\techo 'Hello, World';\ndone\n",
			options: FormatOptions{
				Indentation: 4,
			},
			expected: "    ```text\n    while True;\n    do\n    \techo 'Hello, World';\n    done\n    ```",
		},
		{
			text: "         echo 1;",
			options: FormatOptions{
				Indentation: -2,
			},
			expected: "```text\necho 1;\n```",
		},
	}
	for _, testCase := range testCases {
		shell := Shell(testCase.text)
		assert.Equal(t, testCase.expected, shell.Format(testCase.options))
	}
}
