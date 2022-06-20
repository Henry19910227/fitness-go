package workout_set

import (
	workoutSetService "github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	workoutSetSVC := workoutSetService.NewService(db)
	return New(workoutSetSVC)
}
