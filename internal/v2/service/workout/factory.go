package workout

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/workout"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := workout.New(db)
	return New(repository)
}
