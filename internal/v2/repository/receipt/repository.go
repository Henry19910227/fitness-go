package receipt

import (
	"fmt"
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
	// OrderID 篩選條件
	if input.OrderID != nil {
		db = db.Where("order_id = ?", *input.OrderID)
	}
	// TransactionID 篩選條件
	if input.TransactionID != nil {
		db = db.Where("transaction_id = ?", *input.TransactionID)
	}
	// OriginalTransactionID 篩選條件
	if input.OriginalTransactionID != nil {
		db = db.Where("original_transaction_id = ?", *input.OriginalTransactionID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			if preload.OrderBy != nil {
				db = db.Preload(preload.Field, func(db *gorm.DB) *gorm.DB {
					return db.Order(fmt.Sprintf("%s %s", preload.OrderBy.OrderField, preload.OrderBy.OrderType))
				})
				continue
			}
			db = db.Preload(preload.Field)
		}
	}
	// Count
	db = db.Count(&amount)
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) CreateOrUpdate(item *model.Table) (id *int64, err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "order_id"}, {Name: "original_transaction_id"}, {Name: "transaction_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"receipt_token"}),
	}).Create(&item).Error
	if err != nil {
		return nil, err
	}
	return item.ID, err
}
