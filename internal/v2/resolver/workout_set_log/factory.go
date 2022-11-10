package workout_set_log

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set_log"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	workoutSetLogService := workout_set_log.NewService(db)
	return New(workoutSetLogService)
}
