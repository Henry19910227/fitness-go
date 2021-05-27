package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

var Suit validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 10)
}

var Equipment validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 9)
}

var Place validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 5)
}

var TrainTarget validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 5)
}

var BodyTarget validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 7)
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("suit", Suit)
		v.RegisterValidation("equipment", Equipment)
		v.RegisterValidation("place", Place)
		v.RegisterValidation("train_target", TrainTarget)
		v.RegisterValidation("body_target", BodyTarget)
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
	Suit string `json:"suit" binding:"omitempty,suit,max=5" example:"2,5,7"` // 適用對象(1:女性/2:男性/3:初學者/4:進階者/5:專業/6:長輩/7:運動員/8:孕婦/9:產後/10:其他)
	Equipment string `json:"equipment" binding:"omitempty,equipment,max=5" example:"1,2,5"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Place string `json:"place" binding:"omitempty,place,max=5" example:"1,2,3"` // 適合場地(1:健身房/2:居家/3:空地/4:戶外/5:其他)
	TrainTarget string `json:"train_target" binding:"omitempty,train_target,max=5" example:"1,2,4"` // 訓練目的(1:減脂/2:增肌/3:維持健康/4:鐵人三項/5:其他)
	BodyTarget string `json:"body_target" binding:"omitempty,body_target,max=5" example:"1,2,6"` // 體態目標(1:比基尼身材/2:翹臀/3:健力/4:健美/5:腹肌/6:馬甲線/7:其他)
}

func validateCourseFieldByRange(fl validator.FieldLevel, min int, max int) bool  {
	str, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	results := strings.Split(str, ",")
	var tmp = 0
	for _, item := range results {
		//將string轉換為int
		value, err := strconv.Atoi(item)
		if err != nil {
			return false
		}
		//檢查選項是否在範圍內
		if value < min || value > max {
			return false
		}
		//檢查選項是否由小到大排列 (必須 > tmp)
		if value < tmp {
			return false
		}
		tmp = value
	}
	return true
}
