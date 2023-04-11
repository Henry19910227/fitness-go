package course_category_training_monthly_statistic

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/course_category_training_monthly_statistic"
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
	//加入 category 篩選條件
	if input.Category != nil {
		db = db.Where("category = ?", *input.Category)
	}
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
	//加入 category 篩選條件
	if input.Category != nil {
		db = db.Where("course_category_training_monthly_statistics.category = ?", *input.Category)
	}
	//加入 year 篩選條件
	if input.Year != nil {
		db = db.Where("course_category_training_monthly_statistics.year = ?", *input.Year)
	}
	//加入 month 篩選條件
	if input.Month != nil {
		db = db.Where("course_category_training_monthly_statistics.month = ?", *input.Month)
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
	db = db.Select("course_category_training_monthly_statistics.*")
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
		db = db.Order(fmt.Sprintf("course_category_training_monthly_statistics.%s %s", input.OrderField, input.OrderType))
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
INSERT INTO course_category_training_monthly_statistics (category, year, month, total, male, female, age_13_17, age_18_24, age_25_34, age_35_44, age_45_54, age_55_64, age_65_up)
SELECT
  2 AS category,
  2022 AS year,
  5 AS month,
  COUNT(DISTINCT CONCAT(workout_logs.user_id, '_', courses.id)) AS total,
  COUNT(DISTINCT CASE WHEN users.sex = 'm' THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS male,
  COUNT(DISTINCT CASE WHEN users.sex = 'f' THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS female,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 13 AND 17 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_13_17,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 18 AND 24 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_18_24,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 25 AND 34 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_25_34,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 35 AND 44 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_35_44,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 45 AND 54 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_45_54,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 55 AND 64 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_55_64,
  COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) >= 65 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_65_up
FROM courses
INNER JOIN plans ON courses.id = plans.course_id
INNER JOIN workouts ON plans.id = workouts.plan_id
INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id
INNER JOIN users ON users.id = workout_logs.user_id
WHERE courses.category = 2 AND YEAR(workout_logs.create_at) = 2022 AND MONTH(workout_logs.create_at) = 5 AND (courses.sale_type = 1 OR courses.sale_type = 2 OR courses.sale_type = 3)
ON DUPLICATE KEY UPDATE
  total = VALUES(total),
  male = VALUES(male),
  female = VALUES(female),
  age_13_17 = VALUES(age_13_17),
  age_18_24 = VALUES(age_18_24),
  age_25_34 = VALUES(age_25_34),
  age_35_44 = VALUES(age_35_44),
  age_45_54 = VALUES(age_45_54),
  age_55_64 = VALUES(age_55_64),
  age_65_up = VALUES(age_65_up),
  update_at = CURRENT_TIMESTAMP;
*/
func (r *repository) Statistic(input *model.StatisticInput) (err error) {
	err = r.db.Exec("INSERT INTO course_category_training_monthly_statistics (category, year, month, total, male, female, age_13_17, age_18_24, age_25_34, age_35_44, age_45_54, age_55_64, age_65_up)\n"+
		"SELECT\n  "+
		"? AS category,\n  "+
		"? AS year,\n  "+
		"? AS month,\n  "+
		"COUNT(DISTINCT CONCAT(workout_logs.user_id, '_', courses.id)) AS total,\n  "+
		"COUNT(DISTINCT CASE WHEN users.sex = 'm' THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS male,\n  "+
		"COUNT(DISTINCT CASE WHEN users.sex = 'f' THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS female,\n  "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 13 AND 17 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_13_17,\n  "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 18 AND 24 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_18_24,\n  "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 25 AND 34 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_25_34,\n  "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 35 AND 44 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_35_44,\n  "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 45 AND 54 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_45_54,\n  "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) BETWEEN 55 AND 64 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_55_64,\n  "+
		"COUNT(DISTINCT CASE WHEN TIMESTAMPDIFF(YEAR, users.birthday, LAST_DAY(CONCAT(2022, '-', 6, '-01'))) >= 65 THEN CONCAT(workout_logs.user_id, '_', courses.id) END) AS age_65_up\n"+
		"FROM courses\n"+
		"INNER JOIN plans ON courses.id = plans.course_id\n"+
		"INNER JOIN workouts ON plans.id = workouts.plan_id\n"+
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id\n"+
		"INNER JOIN users ON users.id = workout_logs.user_id\n"+
		"WHERE courses.category = ? AND YEAR(workout_logs.create_at) = ? AND MONTH(workout_logs.create_at) = ? AND (courses.sale_type = 1 OR courses.sale_type = 2 OR courses.sale_type = 3)\n"+
		"ON DUPLICATE KEY UPDATE\n  "+
		"total = VALUES(total),\n  "+
		"male = VALUES(male),\n  "+
		"female = VALUES(female),\n  "+
		"age_13_17 = VALUES(age_13_17),\n  "+
		"age_18_24 = VALUES(age_18_24),\n  "+
		"age_25_34 = VALUES(age_25_34),\n  "+
		"age_35_44 = VALUES(age_35_44),\n  "+
		"age_45_54 = VALUES(age_45_54),\n  "+
		"age_55_64 = VALUES(age_55_64),\n  "+
		"age_65_up = VALUES(age_65_up),\n  "+
		"update_at = CURRENT_TIMESTAMP",
		input.Category, input.Year, input.Month,
		input.Category, input.Year, input.Month).Error
	return err
}
