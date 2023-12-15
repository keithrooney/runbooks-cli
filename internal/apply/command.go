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
		applicable, err := runbook.Mitigation.Evaluate()
		if err != nil {
			return err
		}
		if applicable {
			cmd.Printf("RUNBOOK [%s] => [applicable]\n\n", runbook.Title)
			return runbook.Mitigation.Execute(cmd.Printf)
		}
		cmd.Printf("RUNBOOK [%s] => [skipped]\n\n", runbook.Title)
		return nil
	},
}
