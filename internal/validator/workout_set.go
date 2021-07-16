package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CreateWorkoutSetBody struct {
	ActionIDs []int64 `json:"action_ids" binding:"required,workout_set_action_ids" example:"1,10,15"` // 動作id
}

type WorkoutSetIDUri struct {
	WorkoutSetID int64 `uri:"workout_set_id" binding:"required" example:"1"`
}

type UpdateWorkoutSetBody struct {
	AutoNext *string `json:"auto_next" binding:"omitempty,oneof=Y N" example:"Y"` //自動下一組(Y:是/N:否)
	StartAudio *string `json:"start_audio" binding:"omitempty,max=50" example:"e6d2131w5q.mp3"` //前導語音
	Remark *string `json:"remark" binding:"omitempty,max=40" example:"注意呼吸"` //備註(1~20字元)
	Weight *float64 `json:"weight" binding:"omitempty,min=0.01,max=999.99" example:"10.25"` //重量(0.01~999.99公斤)
	Reps *int `json:"reps" binding:"omitempty,min=1,max=999" example:"10"` //次數(1~999次)
	Distance *float64 `json:"distance" binding:"omitempty,min=0.01,max=999.99" example:"20.25"` //距離(0.01~999.99公里)
	Duration *int `json:"duration" binding:"omitempty,min=1,max=38439" example:"60"` //時長(1~38439秒)
	Incline *float64 `json:"incline" binding:"omitempty,min=0.01,max=999.99" example:"15.5"` //坡度(0.01~999.99)
}

type UpdateWorkoutSetOrderBody struct {
	Orders []WorkoutSetOrder `json:"orders" binding:"required,workout_set_orders"` //訓練組排序
}

type WorkoutSetOrder struct {
	WorkoutSetID int64 `json:"workout_set_id" binding:"omitempty" example:"10"` //訓練組id
	Seq int `json:"seq" binding:"required" example:"1"` //排列序號
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("workout_set_action_ids", WorkoutSetActionIDs)
		_ = v.RegisterValidation("workout_set_orders", WorkoutSetOrders)
	}
}

var WorkoutSetActionIDs validator.Func = func(fl validator.FieldLevel) bool {
	return validateWorkoutSetActionIDs(fl,10)
}

var WorkoutSetOrders validator.Func = func(fl validator.FieldLevel) bool {
	return validateWorkoutSetOrders(fl)
}

func validateWorkoutSetActionIDs(fl validator.FieldLevel, maxCount int) bool {
	actionIDs, ok := fl.Field().Interface().([]int64)
	if !ok {
		return false
	}
	//檢查是否丟空陣列
	if len(actionIDs) == 0 {
		return false
	}
	//檢查個數是否超過上限
	if len(actionIDs) > maxCount {
		return false
	}
	//檢查是否重複，沒重複就把新值加入map備查
	dupMap := make(map[int64]int64)
	for _, item := range actionIDs {
		_, ok := dupMap[item]
		if ok {
			return false
		}
		dupMap[item] = item
	}
	return true
}

func validateWorkoutSetOrders(fl validator.FieldLevel) bool {
	orders, ok := fl.Field().Interface().([]WorkoutSetOrder)
	if !ok {
		return false
	}
	//檢查是否丟空陣列
	if len(orders) == 0 {
		return false
	}
	return true
}