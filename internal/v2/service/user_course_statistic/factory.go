package user_course_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/user_course_statistic"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := user_course_statistic.New(db)
	return New(repository)
}
