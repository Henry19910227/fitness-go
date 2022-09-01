package workout

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/plan"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	planService := plan.NewService(db)
	courseService := course.NewService(db)
	workoutService := workout.NewService(db)
	startAudioTool := uploader.NewWorkoutStartAudioTool()
	endAudioTool := uploader.NewWorkoutEndAudioTool()
	return New(workoutService, planService, courseService, startAudioTool, endAudioTool)
}
