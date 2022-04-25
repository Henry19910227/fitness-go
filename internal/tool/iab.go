package tool

import (
	"encoding/json"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type iab struct {
	setting setting.IAB
}

func NewIAB(setting setting.IAB) IAB {
	return &iab{setting}
}

func (i *iab) TokenURI() string {
	jsonData, err := ioutil.ReadFile(i.setting.GetJsonFilePath())
	if err != nil {
		return ""
	}
	var dict map[string]string
	if err := json.Unmarshal(jsonData, &dict); err != nil {
		return ""
	}
	uri, ok := dict["token_uri"]
	if !ok {
		return ""
	}
	return uri
}

func (i *iab) URL() string {
	return i.setting.GetURL()
}

func (i *iab) Scope() string {
	return i.setting.GetScope()
}

func (i *iab) PackageName() string {
	return i.setting.GetPackageName()
}

func (i *iab) GenerateGoogleOAuth2Token(duration time.Duration) (string, error) {
	jsonData, err := ioutil.ReadFile(i.setting.GetJsonFilePath())
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
		"scope": i.setting.GetScope(),
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
