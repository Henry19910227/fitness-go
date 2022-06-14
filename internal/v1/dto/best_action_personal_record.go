package dto

import "github.com/Henry19910227/fitness-go/internal/v1/model"

type ActionBestPR struct {
	ActionID int64 `json:"action_id" example:"12"`
	MaxRM    struct {
		Value    float64 `json:"value" example:"100"`                     //最大反覆重量(公斤)
		UpdateAt string  `json:"update_at" example:"2021-05-28 11:00:00"` //紀錄更新時間
	} `json:"max_rm"` //最大反覆重量(公斤)
	MaxWeight struct {
		Value    float64 `json:"value" example:"100"`                     //最大重量(公斤)
		UpdateAt string  `json:"update_at" example:"2021-05-28 11:00:00"` //紀錄更新時間
	} `json:"max_weight"` //最大重量(公斤)
	MaxReps struct {
		Value    int    `json:"value" example:"10"`                      //最多次數
		UpdateAt string `json:"update_at" example:"2021-05-28 11:00:00"` //紀錄更新時間
	} `json:"max_reps"` //最多次數
	MinDuration struct {
		Value    int    `json:"value" example:"3600"`                    //最短時長(秒)
		UpdateAt string `json:"update_at" example:"2021-05-28 11:00:00"` //紀錄更新時間
	} `json:"min_duration"` //最短時長
	MaxSpeed struct {
		Value    float64 `json:"value" example:"90"`                      //最高速率
		UpdateAt string  `json:"update_at" example:"2021-05-28 11:00:00"` //紀錄更新時間
	} `json:"max_speed"` //最高速率
	MaxDistance struct {
		Value    float64 `json:"value" example:"15"`                      //最長距離(公里)
		UpdateAt string  `json:"update_at" example:"2021-05-28 11:00:00"` //紀錄更新時間
	} `json:"max_distance"` //最高速率
}

func NewActionBestPR(data *model.ActionBestPR) ActionBestPR {
	pr := ActionBestPR{}
	if data != nil {
		pr.ActionID = data.ActionID
		pr.MaxRM.Value = data.MaxRM
		pr.MaxWeight.Value = data.MaxWeight
		pr.MaxReps.Value = data.MaxReps
		pr.MinDuration.Value = data.MinDuration
		pr.MaxSpeed.Value = data.MaxSpeed
		pr.MaxDistance.Value = data.MaxDistance
		pr.MaxRM.UpdateAt = data.MaxRMUpdateAt
		pr.MaxWeight.UpdateAt = data.MaxWeightUpdateAt
		pr.MaxReps.UpdateAt = data.MaxRepsUpdateAt
		pr.MinDuration.UpdateAt = data.MinDurationUpdateAt
		pr.MaxSpeed.UpdateAt = data.MaxSpeedUpdateAt
		pr.MaxDistance.UpdateAt = data.MaxDistanceUpdateAt
	}
	return pr
}
