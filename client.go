package qBittorent

import (
	"net/http"
)

type Config struct {
	BasePath string
}

type Client struct {
	BasePath string
	client   *http.Client
}

func NewClientWithConfig(conf Config) *Client {
	return &Client{
		BasePath: conf.BasePath,
		client:   &http.Client{},
	}
}

func NewClient() *Client {
	return NewClientWithConfig(Config{BasePath: "http://localhost:8080"})
}
