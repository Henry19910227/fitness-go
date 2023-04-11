package course_category_training_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course_category_training_monthly_statistic"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	statisticService := course_category_training_monthly_statistic.NewService(db)
	return New(statisticService)
}
