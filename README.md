# news-summarizer

## About

Command to retrieve a summary of newsworthy current events for a specific month in history. The source data comes from Wikipedia, and is processed through Claude AI to generate a summary article.

The articles are posted to Github Pages automatically once per month in https://github.com/greenmato/news-summary.

https://greenmato.github.io/news-summary

## How to run

To run locally:
```
ANTHROPIC_API_KEY=${ANTHROPIC_API_KEY} go run main.go --month=3 --year=2025
```

To run with Docker:
```
docker run -e ANTHROPIC_API_KEY=${ANTHROPIC_API_KEY} ghcr.io/greenmato/news-summarizer:latest app --month 10 --year 2024
```
