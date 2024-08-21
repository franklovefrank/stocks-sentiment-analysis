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
		fmt.Println("Running analyze command...")
		symbol, _ := cmd.Flags().GetString("symbol")
		startDate, _ := cmd.Flags().GetString("start-date")
		endDate, _ := cmd.Flags().GetString("end-date")

		if symbol == "" || startDate == "" || endDate == "" {
			var symbol, startDate, endDate string
			fmt.Println("Enter Stock Symbol:")
			fmt.Scanln(&symbol)
			fmt.Println("Enter Start Date (YYYY-MM-DD):")
			fmt.Scanln(&startDate)
			fmt.Println("Enter End Date (YYYY-MM-DD):")
			fmt.Scanln(&endDate)
		}

		// Initialize API clients
		yahooClient := &api.YahooFinanceClient{}
		blueskyClient := &api.BlueskyClient{}
		googleNLPClient := &api.GoogleNLPClient{}

		// Initialize facade
		apiFacade := facade.NewAPIFacade(yahooClient, blueskyClient, googleNLPClient)

		// Initialize services
		stockService := service.NewStockService(apiFacade)

		// Fetch and analyze posts
		postsResponse, err := apiFacade.FetchPosts(symbol, "eyJhbGciOiJFUzI1NksifQ.eyJzY29wZSI6ImNvbS5hdHByb3RvLmFjY2VzcyIsInN1YiI6ImRpZDpwbGM6aW13YTJiY3c2cnY2NDd6cG1oeHc2bXV5IiwiaWF0IjoxNzI0MjQzNTEzLCJleHAiOjE3MjQyNTA3MTMsImF1ZCI6ImRpZDp3ZWI6cG9yY2luaS51cy1lYXN0Lmhvc3QuYnNreS5uZXR3b3JrIn0.EPSlqyPPnxqg-d_Nh3_b3wPyjafPJ3sizuKJbb0gGalXtAZAQeNfLd6Ci_9my-_WI6BjxTUukWA6OcCzxOb0EA", startDate, endDate)
		if err != nil {
			fmt.Printf("Error fetching posts: %v\n", err)
			return
		}

		var positive, neutral, negative int
		for _, post := range postsResponse.Posts {
			postText := post.Record.Text
			if strings.TrimSpace(postText) == "" {
				continue
			}

			result, err := apiFacade.AnalyzeSentiment(postText)
			if err != nil {
				fmt.Printf("Error analyzing sentiment: %v\n", err)
				return
			}

			if result.Score > 0 {
				positive++
			} else if result.Score == 0 {
				neutral++
			} else {
				negative++
			}
		}

		total := positive + neutral + negative
		positivePercent := float64(positive) / float64(total) * 100
		neutralPercent := float64(neutral) / float64(total) * 100
		negativePercent := float64(negative) / float64(total) * 100

		// Calculate stock metrics
		returns, drawdown, err := stockService.CalculateReturnsAndDrawdown(symbol, startDate, endDate)
		if err != nil {
			fmt.Printf("Error calculating returns and drawdown: %v\n", err)
			return
		}

		fmt.Println("\n+----------------------------------------------------+")
		fmt.Println("|  - Summary:                                       |")
		fmt.Printf("|    \"Stock Symbol: %s\"                          |\n", symbol)
		fmt.Printf("|    \"Start Date: %s\"                             |\n", startDate)
		fmt.Printf("|    \"End Date: %s\"                               |\n", endDate)
		fmt.Println("|    - Sentiment Analysis:                         |")
		fmt.Printf("|      - Positive: %.2f%%                          |\n", positivePercent)
		fmt.Printf("|      - Neutral: %.2f%%                           |\n", neutralPercent)
		fmt.Printf("|      - Negative: %.2f%%                          |\n", negativePercent)
		fmt.Println("|    - Financial Metrics:                          |")
		fmt.Printf("|      - Simple Return: %.2f%%                     |\n", returns)
		fmt.Printf("|      - Maximum Drawdown: %.2f%%                  |\n", drawdown)
		fmt.Println("+----------------------------------------------------+")
	},
}

func ExecuteAnalyze() error {
	return analyzeCmd.Execute()
}
