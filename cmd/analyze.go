package cmd

import (
	"fmt"
	"stock-sentiment-cli/internal/model"
	"stock-sentiment-cli/internal/setup"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze stock and sentiment data",
	Run: func(cmd *cobra.Command, args []string) {
		query := getInput(cmd)

		stockService, sentimentService := setup.InitializeServices()

		positive, neutral, negative, err := sentimentService.AnalyzePosts(query)
		if err != nil {
			fmt.Printf("Error fetching or analyzing posts: %v\n", err)
			return
		}

		returns, drawdown, err := stockService.CalculateReturnsAndDrawdown(query)
		if err != nil {
			fmt.Printf("Error calculating returns and drawdown: %v\n", err)
			return
		}

		displayResults(query.Symbol, query.StartDate, query.EndDate, positive, neutral, negative, returns, drawdown)
	},
}

func getInput(cmd *cobra.Command) model.StockQuery {
	symbol, _ := cmd.Flags().GetString("symbol")
	startDate, _ := cmd.Flags().GetString("start-date")
	endDate, _ := cmd.Flags().GetString("end-date")

	if symbol == "" || startDate == "" || endDate == "" {
		fmt.Println("Enter Stock Symbol:")
		fmt.Scanln(&symbol)
		fmt.Println("Enter Start Date (YYYY-MM-DD):")
		fmt.Scanln(&startDate)
		fmt.Println("Enter End Date (YYYY-MM-DD):")
		fmt.Scanln(&endDate)
	}

	return model.StockQuery{
		Symbol:    symbol,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func displayResults(symbol, startDate, endDate string, positive, neutral, negative int, returns, drawdown float64) {
	total := positive + neutral + negative
	positivePercent := float64(positive) / float64(total) * 100
	neutralPercent := float64(neutral) / float64(total) * 100
	negativePercent := float64(negative) / float64(total) * 100

	fmt.Println("\n+-------------------------------------------------------+")
	fmt.Println("  - Summary:                                          ")
	fmt.Printf("    Stock Symbol: %-30s\n", symbol)
	fmt.Printf("    Start Date: %-32s\n", startDate)
	fmt.Printf("    End Date: %-34s\n", endDate)
	fmt.Println("    - Sentiment Analysis:                            ")
	fmt.Printf("      - Positive: %6.2f%%\n", positivePercent)
	fmt.Printf("      - Neutral: %6.2f%%\n", neutralPercent)
	fmt.Printf("      - Negative: %6.2f%%\n", negativePercent)
	fmt.Println("    - Financial Metrics:                             ")
	fmt.Printf("      - Simple Return: %6.2f%%\n", returns)
	fmt.Printf("      - Maximum Drawdown: %6.2f%%\n", drawdown)
	fmt.Println("+-------------------------------------------------------+")
}

func ExecuteAnalyze() error {
	return analyzeCmd.Execute()
}
