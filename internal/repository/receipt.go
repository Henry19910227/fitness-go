package repository

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/global"
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

func (r *receipt) FindReceiptsByOrderID(orderID string, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.Receipt, error) {
	var receipts []*model.Receipt
	db := r.gorm.DB().Table("receipts")
	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	//頁數
	if paging != nil {
		if paging.Offset > 0 {
			db = db.Offset(paging.Offset)
		}
	}
	//筆數
	if paging != nil {
		if paging.Limit > 0 {
			db = db.Limit(paging.Limit)
		}
	}
	if err := db.Where("order_id = ?", orderID).Find(&receipts).Error; err != nil {
		return nil, err
	}
	return receipts, nil
}

func (r *receipt) FindReceiptsByPaymentType(userID int64, paymentType global.PaymentType, orderBy *model.OrderBy, paging *model.PagingParam) ([]*model.Receipt, error) {
	var receipts []*model.Receipt
	db := r.gorm.DB().
		Table("receipts").
		Joins("INNER JOIN orders ON receipts.order_id = orders.id")
	//排序
	if orderBy != nil {
		db = db.Order(fmt.Sprintf("%s %s", orderBy.Field, orderBy.OrderType))
	}
	//頁數
	if paging != nil {
		if paging.Offset > 0 {
			db = db.Offset(paging.Offset)
		}
	}
	//筆數
	if paging != nil {
		if paging.Limit > 0 {
			db = db.Limit(paging.Limit)
		}
	}
	if err := db.Where("orders.user_id = ? AND receipts.payment_type = ?", userID, paymentType).Find(&receipts).Error; err != nil {
		return nil, err
	}
	return receipts, nil
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
		Columns:   []clause.Column{{Name: "order_id"}, {Name: "original_transaction_id"}, {Name: "transaction_id"}},
		DoNothing: true,
	}).Create(&receipt).Error; err != nil {
		return 0, err
	}
	return receipt.ID, nil
}
