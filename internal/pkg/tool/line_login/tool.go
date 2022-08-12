package line_login

import (
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/line_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
)

type tool struct {
	setting line_login.Setting
}

func New(setting line_login.Setting) Tool {
	return &tool{setting: setting}
}

func (t *tool) GetUserID(authCode string) (string, error) {
	//驗證 access_token 合法性
	param := map[string]interface{}{
		"access_token": authCode,
	}
	resp, err := util.SendRequest("GET", t.setting.GetVerifyTokenURL(), nil, nil, param)
	if err != nil {
		return "", err
	}
	errorValue, ok := resp["error"].(string)
	if ok {
		return "", errors.New(errorValue)
	}
	//驗證 expires_in
	exp, ok := resp["expires_in"].(float64)
	if !ok {
		return "", errors.New("invalid token")
	}
	if exp <= 0 {
		return "", errors.New("token expired")
	}
	//驗證 client_id
	clientID, ok := resp["client_id"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	if clientID != t.setting.GetClientID() {
		return "", errors.New("invalid token")
	}
	//獲取profile
	header := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", authCode)}
	profile, err := util.SendRequest("GET", t.setting.GetProfileURL(), header, nil, nil)
	if err != nil {
		return "", err
	}
	userID, ok := profile["userId"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	if len(userID) == 0 {
		return "", errors.New("invalid token")
	}
	return userID, nil
}
