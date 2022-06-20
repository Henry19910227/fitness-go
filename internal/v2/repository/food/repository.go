package food

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/v2/model/food"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/food"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List(input *food.ListInput) (outputs []*food.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{}).
		Select("foods.id AS id", "foods.user_id AS user_id", "foods.food_category_id AS food_category_id",
			"foods.source AS source", "foods.name AS name", "foods.calorie AS calorie",
			"foods.amount_desc AS amount_desc", "foods.create_at AS create_at", "foods.update_at AS update_at")
	//加入 tag 篩選條件
	if input.Tag != nil {
		db = db.Joins("INNER JOIN food_categories ON foods.food_category_id = food_categories.id")
		db = db.Where("food_categories.tag = ?", *input.Tag)
	}
	//加入 name 篩選條件
	if input.Name != nil {
		db = db.Where("foods.name LIKE ?", "%"+*input.Name+"%")
	}
	if input.UserID != nil {
		db = db.Where("(foods.user_id = ? OR foods.user_id IS NULL)", *input.UserID)
	}
	//加入 is_delete 篩選條件
	if input.IsDeleted != nil {
		db = db.Where("foods.is_deleted = ?", *input.IsDeleted)
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
		db = db.Order(fmt.Sprintf("foods.%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) Create(items []*food.Table) (err error) {
	panic("implement me")
}

func (r *repository) Update(items []*food.Table) (err error) {
	panic("implement me")
}
