package review

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/review"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := review.New(db)
	return New(repository)
}
