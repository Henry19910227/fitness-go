package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/v2/service/action"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	workoutSetService := workout_set.NewService(db)
	workoutService := workout.NewService(db)
	actionService := action.NewService(db)
	return New(workoutSetService, workoutService, actionService)
}
