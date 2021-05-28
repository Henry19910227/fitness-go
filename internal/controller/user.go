package controller

import (
	"github.com/Henry19910227/fitness-go/internal/dto/userdto"
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type user struct {
	Base
	userService service.User
}

func NewUser(baseGroup *gin.RouterGroup, userService service.User, userMiddleware gin.HandlerFunc) {
	user := &user{userService: userService}

	userGroup := baseGroup.Group("/user")
	userGroup.Use(userMiddleware)
	userGroup.PATCH("/info", user.UpdateUserInfo)
	userGroup.GET("/info", user.GetUserInfo)
	userGroup.POST("/role/trainer", user.CreateTrainer)
	userGroup.GET("/role/trainer", user.GetTrainerInfo)
}

// UpdateUserInfo 更新個人資訊
// @Summary 更新個人資訊
// @Description 更新個人資訊
// @Tags User
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param json_body body validator.UpdateUserInfoBody true "更新欄位"
// @Success 200 {object} model.SuccessResult{data=userdto.User} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /user/info [PATCH]
func (u *user) UpdateUserInfo(c *gin.Context)  {
	var header validator.TokenHeader
	var body validator.UpdateUserInfoBody
	if err := c.ShouldBindHeader(&header); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	//更新個人資訊
	user, err := u.userService.UpdateUserByToken(c, header.Token, &userdto.UpdateUserParam{
		Nickname: body.Nickname,
		Sex: body.Sex,
		Birthday: body.Birthday,
		Height: body.Height,
		Weight: body.Weight,
		Experience: body.Experience,
		Target: body.Target,
	})
	if err != nil {
		u.JSONErrorResponse(c, err)
		return
	}
	u.JSONSuccessResponse(c, user, "update success!")
}

// GetUserInfo 獲取個人資訊
// @Summary 獲取個人資訊
// @Description 獲取個人資訊
// @Tags User
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Success 200 {object} model.SuccessResult{data=userdto.User} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /user/info [GET]
func (u *user) GetUserInfo(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	user, err := u.userService.GetUserByToken(c, header.Token)
	if err != nil {
		u.JSONErrorResponse(c, err)
		return
	}
	u.JSONSuccessResponse(c, user, "success!")
}

// CreateTrainer 創建我的教練身份
// @Summary 創建我的教練身份
// @Description 創建我的教練身份
// @Tags User
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param json_body body validator.CreateTrainerBody true "輸入欄位"
// @Success 200 {object} model.SuccessResult "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /user/role/trainer [POST]
func (u *user) CreateTrainer(c *gin.Context)  {
	var header validator.TokenHeader
	var body validator.CreateTrainerBody
	if err := c.ShouldBindHeader(&header); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result , err := u.userService.CreateTrainerByToken(c, header.Token, &userdto.CreateTrainerParam{
		Name: body.Name,
		Nickname: body.Nickname,
		Phone: body.Phone,
		Email: body.Email,
	})
	if err != nil {
		u.JSONErrorResponse(c, err)
		return
	}
	u.JSONSuccessResponse(c, result, "create success!")
}

// GetTrainerInfo 取得我的教練身份資訊
// @Summary 取得我的教練身份資訊
// @Description 取得我的教練身份資訊
// @Tags User
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Success 200 {object} model.SuccessResult{data=userdto.Trainer} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /user/role/trainer [GET]
func (u *user) GetTrainerInfo(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := u.userService.GetTrainerInfoByToken(c, header.Token)
	if err != nil {
		u.JSONErrorResponse(c, err)
		return
	}
	u.JSONSuccessResponse(c, result, "success!")
}