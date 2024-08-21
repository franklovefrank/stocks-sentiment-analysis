# Design Document for Stock Sentiment CLI Tool
    
## Summary 
The Stock Sentiment Analysis app allows users to explore the relationship between stock performance and public sentiment through a simple, intuitive CLI. The app utilizes the Yahoo Finance, Reddit, and Google Cloud Natural Language APIs to combine sentiment anaylsis of tweets related to a specific stock symbol with historical data for that stock. Users can use the CLI to view sentiment analysis results, manage previous reports, and export data in various formats.

## Architecture 
### Overview 
- **CLI Interface**: Handles user input and displays results.
- **API Integrations**:
    - **Yahoo Finance API**: Fetch historical stock data.
    - **Reddit API**: Collect recent tweets about the stock.
    - **Google Cloud Natural Language API**: Perform sentiment analysis on tweets.
- **Data Processing**: Analyze stock data and sentiment to provide insights.
- **Report Generation**: Generate and display reports combining stock metrics and sentiment analysis.
### Design Patterns
  - **Facade Pattern**: Simplifies interactions with the Yahoo Finance, Reddit, and Google Cloud Natural Language APIs.
  - **Strategy Pattern**: Allows dynamic application of different sentiment analysis filters or strategies.
  - **Builder Pattern**: Constructs complex reports by combining data from various sources in a modular manner.
  - **Decorator Pattern**: Enhances the functionality of sentiment analysis results, such as adding context or aggregating sentiment data.
### Misc. Considerations
1. **Metrics**
    - Integrate Prometheus or another monitoring tool to track metrics like API call duration, data processing time, and request counts.
    - Example metrics:
        - API Call Duration: Measures how long it takes to fetch data from Yahoo Finance, Reddit, and Google Cloud Natural Language.
        - Request Counts: Tracks how many times different commands or strategies are used.
2. **Testing and Benchmarking**
    - **Unit Tests**: For individual components and functions.
    - **Integration Tests**: For interactions with Yahoo Finance API, Reddit API, and Google Cloud Natural Language API.
    - **Benchmarking**: Measure performance using Go’s benchmarking tools.
3. **Scalability and Performance**
    - Optimize data fetching and processing.
    - Implement efficient algorithms and data structures for risk calculations and sentiment analysis.
    - Use concurrency to enhance performance and responsiveness.
  
## CLI Flow Chart

