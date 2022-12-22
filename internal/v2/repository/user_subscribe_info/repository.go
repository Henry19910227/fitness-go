package user_subscribe_info

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_subscribe_info"
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

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("user_id = ?", *input.UserID)
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
	//加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("user_subscribe_infos.user_id = ?", *input.UserID)
	}
	//加入 status 篩選條件
	if input.Status != nil {
		db = db.Where("user_subscribe_infos.status = ?", *input.Status)
	}
	//加入 order_id 篩選條件
	if input.OrderID != nil {
		db = db.Where("user_subscribe_infos.order_id = ?", *input.OrderID)
	}
	//加入 original_transaction_id 篩選條件
	if input.OriginalTransactionID != nil {
		db = db.Where("user_subscribe_infos.original_transaction_id = ?", *input.OriginalTransactionID)
	}
	//加入 payment_type 篩選條件
	if input.PaymentType != nil {
		db = db.Where("user_subscribe_infos.payment_type = ?", *input.PaymentType)
	}
	// Custom Where
	if len(input.Wheres) > 0 {
		for _, where := range input.Wheres {
			db = db.Where(where.Query, where.Args...)
		}
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	// Count
	db = db.Count(&amount)
	// Select
	if len(input.Selects) > 0 {
		db = db.Select(input.Selects[0].Query, input.Selects[0].Args...)
	} else {
		db = db.Select("user_subscribe_infos.*")
	}
	// Group
	if len(input.Groups) > 0 {
		db = db.Group(input.Groups[0].Name)
	}
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

func (r *repository) Create(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	return err
}

func (r *repository) CreateOrUpdate(item *model.Table) (err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"order_id", "original_transaction_id", "status", "payment_type", "start_date", "expires_date", "update_at"}),
	}).Create(&item).Error
	return err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("user_id = ?", *item.UserID).Save(item).Error
	return err
}

func (r *repository) Updates(items []*model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Save(&items).Error
	return err
}
