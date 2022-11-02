package card

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/card"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := card.New(db)
	return New(repository)
}

