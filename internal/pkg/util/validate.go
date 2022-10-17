package util

import (
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

func ValidateIntRangeField(fl validator.FieldLevel, min int, max int) bool {
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

func ValidateStringRangeField(fl validator.FieldLevel, dict map[string]string, maxCount int) bool {
	str, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	items := strings.Split(str, ",")
	//檢查是否只傳了空白字串，如果是空白字串就直接返回
	if len(items) == 1 && items[0] == "" {
		return true
	}
	//檢查個數是否超過上限
	if len(items) > maxCount {
		return false
	}

	dupMap := make(map[string]string)
	for _, item := range items {
		//檢查是否重複，沒重複就把新值加入map
		if _, ok := dupMap[item]; ok {
			return false
		}
		//檢查選項是否在範圍內
		if _, ok := dict[item]; !ok {
			return false
		}
		dupMap[item] = item
	}
	return true
}
