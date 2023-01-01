package user_course_usage_monthly_statistic

import (
	"fmt"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user_course_usage_monthly_statistic"
	"gorm.io/gorm"
	"time"
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
	// Join
	if len(input.Joins) > 0 {
		for _, join := range input.Joins {
			db = db.Joins(join.Query, join.Args...)
		}
	}
	// 加入 id 篩選條件
	if input.UserID != nil {
		db = db.Where("user_course_usage_monthly_statistics.user_id = ?", *input.UserID)
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
	// Select
	db = db.Select("user_course_usage_monthly_statistics.*")
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
	// user_id 篩選條件
	if input.UserID != nil {
		db = db.Where("user_course_usage_monthly_statistics.user_id = ?", *input.UserID)
	}
	// year 篩選條件
	if input.Year != nil {
		db = db.Where("user_course_usage_monthly_statistics.year = ?", *input.Year)
	}
	// month 篩選條件
	if input.Month != nil {
		db = db.Where("user_course_usage_monthly_statistics.month = ?", *input.Month)
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
	db = db.Select("user_course_usage_monthly_statistics.*")
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
		db = db.Order(fmt.Sprintf("user_course_usage_monthly_statistics.%s %s", input.OrderField, input.OrderType))
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

// Statistic https://kind-bass-788.notion.site/user_course_usage_monthly_statistic-ce495e1933c3483faef17f5fa8c3d2e8
func (r *repository) Statistic() (err error) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	err = r.db.Exec("INSERT INTO user_course_usage_monthly_statistics " +
		"( " +
		"user_id, " +
		"free_usage_count, " +
		"subscribe_usage_count, " +
		"charge_usage_count, " +
		"year, " +
		"month " +
		") " +
		"SELECT " +
		"IFNULL(base.user_id, 0) AS user_id, " +
		"IFNULL(a.free_usage_count, 0) AS free_usage_count, " +
		"IFNULL(b.subscribe_usage_count, 0) AS subscribe_usage_count, " +
		"IFNULL(c.charge_usage_count, 0) AS charge_usage_count, " +
		"IFNULL(base.year, 0) AS year, " +
		"IFNULL(base.month, 0) AS month " +
		"FROM " +
		"( " + tableBase() + " ) AS base " +
		"LEFT JOIN " +
		"( " + tableA() + " ) AS a ON base.user_id = a.user_id " +
		"LEFT JOIN " +
		"( " + tableB() +" ) AS b on base.user_id = b.user_id " +
		"LEFT JOIN " +
		"( " + tableC() + " ) AS c on base.user_id = c.user_id " +
		"ON DUPLICATE KEY UPDATE " +
		"free_usage_count = IFNULL(a.free_usage_count, 0), " +
		"subscribe_usage_count = IFNULL(b.subscribe_usage_count, 0), " +
		"charge_usage_count = IFNULL(c.charge_usage_count, 0)", timeStr, timeStr, timeStr, timeStr, timeStr, timeStr).Error
	return err
}

func tableBase() string {
	t :=  "SELECT MAX(courses.user_id) AS trainer_id, courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM `courses` " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"WHERE courses.sale_type != 4 " +
		"AND DATE_FORMAT(workout_logs.create_at, '%Y-%m') = DATE_FORMAT(?, '%Y-%m') " +
		"GROUP BY courses.id, workout_logs.user_id"
	return "SELECT " +
		"t.trainer_id AS user_id, " +
		"COUNT(*) AS total_count, " +
		"MAX(DATE_FORMAT(?, '%Y')) AS year, " +
		"MAX(DATE_FORMAT(?, '%m')) AS month " +
		"FROM " +
		"( " + t + " ) AS t " +
		"GROUP BY `t`.`trainer_id` "
}

func tableA() string {
	t := "SELECT MAX(courses.user_id) AS trainer_id, courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM `courses` " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"WHERE courses.sale_type = 1 " +
		"AND DATE_FORMAT(workout_logs.create_at, '%Y-%m') = DATE_FORMAT(?, '%Y-%m') " +
		"GROUP BY courses.id, workout_logs.user_id "
	return "SELECT " +
		"t.trainer_id AS user_id, " +
		"COUNT(*) AS free_usage_count " +
		"FROM ( " + t + " ) AS t " +
		"GROUP BY `t`.`trainer_id` "
}

func tableB() string {
	t := "SELECT MAX(courses.user_id) AS trainer_id, courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM `courses` " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"WHERE courses.sale_type = 2 " +
		"AND DATE_FORMAT(workout_logs.create_at, '%Y-%m') = DATE_FORMAT(?, '%Y-%m') " +
		"GROUP BY courses.id, workout_logs.user_id "
	return "SELECT " +
		"t.trainer_id AS user_id, " +
		"COUNT(*) AS subscribe_usage_count " +
		"FROM ( " + t + " ) AS t " +
		"GROUP BY `t`.`trainer_id` "
}

func tableC() string {
	t := "SELECT MAX(courses.user_id) AS trainer_id, courses.id AS course_id, workout_logs.user_id AS user_id " +
		"FROM `courses` " +
		"INNER JOIN plans ON courses.id = plans.course_id " +
		"INNER JOIN workouts ON plans.id = workouts.plan_id " +
		"INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id " +
		"WHERE courses.sale_type = 3 " +
		"AND DATE_FORMAT(workout_logs.create_at, '%Y-%m') = DATE_FORMAT(?, '%Y-%m') " +
		"GROUP BY courses.id, workout_logs.user_id "
	return "SELECT " +
		"t.trainer_id AS user_id, " +
		"COUNT(*) AS charge_usage_count " +
		"FROM ( " + t + " ) AS t " +
		"GROUP BY `t`.`trainer_id` "
}