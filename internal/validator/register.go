package validator


type EmailBody struct {
	Email string `json:"email" binding:"required,email" example:"test@gmail.com"`
}

type RegisterForEmailBody struct {
	Email    string `json:"email" binding:"required,email,max=255" example:"test@gmail.com"`     // 信箱 (最大255字元)
	Password string `json:"password" binding:"required,min=6,max=18" example:"12345678"` // 密碼 (6~18字元)
	Nickname string `json:"nickname" binding:"required,min=1,max=20" example:"henry"`    // 暱稱 (1~20字元)
	OTPCode  string `json:"otp_code" binding:"required,max=16" example:"531476"`         // 信箱驗證碼
}

type ValidateNicknameDupBody struct {
	Nickname string `json:"nickname" binding:"required,min=1,max=20" example:"henry"` // 暱稱 (1~20字元)
}

type ValidateEmailDupBody struct {
	Email string `json:"email" binding:"required,email" example:"henry@gmail.com"` // 信箱
}