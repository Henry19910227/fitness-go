package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type receipt struct {
	gorm tool.Gorm
}

func NewReceipt(gorm tool.Gorm) Receipt {
	return &receipt{gorm: gorm}
}

func (r *receipt) SaveReceipt(tx *gorm.DB, param *model.CreateReceiptParam) (int64, error) {
	if param == nil {
		return 0, nil
	}
	db := r.gorm.DB()
	if tx != nil {
		db = tx
	}
	receipt := entity.Receipt{
		OrderID:               param.OrderID,
		PaymentType:           param.PaymentType,
		ReceiptToken:          param.ReceiptToken,
		OriginalTransactionID: param.OriginalTransactionID,
		TransactionID:         param.TransactionID,
		ProductID:             param.ProductID,
		Quantity:              param.Quantity,
		CreateAt:              time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "order_id"}, {Name: "original_transaction_id"}},
		DoNothing: true,
	}).Create(&receipt).Error; err != nil {
		return 0, err
	}
	return receipt.ID, nil
}
