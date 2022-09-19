package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	courseService := course.NewService(db)
	return New(courseService)
}
