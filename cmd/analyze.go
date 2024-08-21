package cmd

import (
	"fmt"
	"strings"

	"stock-sentiment-cli/internal/api"
	"stock-sentiment-cli/internal/facade"
	"stock-sentiment-cli/internal/service"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze stock and sentiment data",
	Run: func(cmd *cobra.Command, args []string) {
		symbol, startDate, endDate := getInput(cmd)

		apiFacade, stockService := initializeServices()

		positive, neutral, negative, err := fetchAndAnalyzePosts(apiFacade, symbol, startDate, endDate)
		if err != nil {
			fmt.Printf("Error fetching or analyzing posts: %v\n", err)
			return
		}

		returns, drawdown, err := stockService.CalculateReturnsAndDrawdown(symbol, startDate, endDate)
		if err != nil {
			fmt.Printf("Error calculating returns and drawdown: %v\n", err)
			return
		}

		displayResults(symbol, startDate, endDate, positive, neutral, negative, returns, drawdown)
	},
}

func getInput(cmd *cobra.Command) (string, string, string) {
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

	return symbol, startDate, endDate
}

func initializeServices() (facade.APIFacade, *service.StockService) {
	yahooClient := &api.YahooFinanceClient{}
	blueskyClient := &api.BlueskyClient{}
	googleNLPClient := &api.GoogleNLPClient{}

	apiFacade := facade.NewAPIFacade(yahooClient, blueskyClient, googleNLPClient)
	stockService := service.NewStockService(apiFacade)

	return apiFacade, stockService
}

func fetchAndAnalyzePosts(apiFacade facade.APIFacade, symbol, startDate, endDate string) (int, int, int, error) {
	postsResponse, err := apiFacade.FetchPosts(symbol, "eyJhbGciOiJFUzI1NksifQ.eyJzY29wZSI6ImNvbS5hdHByb3RvLmFjY2VzcyIsInN1YiI6ImRpZDpwbGM6aW13YTJiY3c2cnY2NDd6cG1oeHc2bXV5IiwiaWF0IjoxNzI0MjQzNTEzLCJleHAiOjE3MjQyNTA3MTMsImF1ZCI6ImRpZDp3ZWI6cG9yY2luaS51cy1lYXN0Lmhvc3QuYnNreS5uZXR3b3JrIn0.EPSlqyPPnxqg-d_Nh3_b3wPyjafPJ3sizuKJbb0gGalXtAZAQeNfLd6Ci_9my-_WI6BjxTUukWA6OcCzxOb0EA", startDate, endDate)
	if err != nil {
		return 0, 0, 0, err
	}

	var positive, neutral, negative int
	for _, post := range postsResponse.Posts {
		postText := post.Record.Text
		if strings.TrimSpace(postText) == "" {
			continue
		}

		result, err := apiFacade.AnalyzeSentiment(postText)
		if err != nil {
			return 0, 0, 0, err
		}

		if result.Score > 0 {
			positive++
		} else if result.Score == 0 {
			neutral++
		} else {
			negative++
		}
	}

	return positive, neutral, negative, nil
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
