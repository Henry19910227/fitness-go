package fcm

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	fcmModel "github.com/Henry19910227/fitness-go/internal/v2/model/fcm"
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

func TestTool_APISendMessage(t *testing.T) {
	tool := NewTool()
	//產出 auth token
	oauthToken, err := tool.GenerateGoogleOAuth2Token(time.Hour)
	assert.NoError(t, err)
	assert.Equal(t, true, len(oauthToken) > 0)
	//獲取API Token
	token, err := tool.APIGetGooglePlayToken(oauthToken)
	assert.NoError(t, err)
	assert.Equal(t, true, len(token) > 0)
	//parser model
	deviceToken := "dgVouNzFbUydv8HSF85bDC:APA91bH8AwOU5C2iiSiHwkUMmgUIRSc87Xx2BEngNvuanR1c0BdQDqVGXxCpggEKN7WRHaH_8_inyGkcrVADSNLBrAGxkPbhw_lmkfOoUt_sMNMQ4hmmFi8-b4OJTxhfYUO14fZiKdqV"
	model := fcmModel.Input{}
	model.Message = fcmModel.Message{
		Token: deviceToken,
		Notification: fcmModel.Notification{
			Title: "Hello World",
			Body:  "說你好!~~~~",
		},
	}
	message := make(map[string]interface{})
	err = util.Parser(model, &message)
	assert.NoError(t, err)
	//發送推播
	err = tool.APISendMessage(token, message)
	assert.NoError(t, err)
}
