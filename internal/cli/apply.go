package cli

import (
	"github.com/keithrooney/runbooks/cli/internal/core"
	"github.com/spf13/cobra"
)

var (
	arguments map[string]string
)

func init() {
	applyCmd.PersistentFlags().StringToStringVarP(&arguments, "argument", "a", map[string]string{}, "")
}

var applyCmd = &cobra.Command{
	Use:   "apply [flags] filename",
	Short: "Apply the runbook",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.Apply(args[0], arguments)
	},
}
