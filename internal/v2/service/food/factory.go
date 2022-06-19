package food

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/food"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := food.New(db)
	return New(repository)
}
