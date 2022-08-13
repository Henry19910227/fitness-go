package iab

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/config"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/iab"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/iab"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)

type tool struct {
	setting iab.Setting
}

func New(setting iab.Setting) Tool {
	return &tool{setting: setting}
}

func (t *tool) GenerateGoogleOAuth2Token(duration time.Duration) (string, error) {
	jsonData, err := ioutil.ReadFile(config.RootPath() + "/" + t.setting.GetKeyName())
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

func (t *tool) APIGetProducts(productID string, purchaseToken string, token string) (*model.IABProductAPIResponse, error) {
	url := t.setting.GetURL() + "/androidpublisher/v3/applications/" + t.setting.GetPackageName() + "/purchases/products/" + productID + "/tokens/" + purchaseToken
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", token)
	result, err := util.SendRequest("GET", url, header, nil, nil)
	if err != nil {
		return nil, err
	}
	return model.NewIABProductAPIResponse(result), nil
}

func (t *tool) APIGetSubscription(productID string, purchaseToken string, token string) (*model.IABSubscriptionAPIResponse, error) {
	url := t.setting.GetURL() + "/androidpublisher/v3/applications/" + t.setting.GetPackageName() + "/purchases/subscriptions/" + productID + "/tokens/" + purchaseToken
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", token)
	result, err := util.SendRequest("GET", url, header, nil, nil)
	if err != nil {
		return nil, err
	}
	if errResult, ok := result["error"].(map[string]interface{}); ok {
		msg, ok := errResult["message"].(string)
		if !ok {
			return nil, errors.New("api error")
		}
		return nil, errors.New(msg)
	}
	return model.NewIABSubscriptionAPIResponse(result), nil
}


