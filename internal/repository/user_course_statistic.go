package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
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

func (c *userCourseStatistic) SaveUserCourseStatistic(tx *gorm.DB, param *model.SaveUserCourseStatisticParam) (int64, error) {
	db := c.gorm.DB()
	if tx != nil {
		db = tx
	}
	courseStatistic := entity.UserCourseStatistic{
		UserID:            param.UserID,
		CourseID:          param.CourseID,
		Duration:          param.Duration,
		WorkoutCourt:      param.WorkoutCourt,
		TotalWorkoutCourt: param.TotalWorkoutCourt,
		UpdateAt:          time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := db.Create(&courseStatistic).Error; err != nil {
		return 0, err
	}
	return courseStatistic.ID, nil
}
