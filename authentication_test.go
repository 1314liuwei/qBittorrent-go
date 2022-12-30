package qBittorent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()
	client := NewClientWithConfig(Config{BasePath: "http://10.113.75.153:8111"})
	err := client.Login(ctx, "admin", "adminadmin")
	assert.NoError(t, err)
}
