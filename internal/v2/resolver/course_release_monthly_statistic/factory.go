package course_release_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/course_release_monthly_statistic"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	statisticService := course_release_monthly_statistic.NewService(db)
	return New(statisticService)
}
