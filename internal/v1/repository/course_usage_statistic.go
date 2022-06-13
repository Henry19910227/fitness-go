package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
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

func (c *courseUsageStatistic) CalculateAge13to17CountAvg(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	return c.calculateAgeCountAvg(tx, 13, util.PointerInt(17))
}

func (c *courseUsageStatistic) CalculateAge18to24CountAvg(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	return c.calculateAgeCountAvg(tx, 18, util.PointerInt(24))
}

func (c *courseUsageStatistic) CalculateAge25to34CountAvg(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	return c.calculateAgeCountAvg(tx, 25, util.PointerInt(34))
}

func (c *courseUsageStatistic) CalculateAge35to44CountAvg(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	return c.calculateAgeCountAvg(tx, 35, util.PointerInt(44))
}

func (c *courseUsageStatistic) CalculateAge45to54CountAvg(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	return c.calculateAgeCountAvg(tx, 45, util.PointerInt(54))
}

func (c *courseUsageStatistic) CalculateAge55to64CountAvg(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	return c.calculateAgeCountAvg(tx, 55, util.PointerInt(64))
}

func (c *courseUsageStatistic) CalculateAge65UpCountAvg(tx *gorm.DB) ([]*model.CourseUsageStatisticResult, error) {
	return c.calculateAgeCountAvg(tx, 65, nil)
}

func (c *courseUsageStatistic) calculateAgeCountAvg(tx *gorm.DB, start int, end *int) ([]*model.CourseUsageStatisticResult, error) {
	query := "1=1 "
	params := make([]interface{}, 0)
	//加入 start 篩選條件
	query += "AND TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) >= ? "
	params = append(params, start)
	//加入 end 篩選條件
	if end != nil {
		query += "AND TIMESTAMPDIFF(YEAR, users.birthday, CURDATE()) <= ? "
		params = append(params, *end)
	}
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
		Joins("INNER JOIN users ON t.user_id = users.id").
		Where(query, params...).
		Group("t.course_id").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (c *courseUsageStatistic) Save(tx *gorm.DB, ColumnName string, values []*model.CourseUsageStatisticResult) error {
	if len(values) == 0 {
		return nil
	}
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	params := make([]map[string]interface{}, 0)
	for _, v := range values {
		params = append(params, map[string]interface{}{"course_id": v.CourseID, ColumnName: v.Value})
	}
	if err := db.Table("course_usage_statistics").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "course_id"}},
		DoUpdates: clause.AssignmentColumns([]string{ColumnName, "update_at"}),
	}).Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func (c *courseUsageStatistic) FindCourseUsageStatisticOutput(courseID int64, output interface{}) error {
	if err := c.gorm.DB().
		Model(&entity.CourseUsageStatistic{}).
		Where("course_id = ?", courseID).
		First(output).Error; err != nil {
		return err
	}
	return nil
}
