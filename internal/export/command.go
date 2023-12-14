package export

import (
	"errors"
	"text/template"

	"github.com/keithrooney/runbooks/cli/internal/runbook"
	"github.com/spf13/cobra"
)

const templ = `# {{ .Title }}

`

var Command = &cobra.Command{
	Use:           "export [flags] filename",
	Short:         "Export the runbook.",
	Args:          cobra.ExactArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		runbook, err := runbook.Load(filepath)
		if err != nil {
			return err
		}
		T := template.Must(template.New("export").Parse(templ))
		if err := T.Execute(cmd.OutOrStdout(), runbook); err != nil {
			return errors.New("Error: unexpected template error\n")
		}
		return nil
	},
}