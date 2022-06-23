package action

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/action"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := action.New(db)
	return New(repository)
}
