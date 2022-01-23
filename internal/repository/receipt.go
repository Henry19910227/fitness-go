package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type receipt struct {
	gorm  tool.Gorm
}

func NewReceipt(gorm tool.Gorm) Receipt {
	return &receipt{gorm: gorm}
}

func (r *receipt) CreateReceipt(tx *gorm.DB, param *model.CreateReceiptParam) (int64, error) {
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	receipt := entity.Receipt{
		OrderID: param.OrderID,
		PaymentType: param.PaymentType,
		ReceiptToken: param.ReceiptToken,
		OriginalTransactionID: param.OriginalTransactionID,
		TransactionID: param.TransactionID,
		Quantity: param.Quantity,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&receipt).Error; err != nil {
		return 0, err
	}
	return receipt.ID, nil
}