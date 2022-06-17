package meal

import (
	repository "github.com/Henry19910227/fitness-go/internal/v2/repository/meal"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	return New(repository.New(db))
}
