package dto

import "github.com/Henry19910227/fitness-go/internal/model"

type WorkoutSetLog struct {
	ID       int64   `json:"id"`                    //訓練組紀錄id
	Name     string  `json:"name"`                  //訓練名稱
	Weight   float64 `json:"weight" example:"10"`   //重量(公斤)
	Reps     int     `json:"reps" example:"5"`      //次數
	Distance float64 `json:"distance" example:"1"`  //距離(公里)
	Duration int     `json:"duration" example:"30"` //時長(秒)
	Incline  float64 `json:"incline" example:"5"`   //坡度
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
		if data.WorkoutSet.Action != nil {
			workoutSetLog.Name = data.WorkoutSet.Action.Name
		}
	}
	return workoutSetLog
}

type WorkoutSetLogParam struct {
	WorkoutSetID int64   `json:"workout_set_id" example:"1"` //訓練組id
	Weight       float64 `json:"weight" example:"10"`        //重量(公斤)
	Reps         int     `json:"reps" example:"5"`           //次數
	Distance     float64 `json:"distance" example:"1"`       //距離(公里)
	Duration     int     `json:"duration" example:"30"`      //時長(秒)
	Incline      float64 `json:"incline" example:"5"`        //坡度
}
