package user_course_asset

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_asset"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// 加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("user_id = ?", *input.UserID)
	}
	// 加入 course_id 篩選條件
	if input.CourseID != nil {
		db = db.Where("course_id = ?", *input.CourseID)
	}
	// 加入 available 篩選條件
	if input.Available != nil {
		db = db.Where("available = ?", *input.Available)
	}
	// Preload
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

