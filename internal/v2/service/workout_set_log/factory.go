package workout_set_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/repository/workout_set_log"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB) Service {
	repository := workout_set_log.New(db)
	return New(repository)
}
