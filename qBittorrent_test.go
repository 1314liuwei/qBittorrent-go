package qBittorent

import (
	"context"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Host     = "http://127.0.0.1:8080"
	User     = "admin"
	Password = "adminadmin"
)

/*
Authentication functional testing
*/
func TestLogin(t *testing.T) {
	ctx := context.Background()
	client := NewClientWithConfig(Config{BasePath: Host})
	err := client.Login(ctx, User, Password)
	assert.NoError(t, err)
}

/*
Application functional testing
*/
func TestGetApplicationVersion(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	version, err := client.GetApplicationVersion(ctx)
	assert.NoError(t, err)
	assert.Regexp(t, regexp.MustCompile("v[0-9].[0-9].[0-9]"), version)
}

func TestGetAPIVersion(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	version, err := client.GetAPIVersion(ctx)
	assert.NoError(t, err)
	assert.Regexp(t, regexp.MustCompile("^[0-9].[0-9]"), version)
}

func TestGetBuildInfo(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	info, err := client.GetBuildInfo(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

func TestGetApplicationPreferences(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	info, err := client.GetApplicationPreferences(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

func TestGetDefaultSavePath(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	info, err := client.GetDefaultSavePath(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, info)
}

/*
Log functional testing
*/

func TestGetLog(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	info, err := client.GetLog(ctx, &QueryMainLogParam{
		Normal:   true,
		Info:     true,
		Warning:  true,
		Critical: true,
	})
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

func TestGetPeerLog(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	info, err := client.GetPeerLog(ctx, -1)
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

/*
Sync functional testing
*/

func TestGetMainData(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	info, err := client.GetMainData(ctx, 0)
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

func TestGetTorrentPeersData(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	info, err := client.GetTorrentPeersData(ctx, "2a99f2d2f7d25f01746e132a6ecd2ec6573b9a83")
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

/*
Transfer info functional testing
*/

func TestGetTransferInfo(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	info, err := client.GetGlobalTransferInfo(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

func TestGetAlternativeSpeedLimitsState(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	state, err := client.GetAlternativeSpeedLimitsState(ctx)
	assert.NoError(t, err)
	assert.Equal(t, state, false)
}

func TestToggleAlternativeSpeedLimits(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	err = client.ToggleAlternativeSpeedLimits(ctx)
	assert.NoError(t, err)
}

func TestGetGlobalDownloadLimit(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	limit, err := client.GetGlobalDownloadLimit(ctx)
	assert.NoError(t, err)
	assert.NotEqual(t, limit, -1)
}

func TestSetGlobalDownloadLimit(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	err = client.SetGlobalDownloadLimit(ctx, 1024)
	assert.NoError(t, err)
	limit, err := client.GetGlobalDownloadLimit(ctx)
	assert.NoError(t, err)
	assert.Equal(t, limit, 1024)
}

func TestGetGlobalUploadLimit(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	limit, err := client.GetGlobalUploadLimit(ctx)
	assert.NoError(t, err)
	assert.NotEqual(t, limit, -1)
}

func TestSetGlobalUploadLimit(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	err = client.SetGlobalUploadLimit(ctx, 1024)
	assert.NoError(t, err)
	limit, err := client.GetGlobalUploadLimit(ctx)
	assert.NoError(t, err)
	assert.Equal(t, limit, 1024)
}

func TestBanPeers(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	err = client.BanPeers(ctx, "1.1.1.1:8080", "2.2.2.2:80")
	assert.NoError(t, err)
}

/*
Torrent management functional testing
*/

func TestGetTorrentList(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	_, err = client.GetTorrentList(ctx, nil)
	assert.NoError(t, err)
}

func TestGetTorrentGenericProperties(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	info, err := client.GetTorrentGenericProperties(ctx, "2a99f2d2f7d25f01746e132a6ecd2ec6573b9a83")
	assert.NoError(t, err)
	assert.NotNil(t, info)
}

func TestGetTorrentTrackers(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	_, err = client.GetTorrentTrackers(ctx, "2a99f2d2f7d25f01746e132a6ecd2ec6573b9a83")
	assert.NoError(t, err)
}

func TestGetTorrentWebSeeds(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	_, err = client.GetTorrentWebSeeds(ctx, "2a99f2d2f7d25f01746e132a6ecd2ec6573b9a83")
	assert.NoError(t, err)
}

func TestGetTorrentContents(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	_, err = client.GetTorrentContents(ctx, "2a99f2d2f7d25f01746e132a6ecd2ec6573b9a83")
	assert.NoError(t, err)
}

func TestGetTorrentPiecesStates(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	_, err = client.GetTorrentPiecesStates(ctx, "2a99f2d2f7d25f01746e132a6ecd2ec6573b9a83")
	assert.NoError(t, err)
}

func TestGetTorrentPiecesHashes(t *testing.T) {
	ctx := context.Background()
	client, err := New(Host, User, Password)
	assert.NoError(t, err)
	_, err = client.GetTorrentPiecesHashes(ctx, "")
	assert.NoError(t, err)
}
