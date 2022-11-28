package fcm

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/fcm"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type tool struct {
	setting fcm.Setting
}

func New(setting fcm.Setting) Tool {
	return &tool{setting: setting}
}

func (t *tool) Key() string {
	return "fitness.fcm.token"
}

func (t *tool) GetExpire() time.Duration {
	return t.setting.GetExpire()
}

func (t *tool) GenerateGoogleOAuth2Token(duration time.Duration) (string, error) {
	jsonData, err := ioutil.ReadFile(util.RootPath() + "/config/" + t.setting.GetKeyName())
	if err != nil {
		return "", err
	}
	var dict map[string]string
	if err := json.Unmarshal(jsonData, &dict); err != nil {
		return "", err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(dict["private_key"]))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss":   dict["client_email"],
		"sub":   dict["client_email"],
		"aud":   dict["token_uri"],
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(duration).Unix(),
		"scope": t.setting.GetScope(),
	})
	token.Header["alg"] = "RS256"
	token.Header["typ"] = "JWT"
	token.Header["kid"] = dict["private_key_id"]
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *tool) APIGetGooglePlayToken(oauthToken string) (string, error) {
	body := map[string]interface{}{
		"grant_type": "urn:ietf:params:oauth:grant-type:jwt-bearer",
		"assertion":  oauthToken,
	}
	result, err := util.SendRequest("POST", t.setting.GetTokenURL(), nil, body, nil)
	if err != nil {
		return "", err
	}
	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", errors.New("無法取得 access token")
	}
	return accessToken, nil
}

func (t *tool) APISendMessage(token string, message map[string]interface{}) error {
	url := t.setting.GetURL() + "/v1/projects/" + t.setting.GetProjectID() + "/messages:send"
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", token)
	_, err := util.SendRequest("POST", url, header, message, nil)
	return err
}
