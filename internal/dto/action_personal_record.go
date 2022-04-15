package dto

import "github.com/Henry19910227/fitness-go/internal/model"

type ActionPR struct {
	ActionID int64   `json:"action_id" example:"12"`
	Weight   float64 `json:"weight" example:"100"`  //重量(公斤)
	Reps     int     `json:"reps" example:"10"`     //次數
	Distance float64 `json:"distance" example:"5"`  //距離(公里)
	Duration int     `json:"duration" example:"60"` //時長(秒)
	Incline  float64 `json:"incline" example:"7"`   //坡度
}

func NewActionPR(data *model.ActionPR) ActionPR {
	pr := ActionPR{}
	if data != nil {
		pr.ActionID = data.ActionID
		pr.Weight = data.Weight
		pr.Reps = data.Reps
		pr.Distance = data.Distance
		pr.Duration = data.Duration
		pr.Incline = data.Incline
	}
	return pr
}
