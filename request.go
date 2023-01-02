package qBittorent

import (
	"bytes"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Response struct {
	StatusCode int
	Body       []byte
}

func (c *Client) PostFormData(ctx context.Context, p string, data map[string]string) (*Response, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for k, v := range data {
		err := writer.WriteField(k, v)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(http.MethodPost, c.BasePath+p, payload)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.WithContext(ctx)

	res, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	all, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{
		StatusCode: res.StatusCode,
		Body:       all,
	}

	return response, nil
}

func (c *Client) Get(ctx context.Context, p string, data map[string]interface{}) (*Response, error) {
	payload := url.Values{}
	for k, v := range data {
		payload.Set(k, gconv.String(v))
	}

	Url, err := url.Parse(c.BasePath + p)
	if err != nil {
		return nil, err
	}
	Url.RawQuery = payload.Encode()
	urlPath := Url.String()

	resp, err := c.client.Get(urlPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{
		StatusCode: resp.StatusCode,
		Body:       content,
	}
	return response, nil
}
