package action

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("action_source", ActionSource)
		_ = v.RegisterValidation("action_category", ActionCategory)
		_ = v.RegisterValidation("action_body", ActionBody)
		_ = v.RegisterValidation("action_equipment", ActionEquipment)
	}
}

var ActionSource validator.Func = func(fl validator.FieldLevel) bool {
	return validateActionFieldByRange(fl, 1, 3)
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
	}
	return true
}
