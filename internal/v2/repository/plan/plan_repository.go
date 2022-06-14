package plan

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/plan"
	"gorm.io/gorm"
)

type repository struct {
	gorm tool.Gorm
}

func New(gormTool tool.Gorm) Repository {
	return &repository{gorm: gormTool}
}

func (r *repository) List(input *model.ListInput) (output []*model.Table, amount int64, err error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 id 篩選條件
	if input.CourseID != nil {
		query += "AND course_id = ? "
		params = append(params, *input.CourseID)
	}

	db := r.gorm.DB().Model(&model.Table{})
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
	//查詢數據
	err = db.Count(&amount).
		Offset((input.Page - 1) * input.Size).
		Limit(input.Size).
		Order(fmt.Sprintf("%s %s", input.OrderField, input.OrderType)).
		Find(&output).Error
	return output, amount, err
}
