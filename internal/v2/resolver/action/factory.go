package action

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/tool/uploader"
	"github.com/Henry19910227/fitness-go/internal/v2/service/action"
	"github.com/Henry19910227/fitness-go/internal/v2/service/course"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout"
	"github.com/Henry19910227/fitness-go/internal/v2/service/workout_set"
	"gorm.io/gorm"
)

func NewResolver(db *gorm.DB) Resolver {
	actionService := action.NewService(db)
	courseService := course.NewService(db)
	workoutService := workout.NewService(db)
	workoutSetService := workout_set.NewService(db)
	coverUploadTool := uploader.NewActionCoverTool()
	videoUploadTool := uploader.NewActionVideoTool()
	return New(actionService, courseService, workoutService, workoutSetService, coverUploadTool, videoUploadTool)
}
