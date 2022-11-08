package favorite_trainer

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/favorite_trainer"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return err
	}
	return err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	// user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("favorite_trainers.user_id = ?", *input.UserID)
	}
	// trainer_id 篩選條件
	if input.TrainerID != nil {
		db = db.Where("favorite_trainers.trainer_id = ?", *input.TrainerID)
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
	db = db.Select("favorite_trainers.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("favorite_trainers.%s %s", input.OrderField, input.OrderType))
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

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	err = r.db.
		Where("user_id = ?", input.UserID).
		Where("trainer_id = ?", input.TrainerID).
		Delete(&model.Table{}).Error
	return err
}
