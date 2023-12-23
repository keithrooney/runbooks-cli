package core

import (
	"fmt"
	"path/filepath"
	"strings"

	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/interp"
)

type Include string

func (include Include) Execute(runner *interp.Runner) error {
	arguments := map[string]string{}
	runner.Env.Each(func(name string, variable expand.Variable) bool {
		arguments[name] = variable.String()
		return true
	})
	return Apply(string(include), arguments)
}

func (include Include) Format(options FormatOptions) string {
	directory, name := filepath.Split(string(include))
	markdown := filepath.Join(directory, fmt.Sprintf("%s.md", strings.TrimSuffix(name, filepath.Ext(name))))
	return Format(markdown, options)
}
