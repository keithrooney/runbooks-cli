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
		return core.Export(args[0], cmd.OutOrStdout())
	},
}
