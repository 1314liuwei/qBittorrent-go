package qBittorent

import (
	"context"
	"encoding/json"
)

/*
https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#application
*/

type BuildInfo struct {
	QtVersion          string `json:"qt"`
	LibtorrentVersion  string `json:"libtorrent"`
	BoostVersion       string `json:"boost"`
	OpenSSLVersion     string `json:"openssl"`
	ApplicationBitness int    `json:"bitness"`
}

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

func (c *Client) GetBuildInfo(ctx context.Context) (*BuildInfo, error) {
	res, err := c.Get(ctx, "/api/v2/app/buildInfo", nil)
	if err != nil {
		return nil, err
	}

	result := &BuildInfo{}
	err = json.Unmarshal(res.Body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
