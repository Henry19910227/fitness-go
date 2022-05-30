package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type RDA struct {
	Base
	rdaService service.RDA
}

func NewRDA(baseGroup *gin.RouterGroup, rdaService service.RDA, userMiddleware middleware.User)  {
	rda := RDA{rdaService:  rdaService}
	baseGroup.POST("/calculate_rda",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		rda.CalculateRDA)
}

// CalculateRDA 飲食計算機獲取建議飲食攝取量(Recommended Dietary Allowances)
// @Summary 飲食計算機獲取建議飲食攝取量(Recommended Dietary Allowances)
// @Description 飲食計算機獲取建議飲食攝取量(Recommended Dietary Allowances)
// @Tags Diet
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CalculateRDABody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.RDA} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /calculate_rda [POST]
func (r *RDA) CalculateRDA(c *gin.Context) {
	var body validator.CalculateRDABody
	if err := c.ShouldBindJSON(&body); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result := r.rdaService.CalculateRDA(&dto.CalculateRDAParam{
		Sex: body.Sex,
		Birthday: body.Birthday,
		Height: body.Height,
		Weight: body.Weight,
		BodyFat: body.BodyFat,
		ActivityLevel: body.ActivityLevel,
		ExerciseFeq: body.ExerciseFeq,
		Target: body.Target,
		DietType: body.DietType,
	})
	r.JSONSuccessResponse(c, result, "success!")
}
