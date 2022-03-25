package repository

import (
	"github.com/Henry19910227/fitness-go/internal/entity"
	"github.com/Henry19910227/fitness-go/internal/model"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"gorm.io/gorm"
	"time"
)

type workoutSetLog struct {
	gorm tool.Gorm
}

func NewWorkoutSetLog(gorm tool.Gorm) WorkoutSetLog {
	return &workoutSetLog{gorm: gorm}
}

func (w *workoutSetLog) FindWorkoutSetLogsByWorkoutLogID(workoutLogID int64) ([]*model.WorkoutSetLog, error) {
	workoutSetLogs := make([]*model.WorkoutSetLog, 0)
	if err := w.gorm.DB().
		Preload("WorkoutSet.Action").
		Find(&workoutSetLogs, "workout_log_id = ?", workoutLogID).Error; err != nil {
		return nil, err
	}
	return workoutSetLogs, nil
}

func (w *workoutSetLog) CreateWorkoutSetLogs(tx *gorm.DB, params []*model.WorkoutSetLogParam) error {
	db := w.gorm.DB()
	if tx != nil {
		db = tx
	}
	workoutSetLogs := make([]*entity.WorkoutSetLog, 0)
	for _, param := range params {
		workoutSetLog := entity.WorkoutSetLog{
			WorkoutLogID: param.WorkoutLogID,
			WorkoutSetID: param.WorkoutSetID,
			Weight:       param.Weight,
			Reps:         param.Reps,
			Distance:     param.Distance,
			Duration:     param.Duration,
			Incline:      param.Incline,
			CreateAt:     time.Now().Format("2006-01-02 15:04:05"),
		}
		workoutSetLogs = append(workoutSetLogs, &workoutSetLog)
	}
	if err := db.Create(&workoutSetLogs).Error; err != nil {
		return err
	}
	return nil
}
