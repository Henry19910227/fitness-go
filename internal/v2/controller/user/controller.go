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
// @Tags 用戶_v2
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

// CreateRegisterOTP 發送OTP
// @Summary 發送OTP
// @Description 發送OTP
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APICreateRegisterOTPBody true "輸入參數"
// @Success 200 {object} user.APICreateRegisterOTPOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/otp [POST]
func (c *controller) CreateRegisterOTP(ctx *gin.Context) {
	var input model.APICreateRegisterOTPInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APICreateRegisterOTP(&input)
	ctx.JSON(http.StatusOK, output)
}

// RegisterAccountValidate 驗證帳戶是否可使用
// @Summary 驗證帳戶是否可使用
// @Description 驗證帳戶是否可使用
// @Tags 註冊_v2
// @Accept json
// @Produce json
// @Param json_body body user.APIRegisterAccountValidateBody true "輸入參數"
// @Success 200 {object} user.APIRegisterAccountValidateOutput "成功!"
// @Failure 400 {object} base.Output "失敗!"
// @Router /v2/register/account/validate [POST]
func (c *controller) RegisterAccountValidate(ctx *gin.Context) {
	var input model.APIRegisterAccountValidateInput
	if err := ctx.ShouldBindJSON(&input.Body); err != nil {
		ctx.JSON(http.StatusBadRequest, baseModel.BadRequest(util.PointerString(err.Error())))
		return
	}
	output := c.resolver.APIRegisterAccountValidate(&input)
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
