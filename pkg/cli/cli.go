package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stock-sentiment-cli",
	Short: "A CLI for stock sentiment analysis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Stock Sentiment Analysis CLI.")
		showMainMenu()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func showMainMenu() {
	fmt.Println("Choose an option:")
	fmt.Println("1. Analyze Stock and Sentiment")
	fmt.Println("2. View Report")
	fmt.Println("3. Export Data")
	fmt.Println("4. Exit")
	var choice int
	fmt.Scanln(&choice)
	handleChoice(choice)
}

func handleChoice(choice int) {
	switch choice {
	case 1:
		analyzeStockAndSentiment()
	case 2:
		viewReport()
	case 3:
		exportData()
	case 4:
		fmt.Println("Thank you for using the Stock Sentiment Analysis CLI. Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
		showMainMenu()
	}
}

// Define analyzeStockAndSentiment, viewReport, and exportData functions here.
