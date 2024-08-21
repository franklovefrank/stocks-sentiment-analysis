# Stock Sentiment Analysis CLI

## Summary

The Stock Sentiment Analysis app allows users to explore the relationship between stock performance and public sentiment through a simple, intuitive CLI. The app utilizes the Yahoo Finance, Bluesky, and Google Cloud Natural Language APIs to combine sentiment analysis of tweets related to a specific stock symbol with historical data for that stock. Users can use the CLI to view sentiment analysis results, manage previous reports, and export data in various formats.

## Setup Instructions

### 1. Setup Bluesky API

1. **Create a Bluesky Developer Account:**
   - Visit the [Bluesky developer portal](https://docs.bsky.app/docs/get-started) and sign up for an account.

2. **Obtain API Credentials:**
   - After signing up, navigate to the API section and generate an API token.

3. **Configure API Token:**
   - Store your API token securely. You'll need to provide it when running the app.

### 2. Setup Google Cloud Natural Language API

1. **Create a Google Cloud Account:**
   - Sign up for a Google Cloud account at [Google Cloud Console](https://console.cloud.google.com/).

2. **Create a New Project:**
   - In the Google Cloud Console, create a new project.

3. **Enable the Natural Language API:**
   - Navigate to the API & Services dashboard, and enable the Natural Language API for your project.

4. **Generate API Credentials:**
   - Go to the "Credentials" page, create a new API key or service account, and download the JSON key file. Securely store this key file.

5. **Set Up Authentication:**
   - Set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to the path of your JSON key file.

   ```bash
   export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/credentials-file.json"
   ```

### 3. Install Dependencies

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/franklovefrank/stock-sentiment-cli.git
   cd stock-sentiment-cli
   ```

2. **Install Go Modules:**

   ```bash
   go mod tidy
   ```

   This command will download and install the required Go dependencies.

### 4. Run the App

1. **Build the Application:**

   ```bash
   go build -o stock-sentiment-cli
   ```

2. **Run the Application:**

   ```bash
   ./stock-sentiment-cli
   ```

   If you run the application without any arguments, it will display the main menu with options for analysis, viewing reports, exporting data, and exiting.

3. **Using CLI Commands:**

   - **Analyze Stock and Sentiment:**

     ```bash
     ./stock-sentiment-cli analyze --symbol AAPL --start-date 2024-01-01 --end-date 2024-01-31
     ```

   - **View Report:**

     ```bash
     ./stock-sentiment-cli report
     ```

   - **Export Data:**

     ```bash
     ./stock-sentiment-cli export
     ```
