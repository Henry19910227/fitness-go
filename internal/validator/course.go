package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var suit validator.Func = func(fl validator.FieldLevel) bool {
	str, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if str != "1,2,3" {
		return false
	}
	return true
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("suit", suit)
	}
}

type CreateCourseBody struct {
	Name string `json:"course_id" binding:"required,min=1,max=20" example:"Henry課表"`        // 課表名稱(1~20字元)
	Level int   `json:"level" binding:"required,oneof=1 2 3 4" example:"1"`           // 強度(1:初級/2:中級/3:中高級/4:高級)
	Category int `json:"category" binding:"required,oneof=1 2 3 4 5 6" example:"1"`   // 課表類別(1:有氧心肺訓練/2:間歇肌力訓練/3:重量訓練/4:阻力訓練/5:徒手訓練/6:其他)
	CategoryOther string `json:"category_other" binding:"omitempty,max=20" example:"其他訓練"` // 課表其他類別名稱(最大20字元)
	ScheduleType int `json:"schedule_type" binding:"required,oneof=1 2" example:"1"` // 排課類別(1:單一訓練/2:多項計畫)
}

type UpdateCourseBody struct {
	suit string `json:"course_id" binding:"omitempty,suit,max=5" example:"2,5,7"` //適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
}
