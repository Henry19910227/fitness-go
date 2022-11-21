package subscribe_plan

import (
	"fmt"
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

func (r *repository) WithTrx(tx *gorm.DB) Repository {
	return New(tx)
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

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	// enable 篩選條件
	if input.Enable != nil {
		db = db.Where("subscribe_plans.enable = ?", *input.Enable)
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
	db = db.Select("subscribe_plans.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("subscribe_plans.%s %s", input.OrderField, input.OrderType))
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
