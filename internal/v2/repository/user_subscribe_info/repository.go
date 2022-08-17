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
	//加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("user_id = ?", *input.UserID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
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

func (r *repository) CreateOrUpdate(item *model.Table) (err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"order_id", "status", "payment_type", "start_date", "expires_date", "update_at"}),
	}).Create(&item).Error
	return err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("user_id = ?", *item.UserID).Save(item).Error
	return err
}
