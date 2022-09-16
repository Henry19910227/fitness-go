package course_create_monthly_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course_create_monthly_statistic"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := course_create_monthly_statistic.NewResolver(db)
	return New(resolver)
}
