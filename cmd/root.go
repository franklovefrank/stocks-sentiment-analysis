package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stock-sentiment-cli",
	Short: "A CLI for stock sentiment analysis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Stock Sentiment Analysis CLI.")
		ShowMainMenu()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
