package plan

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
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

func (r *repository) Create(item *model.Table) (id int64, err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return 0, err
	}
	return *item.ID, err
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	db := r.db.Model(&model.Output{})
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	//查詢數據
	err = db.First(&output).Error
	return output, err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", *item.ID).Save(item).Error
	return err
}

func (r *repository) List(input *model.ListInput) (output []*model.Output, amount int64, err error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 id 篩選條件
	if input.CourseID != nil {
		query += "AND course_id = ? "
		params = append(params, *input.CourseID)
	}

	db := r.db.Model(&model.Output{})
	db = db.Where(query, params...)
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
		db = db.Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType))
	}
	//查詢數據
	err = db.Find(&output).Error
	return output, amount, err
}
