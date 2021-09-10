package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto"
	"github.com/Henry19910227/fitness-go/internal/global"
	midd "github.com/Henry19910227/fitness-go/internal/middleware"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Trainer struct {
	Base
	trainerService service.Trainer
}

func NewTrainer(baseGroup *gin.RouterGroup, trainerService service.Trainer, userMiddleware gin.HandlerFunc, userMidd midd.User)  {
	baseGroup.StaticFS("/resource/trainer/avatar", http.Dir("./volumes/storage/trainer/avatar"))
	baseGroup.StaticFS("/resource/trainer/card_front_image", http.Dir("./volumes/storage/trainer/card_front_image"))
	baseGroup.StaticFS("/resource/trainer/card_back_image", http.Dir("./volumes/storage/trainer/card_back_image"))
	trainer := &Trainer{trainerService: trainerService}
	trainerGroup := baseGroup.Group("/trainer")
	trainerGroup.Use(userMiddleware)
	trainerGroup.GET("/info", trainer.GetTrainerInfo)
	trainerGroup.POST("/avatar", trainer.UploadMyTrainerAvatar)

	baseGroup.POST("/trainer",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		trainer.CreateTrainer)

	baseGroup.PATCH("/trainer",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.UserStatusPermission([]global.UserStatus{global.UserActivity}),
		trainer.UpdateTrainer)

	baseGroup.POST("/card_front_image",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerDraft}),
		trainer.UploadCardFrontImage)

	baseGroup.POST("/card_back_image",
		userMidd.TokenPermission([]global.Role{global.UserRole}),
		userMidd.TrainerStatusPermission([]global.TrainerStatus{global.TrainerActivity, global.TrainerDraft}),
		trainer.UploadCardBackImage)
}

// CreateTrainer 創建我的教練身份
// @Summary 創建我的教練身份
// @Description 創建我的教練身份
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.CreateTrainerBody true "輸入欄位"
// @Success 200 {object} model.SuccessResult{data=dto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer [POST]
func (t *Trainer) CreateTrainer(c *gin.Context)  {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var body validator.CreateTrainerBody
	if err := c.ShouldBindJSON(&body); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result , err := t.trainerService.CreateTrainer(c, uid, &dto.CreateTrainerParam{
		Name: body.Name,
		Address: body.Address,
		Phone: body.Phone,
		Email: body.Email,
	})
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, result, "create success!")
}

// UpdateTrainer 修改我的教練資訊
// @Summary 修改我的教練資訊
// @Description 修改我的教練資訊
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.UpdateTrainerBody true "輸入欄位"
// @Success 200 {object} model.SuccessResult{data=dto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /trainer [PATCH]
func (t *Trainer) UpdateTrainer(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	var body validator.UpdateTrainerBody
	if err := c.ShouldBindJSON(&body); err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	trainer, err := t.trainerService.UpdateTrainer(c, uid, &dto.UpdateTrainerParam{
		Name: body.Name,
		Nickname: body.Nickname,
		Email: body.Email,
		Phone: body.Phone,
		Address: body.Address,
		Intro: body.Intro,
		Experience: body.Experience,
		Motto: body.Motto,
		FacebookURL: body.FacebookURL,
		InstagramURL: body.InstagramURL,
		YoutubeURL: body.YoutubeURL,
	})
	if err != nil {
		t.JSONErrorResponse(c, err)
		return
	}
	t.JSONSuccessResponse(c, trainer, "update success!")
}

// GetTrainerInfo 取得我的教練身份資訊
// @Summary 取得我的教練身份資訊
// @Description 取得我的教練身份資訊
// @Tags Trainer
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=dto.Trainer} "成功!"
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
// @Security fitness_token
// @Accept mpfd
// @Param avatar formData file true "教練大頭照"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.TrainerAvatar} "成功!"
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

// UploadCardFrontImage 上傳身分證正面
// @Summary 上傳身分證正面
// @Description 查看身分證正面照 : https://www.fitness-app.tk/api/v1/resource/trainer/card_front_image/{圖片名}
// @Tags Trainer
// @Security fitness_token
// @Accept mpfd
// @Param card_front_image formData file true "身分證正面"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.TrainerCardFront} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /card_front_image [POST]
func (t *Trainer) UploadCardFrontImage(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("card_front_image")
	if err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, er := t.trainerService.UploadCardFrontImageByUID(c, uid, fileHeader.Filename, file)
	if er != nil {
		t.JSONErrorResponse(c, er)
		return
	}
	t.JSONSuccessResponse(c, result, "success upload")
}

// UploadCardBackImage 上傳身分證背面
// @Summary 上傳身分證背面
// @Description 查看身分證背面照 : https://www.fitness-app.tk/api/v1/resource/trainer/card_back_image/{圖片名}
// @Tags Trainer
// @Security fitness_token
// @Accept mpfd
// @Param card_back_image formData file true "身分證背面"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.TrainerCardBack} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /card_back_image [POST]
func (t *Trainer) UploadCardBackImage(c *gin.Context) {
	uid, e := t.GetUID(c)
	if e != nil {
		t.JSONValidatorErrorResponse(c, e.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("card_back_image")
	if err != nil {
		t.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, errs := t.trainerService.UploadCardBackImageByUID(c, uid, fileHeader.Filename, file)
	if errs != nil {
		t.JSONErrorResponse(c, errs)
		return
	}
	t.JSONSuccessResponse(c, result, "success upload")
}