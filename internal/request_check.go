package internal

import (
	"net/http"
	"sync"
	"time"
	filemanager "urlChecker2/pkg"
)

type UrlChecker struct {
	client http.Client
}

func NewChecker(cl Client) UrlChecker {
	return UrlChecker{
		client: cl.Client,
	}
}

func Start() {
	logFile := filemanager.InitLog()
	rateLimit, format := filemanager.ReadConfig("config.json")
	urlData := filemanager.ReadUrl("urls")
	client := NewClient()
	checker := NewChecker(client)
	results := checker.CheckUrls(urlData, rateLimit)
	filemanager.SaveResult(results, format)
	filemanager.CloseLog(logFile)
}

// This Go code defines a UrlChecker struct with methods to check multiple URLs concurrently with a specified rate limit using goroutines and channels.
// The CheckUrls method initializes channels for input and results, spawns goroutines to check each URL, and aggregates the results before returning them.

func (uc *UrlChecker) CheckUrls(urls []string, rateLimit int) []RespResult {
	wg := &sync.WaitGroup{}

	inputUrl := make(chan string)
	result := make(chan RespResult)
	wgCount := rateLimit

	go func() {
		defer close(inputUrl)

		for _, url := range urls {
			inputUrl <- url
		}

	}()

	go func() {
		for i := 0; i < wgCount; i++ {
			wg.Add(1)

			go uc.checkUrl(wg, inputUrl, result)
		}
		wg.Wait()
		close(result)
	}()

	reqResult := []RespResult{}

	for res := range result {
		reqResult = append(reqResult, res)
	}

	return reqResult
}

func (uc *UrlChecker) checkUrl(wg *sync.WaitGroup, inputUrl <-chan string, result chan<- RespResult) {
	defer wg.Done()
	for url := range inputUrl {

		startTime := time.Now()

		resp, err := uc.client.Get(url)
		if resp != nil {
			defer resp.Body.Close()
		}

		responseTime := time.Since(startTime).Milliseconds()

		result <- setResult(responseTime, url, err)
	}
}
