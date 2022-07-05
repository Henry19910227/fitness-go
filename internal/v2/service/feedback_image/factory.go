package feedback_image

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/feedback_image"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := feedback_image.New(db)
	return New(repository)
}
