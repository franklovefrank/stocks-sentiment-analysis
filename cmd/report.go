package cmd

import (
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "View the report",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation of the report command
	},
}

func ExecuteReport() error {
	return reportCmd.Execute()
}
