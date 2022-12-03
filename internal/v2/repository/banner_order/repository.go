package banner_order

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/banner_order"
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

func (r *repository) Creates(items []*model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Create(items).Error
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
	db = db.Select("banner_orders.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("banner_orders.%s %s", input.OrderField, input.OrderType))
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
	err = r.db.Where("banner_orders.banner_id = ?", input.BannerID).Delete(&model.Table{}).Error
	return err
}

func (r *repository) DeleteAll() (err error) {
	err = r.db.Where("banner_orders.banner_id > 0").Delete(&model.Table{}).Error
	return err
}
