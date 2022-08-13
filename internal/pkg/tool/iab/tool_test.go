package iab

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTool_GenerateGoogleOAuth2Token(t *testing.T) {
	tool := NewTool()
	token, err := tool.GenerateGoogleOAuth2Token(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(token) > 0)
}

func TestTool_APIGetGooglePlayToken(t *testing.T) {
	tool := NewTool()
	//產出 auth token
	oauthToken, err := tool.GenerateGoogleOAuth2Token(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(oauthToken) > 0)
	//獲取API Token
	token, err := tool.APIGetGooglePlayToken(oauthToken)
	assert.NoError(t, err)
	assert.Equal(t, true, len(token) > 0)
}

func TestTool_APIGetProducts(t *testing.T) {
	tool := NewTool()
	//產出 auth token
	oauthToken, err := tool.GenerateGoogleOAuth2Token(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(oauthToken) > 0)
	//獲取API Token
	token, err := tool.APIGetGooglePlayToken(oauthToken)
	assert.NoError(t, err)
	assert.Equal(t, true, len(token) > 0)
	//獲取商品資訊
	response, err := tool.APIGetProducts("com.fitness.copper_course", "fpcklbdlpjaoophlldccigeb.AO-J1Ox8_8AP9u656lhy2MxbUEGBwkxVU4vIzJPaEmPtBI0O4iGo7Yo_RGVsMtsJcvwK_ZkTFGa8YdWncp7yN-W0PxIqIMeKaa3GIzP9gfd9QwCzDNUFUhY", token)
	assert.NoError(t, err)
	assert.Equal(t, true, response.OrderId == "GPA.3335-7558-9650-14006")
}

func TestTool_APIGetSubscription(t *testing.T) {
	tool := NewTool()
	//產出 auth token
	oauthToken, err := tool.GenerateGoogleOAuth2Token(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(oauthToken) > 0)
	//獲取API Token
	token, err := tool.APIGetGooglePlayToken(oauthToken)
	assert.NoError(t, err)
	assert.Equal(t, true, len(token) > 0)
	//獲取商品資訊
	_, err = tool.APIGetSubscription("com.fitness.gold_member_month", "focfcjeioimcfhljabiobjbb.AO-J1OzNW4lNCqa4AevJvs_WH3b2L3P8DXXoYJ103vgPwLpbZAK70uMHGo6tMV4ZdD9LSiRWN_eLhAAs96I87hK2zeQnykPx7WQiU21xseFlKikKeZheSzI", token)
	assert.NoError(t, err)
}