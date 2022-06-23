package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/trainer"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := trainer.New(db)
	return New(repository)
}
