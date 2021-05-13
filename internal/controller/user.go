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
	userGroup.PATCH("/my/info", user.UpdateMyUserInfo)
}

// UpdateMyUserInfo 更新個人資訊
// @Summary 更新個人資料
// @Description 更新個人資料
// @Tags User
// @Accept json
// @Produce json
// @Security fitness_user_token
// @Param json_body body validator.UpdateMyUserInfoBody true "更新欄位"
// @Success 200 {object} model.SuccessResult{data=userdto.User} "成功!"
// @Failure 400 {object} model.ErrorResult "失敗!"
// @Router /user/my/info [PATCH]
func (u *user) UpdateMyUserInfo(c *gin.Context)  {
	var header validator.TokenHeader
	var body validator.UpdateMyUserInfoBody
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
		//Email: body.Email,
		//Nickname: body.Nickname,
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