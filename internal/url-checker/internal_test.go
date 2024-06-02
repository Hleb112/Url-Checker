package urlchecker

import (
	"errors"
	"net/http"
	"sync"
	"testing"
	"urlChecker2/internal/models"
)

func TestSetResult_Positive(t *testing.T) {
	var err error
	reqDuration := int64(100)
	url := "https://example.com"
	err = nil

	result := models.SetResult(reqDuration, url, err)

	if result.Url != url {
		t.Errorf("Expected URL to be %s, but got %s", url, result.Url)
	}

	if result.Available != true {
		t.Errorf("Expected Available to be true, but got false")
	}

	if result.ReqDuration != reqDuration {
		t.Errorf("Expected ReqDuration to be %d, but got %d", reqDuration, result.ReqDuration)
	}

	if result.Err != nil {
		t.Errorf("Expected Err to be nil, but got %v", result.Err)
	}
}

func TestSetResult_Negative(t *testing.T) {
	reqDuration := int64(0)
	url := "https://example.com"
	err := errors.New("Error message")

	result := models.SetResult(reqDuration, url, err)

	if result.Url != url {
		t.Errorf("Expected URL to be %s, but got %s", url, result.Url)
	}

	if result.Available != false {
		t.Errorf("Expected Available to be false, but got true")
	}

	if result.ReqDuration != reqDuration {
		t.Errorf("Expected ReqDuration to be %d, but got %d", reqDuration, result.ReqDuration)
	}

	if result.Err == nil {
		t.Errorf("Expected Err to be not nil, but got nil")
	}
}

func TestCheckUrl_PositiveCase(t *testing.T) {
	uc := UrlChecker{Client: http.Client{}}
	inputUrl := make(chan string)
	result := make(chan models.RespResult)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go uc.checkUrl()

	inputUrl <- "https://www.google.com"
	close(inputUrl)

	respResult := <-result

	if respResult.Err != nil {
		t.Errorf("Expected no error, got %v", respResult.Err)
	}

	if respResult.Url != "https://www.google.com" {
		t.Errorf("Expected url to be https://www.google.com, got %s", respResult.Url)
	}
}

func TestCheckUrl_NegativeCase(t *testing.T) {
	uc := UrlChecker{Client: http.Client{}}
	inputUrl := make(chan string)
	result := make(chan models.RespResult)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go uc.checkUrl()

	inputUrl <- "invalidurl"
	close(inputUrl)

	respResult := <-result

	if respResult.Err == nil {
		t.Error("Expected an error, got nil")
	}
}

func TestCheckUrls_Positive(t *testing.T) {
	uc := &UrlChecker{}
	urls := []string{"http://example.com"}
	rateLimit := 2
	expected := []models.RespResult{
		{Url: "http://example.com", Available: true},
	}

	results := uc.CheckUrls(urls, rateLimit)
	realResult := results[0].Available

	if realResult != expected[0].Available {
		t.Errorf("Expected %v, but got %v", expected, results)
	}
}

func TestCheckUrls_Negative(t *testing.T) {
	uc := &UrlChecker{}
	urls := []string{"http://invalidurl"}
	rateLimit := 2
	expected := []models.RespResult{
		{Url: "http://invalidurl", Available: false, ReqDuration: 0},
	}

	results := uc.CheckUrls(urls, rateLimit)
	realResult := results[0].Available

	if realResult != expected[0].Available {
		t.Errorf("Expected %v, but got %v", expected, results)
	}
}
