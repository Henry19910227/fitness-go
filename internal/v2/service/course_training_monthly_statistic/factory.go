package course_training_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/course_training_monthly_statistic"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := course_training_monthly_statistic.New(db)
	return New(repository)
}
