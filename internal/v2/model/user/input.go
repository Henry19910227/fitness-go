package user

type GenerateInput struct {
	DataAmount int
}

type FindInput struct {
	IDOptional
}

// APIUpdatePasswordInput /v2/password [PATCH]
type APIUpdatePasswordInput struct {
	IDRequired
	Body APIUpdatePasswordBody
}
type APIUpdatePasswordBody struct {
	OldPassword string `json:"old_password" binding:"required,min=6,max=18" example:"12345678"` // 舊密碼 (6~18字元)
	PasswordRequired
}

// APIRegisterEmailInput /v2/register/email [POST]
type APIRegisterEmailInput struct {
	Body APIRegisterEmailBody
}
type APIRegisterEmailBody struct {
	EmailRequired
	PasswordRequired
	NicknameRequired
	OTPCode string `json:"otp_code" binding:"required,max=16" example:"531476"` // 信箱驗證碼
}
