package review_image

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/review_image"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := review_image.New(db)
	return New(repository)
}
