package feedback

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/feedback"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := feedback.New(db)
	return New(repository)
}
