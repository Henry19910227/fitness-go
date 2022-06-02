package controller

import (
	"github.com/Henry19910227/fitness-go/internal/global"
	"github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Diet struct {
	Base
	dietService service.Diet
}

func NewDiet(baseGroup *gin.RouterGroup, dietService service.Diet, userMiddleware middleware.User) {
	diet := &Diet{dietService: dietService}
	baseGroup.POST("/diet",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		diet.CreateDiet)

	baseGroup.GET("/diet",
		userMiddleware.TokenPermission([]global.Role{global.UserRole}),
		diet.GetDiet)
}

// CreateDiet 創建飲食紀錄
// @Summary 創建飲食紀錄
// @Description 創建飲食紀錄
// @Tags Diet
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CreateDietBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=dto.Diet} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /diet [POST]
func (d *Diet) CreateDiet(c *gin.Context) {
	uid, e := d.GetUID(c)
	if e != nil {
		d.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var body validator.CreateDietBody
	if err := c.ShouldBindJSON(&body); err != nil {
		d.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	diet, err := d.dietService.CreateDiet(c, uid, body.ScheduleAt)
	if err != nil {
		d.JSONErrorResponse(c, err)
		return
	}
	d.JSONSuccessResponse(c, diet, "success!")
}

// GetDiet 以日期獲取飲食紀錄
// @Summary 以日期獲取飲食紀錄
// @Description 以日期獲取飲食紀錄
// @Tags Diet
// @Accept json
// @Produce json
// @Security fitness_token
// @Param schedule_at query string true "排程日"
// @Success 200 {object} model.SuccessResult{data=dto.Diet} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /diet [GET]
func (d *Diet) GetDiet(c *gin.Context) {
	uid, e := d.GetUID(c)
	if e != nil {
		d.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var query validator.GetDietQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		d.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	diet, err := d.dietService.GetDiet(c, uid, query.ScheduleAt)
	if err != nil {
		d.JSONErrorResponse(c, err)
		return
	}
	d.JSONSuccessResponse(c, diet, "success!")
}
