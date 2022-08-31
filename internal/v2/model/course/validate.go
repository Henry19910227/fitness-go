package course

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("suit", Suit)
		_ = v.RegisterValidation("equipment", Equipment)
		_ = v.RegisterValidation("place", Place)
		_ = v.RegisterValidation("train_target", TrainTarget)
		_ = v.RegisterValidation("body_target", BodyTarget)
	}
}

var Suit validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 10, 3)
}

var Equipment validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 9, 3)
}

var Place validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 5, 3)
}

var TrainTarget validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 5, 3)
}

var BodyTarget validator.Func = func(fl validator.FieldLevel) bool {
	return validateCourseFieldByRange(fl, 1, 7, 3)
}

func validateCourseFieldByRange(fl validator.FieldLevel, min int, max int, maxCount int) bool {
	str, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	results := strings.Split(str, ",")
	//檢查是否只傳了空白字串，如果是空白字串就直接返回
	if len(results) == 1 && results[0] == "" {
		return true
	}
	//檢查個數是否超過上限
	if len(results) > maxCount {
		return false
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
