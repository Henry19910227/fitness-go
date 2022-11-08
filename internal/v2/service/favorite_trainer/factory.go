package favorite_trainer

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/favorite_trainer"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := favorite_trainer.New(db)
	return New(repository)
}
