package internal

import (
	"crypto/tls"
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
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	} //ОТКЛЮЧИЛ ВСЮ ЗАЩИТУ НОРМ ???

	client := http.Client{
		Timeout:   defaultClientTimeout,
		Transport: tr,
	}

	c := Client{
		Client: client,
	}

	for _, opt := range opts {
		opt(&c)
	}

	return c
}
