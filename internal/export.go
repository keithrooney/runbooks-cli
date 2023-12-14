package internal

import (
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:           "export [flags] filename",
	Short:         "Export the runbook.",
	Args:          cobra.ExactArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		_, err := Load(filepath)
		if err != nil {
			return err
		}
		cmd.Printf("Exporting %s ...\n", filepath)
		return nil
	},
}
