package cmd

import (
	"fmt"
	"os"
)

func ShowMainMenu() {
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
		fmt.Println("choice 1")
		if err := ExecuteAnalyze(); err != nil {
			fmt.Println("Error executing analyze command:", err)
		}
	case 2:
		if err := ExecuteReport(); err != nil {
			fmt.Println("Error executing report command:", err)
		}
	case 3:
		if err := ExecuteExport(); err != nil {
			fmt.Println("Error executing export command:", err)
		}
	case 4:
		fmt.Println("Thank you for using the Stock Sentiment Analysis CLI. Goodbye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
		ShowMainMenu()
	}
}
