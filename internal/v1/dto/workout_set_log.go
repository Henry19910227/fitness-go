package dto

import "github.com/Henry19910227/fitness-go/internal/v1/model"

type WorkoutSetLog struct {
	ID         int64       `json:"id"`                    //訓練組紀錄id
	Name       string      `json:"name"`                  //訓練名稱
	Weight     float64     `json:"weight" example:"10"`   //重量(公斤)
	Reps       int         `json:"reps" example:"5"`      //次數
	Distance   float64     `json:"distance" example:"1"`  //距離(公里)
	Duration   int         `json:"duration" example:"30"` //時長(秒)
	Incline    float64     `json:"incline" example:"5"`   //坡度
	WorkoutSet *WorkoutSet `json:"workout_set"`           // 訓練組
}

type WorkoutSetLogTag struct {
	ID         int64       `json:"id"`                     //訓練組紀錄id
	Weight     float64     `json:"weight" example:"10"`    //重量(公斤)
	Reps       int         `json:"reps" example:"5"`       //次數
	Distance   float64     `json:"distance" example:"1"`   //距離(公里)
	Duration   int         `json:"duration" example:"30"`  //時長(秒)
	Incline    float64     `json:"incline" example:"5"`    //坡度
	WorkoutSet *WorkoutSet `json:"workout_set"`            // 訓練組
	NewRecord  int         `json:"new_record" example:"1"` //是否是新紀錄(0:否/1:是)
}

type WorkoutSetLogSummary struct {
	ID           int64   `json:"id" example:"1"`                          //訓練組紀錄id
	WorkoutLogID int64   `json:"workout_log_id" example:"5"`              // 訓練歷史id
	WorkoutSetID int64   `json:"workout_set_id" example:"15"`             // 訓練組id
	Weight       float64 `json:"weight" example:"10"`                     //重量(公斤)
	Reps         int     `json:"reps" example:"5"`                        //次數
	Distance     float64 `json:"distance" example:"1"`                    //距離(公里)
	Duration     int     `json:"duration" example:"30"`                   //時長(秒)
	Incline      float64 `json:"incline" example:"5"`                     //坡度
	CreateAt     string  `json:"create_at" example:"2021-05-28 11:00:00"` // 新增日期
}

func NewWorkoutSetLog(data *model.WorkoutSetLog) WorkoutSetLog {
	if data == nil {
		return WorkoutSetLog{}
	}
	workoutSetLog := WorkoutSetLog{
		ID:       data.ID,
		Weight:   data.Weight,
		Reps:     data.Reps,
		Distance: data.Distance,
		Duration: data.Duration,
		Incline:  data.Incline,
	}
	if data.WorkoutSet != nil {
		workoutSetLog.WorkoutSet = &WorkoutSet{
			ID:            data.WorkoutSet.ID,
			Type:          data.WorkoutSet.Type,
			AutoNext:      data.WorkoutSet.AutoNext,
			StartAudio:    data.WorkoutSet.StartAudio,
			ProgressAudio: data.WorkoutSet.ProgressAudio,
			Remark:        data.WorkoutSet.Remark,
			Weight:        data.WorkoutSet.Weight,
			Reps:          data.WorkoutSet.Reps,
			Distance:      data.WorkoutSet.Distance,
			Duration:      data.WorkoutSet.Duration,
			Incline:       data.WorkoutSet.Incline,
		}
		if data.WorkoutSet.Action != nil {
			workoutSetLog.WorkoutSet.Action = &Action{
				ID:        data.WorkoutSet.Action.ID,
				Name:      data.WorkoutSet.Action.Name,
				Source:    data.WorkoutSet.Action.Source,
				Type:      data.WorkoutSet.Action.Type,
				Category:  data.WorkoutSet.Action.Category,
				Body:      data.WorkoutSet.Action.Body,
				Equipment: data.WorkoutSet.Action.Equipment,
				Intro:     data.WorkoutSet.Action.Intro,
				Cover:     data.WorkoutSet.Action.Cover,
				Video:     data.WorkoutSet.Action.Video,
			}
		}
	}
	return workoutSetLog
}

func NewWorkoutSetLogTag(data *model.WorkoutSetLog) WorkoutSetLogTag {
	if data == nil {
		return WorkoutSetLogTag{}
	}
	workoutSetLog := WorkoutSetLogTag{
		ID:       data.ID,
		Weight:   data.Weight,
		Reps:     data.Reps,
		Distance: data.Distance,
		Duration: data.Duration,
		Incline:  data.Incline,
	}
	if data.WorkoutSet != nil {
		workoutSetLog.WorkoutSet = &WorkoutSet{
			ID:            data.WorkoutSet.ID,
			Type:          data.WorkoutSet.Type,
			AutoNext:      data.WorkoutSet.AutoNext,
			StartAudio:    data.WorkoutSet.StartAudio,
			ProgressAudio: data.WorkoutSet.ProgressAudio,
			Remark:        data.WorkoutSet.Remark,
			Weight:        data.WorkoutSet.Weight,
			Reps:          data.WorkoutSet.Reps,
			Distance:      data.WorkoutSet.Distance,
			Duration:      data.WorkoutSet.Duration,
			Incline:       data.WorkoutSet.Incline,
		}
		if data.WorkoutSet.Action != nil {
			workoutSetLog.WorkoutSet.Action = &Action{
				ID:        data.WorkoutSet.Action.ID,
				Name:      data.WorkoutSet.Action.Name,
				Source:    data.WorkoutSet.Action.Source,
				Type:      data.WorkoutSet.Action.Type,
				Category:  data.WorkoutSet.Action.Category,
				Body:      data.WorkoutSet.Action.Body,
				Equipment: data.WorkoutSet.Action.Equipment,
				Intro:     data.WorkoutSet.Action.Intro,
				Cover:     data.WorkoutSet.Action.Cover,
				Video:     data.WorkoutSet.Action.Video,
			}
		}
	}
	return workoutSetLog
}

func NewWorkoutSetLogSummary(data *model.WorkoutSetLogSummary) *WorkoutSetLogSummary {
	w := WorkoutSetLogSummary{
		ID:           data.ID,
		WorkoutLogID: data.WorkoutLogID,
		WorkoutSetID: data.WorkoutSetID,
		Weight:       data.Weight,
		Reps:         data.Reps,
		Distance:     data.Distance,
		Duration:     data.Duration,
		Incline:      data.Incline,
		CreateAt:     data.CreateAt,
	}
	return &w
}

type WorkoutSetLogParam struct {
	WorkoutSetID int64   `json:"workout_set_id" example:"1"` //訓練組id
	Weight       float64 `json:"weight" example:"10"`        //重量(公斤)
	Reps         int     `json:"reps" example:"5"`           //次數
	Distance     float64 `json:"distance" example:"1"`       //距離(公里)
	Duration     int     `json:"duration" example:"30"`      //時長(秒)
	Incline      float64 `json:"incline" example:"5"`        //坡度
}
