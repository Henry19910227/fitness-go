package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/handler"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strings"
	"time"
)

type set struct {
	setRepo    repository.WorkoutSet
	uploader   handler.Uploader
	resHandler handler.Resource
	logger     handler.Logger
	jwtTool    tool.JWT
	errHandler errcode.Handler
}

func NewWorkoutSet(setRepo repository.WorkoutSet,
	uploader handler.Uploader,
	resHandler handler.Resource,
	logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) WorkoutSet {
	return &set{setRepo: setRepo, uploader: uploader, resHandler: resHandler, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (s *set) CreateWorkoutSets(c *gin.Context, workoutID int64, actionIDs []int64) ([]*dto.WorkoutSet, errcode.Error) {
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

func (s *set) CreateRestSet(c *gin.Context, workoutID int64) (*dto.WorkoutSet, errcode.Error) {
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
	return parserWorkoutSet(data), nil
}

func (s *set) DuplicateWorkoutSets(c *gin.Context, setID int64, count int) ([]*dto.WorkoutSet, errcode.Error) {
	data, err := s.setRepo.FindWorkoutSetByID(setID)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	sets := make([]*entity.WorkoutSet, 0)
	for i := 0; i < count; i++ {
		set := entity.WorkoutSet{
			WorkoutID:     data.WorkoutID,
			ActionID:      &data.Action.ID,
			Type:          data.Type,
			AutoNext:      data.AutoNext,
			StartAudio:    data.StartAudio,
			ProgressAudio: data.ProgressAudio,
			Remark:        data.Remark,
			Weight:        data.Weight,
			Reps:          data.Reps,
			Distance:      data.Distance,
			Duration:      data.Duration,
			Incline:       data.Incline,
			CreateAt:      time.Now().Format("2006-01-02 15:04:05"),
			UpdateAt:      time.Now().Format("2006-01-02 15:04:05"),
		}
		sets = append(sets, &set)
	}
	setIDs, err := s.setRepo.CreateWorkoutSetsByWorkoutIDAndSets(sets[0].WorkoutID, sets)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	entities, err := s.setRepo.FindWorkoutSetsByIDs(setIDs)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	return parserWorkoutSets(entities), nil
}

func (s *set) GetWorkoutSets(c *gin.Context, workoutID int64, userID *int64) ([]*dto.WorkoutSet, errcode.Error) {
	datas, err := s.setRepo.FindWorkoutSetsByWorkoutID(workoutID, userID)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	return parserWorkoutSets(datas), nil
}

func (s *set) UpdateWorkoutSet(c *gin.Context, setID int64, param *dto.UpdateWorkoutSetParam) (*dto.WorkoutSet, errcode.Error) {
	if err := s.setRepo.UpdateWorkoutSetByID(setID, &model.UpdateWorkoutSetParam{
		AutoNext:   param.AutoNext,
		StartAudio: param.StartAudio,
		Remark:     param.Remark,
		Weight:     param.Weight,
		Reps:       param.Reps,
		Distance:   param.Distance,
		Duration:   param.Duration,
		Incline:    param.Incline,
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

func (s *set) DeleteWorkoutSet(c *gin.Context, setID int64) (*dto.WorkoutSetID, errcode.Error) {
	if err := s.setRepo.DeleteWorkoutSetByID(setID); err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	return &dto.WorkoutSetID{ID: setID}, nil
}

func (s *set) UpdateWorkoutSetOrders(c *gin.Context, workoutID int64, params []*dto.WorkoutSetOrder) errcode.Error {
	var models []*model.WorkoutSetOrder
	for _, data := range params {
		model := model.WorkoutSetOrder{
			WorkoutID:    workoutID,
			WorkoutSetID: data.WorkoutSetID,
			Seq:          data.Seq,
		}
		models = append(models, &model)
	}
	if err := s.setRepo.UpdateWorkoutSetOrdersByWorkoutID(workoutID, models); err != nil {
		//檢測到不存在此課表的訓練組
		if strings.Contains(err.Error(), "1452") {
			return s.errHandler.DataNotFound()
		}
		//插入多個重複的組與相同的序號
		if strings.Contains(err.Error(), "1062") {
			return s.errHandler.DataAlreadyExists()
		}
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return s.errHandler.SystemError()
	}
	return nil
}

func (s *set) UploadWorkoutSetStartAudio(c *gin.Context, setID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error) {
	newAudioNamed, err := s.uploader.UploadWorkoutSetStartAudio(file, audioNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, s.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, s.errHandler.FileSizeError()
		}
		s.logger.Set(c, handler.Error, "Resource Handler", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	if err := s.setRepo.UpdateWorkoutSetByID(setID, &model.UpdateWorkoutSetParam{
		StartAudio: &newAudioNamed,
	}); err != nil {
		s.logger.Set(c, handler.Error, "WorkoutRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	return &dto.WorkoutAudio{Named: newAudioNamed}, nil
}

func (s *set) UploadWorkoutSetProgressAudio(c *gin.Context, setID int64, audioNamed string, file multipart.File) (*dto.WorkoutAudio, errcode.Error) {
	newAudioNamed, err := s.uploader.UploadWorkoutSetProgressAudio(file, audioNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, s.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, s.errHandler.FileSizeError()
		}
		s.logger.Set(c, handler.Error, "Resource Handler", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	if err := s.setRepo.UpdateWorkoutSetByID(setID, &model.UpdateWorkoutSetParam{
		ProgressAudio: &newAudioNamed,
	}); err != nil {
		s.logger.Set(c, handler.Error, "WorkoutRepo", s.errHandler.SystemError().Code(), err.Error())
		return nil, s.errHandler.SystemError()
	}
	return &dto.WorkoutAudio{Named: newAudioNamed}, nil
}

func (s *set) DeleteWorkoutSetStartAudio(c *gin.Context, setID int64) errcode.Error {
	set, err := s.setRepo.FindWorkoutSetByID(setID)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return s.errHandler.SystemError()
	}
	//移除workout_set與start_audio檔案關聯
	var startAudio = ""
	if err := s.setRepo.UpdateWorkoutSetByID(setID, &model.UpdateWorkoutSetParam{
		StartAudio: &startAudio,
	}); err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return s.errHandler.SystemError()
	}
	count, err := s.setRepo.FindStartAudioCountByAudioName(set.StartAudio)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return s.errHandler.SystemError()
	}
	if count > 0 {
		return nil
	}
	//移除start_audio檔案
	if err := s.resHandler.DeleteWorkoutSetStartAudio(set.StartAudio); err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
	}
	return nil
}

func (s *set) DeleteWorkoutSetProgressAudio(c *gin.Context, setID int64) errcode.Error {
	set, err := s.setRepo.FindWorkoutSetByID(setID)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return s.errHandler.SystemError()
	}
	var progressAudio = ""
	if err := s.setRepo.UpdateWorkoutSetByID(setID, &model.UpdateWorkoutSetParam{
		ProgressAudio: &progressAudio,
	}); err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return s.errHandler.SystemError()
	}
	count, err := s.setRepo.FindProgressAudioCountByAudioName(set.ProgressAudio)
	if err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
		return s.errHandler.SystemError()
	}
	if count > 0 {
		return nil
	}
	if err := s.resHandler.DeleteWorkoutSetProgressAudio(set.ProgressAudio); err != nil {
		s.logger.Set(c, handler.Error, "WorkoutSetRepo", s.errHandler.SystemError().Code(), err.Error())
	}
	return nil
}

func parserWorkoutSet(data *model.WorkoutSet) *dto.WorkoutSet {
	set := dto.WorkoutSet{
		ID:            data.ID,
		Type:          data.Type,
		AutoNext:      data.AutoNext,
		StartAudio:    data.StartAudio,
		ProgressAudio: data.ProgressAudio,
		Remark:        data.Remark,
		Weight:        data.Weight,
		Reps:          data.Reps,
		Distance:      data.Distance,
		Duration:      data.Duration,
		Incline:       data.Incline,
	}
	if data.Action != nil {
		action := dto.Action{
			ID:        data.Action.ID,
			Name:      data.Action.Name,
			Source:    data.Action.Source,
			Type:      data.Action.Type,
			Category:  data.Action.Category,
			Body:      data.Action.Body,
			Equipment: data.Action.Equipment,
			Intro:     data.Action.Intro,
			Cover:     data.Action.Cover,
			Video:     data.Action.Video,
		}
		set.Action = &action
	}
	return &set
}

func parserWorkoutSets(datas []*model.WorkoutSet) []*dto.WorkoutSet {
	sets := make([]*dto.WorkoutSet, 0)
	for _, data := range datas {
		set := dto.NewWorkoutSet(data)
		sets = append(sets, &set)
	}
	return sets
}
