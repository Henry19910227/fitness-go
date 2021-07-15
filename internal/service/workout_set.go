package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto/workoutdto"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"strings"
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

func (s *set) CreateWorkoutSets(c *gin.Context, workoutID int64, actionIDs []int64) ([]*workoutdto.WorkoutSet, errcode.Error) {
	//創建動作組
	setIDs, err := s.setRepo.CreateWorkoutSetsByWorkoutID(workoutID, actionIDs)
	if err != nil {
		if strings.Contains(err.Error(), "1452") {
			s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.ActionNotExist().Code(), err.Error())
			return nil, s.errHandler.ActionNotExist()
		}
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	datas, err := s.setRepo.FindWorkoutSetsByIDs(setIDs)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	return parserWorkoutSets(datas), nil
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

func (s *set) GetWorkoutSets(c *gin.Context, workoutID int64) ([]*workoutdto.WorkoutSet, errcode.Error) {
	datas, err := s.setRepo.FindWorkoutSetsByWorkoutID(workoutID)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	return parserWorkoutSets(datas), nil
}

func (s *set) UpdateWorkoutSet(c *gin.Context, setID int64, param *workoutdto.UpdateWorkoutSetParam) (*workoutdto.WorkoutSet, errcode.Error) {
	if err := s.setRepo.UpdateWorkoutSetByID(setID, &model.UpdateWorkoutSetParam {
		AutoNext: param.AutoNext,
		StartAudio: param.StartAudio,
		Remark: param.Remark,
		Weight: param.Weight,
		Reps: param.Reps,
		Distance: param.Distance,
		Duration: param.Duration,
		Incline: param.Incline,
	}); err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	data, err := s.setRepo.FindWorkoutSetByID(setID)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	return parserWorkoutSet(data), nil
}

func (s *set) DeleteWorkoutSet(c *gin.Context, setID int64) (*workoutdto.WorkoutSetID, errcode.Error) {
	if err := s.setRepo.DeleteWorkoutSetByID(setID); err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	return &workoutdto.WorkoutSetID{ID: setID}, nil
}

func parserWorkoutSet(data *model.WorkoutSetEntity) *workoutdto.WorkoutSet {
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
	if data.Action != nil {
		action := workoutdto.WorkoutSetAction{
			ID: data.Action.ID,
			Name: data.Action.Name,
			Source: data.Action.Source,
			Type: data.Action.Type,
			Intro: data.Action.Intro,
			Cover: data.Action.Cover,
			Video: data.Action.Video,
		}
		set.Action = &action
	}
	return &set
}

func parserWorkoutSets(datas []*model.WorkoutSetEntity) []*workoutdto.WorkoutSet {
	sets := make([]*workoutdto.WorkoutSet, 0)
	for _, data := range datas {
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
		if data.Action != nil {
			action := workoutdto.WorkoutSetAction{
				ID: data.Action.ID,
				Name: data.Action.Name,
				Source: data.Action.Source,
				Type: data.Action.Type,
				Intro: data.Action.Intro,
				Cover: data.Action.Cover,
				Video: data.Action.Video,
			}
			set.Action = &action
		}
		sets = append(sets, &set)
	}
	return sets
}