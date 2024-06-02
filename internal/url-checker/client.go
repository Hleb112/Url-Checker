package urlchecker

import (
	"net/http"
	"time"
)

const (
	defaultClientTimeout = 15 * time.Second
)

type Client struct {
	Client http.Client
}

type Option func(*Client)

func NewClient(opts ...Option) Client {
	client := http.Client{
		Timeout: defaultClientTimeout,
	}

	c := Client{
		Client: client,
	}

	for _, opt := range opts {
		opt(&c)
	}

	return c
}
