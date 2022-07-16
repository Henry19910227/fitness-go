package fb_login

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/fb_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
)

type tool struct {
	setting fb_login.Setting
}

func New(setting fb_login.Setting) Tool {
	return &tool{setting: setting}
}

func (t *tool) GetFbUidByAccessToken(accessToken string) (string, error) {
	param := map[string]interface{}{
		"input_token":  accessToken,
		"access_token": t.setting.GetAppID() + "|" + t.setting.GetAppSecret(),
	}
	dict, err := util.SendRequest("GET", t.setting.GetDebugTokenURL(), nil,nil, param)
	if err != nil {
		return "", err
	}
	data, ok := dict["data"].(map[string]interface{})
	if !ok {
		return "", errors.New("invalid token")
	}
	isValid, ok := data["is_valid"].(bool)
	if !ok {
		return "", errors.New("invalid token")
	}
	if !isValid {
		return "", errors.New("invalid token")
	}
	fbUID, ok := data["user_id"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	return fbUID, nil
}
