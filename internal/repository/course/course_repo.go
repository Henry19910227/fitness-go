package course

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/model/course"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type repository struct {
	gorm tool.Gorm
}

func New(gormTool tool.Gorm) Repository {
	return &repository{gorm: gormTool}
}

func (r *repository) List(input *model.ListParam) (output []*model.Table, amount int64, err error) {
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

	db := r.gorm.DB().Model(&model.Table{})
	db = db.Where(query, params...)
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
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
