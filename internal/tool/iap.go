package tool

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"strings"
	"time"
)

type iap struct {
	setting setting.IAP
}

func NewIAP(setting setting.IAP) IAP {
	return &iap{setting}
}

func (i *iap) SandboxURL() string {
	return i.setting.GetSandboxURL()
}

func (i *iap) ProductURL() string {
	return i.setting.GetProductURL()
}

func (i *iap) Password() string {
	return i.setting.GetPassword()
}

func (i *iap) ParserIAPNotificationResponse(base64String string) (*dto.IAPNotificationResponse, error) {
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