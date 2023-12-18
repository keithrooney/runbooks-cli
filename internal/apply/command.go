package apply

import (
	"github.com/keithrooney/runbooks/cli/internal/core"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "apply [flags] filename",
	Short: "Apply the runbook",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.Apply(args[0])
	},
}
