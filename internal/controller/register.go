package controller

import (
	"github.com/Henry19910227/fitness-go/internal/service"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
)

type Register struct {
	Base
	regService service.Register
}

func NewRegister(baseGroup *gin.RouterGroup, regService service.Register)  {
	register := &Register{regService: regService}
	baseGroup.POST("/register/email/otp", register.SendEmailOTP)
	baseGroup.POST("/register/email", register.RegisterForEmail)
	baseGroup.POST("/register/nickname/validate", register.ValidateNicknameDuplicate)
	baseGroup.POST("/register/email/validate", register.ValidateEmailDuplicate)
}

// SendEmailOTP 發送 Email OTP
// @Summary 發送 Email OTP
// @Description 發送 Email OTP
// @Tags Register
// @Accept json
// @Produce json
// @Param json_body Parameter body validator.EmailBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=registerdto.OTP} "驗證郵件已寄出"
// @Failure 400 {object} model.ErrorResult "發送失敗"
// @Router /register/email/otp [POST]
func (r *Register) SendEmailOTP(c *gin.Context)  {
	var body validator.EmailBody
	if err := c.ShouldBindJSON(&body); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	result, err := r.regService.SendEmailOTP(c, body.Email)
	if err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessResponse(c, result, "驗證郵件已寄出!")
}

// RegisterForEmail 使用信箱註冊
// @Summary 使用信箱註冊
// @Description 使用信箱註冊
// @Tags Register
// @Accept json
// @Produce json
// @Param json_body body validator.RegisterForEmailBody true "輸入參數"
// @Success 200 {object} model.SuccessResult{data=registerdto.Register} "註冊成功"
// @Failure 400 {object} model.ErrorResult "註冊失敗"
// @Router /register/email [POST]
func (r *Register) RegisterForEmail(c *gin.Context)  {
	var body validator.RegisterForEmailBody
	if err := c.ShouldBindJSON(&body); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	//檢查Email是否重複
	if err := r.regService.ValidateEmailDup(c, body.Email); err != nil {
		r.JSONErrorResponse(c, err)
	}
	//檢查暱稱是否重複
	if err := r.regService.ValidateNicknameDup(c, body.Email); err != nil {
		r.JSONErrorResponse(c, err)
	}
	result, err := r.regService.EmailRegister(c, body.OTPCode, body.Email, body.Nickname, body.Password)
	if err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessResponse(c, result, "註冊成功!")
}

// ValidateNicknameDuplicate 驗證暱稱是否可使用
// @Summary 驗證暱稱是否可使用
// @Description 驗證暱稱是否可使用
// @Tags Register
// @Accept json
// @Produce json
// @Param json_body body validator.ValidateNicknameDupBody true "輸入參數"
// @Success 200 {object} model.SuccessResult "此暱稱可使用"
// @Failure 400 {object} model.ErrorResult "該資料已存在"
// @Router /register/nickname/validate [POST]
func (r *Register) ValidateNicknameDuplicate(c *gin.Context) {
	var body validator.ValidateNicknameDupBody
	// 驗證輸入
	if err := c.ShouldBindJSON(&body); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	// 檢查暱稱是否重複
	if err := r.regService.ValidateNicknameDup(c, body.Nickname); err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessResponse(c, nil, "此暱稱可使用")
}

// ValidateEmailDuplicate 驗證信箱是否可使用
// @Summary 驗證信箱是否可使用
// @Description 驗證信箱是否可使用
// @Tags Register
// @Accept json
// @Produce json
// @Param json_body body validator.ValidateEmailDupBody true "輸入參數"
// @Success 200 {object} model.SuccessResult "此暱稱可使用"
// @Failure 400 {object} model.ErrorResult "該資料已存在"
// @Router /register/email/validate [POST]
func (r *Register) ValidateEmailDuplicate(c *gin.Context) {
	var body validator.ValidateEmailDupBody
	// 驗證輸入
	if err := c.ShouldBindJSON(&body); err != nil {
		r.JSONValidatorErrorResponse(c, err.Error())
		return
	}
	// 檢查信箱是否重複
	if err := r.regService.ValidateEmailDup(c, body.Email); err != nil {
		r.JSONErrorResponse(c, err)
		return
	}
	r.JSONSuccessResponse(c, nil, "此信箱可使用")
}