package export

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/keithrooney/runbooks/cli/internal/core"
	"github.com/spf13/cobra"
)

//go:embed templates/markdown.tpl
var templ string

var Command = &cobra.Command{
	Use:   "export [flags] filename",
	Short: "Export the runbook.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		runbook, err := core.Load(filepath)
		if err != nil {
			return err
		}
		T := template.Must(template.New("export").Funcs(map[string]any{"split": strings.Split}).Parse(templ))
		if err := T.Execute(cmd.OutOrStdout(), runbook); err != nil {
			return fmt.Errorf("template error: %w\n", err)
		}
		return nil
	},
}
