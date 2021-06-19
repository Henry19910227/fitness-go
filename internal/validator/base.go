package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

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

type TokenHeader struct {
	Token string `header:"Token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTQ0MDc3NjMsInN1YiI6IjEwMDEzIn0.Z5UeEC8ArCVYej9kI1paXD2f5FMFiTfeLpU6e_CZZw0"`
}

type UserIDUri struct {
	UserID *int64 `uri:"user_id" binding:"required" example:"10001"`
}

type CourseIDUri struct {
	CourseID int64 `uri:"course_id" binding:"required" example:"1"`
}

type PagingQuery struct {
	Page *int `form:"page" binding:"omitempty,min=1" example:"1"`
	Size *int `form:"size" binding:"omitempty,min=1" example:"5"`
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("suit", Suit)
		_ = v.RegisterValidation("equipment", Equipment)
		_ = v.RegisterValidation("place", Place)
		_ = v.RegisterValidation("train_target", TrainTarget)
		_ = v.RegisterValidation("body_target", BodyTarget)
	}
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