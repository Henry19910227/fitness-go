package dto

import "github.com/Henry19910227/fitness-go/internal/model"

type WorkoutLog struct {
	ID             int64            `json:"id" example:"1"`                          // 紀錄id
	Duration       int              `json:"duration" example:"1"`                    // 訓練時長
	Intensity      int              `json:"intensity" example:"1"`                   // 訓練強度(0:未指定/1:輕鬆/2:適中/3:稍難/4:很累)
	Place          int              `json:"place" example:"1"`                       // 地點(0:未指定/1:住家/2:健身房/3:戶外)
	Course         *CourseCustom1   `json:"course"`                                  // 課表
	Workout        *Workout         `json:"workout"`                                 // 訓練
	WorkoutSetLogs []*WorkoutSetLog `json:"workout_set_logs"`                        // 訓練組
	CreateAt       string           `json:"create_at" example:"2021-05-28 11:00:00"` // 創建時間
}

type WorkoutLogSummary struct {
	ID        int64    `json:"id" example:"1"`                          // 紀錄id
	Duration  int      `json:"duration" example:"1"`                    // 訓練時長
	Intensity int      `json:"intensity" example:"1"`                   // 訓練強度(0:未指定/1:輕鬆/2:適中/3:稍難/4:很累)
	Place     int      `json:"place" example:"1"`                       // 地點(0:未指定/1:住家/2:健身房/3:戶外)
	Workout   *Workout `json:"workout"`                                 // 訓練
	CreateAt  string   `json:"create_at" example:"2021-05-28 11:00:00"` // 創建時間
}

func NewWorkoutLog(log *model.WorkoutLog, logSets []*model.WorkoutSetLog) WorkoutLog {
	workoutLog := WorkoutLog{
		Duration:  log.Duration,
		Intensity: log.Intensity,
		Place:     log.Place,
		CreateAt:  log.CreateAt,
	}
	if log.Workout != nil {
		workoutLog.Workout = &Workout{
			ID:              log.Workout.ID,
			Name:            log.Workout.Name,
			Equipment:       log.Workout.Equipment,
			StartAudio:      log.Workout.StartAudio,
			EndAudio:        log.Workout.EndAudio,
			WorkoutSetCount: log.Workout.WorkoutSetCount,
		}
	}
	workoutSetLogs := make([]*WorkoutSetLog, 0)
	for _, set := range logSets {
		logSet := NewWorkoutSetLog(set)
		workoutSetLogs = append(workoutSetLogs, &logSet)
	}
	workoutLog.WorkoutSetLogs = workoutSetLogs
	return workoutLog
}

func NewWorkoutLogSummary(data *model.WorkoutLog) WorkoutLogSummary {
	workoutLog := WorkoutLogSummary{
		ID:        data.ID,
		Duration:  data.Duration,
		Intensity: data.Intensity,
		Place:     data.Place,
		CreateAt:  data.CreateAt,
	}
	if data.Workout != nil {
		workoutLog.Workout = &Workout{
			ID:              data.Workout.ID,
			Name:            data.Workout.Name,
			Equipment:       data.Workout.Equipment,
			StartAudio:      data.Workout.StartAudio,
			EndAudio:        data.Workout.EndAudio,
			WorkoutSetCount: data.Workout.WorkoutSetCount,
		}
	}
	return workoutLog
}

type WorkoutLogParam struct {
	Duration       int
	Intensity      *int
	Place          *int
	WorkoutSetLogs []*WorkoutSetLogParam
}
