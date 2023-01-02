package qBittorent

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/util/gconv"
)

/*
https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#log
*/

type MainLog struct {
	Id        int    `json:"id"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Type      int    `json:"type"`
}

type QueryMainLogParam struct {
	Normal      bool `json:"normal"`
	Info        bool `json:"info"`
	Warning     bool `json:"warning"`
	Critical    bool `json:"critical"`
	LastKnownId int  `json:"last_known_id"`
}

type PeerLog struct {
	Id        int    `json:"id"`
	IP        string `json:"ip"`
	Timestamp int    `json:"timestamp"`
	Blocked   bool   `json:"blocked"`
	Reason    string `json:"reason"`
}

func (c *Client) GetLog(ctx context.Context, params *QueryMainLogParam) ([]MainLog, error) {
	res, err := c.Get(ctx, "/api/v2/log/main", gconv.Map(*params, "json"))
	if err != nil {
		panic(err)
		return nil, err
	}

	var result []MainLog
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetPeerLog(ctx context.Context, lid int) ([]PeerLog, error) {
	res, err := c.Get(ctx, "/api/v2/log/peers", map[string]interface{}{"last_known_id": lid})
	if err != nil {
		panic(err)
		return nil, err
	}

	var result []PeerLog
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
