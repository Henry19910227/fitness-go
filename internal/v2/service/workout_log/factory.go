package workout_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/workout_log"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := workout_log.New(db)
	return New(repository)
}
