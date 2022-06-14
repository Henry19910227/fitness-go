package workout_set

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	workoutSet "github.com/Henry19910227/fitness-go/internal/v2/resolver/workout_set"
)

func NewController(gormTool tool.Gorm) Controller {
	resolver := workoutSet.NewResolver(gormTool)
	return New(resolver)
}
