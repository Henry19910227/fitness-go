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

func (w *workoutSetLog) FindWorkoutSetLogsByWorkoutLogID(tx *gorm.DB, workoutLogID int64) ([]*model.WorkoutSetLog, error) {
	db := w.gorm.DB()
	if tx != nil {
		db = tx
	}
	workoutSetLogs := make([]*model.WorkoutSetLog, 0)
	if err := db.
		Preload("WorkoutSet.Action").
		Find(&workoutSetLogs, "workout_log_id = ?", workoutLogID).Error; err != nil {
		return nil, err
	}
	return workoutSetLogs, nil
}

func (w *workoutSetLog) FindWorkoutSetLogsByWorkoutSetIDs(userID int64, workoutSetIDs []int64) ([]*model.WorkoutSetLog, error) {
	workoutSetLogs := make([]*model.WorkoutSetLog, 0)
	if err := w.gorm.DB().
		Preload("WorkoutSet.Action").
		Joins("INNER JOIN workout_logs ON workout_set_logs.workout_log_id = workout_logs.id").
		Where("workout_logs.user_id = ? AND workout_set_logs.workout_set_id IN (?)", userID, workoutSetIDs).
		Find(&workoutSetLogs).Error; err != nil {
		return nil, err
	}
	return workoutSetLogs, nil
}

func (w *workoutSetLog) FindWorkoutSetLogsByDate(userID int64, actionID int64, startDate string, endDate string) ([]*model.WorkoutSetLogSummary, error) {
	workoutSetLogs := make([]*model.WorkoutSetLogSummary, 0)
	if err := w.gorm.DB().Table("workout_set_logs").
		Joins("INNER JOIN workout_logs ON workout_set_logs.workout_log_id = workout_logs.id").
		Joins("INNER JOIN workout_sets ON workout_set_logs.workout_set_id = workout_sets.id").
		Where("workout_logs.user_id = ? AND workout_sets.action_id = ? AND workout_set_logs.create_at BETWEEN ? AND ?", userID, actionID, startDate, endDate).
		Find(&workoutSetLogs).Error; err != nil {
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

func (w *workoutSetLog) CalculateBestWorkoutSetLog(tx *gorm.DB, userID int64, actionIDs []int64) ([]*model.BestActionSetLog, error) {
	db := w.gorm.DB()
	if tx != nil {
		db = tx
	}
	var logs []*model.BestActionSetLog
	if err := db.
		Table("actions").
		Select("MAX(actions.id) AS action_id",
			"MAX(workout_set_logs.weight) AS weight", "MAX(workout_set_logs.reps) AS reps",
			"MAX(workout_set_logs.distance) AS distance", "MAX(workout_set_logs.duration) AS duration",
			"MAX(workout_set_logs.incline) AS incline").
		Joins("INNER JOIN workout_sets ON workout_sets.action_id = actions.id").
		Joins("INNER JOIN workout_set_logs ON workout_set_logs.workout_set_id = workout_sets.id").
		Joins("INNER JOIN workout_logs ON workout_logs.id = workout_set_logs.workout_log_id").
		Where("actions.id IN (?) AND actions.source = ? AND workout_logs.user_id = ?", actionIDs, 1, userID).
		Group("actions.id").
		Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
