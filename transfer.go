package qBittorent

import (
	"context"
	"encoding/json"
)

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

func (c *Client) GetTransferInfo(ctx context.Context) (*TransferInfo, error) {
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
