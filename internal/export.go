package internal

import (
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [flags] filename",
	Short: "Export the runbook.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		cmd.Printf("Exporting %s ...\n", file)
	},
}
