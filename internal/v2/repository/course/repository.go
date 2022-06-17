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

func (r *repository) Find(input *model.FindInput) (output *model.Table, err error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 id 篩選條件
	if input.ID != nil {
		query += "AND id = ? "
		params = append(params, *input.ID)
	}
	db := r.db.Model(&model.Table{})
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

func (r *repository) List(input *model.ListInput) (output []*model.Table, amount int64, err error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 id 篩選條件
	if input.ID != nil {
		query += "AND id = ? "
		params = append(params, *input.ID)
	}
	//加入 name 篩選條件
	if input.Name != nil {
		query += "AND name LIKE ? "
		params = append(params, "%"+*input.Name+"%")
	}
	//加入 course_status 篩選條件
	if input.CourseStatus != nil {
		query += "AND course_status = ? "
		params = append(params, *input.CourseStatus)
	}
	//加入 trainer_status 篩選條件
	if input.SaleType != nil {
		query += "AND sale_type = ? "
		params = append(params, *input.SaleType)
	}

	db := r.db.Model(&model.Table{})
	db = db.Where(query, params...)
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
	err = db.Find(&output).Error
	return output, amount, err
}
