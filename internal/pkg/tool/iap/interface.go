package iap

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/iap"
	"time"
)

type Tool interface {
	GenerateAppleStoreAPIToken(duration time.Duration) (string, error)
	VerifyAppleReceiptAPI(receiptData string) (*model.IAPVerifyReceiptResponse, error)
	GetSubscribeAPI(originalTransactionId string, token string) (*model.IAPSubscribeAPIResponse, error)
}
