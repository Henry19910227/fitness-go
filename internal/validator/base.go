package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func validateByIntRange(fl validator.FieldLevel, min int, max int, maxCount int) bool {
	items, ok := fl.Field().Interface().([]int)
	if !ok {
		return false
	}
	//檢查是否丟空陣列
	if len(items) == 0 {
		return false
	}
	//檢查個數是否超過上限
	if len(items) > maxCount {
		return false
	}

	dupMap := make(map[int]int)
	for _, item := range items {
		//檢查是否重複，沒重複就把新值加入map
		_, ok := dupMap[item]
		if ok {
			return false
		}
		//檢查選項是否在範圍內
		if item < min || item > max {
			return false
		}
		dupMap[item] = item
	}
	return true
}

func validateByStringRange(fl validator.FieldLevel, dict map[string]string, maxCount int) bool {
	items, ok := fl.Field().Interface().([]string)
	if !ok {
		return false
	}
	//檢查是否丟空陣列
	if len(items) == 0 {
		return false
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

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("level_inspect", LevelInspect)
		_ = v.RegisterValidation("category_inspect", CourseCategoryInspect)
		_ = v.RegisterValidation("suit_inspect", SuitInspect)
		_ = v.RegisterValidation("equipment_inspect", EquipmentInspect)
		_ = v.RegisterValidation("place_inspect", PlaceInspect)
		_ = v.RegisterValidation("trainer_target_inspect", TrainerTargetInspect)
		_ = v.RegisterValidation("body_target_inspect", BodyTargetInspect)
		_ = v.RegisterValidation("sale_type_inspect", SaleTypeInspect)
		_ = v.RegisterValidation("trainer_skill_inspect", TrainerSkillInspect)
		_ = v.RegisterValidation("trainer_sex_inspect", TrainerSexInspect)
	}
}

var LevelInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByIntRange(fl, 1, 4, 3)
}

var CourseCategoryInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByIntRange(fl, 1, 6, 3)
}

var SuitInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByIntRange(fl, 1, 10, 3)
}

var EquipmentInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByIntRange(fl, 1, 9, 3)
}

var PlaceInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByIntRange(fl, 1, 5, 3)
}

var TrainerTargetInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByIntRange(fl, 1, 5, 3)
}

var BodyTargetInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByIntRange(fl, 1, 7, 3)
}

var SaleTypeInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByIntRange(fl, 1, 3, 3)
}

var TrainerSkillInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByIntRange(fl, 1, 14, 2)
}

var TrainerSexInspect validator.Func = func(fl validator.FieldLevel) bool {
	return validateByStringRange(fl, map[string]string{
		"m": "男性",
		"f": "女性",
	}, 2)
}