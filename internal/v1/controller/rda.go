package controller

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/middleware"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
)

type RDA struct {
	Base
	rdaService service.RDA
}

func NewRDA(baseGroup *gin.RouterGroup, rdaService service.RDA, userMiddleware middleware.User) {
	rda := RDA{rdaService: rdaService}
	baseGroup.POST("/calculate_rda",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		rda.CalculateRDA)
	baseGroup.PUT("/rda",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		rda.UpdateRDA)
}

// CalculateRDA 飲食計算機獲取建議飲食攝取量(Recommended Dietary Allowances)
// @Summary 飲食計算機獲取建議飲食攝取量(Recommended Dietary Allowances)
// @Description 飲食計算機獲取建議飲食攝取量(Recommended Dietary Allowances)
// @Tags Diet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CalculateRDABody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.RDA} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/calculate_rda [POST]
func (r *RDA) CalculateRDA(c *gin.Context) {
	var body validator.CalculateRDABody
	if err := c.ShouldBindJSON(&body); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result := r.rdaService.CalculateRDA(&dto.CalculateRDAParam{
		Sex:              body.Sex,
		Birthday:         body.Birthday,
		Height:           body.Height,
		Weight:           body.Weight,
		BodyFat:          body.BodyFat,
		ActivityLevel:    body.ActivityLevel,
		ExerciseFeqLevel: body.ExerciseFeqLevel,
		DietTarget:       body.DietTarget,
		DietType:         body.DietType,
	})
	r.JSONSuccessResponse(c, result, "success!")
}

// UpdateRDA 更新建議飲食攝取量
// @Summary 更新建議飲食攝取量 (API已經過時，更新為 /v2/rda [PUT])
// @Description 更新建議飲食攝取量
// @Tags Diet_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.UpdateRDABody true "輸入參數"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/rda [PUT]
func (r *RDA) UpdateRDA(c *gin.Context) {
	uid, e := r.GetUID(c)
	if e != nil {
		r.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var body validator.UpdateRDABody
	if err := c.ShouldBindJSON(&body); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := r.rdaService.CreateRDA(c, uid, &dto.RDA{
		TDEE:      body.TDEE,
		Calorie:   body.Calorie,
		Protein:   body.Protein,
		Fat:       body.Fat,
		Carbs:     body.Carbs,
		Grain:     body.Grain,
		Vegetable: body.Vegetable,
		Fruit:     body.Fruit,
		Meat:      body.Meat,
		Dairy:     body.Dairy,
		Nut:       body.Nut,
	}); err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessResponse(c, nil, "success!")
}
