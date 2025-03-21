package meal

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/meal"
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

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{}).Select("meals.id AS id", "meals.diet_id AS diet_id",
		"meals.food_id AS food_id", "meals.type AS type", "meals.amount AS amount", "meals.create_at AS create_at")
	//加入 diet id 篩選條件
	if input.DietID != nil {
		db = db.Where("meals.diet_id = ?", *input.DietID)
	}
	//加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Joins("INNER JOIN diets ON meals.diet_id = diets.id")
		db = db.Where("diets.user_id = ?", *input.UserID)
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
		db = db.Order(fmt.Sprintf("meals.%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) Create(items []*model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Create(items).Error
	return err
}

func (r *repository) Update(items []*model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Save(&items).Error
	return err
}

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	db := r.db
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	if input.DietID != nil {
		db = db.Where("diet_id = ?", *input.DietID)
	}
	err = db.Delete(&model.Table{}).Error
	return err
}
