package dto

import "github.com/Henry19910227/fitness-go/internal/v1/model"

type WorkoutSetID struct {
	ID int64 `json:"workout_set_id" example:"10"` //訓練組id
}

type WorkoutSet struct {
	ID            int64   `json:"id" example:"10"`                          //訓練組id
	Type          int     `json:"type" example:"2"`                         //動作類別(1:動作/2:休息)
	Action        *Action `json:"action"`                                   //動作
	AutoNext      string  `json:"auto_next" example:"N"`                    //自動下一組(Y:是/N:否)
	StartAudio    string  `json:"start_audio" example:"1d2w3e51d3w.mp3"`    //前導語音
	ProgressAudio string  `json:"progress_audio" example:"1d2w3e51d3w.mp3"` //進行中語音
	Remark        string  `json:"remark" example:"ˋ注意姿勢"`                   //備註
	Weight        float64 `json:"weight" example:"0"`                       //重量(公斤)
	Reps          int     `json:"reps" example:"0"`                         //次數
	Distance      float64 `json:"distance" example:"0"`                     //距離(公里)
	Duration      int     `json:"duration" example:"30"`                    //時長(秒)
	Incline       float64 `json:"incline" example:"0"`                      //坡度
}

func NewWorkoutSet(data *model.WorkoutSet) WorkoutSet {
	set := WorkoutSet{
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
		action := Action{
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
			Favorite:  data.Action.Favorite,
		}
		set.Action = &action
	}
	return set
}

type UpdateWorkoutSetParam struct {
	AutoNext   *string  `gorm:"column:auto_next"`   //自動下一組(Y:是/N:否)
	StartAudio *string  `gorm:"column:start_audio"` //前導語音
	Remark     *string  `gorm:"column:remark"`      //備註
	Weight     *float64 `gorm:"column:weight"`      //重量(公斤)
	Reps       *int     `gorm:"column:reps"`        //次數
	Distance   *float64 `gorm:"column:distance"`    //距離(公里)
	Duration   *int     `gorm:"column:duration"`    //時長(秒)
	Incline    *float64 `gorm:"column:incline"`     //坡度
}

type WorkoutSetOrder struct {
	WorkoutSetID int64 `gorm:"column:workout_set_id"` //訓練組id
	Seq          int   `gorm:"column:seq"`            //排列序號
}
