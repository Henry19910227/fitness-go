package line_login

import (
	"errors"
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
	param := map[string]interface{}{
		"access_token": authCode,
	}
	dict, err := util.SendRequest("GET", t.setting.GetVerifyTokenURL(), nil, nil, param)
	if err != nil {
		return "", err
	}
	errorValue, ok := dict["error"].(string)
	if ok {
		return "", errors.New(errorValue)
	}
	//驗證效期
	exp, ok := dict["expires_in"].(float64)
	if !ok {
		return "", errors.New("invalid token")
	}
	if exp <= 0 {
		return "", errors.New("token expired")
	}
	//取得 uid
	uid, ok := dict["client_id"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	if len(uid) == 0 {
		return "", errors.New("invalid token")
	}
	return uid, nil
}
