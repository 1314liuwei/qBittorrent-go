package qBittorent

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
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
