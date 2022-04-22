package handler

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"time"
)

type iab struct {
	iabTool tool.IAB
	reqTool tool.HttpRequest
}

func NewIAB(iabTool tool.IAB, reqTool tool.HttpRequest) IAB {
	return &iab{iabTool: iabTool, reqTool: reqTool}
}

func (i *iab) GetGooglePlayApiAccessToken() (string, error) {
	oauthToken, err := i.iabTool.GenerateGoogleOAuth2Token(time.Minute * 3)
	if err != nil {
		return "", err
	}
	param := map[string]interface{}{
		"grant_type": "urn:ietf:params:oauth:grant-type:jwt-bearer",
		"assertion":  oauthToken,
	}
	result, err := i.reqTool.SendPostRequestWithJsonBody("https://oauth2.googleapis.com/token", param)
	if err != nil {
		return "", err
	}
	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", errors.New("無法取得 access token")
	}
	return accessToken, nil
}
