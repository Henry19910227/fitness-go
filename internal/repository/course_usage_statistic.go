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

func (c *courseUsageStatistic) SaveTotalFinishWorkoutCount(tx *gorm.DB, values []*model.CourseUsageStatisticResult) error {
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	params := make([]map[string]interface{}, 0)
	for _, v := range values{
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
