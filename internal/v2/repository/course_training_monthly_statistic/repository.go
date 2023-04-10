package course_training_monthly_statistic

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_training_monthly_statistic"
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
		db = db.Where("course_training_monthly_statistics.year = ?", *input.Year)
	}
	//加入 month 篩選條件
	if input.Month != nil {
		db = db.Where("course_training_monthly_statistics.month = ?", *input.Month)
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
	db = db.Select("course_training_monthly_statistics.*")
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
		db = db.Order(fmt.Sprintf("course_training_monthly_statistics.%s %s", input.OrderField, input.OrderType))
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
INSERT INTO course_training_monthly_statistics (year, month, total, free, subscribe, charge, aerobic, interval_training, weight_training, resistance_training, bodyweight_training, other_training)
SELECT
  2023 AS year,
  3 AS month,
  COUNT(DISTINCT CONCAT(workout_logs.user_id, '_', courses.id)) AS total,
  COUNT(DISTINCT CASE WHEN courses.sale_type = 1 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS free,
  COUNT(DISTINCT CASE WHEN courses.sale_type = 2 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS subscribe,
  COUNT(DISTINCT CASE WHEN courses.sale_type = 3 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS charge,
  COUNT(DISTINCT CASE WHEN courses.category = 1 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS aerobic,
  COUNT(DISTINCT CASE WHEN courses.category = 2 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS interval_training,
  COUNT(DISTINCT CASE WHEN courses.category = 3 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS weight_training,
  COUNT(DISTINCT CASE WHEN courses.category = 4 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS resistance_training,
  COUNT(DISTINCT CASE WHEN courses.category = 5 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS bodyweight_training,
  COUNT(DISTINCT CASE WHEN courses.category = 6 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS other_training
FROM courses
INNER JOIN plans ON courses.id = plans.course_id
INNER JOIN workouts ON plans.id = workouts.plan_id
INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id
WHERE YEAR(workout_logs.create_at) = 2023 AND MONTH(workout_logs.create_at) = 3 AND (courses.sale_type = 1 OR courses.sale_type = 2 OR courses.sale_type = 3)
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
	err = r.db.Exec("INSERT INTO course_training_monthly_statistics (year, month, total, free, subscribe, charge, aerobic, interval_training, weight_training, resistance_training, bodyweight_training, other_training)\n"+
		"SELECT\n  "+
		"? AS year,\n  "+
		"? AS month,\n  "+
		"COUNT(DISTINCT CONCAT(workout_logs.user_id, '_', courses.id)) AS total,\n  "+
		"COUNT(DISTINCT CASE WHEN courses.sale_type = 1 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS free,\n  "+
		"COUNT(DISTINCT CASE WHEN courses.sale_type = 2 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS subscribe,\n  "+
		"COUNT(DISTINCT CASE WHEN courses.sale_type = 3 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS charge,\n  "+
		"COUNT(DISTINCT CASE WHEN courses.category = 1 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS aerobic,\n  "+
		"COUNT(DISTINCT CASE WHEN courses.category = 2 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS interval_training,\n  "+
		"COUNT(DISTINCT CASE WHEN courses.category = 3 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS weight_training,\n  "+
		"COUNT(DISTINCT CASE WHEN courses.category = 4 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS resistance_training,\n  "+
		"COUNT(DISTINCT CASE WHEN courses.category = 5 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS bodyweight_training,\n  "+
		"COUNT(DISTINCT CASE WHEN courses.category = 6 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS other_training\n"+
		"FROM courses\n"+
		"INNER JOIN plans ON courses.id = plans.course_id\n"+
		"INNER JOIN workouts ON plans.id = workouts.plan_id\n"+
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id\n"+
		"WHERE YEAR(workout_logs.create_at) = ? AND MONTH(workout_logs.create_at) = ? AND (courses.sale_type = 1 OR courses.sale_type = 2 OR courses.sale_type = 3)\n"+
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
