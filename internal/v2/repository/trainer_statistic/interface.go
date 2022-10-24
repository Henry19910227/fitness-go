package trainer_statistic

import (
	model "github.com/Henry19910227/fitness-go/internal/v2/model/trainer_statistic"
	"gorm.io/gorm"
)

type Repository interface {
	WithTrx(tx *gorm.DB) Repository
	StatisticReviewScore(input *model.StatisticReviewScoreInput) (err error)
}
