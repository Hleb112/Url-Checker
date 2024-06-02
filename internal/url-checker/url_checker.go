package urlchecker

import (
	"net/http"
	"sync"
	"time"
	"urlChecker2/internal/models"
	filemanager "urlChecker2/pkg/file-manager"
)

type UrlChecker struct {
	Client   http.Client
	wg       sync.WaitGroup
	inputUrl chan string
	result   chan models.RespResult
}

func NewUrlChecker(cl Client) UrlChecker {
	return UrlChecker{
		Client:   cl.Client,
		wg:       sync.WaitGroup{},
		inputUrl: make(chan string),
		result:   make(chan models.RespResult),
	}
}

func Start() {
	logFile := filemanager.InitLog()
	config := filemanager.ReadConfig("config.json")
	urlData := filemanager.ReadUrl("urls")
	client := NewClient()
	checker := NewUrlChecker(client)
	results := checker.CheckUrls(urlData, config.Limit)
	filemanager.SaveResult(results, config.Format)
	filemanager.CloseLog(logFile)
}

// This Go code defines a UrlChecker struct with methods to check multiple URLs concurrently with a specified rate limit using goroutines and channels.
// The CheckUrls method initializes channels for input and results, spawns goroutines to check each URL, and aggregates the results before returning them.

func (uc *UrlChecker) CheckUrls(urls []string, rateLimit int) []models.RespResult {
	wgCount := rateLimit

	go func() {
		defer close(uc.inputUrl)

		for _, url := range urls {
			uc.inputUrl <- url
		}

	}()

	go func() {
		for i := 0; i < wgCount; i++ {
			uc.wg.Add(1)

			go uc.checkUrl()
		}
		uc.wg.Wait()
		close(uc.result)
	}()

	reqResult := make([]models.RespResult, 0, len(urls))

	for res := range uc.result {
		reqResult = append(reqResult, res)
	}

	return reqResult
}

func (uc *UrlChecker) checkUrl() {
	defer uc.wg.Done()
	for url := range uc.inputUrl {

		startTime := time.Now()

		resp, err := uc.Client.Get(url)
		if resp != nil {
			defer resp.Body.Close()
		}

		responseTime := time.Since(startTime).Milliseconds()

		uc.result <- models.SetResult(responseTime, url, err)
	}
}
