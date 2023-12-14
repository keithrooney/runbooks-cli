package internal

import (
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the runbook",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		_, err := Load(filepath)
		if err != nil {
			return err
		}
		cmd.Printf("Applying %s ...\n", filepath)
		return nil
	},
}
