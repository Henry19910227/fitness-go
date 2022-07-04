package trainer

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 id 篩選條件
	if input.UserID != nil {
		db = db.Where("user_id = ?", *input.UserID)
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) FavoriteList(input *model.FavoriteListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{}).Joins("INNER JOIN favorite_trainers ON trainers.user_id = favorite_trainers.trainer_id")
	// id 篩選條件
	if input.UserID != nil {
		db = db.Where("favorite_trainers.user_id = ?", *input.UserID)
	}
	// Count
	db = db.Count(&amount)
	// Select
	db = db.Select("trainers.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("favorite_trainers.%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("user_id = ?", *item.UserID).Save(item).Error
	return err
}
