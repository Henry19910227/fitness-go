package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/repository/workout_set"
)

func NewService(gormTool tool.Gorm) Service {
	repository := workoutSet.New(gormTool)
	return New(repository)
}
