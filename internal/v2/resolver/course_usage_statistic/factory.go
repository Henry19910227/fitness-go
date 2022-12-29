package course_usage_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course_usage_statistic"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	statisticService := course_usage_statistic.NewService(db)
	return New(statisticService)
}
