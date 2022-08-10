package iap

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/Henry19910227/fitness-go/config"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/iap"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/iap"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"strconv"
	"time"
)

type tool struct {
	setting iap.Setting
}

func New(setting iap.Setting) Tool {
	return &tool{setting: setting}
}

func (t *tool) GenerateAppleStoreAPIToken(duration time.Duration) (string, error) {
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
		"iss":   t.setting.GetIssuer(),
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(duration).Unix(),
		"aud":   "appstoreconnect-v1",
		"nonce": uuid.New().String(),
		"bid":   t.setting.GetBundleID(),
	})
	token.Header["kid"] = t.setting.GetKeyID()
	tokenString, err := token.SignedString(ecdsaPrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *tool) VerifyAppleReceiptAPI(receiptData string) (*model.IAPVerifyReceiptResponse, error) {
	//apple server 正式區驗證收據
	body := map[string]interface{}{
		"receipt-data":             receiptData,
		"password":                 t.setting.GetPassword(),
		"exclude-old-transactions": 1,
	}
	header := map[string]string{"Content-Type": "application/json"}
	result, err := util.SendRequest("POST", t.setting.GetProductURL(), header, body, nil)
	if err != nil {
		return nil, err
	}
	var response model.IAPVerifyReceiptResponse
	if err := t.ParserAppleReceipt(result, &response); err != nil {
		return nil, err
	}
	//apple server 測試區驗證收據
	if response.Status == 21007 {
		result, err := util.SendRequest("POST", t.setting.GetSandboxURL(), header, body, nil)
		if err != nil {
			return nil, err
		}
		if err := t.ParserAppleReceipt(result, &response); err != nil {
			return nil, err
		}
	}
	return &response, nil
}

func (t *tool) GetSubscribeAPI(originalTransactionId string, token string) (*model.IAPSubscribeAPIResponse, error) {
	url := fmt.Sprintf("%s%s%s", t.setting.GetAppServerAPIURL(), "/inApps/v1/subscriptions/", originalTransactionId)
	header := map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)}
	dict, err := util.SendRequest("GET", url, header, nil, nil)
	if err != nil {
		return nil, err
	}
	return model.NewIAPSubscribeAPIResponse(dict), nil
}

func (t *tool) ParserAppleReceipt(dict map[string]interface{}, receipt *model.IAPVerifyReceiptResponse) error {
	if err := mapstructure.Decode(dict, &receipt); err != nil {
		return err
	}
	for _, item := range receipt.Receipt.InApp {
		t.parserAppleReceiptInfo(item)
	}
	for _, item := range receipt.LatestReceiptInfo {
		t.parserAppleReceiptInfo(item)
	}
	return nil
}

func (t *tool) parserAppleReceiptInfo(receiptInfo *model.ReceiptInfo) {
	originalPurchaseDate, err := t.parserAppleReceiptDate(receiptInfo.OriginalPurchaseDateMS)
	if err == nil {
		receiptInfo.OriginalPurchaseDate = originalPurchaseDate
	}
	purchaseDate, err := t.parserAppleReceiptDate(receiptInfo.PurchaseDateMS)
	if err == nil {
		receiptInfo.PurchaseDate = purchaseDate
	}
	expiresDate, err := t.parserAppleReceiptDate(receiptInfo.ExpiresDateMS)
	if err == nil {
		receiptInfo.ExpiresDate = expiresDate
	}
}

func (t *tool) parserAppleReceiptDate(unixMS string) (*time.Time, error) {
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
