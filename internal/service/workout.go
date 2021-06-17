package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/workoutdto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type workout struct {
	workoutRepo repository.Workout
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewWorkout(workoutRepo repository.Workout, logger handler.Logger, jwtTool tool.JWT, errHandler errcode.Handler) Workout {
	return &workout{workoutRepo: workoutRepo, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (w *workout) CreateWorkout(c *gin.Context, planID int64, name string) (*workoutdto.WorkoutID, errcode.Error) {
	workoutID, err := w.workoutRepo.CreateWorkout(planID, name)
	if err != nil {
		w.logger.Set(c, handler.Error, "CourseRepo", w.errHandler.SystemError().Code(), err.Error())
		return nil, w.errHandler.SystemError()
	}
	return &workoutdto.WorkoutID{ID: workoutID}, nil
}