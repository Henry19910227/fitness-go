package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/workoutdto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
)

type set struct {
	setRepo repository.WorkoutSet
	uploader handler.Uploader
	logger    handler.Logger
	jwtTool   tool.JWT
	errHandler errcode.Handler
}

func NewWorkoutSet(setRepo repository.WorkoutSet,
	uploader handler.Uploader,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) WorkoutSet {
	return &set{setRepo: setRepo, uploader: uploader, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (s *set) CreateRestSet(c *gin.Context, workoutID int64) (*workoutdto.WorkoutSet, errcode.Error) {
	setID, err := s.setRepo.CreateRestSetByWorkoutID(workoutID)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	data, err := s.setRepo.FindWorkoutSetByID(setID)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	set := workoutdto.WorkoutSet{
		ID: data.ID,
		Type: data.Type,
		AutoNext: data.AutoNext,
		StartAudio: data.StartAudio,
		ProgressAudio: data.ProgressAudio,
		Remark: data.Remark,
		Weight: data.Weight,
		Reps: data.Reps,
		Distance: data.Distance,
		Duration: data.Duration,
		Incline: data.Incline,
	}
	return &set, nil
}
