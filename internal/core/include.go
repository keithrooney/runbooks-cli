package core

import (
	"fmt"
	"path/filepath"
	"strings"

	"mvdan.cc/sh/v3/interp"
)

type Include string

func (include Include) Execute(runner *interp.Runner) error {
	return Apply(string(include))
}

func (include Include) Format(options FormatOptions) string {
	filename := string(include)
	directory, name := filepath.Split(filename)
	markdown := filepath.Join(directory, fmt.Sprintf("%s.md", strings.TrimSuffix(name, filepath.Ext(name))))
	return Format(markdown, options)
}
