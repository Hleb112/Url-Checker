package filemanager

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	url "net/url"
	"os"
	"strings"
)

// This Go code defines functions for initializing and closing a log file.

func InitLog() *os.File {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)

	return file
}

func CloseLog(file *os.File) {
	if file != nil {
		file.Close()
	}
}

// ReadUrl defines a function ReadUrl that reads a file, extracts URLs from it, and returns them as a string array.
// It checks if each extracted value is a valid URL using the IsUrl function and logs any invalid URLs.
// If the file cannot be opened, it logs an error and returns nil.

func ReadUrl(fileName string) []string {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		log.Println("Unable to open file:", err)
		return nil
	}

	rawData, err := io.ReadAll(file)
	stringArray := strings.Split(string(rawData), "\r\n")

	var urlArray []string

	for _, value := range stringArray {
		if !IsUrl(value) {
			log.Println(fmt.Sprintf("%v is not Url", value))
			continue
		}
		urlArray = append(urlArray, value)
	}

	return urlArray
}

// ReadConfig defines a function that reads a configuration file,
// parses the data into RateLimit and Format structs using JSON unmarshalling,
// and returns the limit from RateLimit and the format from Format.

func ReadConfig(fileName string) (int, string) {
	var rateLimit RateLimit
	var format Format

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		log.Println("Unable to open config:", err)
		return 0, ""
	}

	rawData, err := io.ReadAll(file)
	err = json.Unmarshal(rawData, &rateLimit)
	if err != nil {
		log.Println("Unable to parse RateLimit:", err)
		return 0, ""
	}

	err = json.Unmarshal(rawData, &format)
	if err != nil {
		log.Println("Unable to parse Format:", err)
		return 0, ""
	}

	return rateLimit.Limit, format.Format
}

// SaveResult saves the given response result to a file or prints it to the console
// based on the specified format.

func SaveResult(respResult interface{}, format string) {
	if format == "JSON" {
		file, _ := json.MarshalIndent(respResult, "", " ")
		_ = os.WriteFile("results.json", file, 0644)
	} else {
		log.Println(respResult)
	}
}

// IsUrl defines a function IsUrl that checks if a given string is a valid URL by parsing it using the url package.
// It returns true if the string is a valid URL, and false otherwise.

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
