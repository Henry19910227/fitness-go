package review

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/review"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	//加入 courses.name 篩選條件
	if input.Name != nil {
		db = db.Joins("INNER JOIN courses ON reviews.course_id = courses.id")
		db = db.Where("courses.name LIKE ?", "%"+*input.Name+"%")
	}
	//加入 users.nickname 篩選條件
	if input.Nickname != nil {
		db = db.Joins("INNER JOIN users ON reviews.user_id = users.id")
		db = db.Where("users.nickname LIKE ?", "%"+*input.Nickname+"%")
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
	db = db.Select("reviews.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("reviews.%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}
