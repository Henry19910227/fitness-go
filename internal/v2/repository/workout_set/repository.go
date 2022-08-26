package workout_set

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
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
	for _, item := range items{
		if item.ID != nil {
			ids = append(ids, *item.ID)
		}
	}
	return ids, err
}

func (r *repository) List(input *model.ListInput) (output []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// WorkoutID 篩選條件
	if input.WorkoutID != nil {
		db = db.Where("workout_sets.workout_id = ?", *input.WorkoutID)
	}
	// Type 篩選條件
	if input.Type != nil {
		db = db.Where("workout_sets.type = ?", *input.Type)
	}
	// Join
	db = db.Joins("LEFT JOIN workout_set_orders ON workout_sets.id = workout_set_orders.workout_set_id")
	// Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	// Count
	db = db.Count(&amount)
	// Select
	db = db.Select("workout_sets.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	db = db.Order(fmt.Sprintf("workout_set_orders.seq IS NULL ASC, workout_set_orders.seq ASC, workout_sets.create_at ASC"))
	// 查詢數據
	err = db.Find(&output).Error
	return output, amount, err
}
