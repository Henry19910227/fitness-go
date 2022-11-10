package action

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/action"
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
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	// Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
		}
	}
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
	// ID List 篩選條件
	if len(input.IDs) > 0 {
		db = db.Where("actions.id IN (?)", input.IDs)
	}
	// Name 篩選條件
	if input.Name != nil {
		db = db.Where("actions.name LIKE ?", "%"+*input.Name+"%")
	}
	// Source 篩選條件
	if len(input.SourceList) > 0 {
		db = db.Where("actions.source IN (?)", input.SourceList)
	}
	// Category 篩選條件
	if len(input.CategoryList) > 0 {
		db = db.Where("actions.category IN (?)", input.CategoryList)
	}
	// Body 篩選條件
	if len(input.BodyList) > 0 {
		db = db.Where("actions.body IN (?)", input.BodyList)
	}
	// Equipment 篩選條件
	if len(input.EquipmentList) > 0 {
		db = db.Where("actions.equipment IN (?)", input.EquipmentList)
	}
	// UserID 篩選條件
	if input.UserID != nil {
		db = db.Where("actions.user_id = ? OR actions.user_id IS NULL", *input.UserID)
	}
	// course_id 篩選條件
	if input.CourseID != nil {
		db = db.Where("actions.course_id = ?", *input.CourseID)
	}
	// Type 篩選條件
	if input.Type != nil {
		db = db.Where("actions.type = ?", *input.Type)
	}
	// Source 篩選條件
	if input.Source != nil {
		db = db.Where("actions.source = ?", *input.Source)
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
	db = db.Select("actions.*")
	// Paging
	if input.Page > 0 && input.Size > 0 {
		db = db.Offset((input.Page - 1) * input.Size).Limit(input.Size)
	}
	// Order
	if len(input.OrderField) > 0 && len(input.OrderType) > 0 {
		db = db.Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType))
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

func (r *repository) Delete(input *model.DeleteInput) (err error) {
	err = r.db.Where("id = ?", input.ID).Delete(&model.Table{}).Error
	return err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", *item.ID).Save(item).Error
	return err
}
