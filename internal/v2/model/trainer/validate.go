package trainer

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("trainer_skill", Skill)
	}
}

var Skill validator.Func = func(fl validator.FieldLevel) bool {
	return util.ValidateIntRangeField(fl, 1, 14)
}
