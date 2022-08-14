package google_login

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/google_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"strconv"
	"time"
)

type tool struct {
	setting google_login.Setting
}

func New(setting google_login.Setting) Tool {
	return &tool{setting: setting}
}

func (t *tool) GetUserID(authCode string) (string, error) {
	param := map[string]interface{}{
		"id_token": authCode,
	}
	dict, err := util.SendRequest("GET", t.setting.GetDebugTokenURL(), nil, nil, param)
	if err != nil {
		return "", err
	}
	errorValue, ok := dict["error"].(string)
	if ok {
		return "", errors.New(errorValue)
	}
	//驗證 iss
	iss, ok := dict["iss"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	if iss != t.setting.GetIss() {
		return "", errors.New("iss error")
	}
	//驗證 aud
	aud, ok := dict["aud"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	if aud != t.setting.GetIOSClientID() && aud != t.setting.GetAndroidClientID() {
		return "", errors.New("aud error")
	}
	//驗證效期
	exp, ok := dict["exp"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	expTimestamp, err := strconv.ParseInt(exp, 10, 64)
	if err != nil {
		return "", errors.New("exp error")
	}
	if time.Now().Unix() >= expTimestamp {
		return "", errors.New("token expired")
	}
	//取得 uid
	uid, ok := dict["sub"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	if len(uid) == 0 {
		return "", errors.New("invalid token")
	}
	return uid, nil
}
