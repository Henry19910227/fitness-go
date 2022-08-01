package user

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/util"
	baseModel "github.com/Henry19910227/fitness-go/internal/v2/model/base"
	model "github.com/Henry19910227/fitness-go/internal/v2/model/user"
	"github.com/Henry19910227/fitness-go/internal/v2/resolver/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	resolver user.Resolver
}

func New(resolver user.Resolver) Controller {
	return &controller{resolver: resolver}
}

// UpdatePassword 修改密碼
// @Summary 修改密碼
// @Description 修改密碼
// @Tags 用戶個人_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body user.APIUpdatePasswordBody true "輸入參數"
// @Success 200 {object} user.APIUpdatePasswordOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/password [PATCH]
func (c *controller) UpdatePassword(ctx *gin.Context) {
	var input model.APIUpdatePasswordInput
	input.ID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdatePassword(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateUserProfile 修改用戶個人資訊
// @Summary 修改用戶個人資訊
// @Description 修改用戶個人資訊
// @Tags 用戶個人_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Param json_body body user.APIUpdateUserProfileBody true "輸入參數"
// @Success 200 {object} user.APIUpdateUserProfileOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/profile [PATCH]
func (c *controller) UpdateUserProfile(ctx *gin.Context) {
	var input model.APIUpdateUserProfileInput
	input.ID = ctx.MustGet("uid").(int64)
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIUpdateUserProfile(&input)
	ctx.JSON(http.StatusOK, output)
}

// UpdateUserAvatar 更新用戶個人大頭照
// @Summary 更新用戶個人大頭照
// @Description 更新用戶個人大頭照 : {Base URL}/v2/resource/user/avatar/{Filename}
// @Tags 用戶個人_v2
// @Security fitness_token
// @Accept mpfd
// @Param avatar formData file true "用戶大頭照"
// @Produce json
// @Success 200 {object} user.APIUpdateUserAvatarOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/avatar [PATCH]
func (c *controller) UpdateUserAvatar(ctx *gin.Context) {
	file, fileHeader, err := ctx.Request.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	input := model.APIUpdateUserAvatarInput{}
	input.ID = ctx.MustGet("uid").(int64)
	input.CoverNamed = fileHeader.Filename
	input.File = file
	output := c.resolver.APIUpdateUserAvatar(&input)
	ctx.JSON(http.StatusOK, output)
}

// GetUserProfile 獲取用戶個人資訊
// @Summary 獲取用戶個人資訊
// @Description 用於個人設定頁面，獲取個人資訊與帳號資訊
// @Tags 用戶個人_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} user.APIGetUserProfileOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/user/profile [GET]
func (c *controller) GetUserProfile(ctx *gin.Context) {
	var input model.APIGetUserProfileInput
	input.ID = ctx.MustGet("uid").(int64)
	output := c.resolver.APIGetUserProfile(&input)
	ctx.JSON(http.StatusOK, output)
}

// LoginForEmail 使用信箱登入
// @Summary 使用信箱登入
// @Description 使用信箱登入
// @Tags 登入_v2
// @Accept json
// @Produce json
// @Param json_body body user.APILoginForEmailBody true "輸入參數"
// @Success 200 {object} user.APILoginForEmailOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/login/email [POST]
func (c *controller) LoginForEmail(ctx *gin.Context) {
	var input model.APILoginForEmailInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APILoginForEmail(&input)
	ctx.JSON(http.StatusOK, output)
}

// LoginForFacebook 使用Facebook登入
// @Summary 使用Facebook登入
// @Description 使用Facebook登入
// @Tags 登入_v2
// @Accept json
// @Produce json
// @Param json_body body user.APILoginForFacebookBody true "輸入參數"
// @Success 200 {object} user.APILoginForFacebookOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/login/facebook [POST]
func (c *controller) LoginForFacebook(ctx *gin.Context) {
	var input model.APILoginForFacebookInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APILoginForFacebook(&input)
	ctx.JSON(http.StatusOK, output)
}

// LoginForGoogle 使用Google登入
// @Summary 使用Google登入
// @Description 使用Google登入
// @Tags 登入_v2
// @Accept json
// @Produce json
// @Param json_body body user.APILoginForGoogleBody true "輸入參數"
// @Success 200 {object} user.APILoginForGoogleOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/login/google [POST]
func (c *controller) LoginForGoogle(ctx *gin.Context) {
	var input model.APILoginForGoogleInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APILoginForGoogle(&input)
	ctx.JSON(http.StatusOK, output)
}

// LoginForLine 使用Line登入
// @Summary 使用Line登入
// @Description 使用Line登入
// @Tags 登入_v2
// @Accept json
// @Produce json
// @Param json_body body user.APILoginForLineBody true "輸入參數"
// @Success 200 {object} user.APILoginForLineOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/login/line [POST]
func (c *controller) LoginForLine(ctx *gin.Context) {
	var input model.APILoginForLineInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APILoginForLine(&input)
	ctx.JSON(http.StatusOK, output)
}

// LoginForApple 使用Apple登入
// @Summary 使用Apple登入
// @Description 使用Apple登入
// @Tags 登入_v2
// @Accept json
// @Produce json
// @Param json_body body user.APILoginForAppleBody true "輸入參數"
// @Success 200 {object} user.APILoginForAppleOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/login/apple [POST]
func (c *controller) LoginForApple(ctx *gin.Context) {
	var input model.APILoginForAppleInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APILoginForApple(&input)
	ctx.JSON(http.StatusOK, output)
}

// Logout 登出
// @Summary 登出
// @Description 登出
// @Tags 登入_v2
// @Accept json
// @Produce json
// @Security fitness_token
// @Success 200 {object} user.APILogoutOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/logout [POST]
func (c *controller) Logout(ctx *gin.Context) {
	var input model.APILogoutInput
	input.ID = ctx.MustGet("uid").(int64)
	output := c.resolver.APILogout(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterForEmail 使用信箱註冊
// @Summary 使用信箱註冊
// @Description 使用信箱註冊
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterForEmailBody true "輸入參數"
// @Success 200 {object} user.APIRegisterForEmailOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/email [POST]
func (c *controller) RegisterForEmail(ctx *gin.Context) {
	var input model.APIRegisterForEmailInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterForEmail(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterForFacebook 使用Facebook註冊
// @Summary 使用Facebook註冊
// @Description 使用Facebook註冊
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterForFacebookBody true "輸入參數"
// @Success 200 {object} user.APIRegisterForFacebookOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/facebook [POST]
func (c *controller) RegisterForFacebook(ctx *gin.Context) {
	var input model.APIRegisterForFacebookInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterForFacebook(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterForGoogle 使用Google註冊
// @Summary 使用Google註冊
// @Description 使用Google註冊
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterForGoogleBody true "輸入參數"
// @Success 200 {object} user.APIRegisterForGoogleOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/google [POST]
func (c *controller) RegisterForGoogle(ctx *gin.Context) {
	var input model.APIRegisterForGoogleInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterForGoogle(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterForLine 使用Line註冊
// @Summary 使用Line註冊
// @Description 使用Line註冊
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterForLineBody true "輸入參數"
// @Success 200 {object} user.APIRegisterForLineOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/line [POST]
func (c *controller) RegisterForLine(ctx *gin.Context) {
	var input model.APIRegisterForLineInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterForLine(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterForApple 使用Apple註冊
// @Summary 使用Apple註冊
// @Description 使用Apple註冊
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterForAppleBody true "輸入參數"
// @Success 200 {object} user.APIRegisterForAppleOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/apple [POST]
func (c *controller) RegisterForApple(ctx *gin.Context) {
	var input model.APIRegisterForAppleInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterForApple(&input)
	ctx.JSON(http.StatusOK, output)
}

// CreateRegisterOTP 發送OTP
// @Summary 發送OTP
// @Description 發送OTP
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APICreateOTPBody true "輸入參數"
// @Success 200 {object} user.APICreateRegisterOTPOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/otp [POST]
func (c *controller) CreateRegisterOTP(ctx *gin.Context) {
	var input model.APICreateOTPInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateRegisterOTP(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterEmailAccountValidate 驗證Email帳戶是否可使用
// @Summary 驗證Email帳戶是否可使用
// @Description 驗證Email帳戶是否可使用
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterEmailAccountValidateBody true "輸入參數"
// @Success 200 {object} user.APIRegisterEmailAccountValidateOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/email_account/validate [POST]
func (c *controller) RegisterEmailAccountValidate(ctx *gin.Context) {
	var input model.APIRegisterEmailAccountValidateInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterEmailAccountValidate(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterFacebookAccountValidate 驗證facebook帳戶是否可使用
// @Summary 驗證facebook帳戶是否可使用
// @Description 驗證facebook帳戶是否可使用
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterFacebookAccountValidateBody true "輸入參數"
// @Success 200 {object} user.APIRegisterFacebookAccountValidateOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/facebook_account/validate [POST]
func (c *controller) RegisterFacebookAccountValidate(ctx *gin.Context) {
	var input model.APIRegisterFacebookAccountValidateInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterFacebookAccountValidate(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterLineAccountValidate 驗證line帳戶是否可使用
// @Summary 驗證line帳戶是否可使用
// @Description 驗證line帳戶是否可使用
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterLineAccountValidateBody true "輸入參數"
// @Success 200 {object} user.APIRegisterLineAccountValidateOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/line_account/validate [POST]
func (c *controller) RegisterLineAccountValidate(ctx *gin.Context) {
	var input model.APIRegisterLineAccountValidateInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterLineAccountValidate(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterGoogleAccountValidate 驗證Google帳戶是否可使用
// @Summary 驗證Google帳戶是否可使用
// @Description 驗證Google帳戶是否可使用
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterGoogleAccountValidateBody true "輸入參數"
// @Success 200 {object} user.APIRegisterGoogleAccountValidateOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/google_account/validate [POST]
func (c *controller) RegisterGoogleAccountValidate(ctx *gin.Context) {
	var input model.APIRegisterGoogleAccountValidateInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterGoogleAccountValidate(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterNicknameValidate 驗證暱稱是否可使用
// @Summary 驗證暱稱是否可使用
// @Description 驗證暱稱是否可使用
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterNicknameValidateBody true "輸入參數"
// @Success 200 {object} user.APIRegisterNicknameValidateOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/nickname/validate [POST]
func (c *controller) RegisterNicknameValidate(ctx *gin.Context) {
	var input model.APIRegisterNicknameValidateInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterNicknameValidate(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterEmailValidate 驗證Email是否可使用
// @Summary 驗證Email是否可使用
// @Description 驗證Email是否可使用
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterEmailValidateBody true "輸入參數"
// @Success 200 {object} user.APIRegisterEmailValidateOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/email/validate [POST]
func (c *controller) RegisterEmailValidate(ctx *gin.Context) {
	var input model.APIRegisterEmailValidateInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterEmailValidate(&input)
	ctx.JSON(http.StatusOK, output)
}
