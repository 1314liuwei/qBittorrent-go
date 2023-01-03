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

type Preferences struct {
	AddTrackers                        string   `json:"add_trackers"`
	AddTrackersEnabled                 bool     `json:"add_trackers_enabled"`
	AltDlLimit                         int      `json:"alt_dl_limit"`
	AltUpLimit                         int      `json:"alt_up_limit"`
	AlternativeWebuiEnabled            bool     `json:"alternative_webui_enabled"`
	AlternativeWebuiPath               string   `json:"alternative_webui_path"`
	AnnounceIp                         string   `json:"announce_ip"`
	AnnounceToAllTiers                 bool     `json:"announce_to_all_tiers"`
	AnnounceToAllTrackers              bool     `json:"announce_to_all_trackers"`
	AnonymousMode                      bool     `json:"anonymous_mode"`
	AsyncIoThreads                     int      `json:"async_io_threads"`
	AutoDeleteMode                     int      `json:"auto_delete_mode"`
	AutoTmmEnabled                     bool     `json:"auto_tmm_enabled"`
	AutorunEnabled                     bool     `json:"autorun_enabled"`
	AutorunProgram                     string   `json:"autorun_program"`
	BannedIPs                          string   `json:"banned_IPs"`
	BittorrentProtocol                 int      `json:"bittorrent_protocol"`
	BlockPeersOnPrivilegedPorts        bool     `json:"block_peers_on_privileged_ports"`
	BypassAuthSubnetWhitelist          string   `json:"bypass_auth_subnet_whitelist"`
	BypassAuthSubnetWhitelistEnabled   bool     `json:"bypass_auth_subnet_whitelist_enabled"`
	BypassLocalAuth                    bool     `json:"bypass_local_auth"`
	CategoryChangedTmmEnabled          bool     `json:"category_changed_tmm_enabled"`
	CheckingMemoryUse                  int      `json:"checking_memory_use"`
	CreateSubfolderEnabled             bool     `json:"create_subfolder_enabled"`
	CurrentInterfaceAddress            string   `json:"current_interface_address"`
	CurrentNetworkInterface            string   `json:"current_network_interface"`
	Dht                                bool     `json:"dht"`
	DiskCache                          int      `json:"disk_cache"`
	DiskCacheTtl                       int      `json:"disk_cache_ttl"`
	DlLimit                            int      `json:"dl_limit"`
	DontCountSlowTorrents              bool     `json:"dont_count_slow_torrents"`
	DyndnsDomain                       string   `json:"dyndns_domain"`
	DyndnsEnabled                      bool     `json:"dyndns_enabled"`
	DyndnsPassword                     string   `json:"dyndns_password"`
	DyndnsService                      int      `json:"dyndns_service"`
	DyndnsUsername                     string   `json:"dyndns_username"`
	EmbeddedTrackerPort                int      `json:"embedded_tracker_port"`
	EnableCoalesceReadWrite            bool     `json:"enable_coalesce_read_write"`
	EnableEmbeddedTracker              bool     `json:"enable_embedded_tracker"`
	EnableMultiConnectionsFromSameIp   bool     `json:"enable_multi_connections_from_same_ip"`
	EnableOsCache                      bool     `json:"enable_os_cache"`
	EnablePieceExtentAffinity          bool     `json:"enable_piece_extent_affinity"`
	EnableUploadSuggestions            bool     `json:"enable_upload_suggestions"`
	Encryption                         int      `json:"encryption"`
	ExportDir                          string   `json:"export_dir"`
	ExportDirFin                       string   `json:"export_dir_fin"`
	FilePoolSize                       int      `json:"file_pool_size"`
	HashingThreads                     int      `json:"hashing_threads"`
	IncompleteFilesExt                 bool     `json:"incomplete_files_ext"`
	IpFilterEnabled                    bool     `json:"ip_filter_enabled"`
	IpFilterPath                       string   `json:"ip_filter_path"`
	IpFilterTrackers                   bool     `json:"ip_filter_trackers"`
	LimitLanPeers                      bool     `json:"limit_lan_peers"`
	LimitTcpOverhead                   bool     `json:"limit_tcp_overhead"`
	LimitUtpRate                       bool     `json:"limit_utp_rate"`
	ListenPort                         int      `json:"listen_port"`
	Locale                             string   `json:"locale"`
	Lsd                                bool     `json:"lsd"`
	MailNotificationAuthEnabled        bool     `json:"mail_notification_auth_enabled"`
	MailNotificationEmail              string   `json:"mail_notification_email"`
	MailNotificationEnabled            bool     `json:"mail_notification_enabled"`
	MailNotificationPassword           string   `json:"mail_notification_password"`
	MailNotificationSender             string   `json:"mail_notification_sender"`
	MailNotificationSmtp               string   `json:"mail_notification_smtp"`
	MailNotificationSslEnabled         bool     `json:"mail_notification_ssl_enabled"`
	MailNotificationUsername           string   `json:"mail_notification_username"`
	MaxActiveDownloads                 int      `json:"max_active_downloads"`
	MaxActiveTorrents                  int      `json:"max_active_torrents"`
	MaxActiveUploads                   int      `json:"max_active_uploads"`
	MaxConcurrentHttpAnnounces         int      `json:"max_concurrent_http_announces"`
	MaxConnec                          int      `json:"max_connec"`
	MaxConnecPerTorrent                int      `json:"max_connec_per_torrent"`
	MaxRatio                           int      `json:"max_ratio"`
	MaxRatioAct                        int      `json:"max_ratio_act"`
	MaxRatioEnabled                    bool     `json:"max_ratio_enabled"`
	MaxSeedingTime                     int      `json:"max_seeding_time"`
	MaxSeedingTimeEnabled              bool     `json:"max_seeding_time_enabled"`
	MaxUploads                         int      `json:"max_uploads"`
	MaxUploadsPerTorrent               int      `json:"max_uploads_per_torrent"`
	OutgoingPortsMax                   int      `json:"outgoing_ports_max"`
	OutgoingPortsMin                   int      `json:"outgoing_ports_min"`
	PeerTurnover                       int      `json:"peer_turnover"`
	PeerTurnoverCutoff                 int      `json:"peer_turnover_cutoff"`
	PeerTurnoverInterval               int      `json:"peer_turnover_interval"`
	Pex                                bool     `json:"pex"`
	PreallocateAll                     bool     `json:"preallocate_all"`
	ProxyAuthEnabled                   bool     `json:"proxy_auth_enabled"`
	ProxyIp                            string   `json:"proxy_ip"`
	ProxyPassword                      string   `json:"proxy_password"`
	ProxyPeerConnections               bool     `json:"proxy_peer_connections"`
	ProxyPort                          int      `json:"proxy_port"`
	ProxyTorrentsOnly                  bool     `json:"proxy_torrents_only"`
	ProxyType                          int      `json:"proxy_type"`
	ProxyUsername                      string   `json:"proxy_username"`
	QueueingEnabled                    bool     `json:"queueing_enabled"`
	RandomPort                         bool     `json:"random_port"`
	RecheckCompletedTorrents           bool     `json:"recheck_completed_torrents"`
	ResolvePeerCountries               bool     `json:"resolve_peer_countries"`
	RssAutoDownloadingEnabled          bool     `json:"rss_auto_downloading_enabled"`
	RssDownloadRepackProperEpisodes    bool     `json:"rss_download_repack_proper_episodes"`
	RssMaxArticlesPerFeed              int      `json:"rss_max_articles_per_feed"`
	RssProcessingEnabled               bool     `json:"rss_processing_enabled"`
	RssRefreshInterval                 int      `json:"rss_refresh_interval"`
	RssSmartEpisodeFilters             string   `json:"rss_smart_episode_filters"`
	SavePath                           string   `json:"save_path"`
	SavePathChangedTmmEnabled          bool     `json:"save_path_changed_tmm_enabled"`
	SaveResumeDataInterval             int      `json:"save_resume_data_interval"`
	ScanDirs                           struct{} `json:"scan_dirs"`
	ScheduleFromHour                   int      `json:"schedule_from_hour"`
	ScheduleFromMin                    int      `json:"schedule_from_min"`
	ScheduleToHour                     int      `json:"schedule_to_hour"`
	ScheduleToMin                      int      `json:"schedule_to_min"`
	SchedulerDays                      int      `json:"scheduler_days"`
	SchedulerEnabled                   bool     `json:"scheduler_enabled"`
	SendBufferLowWatermark             int      `json:"send_buffer_low_watermark"`
	SendBufferWatermark                int      `json:"send_buffer_watermark"`
	SendBufferWatermarkFactor          int      `json:"send_buffer_watermark_factor"`
	SlowTorrentDlRateThreshold         int      `json:"slow_torrent_dl_rate_threshold"`
	SlowTorrentInactiveTimer           int      `json:"slow_torrent_inactive_timer"`
	SlowTorrentUlRateThreshold         int      `json:"slow_torrent_ul_rate_threshold"`
	SocketBacklogSize                  int      `json:"socket_backlog_size"`
	StartPausedEnabled                 bool     `json:"start_paused_enabled"`
	StopTrackerTimeout                 int      `json:"stop_tracker_timeout"`
	TempPath                           string   `json:"temp_path"`
	TempPathEnabled                    bool     `json:"temp_path_enabled"`
	TorrentChangedTmmEnabled           bool     `json:"torrent_changed_tmm_enabled"`
	UpLimit                            int      `json:"up_limit"`
	UploadChokingAlgorithm             int      `json:"upload_choking_algorithm"`
	UploadSlotsBehavior                int      `json:"upload_slots_behavior"`
	Upnp                               bool     `json:"upnp"`
	UpnpLeaseDuration                  int      `json:"upnp_lease_duration"`
	UseHttps                           bool     `json:"use_https"`
	UtpTcpMixedMode                    int      `json:"utp_tcp_mixed_mode"`
	ValidateHttpsTrackerCertificate    bool     `json:"validate_https_tracker_certificate"`
	WebUiAddress                       string   `json:"web_ui_address"`
	WebUiBanDuration                   int      `json:"web_ui_ban_duration"`
	WebUiClickjackingProtectionEnabled bool     `json:"web_ui_clickjacking_protection_enabled"`
	WebUiCsrfProtectionEnabled         bool     `json:"web_ui_csrf_protection_enabled"`
	WebUiCustomHttpHeaders             string   `json:"web_ui_custom_http_headers"`
	WebUiDomainList                    string   `json:"web_ui_domain_list"`
	WebUiHostHeaderValidationEnabled   bool     `json:"web_ui_host_header_validation_enabled"`
	WebUiHttpsCertPath                 string   `json:"web_ui_https_cert_path"`
	WebUiHttpsKeyPath                  string   `json:"web_ui_https_key_path"`
	WebUiMaxAuthFailCount              int      `json:"web_ui_max_auth_fail_count"`
	WebUiPort                          int      `json:"web_ui_port"`
	WebUiSecureCookieEnabled           bool     `json:"web_ui_secure_cookie_enabled"`
	WebUiSessionTimeout                int      `json:"web_ui_session_timeout"`
	WebUiUpnp                          bool     `json:"web_ui_upnp"`
	WebUiUseCustomHttpHeadersEnabled   bool     `json:"web_ui_use_custom_http_headers_enabled"`
	WebUiUsername                      string   `json:"web_ui_username"`
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

func (c *Client) GetApplicationPreferences(ctx context.Context) (*Preferences, error) {
	res, err := c.Get(ctx, "/api/v2/app/preferences", nil)
	if err != nil {
		return nil, err
	}

	result := &Preferences{}
	err = json.Unmarshal(res.Body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// TODO: SetApplicationPreferences

func (c *Client) GetDefaultSavePath(ctx context.Context) (string, error) {
	res, err := c.Get(ctx, "/api/v2/app/defaultSavePath", nil)
	if err != nil {
		return "", err
	}

	return string(res.Body), nil
}
