package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	workoutSetService "github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
)

func NewResolver(gormTool tool.Gorm) Resolver {
	workoutSetSVC := workoutSetService.NewService(gormTool)
	return New(workoutSetSVC)
}
