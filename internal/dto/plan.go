package dto

type PlanID struct {
	ID int64 `json:"plan_id" example:"1"` //計畫id
}

type Plan struct {
	ID int64 `json:"id" example:"1"` //計畫id
	Name string `json:"name" example:"第一週增肌計畫"` //計畫名稱
	WorkoutCount int `json:"workout_count" example:"10"` //包含訓練數量
}
