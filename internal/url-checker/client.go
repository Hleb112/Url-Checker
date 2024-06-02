package urlchecker

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	defaultClientTimeout = 15 * time.Second
	maxRetries           = 2
	maxBackoffTime       = 15 * time.Second
)

type Client struct {
	Client *http.Client
}

func NewClient() *Client {
	client := &http.Client{
		Timeout: defaultClientTimeout,
	}

	return &Client{
		Client: client,
	}
}

func (uc *UrlChecker) DoWithRetry(req *http.Request) (*http.Response, error) {
	retryResult, err := uc.retryWithBackoff(req)
	if err != nil {
		log.Printf("Request failed: %v", err)
	}

	return retryResult, nil
}

func (uc *UrlChecker) retryWithBackoff(req *http.Request) (*http.Response, error) {
	backoff := 1 * time.Second
	for i := 0; i < maxRetries; i++ {
		resp, err := uc.Client.Do(req)
		if err == nil && (resp.StatusCode < 500 || resp.StatusCode == 0) {
			return resp, nil
		}

		if resp != nil {
			resp.Body.Close()
		}

		log.Printf("Request failed (attempt %d/%d): %v", i+1, maxRetries, err)

		if backoff > maxBackoffTime {
			backoff = maxBackoffTime
		}

		jitter := time.Duration(rand.Int63n(int64(backoff)))
		time.Sleep(backoff + jitter)
		backoff *= 2
	}

	return nil, fmt.Errorf("request failed after %d retries", maxRetries)
}
