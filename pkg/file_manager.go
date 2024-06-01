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

func SaveResult(respResult interface{}, format string) {
	if format == "JSON" {
		file, _ := json.MarshalIndent(respResult, "", " ")
		_ = os.WriteFile("results.json", file, 0644)
	} else {
		log.Println(respResult)
	}
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
