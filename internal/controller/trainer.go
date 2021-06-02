package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto/trainerdto"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Trainer struct {
	Base
	trainerService service.Trainer
}

func NewTrainer(baseGroup *gin.RouterGroup, trainerService service.Trainer, userMiddleware gin.HandlerFunc)  {
	baseGroup.StaticFS("/resource/trainer/avatar", http.Dir("./volumes/storage/trainer/avatar"))
	trainer := &Trainer{trainerService: trainerService}
	trainerGroup := baseGroup.Group("/trainer")
	trainerGroup.Use(userMiddleware)
	trainerGroup.POST("", trainer.CreateTrainer)
	trainerGroup.GET("/info", trainer.GetTrainerInfo)
	trainerGroup.POST("/avatar", trainer.UploadMyTrainerAvatar)
}

// CreateTrainer 創建我的教練身份
// @Summary 創建我的教練身份
// @Description 創建我的教練身份
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param json_body body validator.CreateTrainerBody true "輸入欄位"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer [POST]
func (t *Trainer) CreateTrainer(c *gin.Context)  {
	var header validator.TokenHeader
	var body validator.CreateTrainerBody
	if err := c.ShouldBindHeader(&header); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result , err := t.trainerService.CreateTrainerByToken(c, header.Token, &trainerdto.CreateTrainerParam{
		Name: body.Name,
		Nickname: body.Nickname,
		Phone: body.Phone,
		Email: body.Email,
	})
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, result, "create success!")
}

// GetTrainerInfo 取得我的教練身份資訊
// @Summary 取得我的教練身份資訊
// @Description 取得我的教練身份資訊
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Success 200 {object} model.SuccessResult{data=trainerdto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer/info [GET]
func (t *Trainer) GetTrainerInfo(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := t.trainerService.GetTrainerInfoByToken(c, header.Token)
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, result, "success!")
}

// UploadMyTrainerAvatar 上傳我的教練大頭照
// @Summary 上傳我的教練大頭照
// @Description 查看教練大頭照 : https://www.fitness-app.tk/api/v1/resource/trainer/avatar/{圖片名}
// @Tags Trainer
// @Security fitness_user_token
// @Accept mpfd
// @Param avatar formData file true "教練大頭照"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=trainerdto.Avatar} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer/avatar [POST]
func (t *Trainer) UploadMyTrainerAvatar(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("avatar")
	if err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := t.trainerService.UploadTrainerAvatarByToken(c, header.Token, fileHeader.Filename, file)
	if e != nil {
		t.JSONErrorResponse(c, e)
		return
	}
	t.JSONSuccessResponse(c, result, "success upload")
}