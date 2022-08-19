package receipt

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/receipt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) WithTrx(tx *gorm.DB) Repository {
	return New(tx)
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	//加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Joins("INNER JOIN orders ON orders.id = receipts.order_id")
		db = db.Where("orders.user_id = ?", *input.UserID)
	}
	// OrderID 篩選條件
	if input.OrderID != nil {
		db = db.Where("receipts.order_id = ?", *input.OrderID)
	}
	// TransactionID 篩選條件
	if input.TransactionID != nil {
		db = db.Where("receipts.transaction_id = ?", *input.TransactionID)
	}
	// OriginalTransactionID 篩選條件
	if input.OriginalTransactionID != nil {
		db = db.Where("receipts.original_transaction_id = ?", *input.OriginalTransactionID)
	}
	// PaymentType 篩選條件
	if input.PaymentType != nil {
		db = db.Where("receipts.payment_type = ?", *input.PaymentType)
	}
	// HaveReceiptToken 篩選條件
	if input.HaveReceiptToken != nil {
		if *input.HaveReceiptToken > 0 {
			db = db.Where("LENGTH(receipts.receipt_token) > 0")
		} else {
			db = db.Where("LENGTH(receipts.receipt_token) = 0")
		}
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			if preload.OrderBy != nil {
				db = db.Preload(preload.Field, func(db *gorm.DB) *gorm.DB {
					return db.Order(fmt.Sprintf("receipts.%s %s", preload.OrderBy.OrderField, preload.OrderBy.OrderType))
				})
				continue
			}
			db = db.Preload(preload.Field)
		}
	}
	// Count
	db = db.Count(&amount)
	// Select
	db = db.Select("receipts.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("receipts.%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) CreateOrUpdate(item *model.Table) (id *int64, err error) {
	values := make([]string, 0)
	if len(util.OnNilJustReturnString(item.ReceiptToken, "")) > 0 {
		values = append(values, "receipt_token")
	}
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "order_id"}, {Name: "original_transaction_id"}, {Name: "transaction_id"}},
		DoUpdates: clause.AssignmentColumns(values),
	}).Create(&item).Error
	if err != nil {
		return nil, err
	}
	return item.ID, err
}
