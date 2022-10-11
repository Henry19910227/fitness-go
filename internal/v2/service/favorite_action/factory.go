package favorite_action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/favorite_action"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := favorite_action.New(db)
	return New(repository)
}
