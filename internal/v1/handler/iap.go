package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type iap struct {
	iapTool    tool.IAP
	reqTool    tool.HttpRequest
	tokenCache map[string]string
}

func NewIAP(iapTool tool.IAP, reqTool tool.HttpRequest) IAP {
	return &iap{iapTool: iapTool, reqTool: reqTool}
}

func (i *iap) SandboxURL() string {
	return i.iapTool.SandboxURL()
}

func (i *iap) ProductURL() string {
	return i.iapTool.ProductURL()
}

func (i *iap) Password() string {
	return i.iapTool.Password()
}

func (i *iap) ParserIAPNotificationType(notificationType string, subtype string) global.SubscribeLogType {
	if notificationType == "SUBSCRIBED" && subtype == "INITIAL_BUY" {
		return global.InitialBuy
	} else if notificationType == "SUBSCRIBED" && subtype == "RESUBSCRIBE" {
		return global.Resubscribe
	} else if notificationType == "DID_RENEW" {
		return global.Renew
	} else if notificationType == "EXPIRED" {
		return global.Expired
	} else if notificationType == "DID_CHANGE_RENEWAL_PREF" && subtype == "UPGRADE" {
		return global.Upgrade
	} else if notificationType == "DID_CHANGE_RENEWAL_PREF" && subtype == "DOWNGRADE" {
		return global.Downgrade
	} else if notificationType == "DID_CHANGE_RENEWAL_PREF" && subtype == "" {
		return global.DowngradeCancel
	} else if notificationType == "REFUND" {
		return global.Refund
	} else if notificationType == "DID_CHANGE_RENEWAL_STATUS" && subtype == "AUTO_RENEW_ENABLED" {
		return global.RenewEnable
	} else if notificationType == "DID_CHANGE_RENEWAL_STATUS" && subtype == "AUTO_RENEW_DISABLED" {
		return global.RenewDisable
	}
	return global.Unknown
}

func (i *iap) ParserAppleReceipt(dict map[string]interface{}, receipt *dto.IAPVerifyReceiptResponse) error {
	if err := mapstructure.Decode(dict, &receipt); err != nil {
		return err
	}
	for _, item := range receipt.Receipt.InApp {
		i.parserAppleReceiptInfo(item)
	}
	for _, item := range receipt.LatestReceiptInfo {
		i.parserAppleReceiptInfo(item)
	}
	return nil
}

func (i *iap) GetAppleStoreAPIAccessToken() (string, error) {
	accessToken, err := i.iapTool.GenerateAppleStoreAPIToken(time.Hour)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (i *iap) VerifyAppleReceiptAPI(receiptData string) (*dto.IAPVerifyReceiptResponse, error) {
	//apple server 正式區驗證收據
	param := map[string]interface{}{
		"receipt-data":             receiptData,
		"password":                 i.Password(),
		"exclude-old-transactions": 1,
	}
	result, err := i.reqTool.SendRequest("POST", i.ProductURL(), nil, param)
	if err != nil {
		return nil, err
	}
	var response dto.IAPVerifyReceiptResponse
	if err := i.ParserAppleReceipt(result, &response); err != nil {
		return nil, err
	}
	//apple server 測試區驗證收據
	if response.Status == 21007 {
		result, err := i.reqTool.SendRequest("POST", i.SandboxURL(), nil, param)
		if err != nil {
			return nil, err
		}
		if err := i.ParserAppleReceipt(result, &response); err != nil {
			return nil, err
		}
	}
	return &response, nil
}

func (i *iap) GetSubscribeAPI(originalTransactionId string) (*dto.IAPSubscribeAPIResponse, error) {
	url := fmt.Sprintf("%s%s%s", i.iapTool.AppServerAPIURL(), "/inApps/v1/subscriptions/", originalTransactionId)
	token, err := i.iapTool.GenerateAppleStoreAPIToken(time.Hour)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var dict map[string]interface{}
	if err := json.Unmarshal(body, &dict); err != nil {
		return nil, err
	}
	return dto.NewIAPSubscribeAPIResponse(dict), nil
}

func (i *iap) GetHistoryAPI(originalTransactionId string) (*dto.IAPHistoryAPIResponse, error) {
	url := fmt.Sprintf("%s%s%s", i.iapTool.AppServerAPIURL(), "/inApps/v1/history/", originalTransactionId)
	token, err := i.iapTool.GenerateAppleStoreAPIToken(time.Hour)
	if err != nil {
		return nil, err
	}
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", token)
	dict, err := i.reqTool.SendRequest("GET", url, header, nil)
	if err != nil {
		return nil, err
	}
	//req, err := http.NewRequest("GET", url, nil)
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	//if err != nil {
	//	return nil, err
	//}
	//client := &http.Client{}
	//res, err := client.Do(req)
	//if err != nil {
	//	return nil, err
	//}
	//defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	return nil, err
	//}
	//var dict map[string]interface{}
	//if err := json.Unmarshal(body, &dict); err != nil {
	//	return nil, err
	//}
	return dto.NewIAPHistoryAPIResponse(dict), nil
}

func (i *iap) decodeBase64StringToMap(base64String string) (map[string]interface{}, error) {
	payloadString, err := base64.RawURLEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	var dict map[string]interface{}
	if err := json.Unmarshal(payloadString, &dict); err != nil {
		return nil, err
	}
	return dict, nil
}

func (i *iap) parserAppleReceiptInfo(receiptInfo *dto.ReceiptInfo) {
	originalPurchaseDate, err := i.parserAppleReceiptDate(receiptInfo.OriginalPurchaseDateMS)
	if err == nil {
		receiptInfo.OriginalPurchaseDate = originalPurchaseDate
	}
	purchaseDate, err := i.parserAppleReceiptDate(receiptInfo.PurchaseDateMS)
	if err == nil {
		receiptInfo.PurchaseDate = purchaseDate
	}
	expiresDate, err := i.parserAppleReceiptDate(receiptInfo.ExpiresDateMS)
	if err == nil {
		receiptInfo.ExpiresDate = expiresDate
	}
}

func (i *iap) parserAppleReceiptDate(unixMS string) (*time.Time, error) {
	msTime, err := strconv.ParseInt(unixMS, 10, 64)
	if err != nil {
		return nil, err
	}
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil, err
	}
	date, err := time.ParseInLocation("2006-01-02 15:04:05", time.Unix(msTime/1000, 0).Format("2006-01-02 15:04:05"), location)
	if err != nil {
		return nil, err
	}
	return &date, nil
}
