package cmd

import (
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export data",
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation of the export command
	},
}

func ExecuteExport() error {
	return exportCmd.Execute()
}
