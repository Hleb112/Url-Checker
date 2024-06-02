# URL Checker

URL Checker is a Go application that checks the status of a list of URLs and saves the results to a file or prints them to the console. It supports configurable rate limiting and output formatting.

## Features

- Check the status of multiple URLs concurrently
- Configurable rate limiting to avoid overwhelming the target servers
- Output results in JSON or plain text format
- Log errors and invalid URLs to a separate log file

## Installation

1. Clone the repository:

git clone https://github.com/Hleb112/url-checker.git


2. Navigate to the project directory:

cd url-check


## Configuration

The application reads its configuration from a JSON file named `config.json`. The configuration file should be placed in the same directory as the executable.

Example `config.json`:

{
  "rate_limit": 15,
  "format": "JSON"
}

rateLimit.limit: The maximum number of concurrent requests to make (default: 10).
format.output: The output format for the results. Supported values are "json" and saving to log file (default: "json").

## Usage

1) execute main.go .

2) The application expects a file named urls.txt in the same directory, containing a list of URLs to check, one per line.
Example urls.txt:

https://example.com
https://google.com
https://invalid-url

The results will be saved to a file named results.json or printed to the console, depending on the configured output format.

## Testing and comments

project has high test coverage and a lot of comments to make navigation easy








