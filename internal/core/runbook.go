package core

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"text/template"
)

type Runbook struct {
	Title      string     `yaml:"title"`
	Details    string     `yaml:"details"`
	Impact     string     `yaml:"impact"`
	Mitigation Mitigation `yaml:"mitigation"`
}

func (runbook *Runbook) Apply(writer io.Writer) error {
	fmt.Fprintf(writer, "RUNBOOK [%s]\n\n", runbook.Title)
	return runbook.Mitigation.Execute(writer)
}

//go:embed templates/markdown.tpl
var templ string

func (runbook *Runbook) Export(output io.Writer) error {
	T := template.Must(template.New("export").Funcs(map[string]any{"split": strings.Split}).Parse(templ))
	if err := T.Execute(output, runbook); err != nil {
		return fmt.Errorf("template error: %w\n", err)
	}
	return nil
}
