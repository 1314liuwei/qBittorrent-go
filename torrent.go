package qBittorent

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

/*
https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#torrent-management
*/

type GetTorrentListQuery struct {
	Filter   string   `json:"filter"`
	Category string   `json:"category"`
	Tag      string   `json:"tag"`
	Sort     string   `json:"sort"`
	Reverse  bool     `json:"reverse"`
	Limit    int      `json:"limit"`
	Offset   int      `json:"offset"`
	Hashes   []string `json:"hashes"`
}

type TorrentInfo struct {
	AddedOn           int     `json:"added_on"`
	AmountLeft        int64   `json:"amount_left"`
	AutoTmm           bool    `json:"auto_tmm"`
	Availability      float64 `json:"availability"`
	Category          string  `json:"category"`
	Completed         int64   `json:"completed"`
	CompletionOn      int     `json:"completion_on"`
	ContentPath       string  `json:"content_path"`
	DlLimit           int     `json:"dl_limit"`
	Dlspeed           int     `json:"dlspeed"`
	DownloadPath      string  `json:"download_path"`
	Downloaded        int64   `json:"downloaded"`
	DownloadedSession int     `json:"downloaded_session"`
	Eta               int     `json:"eta"`
	FLPiecePrio       bool    `json:"f_l_piece_prio"`
	ForceStart        bool    `json:"force_start"`
	Hash              string  `json:"hash"`
	InfohashV1        string  `json:"infohash_v1"`
	InfohashV2        string  `json:"infohash_v2"`
	LastActivity      int     `json:"last_activity"`
	MagnetUri         string  `json:"magnet_uri"`
	MaxRatio          int     `json:"max_ratio"`
	MaxSeedingTime    int     `json:"max_seeding_time"`
	Name              string  `json:"name"`
	NumComplete       int     `json:"num_complete"`
	NumIncomplete     int     `json:"num_incomplete"`
	NumLeechs         int     `json:"num_leechs"`
	NumSeeds          int     `json:"num_seeds"`
	Priority          int     `json:"priority"`
	Progress          float64 `json:"progress"`
	Ratio             float64 `json:"ratio"`
	RatioLimit        int     `json:"ratio_limit"`
	SavePath          string  `json:"save_path"`
	SeedingTime       int     `json:"seeding_time"`
	SeedingTimeLimit  int     `json:"seeding_time_limit"`
	SeenComplete      int     `json:"seen_complete"`
	SeqDl             bool    `json:"seq_dl"`
	Size              int64   `json:"size"`
	State             string  `json:"state"`
	SuperSeeding      bool    `json:"super_seeding"`
	Tags              string  `json:"tags"`
	TimeActive        int     `json:"time_active"`
	TotalSize         int64   `json:"total_size"`
	Tracker           string  `json:"tracker"`
	TrackersCount     int     `json:"trackers_count"`
	UpLimit           int     `json:"up_limit"`
	Uploaded          int     `json:"uploaded"`
	UploadedSession   int     `json:"uploaded_session"`
	Upspeed           int     `json:"upspeed"`
}

func (c *Client) GetTorrentList(ctx context.Context, params *GetTorrentListQuery) ([]TorrentInfo, error) {
	paramsMap := gconv.Map(params)
	if params != nil && len(params.Hashes) != 0 {
		hashes := strings.Join(params.Hashes, "|")
		paramsMap["hashes"] = hashes
	}

	res, err := c.Get(ctx, "/api/v2/torrents/info", paramsMap)
	if err != nil {
		return nil, err
	}

	var result []TorrentInfo
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
