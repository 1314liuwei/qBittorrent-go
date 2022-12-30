package qBittorent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Host     = "http://10.113.75.153:8111"
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
	assert.Contains(t, version, "v")
}
