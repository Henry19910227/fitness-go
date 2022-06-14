package workout_set

import (
	"fmt"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/workout_set"
	"gorm.io/gorm"
)

type repository struct {
	gorm tool.Gorm
}

func New(gormTool tool.Gorm) Repository {
	return &repository{gorm: gormTool}
}

func (r repository) List(input *model.ListInput) (output []*model.Table, amount int64, err error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//篩選條件
	if input.WorkoutID != nil {
		query += "AND workout_sets.workout_id = ? "
		params = append(params, *input.WorkoutID)
	}
	db := r.gorm.DB().Model(&model.Table{})
	db = db.Joins("INNER JOIN workout_set_orders ON workout_sets.id = workout_set_orders.workout_set_id").Where(query, params...)
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
		Order(fmt.Sprintf("%s %s", "workout_set_orders.seq", "ASC")).
		Find(&output).Error
	return output, amount, err
}
