package handler

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/dto"
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

func (i *iab) GetProductsAPI(productID string, purchaseToken string) (*dto.IABProductAPIResponse, error) {
	accessToken, err := i.GetGooglePlayApiAccessToken()
	if err != nil {
		return nil, err
	}
	url := i.iabTool.URL() + "/androidpublisher/v3/applications/" + i.iabTool.PackageName() + "/purchases/products/" + productID + "/tokens/" + purchaseToken
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)
	result, err := i.reqTool.SendRequest("GET", url, header, nil)
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
	return dto.NewIABProductAPIResponse(result), nil
}

func (i *iab) GetSubscriptionAPI(productID string, purchaseToken string) (*dto.IABSubscriptionAPIResponse, error) {
	accessToken, err := i.GetGooglePlayApiAccessToken()
	if err != nil {
		return nil, err
	}
	url := i.iabTool.URL() + "/androidpublisher/v3/applications/" + i.iabTool.PackageName() + "/purchases/subscriptions/" + productID + "/tokens/" + purchaseToken
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)
	result, err := i.reqTool.SendRequest("GET", url, header, nil)
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
	return dto.NewIABSubscriptionAPIResponse(result), nil
}

func (i *iab) GetGooglePlayApiAccessToken() (string, error) {
	oauthToken, err := i.iabTool.GenerateGoogleOAuth2Token(time.Minute * 30)
	if err != nil {
		return "", err
	}
	param := map[string]interface{}{
		"grant_type": "urn:ietf:params:oauth:grant-type:jwt-bearer",
		"assertion":  oauthToken,
	}
	result, err := i.reqTool.SendRequest("POST", "https://oauth2.googleapis.com/token", nil, param)
	if err != nil {
		return "", err
	}
	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", errors.New("無法取得 access token")
	}
	return accessToken, nil
}

func (i *iab) DecodeIAPNotificationResponse(base64String string) (map[string]interface{}, error) {
	payloadDict, err := iabDecodeBase64StringToMap(base64String)
	if err != nil {
		return nil, err
	}
	return payloadDict, err
}

func iabDecodeBase64StringToMap(base64String string) (map[string]interface{}, error) {
	payloadString, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	var dict map[string]interface{}
	if err := json.Unmarshal(payloadString, &dict); err != nil {
		return nil, err
	}
	return dict, nil
}
