package repository

import (
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type course struct {
	gorm  tool.Gorm
}

func NewCourse(gorm tool.Gorm) Course {
	return &course{gorm: gorm}
}

func (c *course) CreateCourse(uid int64, param *model.CreateCourseParam) (int64, error) {
	course := model.Course{
		UserID: uid,
		Name: param.Name,
		Level: param.Level,
		Category: param.Category,
		CategoryOther: param.CategoryOther,
		ScheduleType: param.ScheduleType,
	}
	if err := c.gorm.DB().Create(&course).Error; err != nil {
		return 0, err
	}
	return course.ID, nil
}
