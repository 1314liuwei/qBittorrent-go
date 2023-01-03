package qBittorent

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"net/http"
	"strings"
)

/*
https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#transfer-info
*/

type TransferInfo struct {
	ConnectionStatus string `json:"connection_status"`
	DhtNodes         int    `json:"dht_nodes"`
	DlInfoData       int64  `json:"dl_info_data"`
	DlInfoSpeed      int    `json:"dl_info_speed"`
	DlRateLimit      int    `json:"dl_rate_limit"`
	UpInfoData       int    `json:"up_info_data"`
	UpInfoSpeed      int    `json:"up_info_speed"`
	UpRateLimit      int    `json:"up_rate_limit"`
}

func (c *Client) GetGlobalTransferInfo(ctx context.Context) (*TransferInfo, error) {
	res, err := c.Get(ctx, "/api/v2/transfer/info", nil)
	if err != nil {
		return nil, err
	}

	result := &TransferInfo{}
	err = json.Unmarshal(res.Body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetAlternativeSpeedLimitsState(ctx context.Context) (bool, error) {
	res, err := c.Get(ctx, "/api/v2/transfer/speedLimitsMode", nil)
	if err != nil {
		return false, err
	}

	if bytes.Contains(res.Body, []byte{'1'}) {
		return true, nil
	} else {
		return false, nil
	}
}

func (c *Client) ToggleAlternativeSpeedLimits(ctx context.Context) error {
	res, err := c.PostFormData(ctx, "/api/v2/transfer/toggleSpeedLimitsMode", nil)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("toggle failed")
	}
	return nil
}

func (c *Client) GetGlobalDownloadLimit(ctx context.Context) (int, error) {
	res, err := c.Get(ctx, "/api/v2/transfer/downloadLimit", nil)
	if err != nil {
		return -1, err
	}

	return gconv.Int(strings.TrimSpace(string(res.Body))), nil
}

func (c *Client) SetGlobalDownloadLimit(ctx context.Context, limit int) error {
	res, err := c.PostFormData(ctx, "/api/v2/transfer/setDownloadLimit", map[string]interface{}{
		"limit": limit,
	})
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("set failed")
	}
	return nil
}

func (c *Client) GetGlobalUploadLimit(ctx context.Context) (int, error) {
	res, err := c.Get(ctx, "/api/v2/transfer/uploadLimit", nil)
	if err != nil {
		return -1, err
	}

	return gconv.Int(strings.TrimSpace(string(res.Body))), nil
}
