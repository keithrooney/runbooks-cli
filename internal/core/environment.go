package core

import (
	"os"
	"strings"

	"mvdan.cc/sh/v3/expand"
)

func Environment(arguments map[string]string) expand.Environ {
	environ := os.Environ()
	for key, value := range arguments {
		environ = append(environ, strings.Join([]string{key, value}, "="))
	}
	return expand.ListEnviron(environ...)
}
