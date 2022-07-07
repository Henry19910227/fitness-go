package receipt

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/receipt"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := receipt.New(db)
	return New(repository)
}
