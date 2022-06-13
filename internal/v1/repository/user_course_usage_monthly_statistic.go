package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userCourseUsageMonthlyStatistic struct {
	gorm tool.Gorm
}

func NewUserCourseUsageMonthlyStatistic(gorm tool.Gorm) UserCourseUsageMonthlyStatistic {
	return &userCourseUsageMonthlyStatistic{gorm: gorm}
}

// CalculateCourseUsageMonthlyCount SQL
//SELECT t.trainer_id AS user_id, COUNT(*) AS `value`
//FROM (
//SELECT MAX(courses.user_id) AS trainer_id, courses.id AS course_id, workout_logs.user_id AS user_id
//FROM `courses`
//INNER JOIN plans ON courses.id = plans.course_id
//INNER JOIN workouts ON plans.id = workouts.plan_id
//INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id
//WHERE courses.sale_type = 2
//AND DATE_FORMAT(workout_logs.create_at, '%Y-%m') = DATE_FORMAT('2022-04-05 11:00:00', '%Y-%m')
//GROUP BY courses.id, workout_logs.user_id
//) AS t
//GROUP BY `t`.`trainer_id`
func (u *userCourseUsageMonthlyStatistic) CalculateCourseUsageMonthlyCount(tx *gorm.DB, saleType global.SaleType, date string) ([]*model.UserCourseUsageMonthlyStatisticResult, error) {
	db := u.gorm.DB()
	if tx != nil {
		db = tx
	}
	results := make([]*model.UserCourseUsageMonthlyStatisticResult, 0)
	sub := db.Table("courses").
		Select("MAX(courses.user_id) AS trainer_id",
			"courses.id AS course_id",
			"workout_logs.user_id AS user_id",
			"MAX(DATE_FORMAT(workout_logs.create_at, '%Y')) AS `year`",
			"MAX(DATE_FORMAT(workout_logs.create_at, '%c')) AS `month`").
		Joins("INNER JOIN plans ON courses.id = plans.course_id").
		Joins("INNER JOIN workouts ON plans.id = workouts.plan_id").
		Joins("INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id").
		Where("courses.sale_type = ? AND DATE_FORMAT(workout_logs.create_at, '%Y-%m') = DATE_FORMAT(?, '%Y-%m')", saleType, date).
		Group("courses.id, workout_logs.user_id")
	if err := db.Table("(?) AS t", sub).
		Select("t.trainer_id AS user_id", "COUNT(*) AS value", "MAX(t.`year`) AS `year`", "MAX(t.`month`) AS `month`").
		Group("t.trainer_id").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (u *userCourseUsageMonthlyStatistic) Save(tx *gorm.DB, ColumnName string, values []*model.UserCourseUsageMonthlyStatisticResult) error {
	if len(values) == 0 {
		return nil
	}
	db := u.gorm.DB()
	if tx != nil {
		db = tx
	}
	params := make([]map[string]interface{}, 0)
	for _, v := range values {
		params = append(params, map[string]interface{}{"user_id": v.UserID, ColumnName: v.Value, "month": v.Month, "year": v.Year})
	}
	if err := db.Table("user_course_usage_monthly_statistics").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "year"}, {Name: "month"}},
		DoUpdates: clause.AssignmentColumns([]string{ColumnName, "update_at"}),
	}).Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func (u *userCourseUsageMonthlyStatistic) Find(userID int64, output interface{}) error {
	if err := u.gorm.DB().
		Model(&entity.UserCourseUsageMonthlyStatistic{}).
		Where("user_id = ?", userID).
		Take(output).Error; err != nil {
		return err
	}
	return nil
}
