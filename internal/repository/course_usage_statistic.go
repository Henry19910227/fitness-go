package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type courseUsageStatistic struct {
	gorm tool.Gorm
}

func NewCourseUsageStatistic(gorm tool.Gorm) CourseUsageStatistic {
	return &courseUsageStatistic{gorm: gorm}
}

func (c *courseUsageStatistic) CalculateTotalFinishWorkoutCount(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	results := make([]*model.CourseUsageStatisticResult, 0)
	if err := db.
		Table("user_course_statistics").
		Select("MAX(course_id) AS course_id", "SUM(total_finish_workout_count) AS value").
		Group("course_id").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (c *courseUsageStatistic) CalculateUserFinishCount(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	results := make([]*model.CourseUsageStatisticResult, 0)
	sub := db.Table("courses").
		Select("courses.id AS course_id", "workout_logs.user_id AS user_id").
		Joins("INNER JOIN plans ON courses.id = plans.course_id").
		Joins("INNER JOIN workouts ON plans.id = workouts.plan_id").
		Joins("INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id").
		Group("courses.id, workout_logs.user_id")
	if err := db.Table("(?) AS t", sub).
		Select("t.course_id AS course_id", "COUNT(*) AS value").
		Group("t.course_id").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (c *courseUsageStatistic) CalculateMaleFinishCount(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	results := make([]*model.CourseUsageStatisticResult, 0)
	sub := db.Table("courses").
		Select("courses.id AS course_id", "workout_logs.user_id AS user_id").
		Joins("INNER JOIN plans ON courses.id = plans.course_id").
		Joins("INNER JOIN workouts ON plans.id = workouts.plan_id").
		Joins("INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id").
		Joins("INNER JOIN users ON workout_logs.user_id = users.id").
		Where("users.sex = ?", "m").
		Group("courses.id, workout_logs.user_id")
	if err := db.Table("(?) AS t", sub).
		Select("t.course_id AS course_id", "COUNT(*) AS value").
		Group("t.course_id").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (c *courseUsageStatistic) CalculateFemaleFinishCount(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	results := make([]*model.CourseUsageStatisticResult, 0)
	sub := db.Table("courses").
		Select("courses.id AS course_id", "workout_logs.user_id AS user_id").
		Joins("INNER JOIN plans ON courses.id = plans.course_id").
		Joins("INNER JOIN workouts ON plans.id = workouts.plan_id").
		Joins("INNER JOIN workout_logs ON workouts.id = workout_logs.workout_id").
		Joins("INNER JOIN users ON workout_logs.user_id = users.id").
		Where("users.sex = ?", "f").
		Group("courses.id, workout_logs.user_id")
	if err := db.Table("(?) AS t", sub).
		Select("t.course_id AS course_id", "COUNT(*) AS value").
		Group("t.course_id").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (c *courseUsageStatistic) CalculateFinishCountAvg(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	results := make([]*model.CourseUsageStatisticResult, 0)
	if err := db.
		Table("user_course_statistics").
		Select("course_id", "ROUND(AVG(finish_workout_count)) AS `value`").
		Group("course_id").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (c *courseUsageStatistic) SaveTotalFinishWorkoutCount(tx *gorm.DB, values []*model.CourseUsageStatisticResult) error {
	if len(values) == 0 {
		return nil
	}
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	params := make([]map[string]interface{}, 0)
	for _, v := range values {
		params = append(params, map[string]interface{}{"course_id": v.CourseID, "total_finish_workout_count": v.Value})
	}
	if err := db.Table("course_usage_statistics").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "course_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"total_finish_workout_count", "update_at"}),
	}).Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func (c *courseUsageStatistic) SaveUserFinishCount(tx *gorm.DB, values []*model.CourseUsageStatisticResult) error {
	if len(values) == 0 {
		return nil
	}
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	params := make([]map[string]interface{}, 0)
	for _, v := range values {
		params = append(params, map[string]interface{}{"course_id": v.CourseID, "user_finish_count": v.Value})
	}
	if err := db.Table("course_usage_statistics").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "course_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"user_finish_count", "update_at"}),
	}).Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func (c *courseUsageStatistic) SaveMaleFinishCount(tx *gorm.DB, values []*model.CourseUsageStatisticResult) error {
	if len(values) == 0 {
		return nil
	}
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	params := make([]map[string]interface{}, 0)
	for _, v := range values {
		params = append(params, map[string]interface{}{"course_id": v.CourseID, "male_finish_count": v.Value})
	}
	if err := db.Table("course_usage_statistics").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "course_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"male_finish_count", "update_at"}),
	}).Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func (c *courseUsageStatistic) SaveFemaleFinishCount(tx *gorm.DB, values []*model.CourseUsageStatisticResult) error {
	if len(values) == 0 {
		return nil
	}
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	params := make([]map[string]interface{}, 0)
	for _, v := range values {
		params = append(params, map[string]interface{}{"course_id": v.CourseID, "male_finish_count": v.Value})
	}
	if err := db.Table("course_usage_statistics").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "course_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"female_finish_count", "update_at"}),
	}).Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func (c *courseUsageStatistic) SaveFinishCountAvg(tx *gorm.DB, values []*model.CourseUsageStatisticResult) error {
	if len(values) == 0 {
		return nil
	}
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	params := make([]map[string]interface{}, 0)
	for _, v := range values {
		params = append(params, map[string]interface{}{"course_id": v.CourseID, "finish_count_avg": v.Value})
	}
	if err := db.Table("course_usage_statistics").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "course_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"finish_count_avg", "update_at"}),
	}).Create(&params).Error; err != nil {
		return err
	}
	return nil
}
