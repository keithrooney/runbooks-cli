package export

import (
	_ "embed"

	"github.com/keithrooney/runbooks/cli/internal/core"
	"github.com/spf13/cobra"
)

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
		return runbook.Export(cmd.OutOrStdout())
	},
}
