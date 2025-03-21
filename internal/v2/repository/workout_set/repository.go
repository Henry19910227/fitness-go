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

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 id 篩選條件
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	// Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
		}
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
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

func (r *repository) List(input *model.ListInput) (output []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	db = db.Joins("LEFT JOIN workout_set_orders ON workout_sets.id = workout_set_orders.workout_set_id")
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	// WorkoutID 篩選條件
	if input.WorkoutID != nil {
		db = db.Where("workout_sets.workout_id = ?", *input.WorkoutID)
	}
	// action_id 篩選條件
	if input.ActionID != nil {
		db = db.Where("workout_sets.action_id = ?", *input.ActionID)
	}
	// Type 篩選條件
	if input.Type != nil {
		db = db.Where("workout_sets.type = ?", *input.Type)
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
	db = db.Select("workout_sets.*")
	// Paging
	if input.Page != nil && input.Size != nil {
		db = db.Offset((*input.Page - 1) * *input.Size).Limit(*input.Size)
	} else if input.Page != nil {
		db = db.Offset(0)
	} else if input.Size != nil {
		db = db.Limit(*input.Size)
	}
	// Order
	db = db.Order(fmt.Sprintf("workout_set_orders.seq IS NULL ASC, workout_set_orders.seq ASC, workout_sets.create_at ASC"))
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("workout_sets.%s %s", input.OrderField, input.OrderType))
	}
	// 查詢數據
	err = db.Find(&output).Error
	return output, amount, err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", *item.ID).Save(item).Error
	return err
}

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	err = r.db.Where("id = ?", input.ID).Delete(&model.Table{}).Error
	return err
}
