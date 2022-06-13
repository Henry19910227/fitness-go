package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

type ActionIDUri struct {
	ActionID int64 `uri:"action_id" binding:"required" example:"1"`
}

type CreateActionForm struct {
	Name string `form:"name" binding:"required,min=1,max=20" example:"槓鈴臥推"` //動作名稱(1~20字元)
	Type int `form:"type" binding:"required,oneof=1 2 3 4 5" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Category int `form:"category" binding:"required,oneof=1 2 3 4 5" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body int `form:"body" binding:"required,oneof=1 2 3 4 5 6 7 8" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment int `form:"equipment" binding:"required,oneof=1 2 3 4 5 6 7 8" example:"1"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Intro string `form:"intro" binding:"required,min=1,max=400" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹(1~400字元)
}

type UpdateActionForm struct {
	Name *string `form:"name" binding:"omitempty,min=1,max=20" example:"槓鈴臥推"` //動作名稱(1~20字元)
	Category *int `form:"category" binding:"omitempty,oneof=1 2 3 4 5" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body *int `form:"body" binding:"omitempty,oneof=1 2 3 4 5 6 7 8" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment *int `form:"equipment" binding:"omitempty,oneof=1 2 3 4 5 6 7 8" example:"1"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	Intro *string `form:"intro" binding:"omitempty,min=1,max=400" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹(1~400字元)
}

type SearchActionsQuery struct {
	Name *string `form:"name" binding:"omitempty,min=1,max=20" example:"槓鈴臥推"` //動作名稱(1~20字元)
	Source *string `form:"source" binding:"omitempty,action_source" example:"1,2"` //動作來源(1:平台動作/2:教練動作)
	Category *string `form:"category" binding:"omitempty,action_category" example:"1,2,3,4,5"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body *string `form:"body" binding:"omitempty,action_body" example:"2,4,6"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment *string `form:"equipment" binding:"omitempty,action_equipment" example:"1,3,5"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}

type GetWorkoutSetLogsQuery struct {
	StartDate string `form:"start_date" binding:"required,datetime=2006-01-02" example:"區間開始日期"` //區間開始日期
	EndDate string `form:"end_date" binding:"required,datetime=2006-01-02" example:"區間結束日期"` //區間開始日期)
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("action_source", ActionSource)
		_ = v.RegisterValidation("action_category", ActionCategory)
		_ = v.RegisterValidation("action_body", ActionBody)
		_ = v.RegisterValidation("action_equipment", ActionEquipment)
	}
}

var ActionSource validator.Func = func(fl validator.FieldLevel) bool {
	return validateActionFieldByRange(fl, 1, 2)
}

var ActionCategory validator.Func = func(fl validator.FieldLevel) bool {
	return validateActionFieldByRange(fl, 1, 5)
}

var ActionBody validator.Func = func(fl validator.FieldLevel) bool {
	return validateActionFieldByRange(fl, 1, 9)
}

var ActionEquipment validator.Func = func(fl validator.FieldLevel) bool {
	return validateActionFieldByRange(fl, 1, 9)
}


func validateActionFieldByRange(fl validator.FieldLevel, min int, max int) bool {
	str, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	results := strings.Split(str, ",")
	//檢查是否只傳了空白字串，如果是空白字串就直接返回
	if len(results) == 1 && results[0] == "" {
		return true
	}
	var tmp = 0
	dupMap := make(map[string]string)
	for _, item := range results {
		//檢查是否重複，沒重複就把新值加入map備查
		_, ok := dupMap[item]
		if ok {
			return false
		}
		dupMap[item] = item
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