package core

import (
	"fmt"
	"strings"
)

func Format(text string, options FormatOptions) string {
	indentation := options.Indentation
	if indentation < 0 {
		indentation = 0
	}
	offset := 1
	lines := strings.Split(strings.TrimSpace(text), "\n")
	formatted := make([]string, len(lines)+2)
	formatted[0] = fmt.Sprintf("%s%s", strings.Repeat(" ", indentation), "```text")
	for index, line := range lines {
		formatted[index+offset] = fmt.Sprintf("%s%s", strings.Repeat(" ", indentation), line)
	}
	formatted[len(lines)+1] = fmt.Sprintf("%s%s", strings.Repeat(" ", indentation), "```")
	return strings.Join(formatted, "\n")
}
