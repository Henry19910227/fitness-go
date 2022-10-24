package trainer_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/trainer_statistic"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := trainer_statistic.New(db)
	return New(repository)
}
