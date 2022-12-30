package qBittorent

import (
	"context"
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

func New(host, user, password string) (*Client, error) {
	client := NewClientWithConfig(Config{BasePath: host})
	err := client.Login(context.Background(), user, password)
	if err != nil {
		return nil, err
	}
	return client, nil
}
