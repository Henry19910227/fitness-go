package model

import "github.com/Henry19910227/fitness-go/internal/global"

type CreatePurchaseLogParam struct {
	UserID int64   // 用戶id
	OrderID string // 訂單id
	Type global.PurchaseLogType // 訂單類型(1:購買/2:退費)
}
