package models

import "time"

type Result struct {
	Url       string
	Reachable bool
	Time      int64
	Error     error
}

type Config struct {
	RateLimit time.Time `json:"rate_limit"`
}

type WorkerResults struct {
	Result Result
}
