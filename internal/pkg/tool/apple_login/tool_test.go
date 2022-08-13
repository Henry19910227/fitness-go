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

func TestTool_APIGetRefreshToken(t *testing.T) {
	tool := NewTool()
	secret, err := tool.GenerateClientSecret(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(secret) > 0)
	refreshToken, err := tool.APIGetRefreshToken("r7fe363c793454c18b56aa0021a00afea.0.sx.o6asOR60ptLYM83e4VcdHA", secret)
	assert.NoError(t, err)
	assert.Equal(t, true, len(refreshToken) > 0)
	fmt.Println(refreshToken)
}

func TestTool_APIGetUserID(t *testing.T) {
	tool := NewTool()
	secret, err := tool.GenerateClientSecret(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(secret) > 0)
	userID, err := tool.APIGetUserID("r85253335e4ea43d28f569f0033ff5615.0.sx.XRei186TVnFH6w0v2bo33Q", secret)
	assert.NoError(t, err)
	assert.Equal(t, "000007.2b92c7d2952e4b3398cf42061249716e.1728", userID)
}