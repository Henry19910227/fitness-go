package trainer_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/trainer_statistic"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	statisticService := trainer_statistic.NewService(db)
	return New(statisticService)
}
