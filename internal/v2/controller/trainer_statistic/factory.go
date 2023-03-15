package trainer_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/trainer_statistic"
	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	resolver := trainer_statistic.NewResolver(db)
	return New(resolver)
}
