package course_create_monthly_statistic

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_create_monthly_statistic"
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
	//加入 year 篩選條件
	if input.Year != nil {
		db = db.Where("year = ?", *input.Year)
	}
	//加入 month 篩選條件
	if input.Month != nil {
		db = db.Where("month = ?", *input.Month)
	}
	//Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field)
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
	//加入 year 篩選條件
	if input.Year != nil {
		db = db.Where("course_create_monthly_statistics.year = ?", *input.Year)
	}
	//加入 month 篩選條件
	if input.Month != nil {
		db = db.Where("course_create_monthly_statistics.month = ?", *input.Month)
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
	db = db.Select("course_create_monthly_statistics.*")
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
		db = db.Order(fmt.Sprintf("course_create_monthly_statistics.%s %s", input.OrderField, input.OrderType))
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

// Statistic SQL
/*
INSERT INTO course_create_monthly_statistics (year, month, total, free, subscribe, charge, aerobic, interval_training, weight_training, resistance_training, bodyweight_training, other_training)
SELECT
  2021 AS year,
  6 AS month,
  COUNT(*) AS total,
  COUNT(CASE WHEN courses.sale_type = 1 THEN 0 END) AS free,
  COUNT(CASE WHEN courses.sale_type = 2 THEN 0 END) AS subscribe,
  COUNT(CASE WHEN courses.sale_type = 3 THEN 0 END) AS charge,
  COUNT(CASE WHEN courses.category = 1 THEN 0 END) AS aerobic,
  COUNT(CASE WHEN courses.category = 2 THEN 0 END) AS interval_training,
  COUNT(CASE WHEN courses.category = 3 THEN 0 END) AS weight_training,
  COUNT(CASE WHEN courses.category = 4 THEN 0 END) AS resistance_training,
  COUNT(CASE WHEN courses.category = 5 THEN 0 END) AS bodyweight_training,
  COUNT(CASE WHEN courses.category = 6 THEN 0 END) AS other_training
FROM courses
WHERE YEAR(courses.create_at) = 2021 AND MONTH(courses.create_at) = 6
AND courses.sale_type = 1 OR courses.sale_type = 2 OR courses.sale_type = 3
ON DUPLICATE KEY UPDATE
  total = VALUES(total),
  free = VALUES(free),
  subscribe = VALUES(subscribe),
  charge = VALUES(charge),
  aerobic = VALUES(aerobic),
  interval_training = VALUES(interval_training),
  weight_training = VALUES(weight_training),
  resistance_training = VALUES(resistance_training),
  bodyweight_training = VALUES(bodyweight_training),
  other_training = VALUES(other_training),
  update_at = CURRENT_TIMESTAMP;
*/
func (r *repository) Statistic(input *model.StatisticInput) (err error) {
	err = r.db.Exec("INSERT INTO course_create_monthly_statistics (year, month, total, free, subscribe, charge, aerobic, interval_training, weight_training, resistance_training, bodyweight_training, other_training)\n "+
		"SELECT\n  "+
		"? AS year,\n  "+
		"? AS month,\n  "+
		"COUNT(*) AS total,\n  "+
		"COUNT(CASE WHEN courses.sale_type = 1 THEN 0 END) AS free,\n  "+
		"COUNT(CASE WHEN courses.sale_type = 2 THEN 0 END) AS subscribe,\n  "+
		"COUNT(CASE WHEN courses.sale_type = 3 THEN 0 END) AS charge,\n  "+
		"COUNT(CASE WHEN courses.category = 1 THEN 0 END) AS aerobic,\n  "+
		"COUNT(CASE WHEN courses.category = 2 THEN 0 END) AS interval_training,\n  "+
		"COUNT(CASE WHEN courses.category = 3 THEN 0 END) AS weight_training,\n  "+
		"COUNT(CASE WHEN courses.category = 4 THEN 0 END) AS resistance_training,\n  "+
		"COUNT(CASE WHEN courses.category = 5 THEN 0 END) AS bodyweight_training,\n  "+
		"COUNT(CASE WHEN courses.category = 6 THEN 0 END) AS other_training\n "+
		"FROM courses\n "+
		"WHERE YEAR(courses.create_at) = ? AND MONTH(courses.create_at) = ?\n "+
		"AND (courses.sale_type = 1 OR courses.sale_type = 2 OR courses.sale_type = 3)\n "+
		"ON DUPLICATE KEY UPDATE\n  "+
		"total = VALUES(total),\n  "+
		"free = VALUES(free),\n  "+
		"subscribe = VALUES(subscribe),\n  "+
		"charge = VALUES(charge),\n  "+
		"aerobic = VALUES(aerobic),\n  "+
		"interval_training = VALUES(interval_training),\n  "+
		"weight_training = VALUES(weight_training),\n  "+
		"resistance_training = VALUES(resistance_training),\n  "+
		"bodyweight_training = VALUES(bodyweight_training),\n  "+
		"other_training = VALUES(other_training),\n  "+
		"update_at = CURRENT_TIMESTAMP",
		input.Year, input.Month,
		input.Year, input.Month).Error
	return err
}
