package repository

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/entity"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type userCourseStatistic struct {
	gorm tool.Gorm
}

func NewUserCourseStatistic(gorm tool.Gorm) UserCourseStatistic {
	return &userCourseStatistic{gorm: gorm}
}

func (c *userCourseStatistic) FindUserCourseStatistic(userID int64, courseID int64) (*model.UserCourseStatistic, error) {
	var courseStatistic *model.UserCourseStatistic
	if err := c.gorm.DB().
		Where("user_id = ? AND course_id = ?", userID, courseID).
		Take(&courseStatistic).Error; err != nil {
		return nil, err
	}
	return courseStatistic, nil
}

func (c *userCourseStatistic) FindUserCourseStatisticByWorkoutID(workoutID int64, userID int64) (*model.UserCourseStatistic, error) {
	var courseStatistic *model.UserCourseStatistic
	if err := c.gorm.DB().
		Table("user_course_statistics AS cs").
		Joins("INNER JOIN plans ON cs.course_id = plans.course_id").
		Joins("INNER JOIN workouts ON plans.id = workouts.plan_id").
		Where("cs.user_id = ? AND workouts.id = ?", userID, workoutID).
		Take(&courseStatistic).Error; err != nil {
		return nil, err
	}
	return courseStatistic, nil
}

func (c *userCourseStatistic) SaveUserCourseStatistic(tx *gorm.DB, param *model.SaveUserCourseStatisticParam) (int64, error) {
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	courseStatistic := entity.UserCourseStatistic{
		UserID:                  param.UserID,
		CourseID:                param.CourseID,
		Duration:                param.Duration,
		FinishWorkoutCourt:      param.FinishWorkoutCount,
		TotalFinishWorkoutCourt: param.TotalFinishWorkoutCount,
		UpdateAt:                time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "course_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"finish_workout_count", "total_finish_workout_count", "duration", "update_at"}),
	}).Create(&courseStatistic).Error; err != nil {
		return 0, err
	}
	return courseStatistic.ID, nil
}
