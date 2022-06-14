package service

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/gin-gonic/gin"
)

type workoutSetLog struct {
	workoutSetLogRepo repository.WorkoutSetLog
	errHandler        errcode.Handler
}

func NewWorkoutSetLog(workoutSetLogRepo repository.WorkoutSetLog, errHandler errcode.Handler) WorkoutSetLog {
	return &workoutSetLog{workoutSetLogRepo: workoutSetLogRepo, errHandler: errHandler}
}

func (w *workoutSetLog) GetWorkoutSetLogSummaries(c *gin.Context, userID int64, actionID int64, startDate string, endDate string) ([]*dto.WorkoutSetLogSummary, errcode.Error) {
	datas, err := w.workoutSetLogRepo.FindWorkoutSetLogsByDate(userID, actionID, startDate, endDate)
	if err != nil {
		return nil, w.errHandler.Set(c, "workout set log repo", err)
	}
	logs := make([]*dto.WorkoutSetLogSummary, 0)
	for _, data := range datas {
		log := dto.NewWorkoutSetLogSummary(data)
		logs = append(logs, log)
	}
	return logs, nil
}
