package subscribe_log

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_log"
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

func (r *repository) CreateOrUpdate(item *model.Table) (id *int64, err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "original_transaction_id"}, {Name: "transaction_id"}, {Name: "type"}},
		DoNothing: true,
	}).Create(&item).Error
	if err != nil {
		return nil, err
	}
	return item.ID, err
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 id 篩選條件
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	//加入 original_transaction_id 篩選條件
	if input.OriginalTransactionID != nil {
		db = db.Where("original_transaction_id = ?", *input.OriginalTransactionID)
	}
	//加入 transaction_id 篩選條件
	if input.TransactionID != nil {
		db = db.Where("transaction_id = ?", *input.TransactionID)
	}
	//加入 transaction_id 篩選條件
	if input.Type != nil {
		db = db.Where("transaction_id = ?", *input.TransactionID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	// 加入 original_transaction_id 篩選條件
	if input.OriginalTransactionID != nil {
		db = db.Where("subscribe_logs.original_transaction_id = ?", *input.OriginalTransactionID)
	}
	// 加入 original_transaction_id 篩選條件
	if input.TransactionID != nil {
		db = db.Where("subscribe_logs.transaction_id = ?", *input.TransactionID)
	}
	// 加入 type 篩選條件
	if input.TransactionID != nil {
		db = db.Where("subscribe_logs.type = ?", *input.Type)
	}
	// Custom Where
	if len(input.Wheres) > 0 {
		for _, where := range input.Wheres {
			db = db.Where(where.Query, where.Args...)
		}
	}
	// Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
		}
	}
	// Count
	db = db.Count(&amount)
	// Select
	db = db.Select("subscribe_logs.*")
	// Paging
	if input.Page != nil && input.Size != nil {
		db = db.Offset((*input.Page - 1) * *input.Size).Limit(*input.Size)
	} else if input.Page != nil {
		db = db.Offset(0)
	} else if input.Size != nil {
		db = db.Limit(*input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType))
	}
	// Custom Order
	if input.Orders != nil {
		for _, orderBy := range input.Orders {
			db = db.Order(orderBy.Value)
		}
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) Create(item *model.Table) (id int64, err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return 0, err
	}
	return *item.ID, err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", *item.ID).Save(item).Error
	return err
}
