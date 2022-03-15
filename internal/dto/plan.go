package dto

import "github.com/Henry19910227/fitness-go/internal/model"

type PlanID struct {
	ID int64 `json:"plan_id" example:"1"` //計畫id
}

type Plan struct {
	ID           int64  `json:"id" example:"1"`             //計畫id
	Name         string `json:"name" example:"第一週增肌計畫"`     //計畫名稱
	WorkoutCount int    `json:"workout_count" example:"10"` //包含訓練數量
}

type PlanStructure struct {
	ID           int64               `json:"id" example:"1"`             //計畫id
	Name         string              `json:"name" example:"第一週增肌計畫"`     //計畫名稱
	WorkoutCount int                 `json:"workout_count" example:"10"` //訓練數量
	Workouts     []*WorkoutStructure `json:"workouts"`                   //訓練列表
}

type PlanAsset struct {
	ID                 int64  `json:"id" example:"1"`                   //計畫id
	Name               string `json:"name" example:"第一週增肌計畫"`           //計畫名稱
	WorkoutCount       int    `json:"workout_count" example:"10"`       //訓練數量
	FinishWorkoutCount int    `json:"finish_workout_count" example:"5"` //完成訓練數量
}

type PlanAssetStructure struct {
	ID                 int64                    `json:"id" example:"1"`                   //計畫id
	Name               string                   `json:"name" example:"第一週增肌計畫"`           //計畫名稱
	WorkoutCount       int                      `json:"workout_count" example:"10"`       //訓練數量
	FinishWorkoutCount int                      `json:"finish_workout_count" example:"5"` //完成訓練數量
	Workouts           []*WorkoutAssetStructure `json:"workouts"`                         //訓練列表
}

func NewPlanStructure(data *model.Plan) PlanStructure {
	plan := PlanStructure{
		ID:           data.ID,
		Name:         data.Name,
		WorkoutCount: data.WorkoutCount,
	}
	plan.Workouts = make([]*WorkoutStructure, 0)
	return plan
}

func NewPlanAssetStructure(data *model.PlanAsset) PlanAssetStructure {
	plan := PlanAssetStructure{
		ID:                 data.ID,
		Name:               data.Name,
		WorkoutCount:       data.WorkoutCount,
		FinishWorkoutCount: data.FinishWorkoutCount,
	}
	plan.Workouts = make([]*WorkoutAssetStructure, 0)
	return plan
}
