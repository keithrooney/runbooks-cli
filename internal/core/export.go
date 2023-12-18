package core

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"text/template"
)

//go:embed templates/markdown.tpl
var exportTemplate string

func Export(file string, writer io.Writer) error {
	runbook, err := Load(file)
	if err != nil {
		return err
	}
	T := template.Must(template.New("export").Funcs(map[string]any{"split": strings.Split}).Parse(exportTemplate))
	if err := T.Execute(writer, runbook); err != nil {
		return fmt.Errorf("template error: %w\n", err)
	}
	return nil
}
