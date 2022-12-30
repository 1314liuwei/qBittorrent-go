package qBittorent

import (
	"context"
)

func (c *Client) GetApplicationVersion(ctx context.Context) (string, error) {
	res, err := c.Get(ctx, "/api/v2/app/version", nil)
	if err != nil {
		return "", err
	}
	return string(res.Body), nil
}

func (c *Client) GetAPIVersion(ctx context.Context) (string, error) {
	res, err := c.Get(ctx, "/api/v2/app/webapiVersion", nil)
	if err != nil {
		return "", err
	}
	return string(res.Body), nil
}
