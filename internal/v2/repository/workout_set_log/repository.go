package workout_set_log

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set_log"
	"gorm.io/gorm"
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

func (r *repository) Create(items []*model.Table) (ids []int64, err error) {
	err = r.db.Model(&model.Table{}).Create(items).Error
	for _, item := range items {
		if item.ID != nil {
			ids = append(ids, *item.ID)
		}
	}
	return ids, err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	//加入 workout_log_id 篩選條件
	if input.WorkoutLogID != nil {
		db = db.Where("workout_log_id = ?", *input.WorkoutLogID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
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
