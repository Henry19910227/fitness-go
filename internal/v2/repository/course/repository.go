package course

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Find(input *model.FindInput) (output *model.Output, err error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 id 篩選條件
	if input.ID != nil {
		query += "AND id = ? "
		params = append(params, *input.ID)
	}
	db := r.db.Model(&model.Output{})
	db = db.Where(query, params...)
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
		}
	}
	//查詢數據
	err = db.Take(&output).Error
	return output, err
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, amount int64, err error) {
	db := r.db.Model(&model.Output{})
	// id 篩選條件
	if input.ID != nil {
		db = db.Where("id = ?", *input.ID)
	}
	// name 篩選條件
	if input.Name != nil {
		db = db.Where("name = ?", "%"+*input.Name+"%")
	}
	// course_status 篩選條件
	if input.CourseStatus != nil {
		db = db.Where("course_status = ?", *input.CourseStatus)
	}
	// trainer_status 篩選條件
	if input.SaleType != nil {
		db = db.Where("sale_type = ?", *input.SaleType)
	}
	if len(input.IDs) > 0 {
		db = db.Where("id IN (?)", input.IDs)
	}
	if len(input.IgnoredCourseStatus) > 0 {
		db = db.Where("course_status NOT IN ?", input.IgnoredCourseStatus)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			if preload.OrderBy != nil {
				db = db.Preload(preload.Field, func(db *gorm.DB) *gorm.DB {
					return db.Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType))
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
	err = db.Find(&outputs).Error
	return outputs, amount, err
}

func (r *repository) Updates(items []*model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Save(&items).Error
	return err
}

func (r *repository) Update(item *model.Table) (err error) {
	err = r.db.Model(&model.Table{}).Where("id = ?", *item.ID).Save(item).Error
	return err
}
