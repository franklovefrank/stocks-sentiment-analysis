# Stock Sentiment Analysis CLI
<img width="527" alt="Screenshot 2024-08-21 at 17 53 03" src="https://github.com/user-attachments/assets/48cc0366-5f42-46a1-89d1-6cfaedf70c1b">

## Summary

The Stock Sentiment Analysis app allows users to explore the relationship between stock performance and public sentiment through a simple, intuitive CLI. The app utilizes the Yahoo Finance, Bluesky, and Google Cloud Natural Language APIs to combine sentiment analysis of tweets related to a specific stock symbol with historical data for that stock. Users can use the CLI to view sentiment analysis results, manage previous reports, and export data in various formats.

## Setup Instructions

### Setup Bluesky API

1. **Obtain Bluesky Access Token**
   - Bluesky does not have a webpage to get API keys yet. To generate a token, you can follow the instructions [here](https://docs.bsky.app/docs/api/com-atproto-server-create-session)

### Setup Google Cloud Natural Language API

1. **Create a Google Cloud Account:**
   - Sign up for a Google Cloud account at [Google Cloud Console](https://console.cloud.google.com/).

2. **Create a New Project:**
   - In the Google Cloud Console, create a new project.

3. **Enable the Natural Language API:**
   - Navigate to the API & Services dashboard, and enable the Natural Language API for your project.

4. **Generate API Credentials:**
   - Go to the "Credentials" page, create a new API key or service account, and download the JSON key file. Securely store this key file.


### Environment Setup 

1. **Copy the `.env.example` file:**
   - `cp .env.example .env`

2. **Set environmental variables:**
   - In the newly created `.env` file, add your Bluesky API token and the path of your JSON key file 

### Install Dependencies

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

### Run the App

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
