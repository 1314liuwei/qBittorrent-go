package qBittorent

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"net/http"
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

type TorrentGenericProperty struct {
	AdditionDate           int     `json:"addition_date"`
	Comment                string  `json:"comment"`
	CompletionDate         int     `json:"completion_date"`
	CreatedBy              string  `json:"created_by"`
	CreationDate           int     `json:"creation_date"`
	DlLimit                int     `json:"dl_limit"`
	DlSpeed                int     `json:"dl_speed"`
	DlSpeedAvg             int     `json:"dl_speed_avg"`
	Eta                    int     `json:"eta"`
	LastSeen               int     `json:"last_seen"`
	NbConnections          int     `json:"nb_connections"`
	NbConnectionsLimit     int     `json:"nb_connections_limit"`
	Peers                  int     `json:"peers"`
	PeersTotal             int     `json:"peers_total"`
	PieceSize              int     `json:"piece_size"`
	PiecesHave             int     `json:"pieces_have"`
	PiecesNum              int     `json:"pieces_num"`
	Reannounce             int     `json:"reannounce"`
	SavePath               string  `json:"save_path"`
	SeedingTime            int     `json:"seeding_time"`
	Seeds                  int     `json:"seeds"`
	SeedsTotal             int     `json:"seeds_total"`
	ShareRatio             float64 `json:"share_ratio"`
	TimeElapsed            int     `json:"time_elapsed"`
	TotalDownloaded        int     `json:"total_downloaded"`
	TotalDownloadedSession int     `json:"total_downloaded_session"`
	TotalSize              int     `json:"total_size"`
	TotalUploaded          int     `json:"total_uploaded"`
	TotalUploadedSession   int     `json:"total_uploaded_session"`
	TotalWasted            int     `json:"total_wasted"`
	UpLimit                int     `json:"up_limit"`
	UpSpeed                int     `json:"up_speed"`
	UpSpeedAvg             int     `json:"up_speed_avg"`
}

type Tracker struct {
	Msg           string `json:"msg"`
	NumDownloaded int    `json:"num_downloaded"`
	NumLeeches    int    `json:"num_leeches"`
	NumPeers      int    `json:"num_peers"`
	NumSeeds      int    `json:"num_seeds"`
	Status        int    `json:"status"`
	Tier          int    `json:"tier"`
	Url           string `json:"url"`
}

type TorrentContent struct {
	Index        int     `json:"index"`
	IsSeed       bool    `json:"is_seed"`
	Name         string  `json:"name"`
	PieceRange   []int   `json:"piece_range"`
	Priority     int     `json:"priority"`
	Progress     int     `json:"progress"`
	Size         int     `json:"size"`
	Availability float64 `json:"availability"`
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

func (c *Client) GetTorrentGenericProperties(ctx context.Context, hash string) (*TorrentGenericProperty, error) {
	if err := IsValidHash(hash); err != nil {
		return nil, err
	}

	res, err := c.Get(ctx, "/api/v2/torrents/properties", map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("torrent hash was not found")
	}

	result := new(TorrentGenericProperty)
	err = json.Unmarshal(res.Body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetTorrentTrackers(ctx context.Context, hash string) ([]Tracker, error) {
	if err := IsValidHash(hash); err != nil {
		return nil, err
	}

	res, err := c.Get(ctx, "/api/v2/torrents/trackers", map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("torrent hash was not found")
	}

	var result []Tracker
	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetTorrentWebSeeds(ctx context.Context, hash string) ([]string, error) {
	if err := IsValidHash(hash); err != nil {
		return nil, err
	}

	res, err := c.Get(ctx, "/api/v2/torrents/webseeds", map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("torrent hash was not found")
	}

	var (
		result []map[string]string
		urls   []string
	)

	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, err
	}

	for _, data := range result {
		if url, ok := data["url"]; ok {
			urls = append(urls, url)
		}
	}

	return urls, nil
}

func (c *Client) GetTorrentContents(ctx context.Context, hash string) ([]TorrentContent, error) {
	if err := IsValidHash(hash); err != nil {
		return nil, err
	}

	res, err := c.Get(ctx, "/api/v2/torrents/files", map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("torrent hash was not found")
	}

	var (
		result []TorrentContent
	)

	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetTorrentPiecesStates(ctx context.Context, hash string) ([]int, error) {
	if err := IsValidHash(hash); err != nil {
		return nil, err
	}

	res, err := c.Get(ctx, "/api/v2/torrents/pieceStates", map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("torrent hash was not found")
	}

	var (
		result []int
	)

	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetTorrentPiecesHashes(ctx context.Context, hash string) ([]string, error) {
	if err := IsValidHash(hash); err != nil {
		return nil, err
	}

	res, err := c.Get(ctx, "/api/v2/torrents/pieceHashes", map[string]interface{}{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("torrent hash was not found")
	}

	var (
		result []string
	)

	err = json.Unmarshal(res.Body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
