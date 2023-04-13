package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course_training_avg_statistic"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	courseService := course.NewService(db)
	statisticService := course_training_avg_statistic.NewService(db)
	return New(courseService, statisticService)
}
