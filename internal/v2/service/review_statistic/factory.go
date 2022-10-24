package review_statistic

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/review_statistic"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := review_statistic.New(db)
	return New(repository)
}