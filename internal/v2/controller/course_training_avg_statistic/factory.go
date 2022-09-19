package course_training_avg_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/course_training_avg_statistic"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := course_training_avg_statistic.NewResolver(db)
	return New(resolver)
}
