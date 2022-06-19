package diet

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/diet"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	return New(diet.New(db))
}
