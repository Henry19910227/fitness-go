package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CreateWorkoutSetBody struct {
	ActionIDs []int64 `json:"action_ids" binding:"required,workout_set_action_ids" example:"1,10,15"` // 動作id
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("workout_set_action_ids", WorkoutSetActionIDs)
	}
}

var WorkoutSetActionIDs validator.Func = func(fl validator.FieldLevel) bool {
	return validateWorkoutSetActionIDs(fl,10)
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