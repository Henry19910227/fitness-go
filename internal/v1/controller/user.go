package controller

import (
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/service"
	"github.com/Henry19910227/fitness-go/internal/v1/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Base
	userService service.User
}

func NewUser(baseGroup *gin.RouterGroup, userService service.User, userMiddleware gin.HandlerFunc) {
	baseGroup.StaticFS("/resource/user/avatar", http.Dir("./volumes/storage/user/avatar"))
	user := &user{userService: userService}
	userGroup := baseGroup.Group("/user")
	userGroup.Use(userMiddleware)
	userGroup.PATCH("/info", user.UpdateUserInfo)
	userGroup.GET("/info", user.GetUserInfo)
	userGroup.POST("/avatar", user.UploadMyUserAvatar)
}

// UpdateUserInfo 更新個人資訊
// @Summary 更新個人資訊
// @Description 更新個人資訊
// @Tags User_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body validator.UpdateUserInfoBody true "更新欄位"
// @Success 200 {object} model.SuccessResult{data=dto.User} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/user/info [PATCH]
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
	user, err := u.userService.UpdateUserByToken(c, header.Token, &dto.UpdateUserParam{
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
// @Tags User_v1
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} model.SuccessResult{data=dto.User} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/user/info [GET]
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

// UploadMyUserAvatar 上傳我的大頭照
// @Summary 上傳我的大頭照
// @Description 查看我的大頭照 : https://www.fitopia-hub.tk/api/v1/resource/user/avatar/{圖片名}
// @Tags User_v1
// @Security fitness_token
// @Accept mpfd
// @Param avatar formData file true "用戶大頭照"
// @Produce json
// @Success 200 {object} model.SuccessResult{data=dto.UserAvatar} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /v1/user/avatar [POST]
func (u *user) UploadMyUserAvatar(c *gin.Context) {
	var header validator.TokenHeader
	if err := c.ShouldBindHeader(&header); err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	file, fileHeader, err := c.Request.FormFile("avatar")
	if err != nil {
		u.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, e := u.userService.UploadUserAvatarByToken(c, header.Token, fileHeader.Filename, file)
	if e != nil {
		u.JSONErrorResponse(c, e)
		return
	}
	u.JSONSuccessResponse(c, result, "success upload")
}