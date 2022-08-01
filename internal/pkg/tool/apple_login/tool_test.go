package apple_login

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTool_GenerateClientSecret(t *testing.T) {
	tool := NewTool()
	secret, err := tool.GenerateClientSecret(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(secret) > 0)
	fmt.Println(secret)
}

func TestTool_GetUserIDByAccessToken(t *testing.T) {
	tool := NewTool()
	secret, err := tool.GenerateClientSecret(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(secret) > 0)
	userID, err := tool.GetUserID("ca5904cbb7ecb4837983acaa6a07b5544.0.sx.ChIulRheNw1KKXvxdDaE0Q", secret)
	assert.NoError(t, err)
	assert.Equal(t, "000007.2b92c7d2952e4b3398cf42061249716e.1728", userID)
}