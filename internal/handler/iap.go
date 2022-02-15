package handler

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type iap struct {
	iapTool tool.IAP
	tokenCache map[string]string
}

func NewIAP(iapTool tool.IAP) IAP {
	return &iap{iapTool: iapTool}
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

func (i *iap) DecodeIAPNotificationResponse(base64String string) (*dto.IAPNotificationResponse, error) {
	payloadDict, err := decodeBase64StringToMap(strings.Split(base64String, ".")[1])
	if err != nil {
		return nil, err
	}
	payloadDataDict, ok := payloadDict["data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	base64SignedRenewalInfoString, ok := payloadDataDict["signedRenewalInfo"].(string)
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	base64SignedTransactionInfoString, ok := payloadDataDict["signedTransactionInfo"].(string)
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	renewalInfoDict, err := decodeBase64StringToMap(strings.Split(base64SignedRenewalInfoString, ".")[1])
	if err != nil {
		return nil, err
	}
	transactionInfo, err := decodeBase64StringToMap(strings.Split(base64SignedTransactionInfoString, ".")[1])
	if err != nil {
		return nil, err
	}
	// parser dto
	var response dto.IAPNotificationResponse
	if err := mapstructure.Decode(payloadDict, &response); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(payloadDataDict, &response.Data); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(renewalInfoDict, &response.Data.SignedRenewalInfo); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(transactionInfo, &response.Data.SignedTransactionInfo); err != nil {
		return nil, err
	}
	return &response, nil
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

func (i *iap) ParserAppleReceipt(dict map[string]interface{}, receipt *dto.AppleReceiptResponse) error {
	if err := mapstructure.Decode(dict, &receipt); err != nil {
		return err
	}
	for _, item := range receipt.Receipt.InApp {
		parserAppleReceiptInfo(item)
	}
	for _, item := range receipt.LatestReceiptInfo {
		parserAppleReceiptInfo(item)
	}
	return nil
}

func (i *iap) GetSubscriptionAPI(originalTransactionId string) (*dto.IAPSubscribeResponse, error) {
	url := fmt.Sprintf("%s%s%s",i.iapTool.AppServerAPIURL(), "/inApps/v1/subscriptions/", originalTransactionId)
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
	response, err := parserSubscriptionAPIRequest(dict)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func decodeBase64StringToMap(base64String string) (map[string]interface{}, error) {
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

func parserAppleReceiptInfo(receiptInfo *dto.ReceiptInfo) {
	originalPurchaseDate, err := parserAppleReceiptDate(receiptInfo.OriginalPurchaseDateMS)
	if err == nil {
		receiptInfo.OriginalPurchaseDate = originalPurchaseDate
	}
	purchaseDate, err := parserAppleReceiptDate(receiptInfo.PurchaseDateMS)
	if err == nil {
		receiptInfo.PurchaseDate = purchaseDate
	}
	expiresDate, err := parserAppleReceiptDate(receiptInfo.ExpiresDateMS)
	if err == nil {
		receiptInfo.ExpiresDate = expiresDate
	}
}

func parserSubscriptionAPIRequest(data map[string]interface{}) (*dto.IAPSubscribeResponse, error) {
	dataArray, ok := data["data"].([]interface{})
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	if len(dataArray) == 0 {
		return nil, errors.New("格式錯誤")
	}
	dataItem, ok := dataArray[0].(map[string]interface{})
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	lastTransactions, ok := dataItem["lastTransactions"].([]interface{})
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	if len(lastTransactions) == 0 {
		return nil, errors.New("格式錯誤")
	}
	transaction, ok := lastTransactions[0].(map[string]interface{})
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	status, ok := transaction["status"].(float64)
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	originalTransactionId, ok := transaction["originalTransactionId"].(string)
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	base64SignedRenewalInfoString, ok := transaction["signedRenewalInfo"].(string)
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	base64SignedTransactionInfoString, ok := transaction["signedTransactionInfo"].(string)
	if !ok {
		return nil, errors.New("格式錯誤")
	}
	var response dto.IAPSubscribeResponse
	response.Status = int(status)
	response.OriginalTransactionID = originalTransactionId
	renewalInfoDict, err := decodeBase64StringToMap(strings.Split(base64SignedRenewalInfoString, ".")[1])
	if err != nil {
		return nil, err
	}
	transactionInfo, err := decodeBase64StringToMap(strings.Split(base64SignedTransactionInfoString, ".")[1])
	if err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(renewalInfoDict, &response.SignedRenewalInfo); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(transactionInfo, &response.SignedTransactionInfo); err != nil {
		return nil, err
	}
	return &response, nil
}

func parserAppleReceiptDate(unixMS string) (*time.Time, error) {
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