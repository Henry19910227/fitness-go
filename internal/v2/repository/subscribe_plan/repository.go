package subscribe_plan

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Find(input *subscribe_plan.FindInput) (output *subscribe_plan.Output, err error) {
	db := r.db.Model(&model.Output{})
	//加入 subscribe_plans.id 篩選條件
	if input.ID != nil {
		db = db.Where("subscribe_plans.id = ?", *input.ID)
	}
	//加入 product_labels.product_id 篩選條件
	if input.ProductID != nil {
		db = db.Joins("INNER JOIN product_labels ON subscribe_plans.product_label_id = product_labels.id")
		db = db.Where("product_labels.product_id = ?", *input.ProductID)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	// Select
	db = db.Select("subscribe_plans.*")
	//查詢數據
	err = db.First(&output).Error
	return output, err
}
