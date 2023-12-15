package apply

import (
	"github.com/keithrooney/runbooks/cli/internal/runbook"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "apply [flags] filename",
	Short: "Apply the runbook",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		runbook, err := runbook.Load(filepath)
		if err != nil {
			return err
		}
		cmd.Printf("RUNBOOK [%s]\n\n", runbook.Title)
		for _, step := range runbook.Mitigation.Steps {
			cmd.Printf("STEP [%s] => [%s]\n\n", step.Name, step.Shell)
			if err := step.Execute(); err != nil {
				return err
			}
		}
		return nil
	},
}
