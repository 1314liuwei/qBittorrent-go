package qBittorent

import (
	"context"
	"encoding/json"
)

type SyncMainData struct {
	Categories struct {
	} `json:"categories"`
	FullUpdate  bool `json:"full_update"`
	Rid         int  `json:"rid"`
	ServerState struct {
		AlltimeDl            int    `json:"alltime_dl"`
		AlltimeUl            int    `json:"alltime_ul"`
		AverageTimeQueue     int    `json:"average_time_queue"`
		ConnectionStatus     string `json:"connection_status"`
		DhtNodes             int    `json:"dht_nodes"`
		DlInfoData           int    `json:"dl_info_data"`
		DlInfoSpeed          int    `json:"dl_info_speed"`
		DlRateLimit          int    `json:"dl_rate_limit"`
		FreeSpaceOnDisk      int64  `json:"free_space_on_disk"`
		GlobalRatio          string `json:"global_ratio"`
		QueuedIoJobs         int    `json:"queued_io_jobs"`
		Queueing             bool   `json:"queueing"`
		ReadCacheHits        string `json:"read_cache_hits"`
		ReadCacheOverload    string `json:"read_cache_overload"`
		RefreshInterval      int    `json:"refresh_interval"`
		TotalBuffersSize     int    `json:"total_buffers_size"`
		TotalPeerConnections int    `json:"total_peer_connections"`
		TotalQueuedSize      int    `json:"total_queued_size"`
		TotalWastedSession   int    `json:"total_wasted_session"`
		UpInfoData           int    `json:"up_info_data"`
		UpInfoSpeed          int    `json:"up_info_speed"`
		UpRateLimit          int    `json:"up_rate_limit"`
		UseAltSpeedLimits    bool   `json:"use_alt_speed_limits"`
		WriteCacheOverload   string `json:"write_cache_overload"`
	} `json:"server_state"`
	Tags     []interface{} `json:"tags"`
	Torrents struct {
	} `json:"torrents"`
	Trackers struct {
	} `json:"trackers"`
}

type SyncTorrentPeersData struct {
	FullUpdate bool `json:"full_update"`
	Peers      map[string]struct {
		Client       string  `json:"client"`
		Connection   string  `json:"connection"`
		Country      string  `json:"country"`
		CountryCode  string  `json:"country_code"`
		DlSpeed      int     `json:"dl_speed"`
		Downloaded   int     `json:"downloaded"`
		Files        string  `json:"files"`
		Flags        string  `json:"flags"`
		FlagsDesc    string  `json:"flags_desc"`
		Ip           string  `json:"ip"`
		PeerIdClient string  `json:"peer_id_client"`
		Port         int     `json:"port"`
		Progress     float64 `json:"progress"`
		Relevance    float64 `json:"relevance"`
		UpSpeed      float64 `json:"up_speed"`
		Uploaded     float64 `json:"uploaded"`
	} `json:"peers"`
	Rid       int  `json:"rid"`
	ShowFlags bool `json:"show_flags"`
}

func (c *Client) GetMainData(ctx context.Context, rid int) (*SyncMainData, error) {
	res, err := c.Get(ctx, "/api/v2/sync/maindata", map[string]interface{}{"rid": rid})
	if err != nil {
		return nil, err
	}

	result := &SyncMainData{}
	err = json.Unmarshal(res.Body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetTorrentPeersData(ctx context.Context, hash string) (*SyncTorrentPeersData, error) {
	res, err := c.Get(ctx, "/api/v2/sync/torrentPeers", map[string]interface{}{"hash": hash})
	if err != nil {
		return nil, err
	}

	result := &SyncTorrentPeersData{}
	err = json.Unmarshal(res.Body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
