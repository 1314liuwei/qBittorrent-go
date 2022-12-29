package qBittorent

import (
	"context"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

func (c *Client) Login(ctx context.Context, username, password string) error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	c.client.Jar = jar

	res, err := c.PostFormData(ctx, "/api/v2/auth/login", map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK && strings.Contains(string(res.Body), "Ok.") {
		return nil
	}

	return errors.New("login failed")
}
