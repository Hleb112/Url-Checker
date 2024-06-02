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


## Project Structure

![image](https://github.com/Hleb112/Url-Checker/assets/54846233/092c6695-d15f-4452-ba78-7fefc39b4b77)


cmd/main.go: The entry point of the application.

internal/client.go: Contains the HTTP client configuration.

internal/request_check.go: Contains the implementation for checking URLs.

pkg/file_manager.go: Contains functions for reading and writing files.

pkg/url_checker.go: Contains the main logic for checking URLs and handling results.

config.json: The configuration file for the application.

urls.txt: The input file containing the list of URLs to check.

go.mod and go.sum: Go module files for managing dependencies.

## Testing and comments

project has high test coverage and a lot of comments to make navigation easy








