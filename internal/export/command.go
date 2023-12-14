package export

import (
	"github.com/keithrooney/runbooks/cli/internal/runbook"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:           "export [flags] filename",
	Short:         "Export the runbook.",
	Args:          cobra.ExactArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		_, err := runbook.Load(filepath)
		if err != nil {
			return err
		}
		cmd.Printf("Exporting %s ...\n", filepath)
		return nil
	},
}
