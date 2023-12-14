package main

import (
	"fmt"
	"os"

	"github.com/keithrooney/runbooks/cli/internal/apply"
	"github.com/keithrooney/runbooks/cli/internal/export"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apply.Command, export.Command)
}

var rootCmd = &cobra.Command{
	Use: "runbooks",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}
