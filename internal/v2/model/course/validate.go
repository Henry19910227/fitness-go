package course

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("course_level", Level)
		_ = v.RegisterValidation("course_category", Category)
		_ = v.RegisterValidation("course_suit", Suit)
		_ = v.RegisterValidation("course_equipment", Equipment)
		_ = v.RegisterValidation("course_place", Place)
		_ = v.RegisterValidation("course_trainer_target", TrainerTarget)
		_ = v.RegisterValidation("course_body_target", BodyTarget)
		_ = v.RegisterValidation("course_sale_type", SaleType)
		_ = v.RegisterValidation("course_trainer_skill", TrainerSkill)
		_ = v.RegisterValidation("course_trainer_sex", TrainerSex)
	}
}

var Level validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 4)
}

var Category validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 6)
}

var Suit validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 10)
}

var Equipment validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 9)
}

var Place validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 5)
}

var TrainerTarget validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 5)
}

var BodyTarget validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 7)
}

var SaleType validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 3)
}

var TrainerSkill validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 14)
}

var TrainerSex validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateStringRangeField(fl, map[string]string{
		"m": "男性",
		"f": "女性",
	}, 2)
}
