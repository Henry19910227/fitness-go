package trainer_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_statistic"
	"gorm.io/gorm"
)

type Service interface {
	Tx(tx *gorm.DB) Service
	StatisticReviewScore(input *model.StatisticReviewScoreInput) (err error)
}