```
+---------------------------+
|  Start CLI Application    |
|  - Welcome message:       |
|    "Welcome to the Stock   |
|    Sentiment Analysis CLI. |
|    Choose an option to     |
|    get started."           |
+------------+--------------+
             |
             v
+------------+--------------+
|  Display Main Menu         |
|  - Options:                |
|    1. Analyze Stock and    |
|       Sentiment            |
|    2. View Report          |
|    3. Export Data          |
|    4. Exit                 |
+------------+--------------+
             |
             v
+------------+--------------+
|  User Chooses Option 1     |
|  "Analyze Stock and        |
|  Sentiment"                |
|  - Input Prompt:           |
|    "Enter Stock Symbol:    |
|    [User Input]"           |
|    "Enter Start Date (YYYY-MM-DD): [User Input]" |
|    "Enter End Date (YYYY-MM-DD): [User Input]"   |
+------------+--------------+
             |
             v
+------------+--------------+
|  Fetch and Analyze Data    |
|  - "Fetching recent tweets..."|
|  - "Analyzing sentiment..."|
+------------+--------------+
             |
             v
+------------+--------------+
|  Display Initial Results   |
|  - Summary:                |
|    "Stock Symbol: [Symbol]"|
|    "Start Date: [Date]"    |
|    "End Date: [Date]"      |
|    - Sentiment Analysis:   |
|      - Positive: [X%]      |
|      - Neutral: [Y%]       |
|      - Negative: [Z%]      |
+------------+--------------+
             |
             v
+------------+--------------+
|  Display Main Menu         |
|  - Prompt: "What would you |
|    like to do next?"       |
|    1. Analyze Stock and    |
|       Sentiment            |
|    2. View Report          |
|    3. Export Data          |
|    4. Exit                 |
+------------+--------------+
             |
             v
+------------+--------------+
|  User Chooses Option 2     |
|  "View Report"             |
|  - Prompt: "Select report   |
|    to view: [List of reports or most recent]"|
+------------+--------------+
             |
             v
+------------+--------------+
|  Display Detailed Report   |
|  - Summary:                |
|    "Detailed Sentiment Report" |
|    - Sentiment Breakdown:  |
|      - Positive: [X%]      |
|      - Neutral: [Y%]       |
|      - Negative: [Z%]      |
+------------+--------------+
             |
             v
+------------+--------------+
|  Display Main Menu         |
|  - Prompt: "What would you |
|    like to do next?"       |
|    1. Analyze Stock and    |
|       Sentiment            |
|    2. View Report          |
|    3. Export Data          |
|    4. Exit                 |
+------------+--------------+
             |
             v
+------------+--------------+
|  User Chooses Option 3     |
|  "Export Data"             |
|  - Prompt: "Choose export   |
|    format:                 |
|    1. CSV                  |
|    2. JSON                 |
|    3. Cancel               |
+------------+--------------+
             |
             v
+------------+--------------+
|  Export Data               |
|  - "Exporting to [Format]..."|
|  - "Data exported successfully!"|
+------------+--------------+
             |
             v
+------------+--------------+
|  Display Main Menu         |
|  - Prompt: "What would you |
|    like to do next?"       |
|    1. Analyze Stock and    |
|       Sentiment            |
|    2. View Report          |
|    3. Export Data          |
|    4. Exit                 |
+------------+--------------+
             |
             v
+------------+--------------+
|  User Chooses Option 4     |
|  "Exit"                    |
|  - Goodbye message:        |
|    "Thank you for using the|
|    Stock Sentiment Analysis|
|    CLI. Goodbye!"          |
+---------------------------+


+---------------------------+

```

## Project Structure
```
stock-sentiment-cli/
│
├── cmd/
│   ├── main.go                          # Entry point of the application
│
├── pkg/
│   ├── api/
│   │   ├── yahoo_finance.go             # Interaction with Yahoo Finance API
│   │   ├── bluesky.go                   # Interaction with Bluesky API
│   │   ├── google_nlp.go                # Interaction with Google Cloud Natural Language API
│   ├── analysis/
│   │   ├── sentiment_analyzer.go        # Sentiment analysis logic using Google Cloud NLP
│   ├── cli/
│   │   ├── cli.go                       # CLI interface and input handling
│   ├── report/
│   │   ├── report_generator.go          # Logic for generating and displaying reports
│   ├── export/
│   │   ├── export.go                    # Export functionalities (e.g., to CSV or JSON)
│   ├── model/
│   │   ├── tweet.go                    # Models for tweet data
│   │   ├── stock_data.go                # Models for stock data
│   │   ├── sentiment_result.go          # Definition of sentiment analysis results
│
├── test/
│   ├── api/
│   │   ├── yahoo_finance_test.go        # Tests for Yahoo Finance API interactions
│   │   ├── reddit_test.go              # Tests for Reddit API interactions
│   │   ├── google_nlp_test.go           # Tests for Google Cloud Natural Language API interactions
│   ├── analysis/
│   │   ├── sentiment_analyzer_test.go   # Tests for sentiment analysis logic
│   ├── cli/
│   │   ├── cli_test.go                  # Tests for CLI interface and input handling
│   ├── report/
│   │   ├── report_generator_test.go      # Tests for report generation and display
│   ├── export/
│   │   ├── export_test.go               # Tests for data export functionalities
│
├── Dockerfile                           # Dockerfile for containerization
├── go.mod                               # Go module file
├── go.sum                               # Go module checksums
└── README.md                            # Project documentation
```