package apple_login

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/Henry19910227/fitness-go/config"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/apple_login"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type tool struct {
	setting apple_login.Setting
}

func New(setting apple_login.Setting) Tool {
	return &tool{setting: setting}
}

func (t *tool) GenerateClientSecret(duration time.Duration) (string, error) {
	p8bytes, err := ioutil.ReadFile(config.RootPath() + "/" + t.setting.GetKeyName())
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(p8bytes)
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("generate apple token error")
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	ecdsaPrivateKey, ok := parsedKey.(*ecdsa.PrivateKey)
	if !ok {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss":   t.setting.GetTeamID(),
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(duration).Unix(),
		"aud":   "https://appleid.apple.com",
		"sub":   t.setting.GetBundleID(),
	})
	token.Header["alg"] = "ES256"
	token.Header["kid"] = t.setting.GetKeyID()
	secret, err := token.SignedString(ecdsaPrivateKey)
	if err != nil {
		return "", err
	}
	return secret, nil
}

func (t *tool) GetUserID(authCode string, clientSecret string) (string, error) {
	param := map[string]interface{}{
		"client_id":     t.setting.GetBundleID(),
		"client_secret": clientSecret,
		"code":          authCode,
		"grant_type":    "authorization_code",
	}
	header := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	dict, err := util.SendRequest("POST", t.setting.GetDebugTokenURL(), header, nil, param)
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
	//以 IdToken 取得 jwtClaims
	idToken, ok := dict["id_token"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	jwtClaims, _ := jwt.Parse(idToken, nil)
	claims, ok := jwtClaims.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("id_token decode error")
	}
	//取得 UserID
	userID, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}
	if len(userID) == 0 {
		return "", errors.New("invalid token")
	}
	return userID, nil
}