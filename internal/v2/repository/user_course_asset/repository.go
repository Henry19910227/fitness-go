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

func (r *repository) WithTrx(tx *gorm.DB) Repository {
	return New(tx)
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	if input.UserID != nil {
		db = db.Where("user_id = ?", *input.UserID)
	}
	if input.CourseID != nil {
		db = db.Where("course_id = ?", *input.CourseID)
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	// 加入 user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("user_course_assets.user_id = ?", *input.UserID)
	}
	// 加入 course_id 篩選條件
	if input.CourseID != nil {
		db = db.Where("user_course_assets.course_id = ?", *input.CourseID)
	}
	// 加入 available 篩選條件
	if input.Available != nil {
		db = db.Where("user_course_assets.available = ?", *input.Available)
	}
	// 加入 source 篩選條件
	if input.Source != nil {
		db = db.Where("user_course_assets.source = ?", *input.Source)
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
	db = db.Select("user_course_assets.*")
	// Paging
	if input.Page != nil && input.Size != nil {
		db = db.Offset((*input.Page - 1) * *input.Size).Limit(*input.Size)
	} else if input.Page != nil {
		db = db.Offset(0)
	} else if input.Size != nil {
		db = db.Limit(*input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("user_course_assets.%s %s", input.OrderField, input.OrderType))
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

func (r *repository) Create(item *model.Table) (id int64, err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return 0, err
	}
	return *item.ID, err
}

func (r *repository) Creates(items []*model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Create(&items).Error
	return err
}

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	db := r.db
	db = db.Where("id = ?", input.ID)
	err = db.Delete(&model.Table{}).Error
	return err
}
