package validator


type EmailBody struct {
	Email string `json:"email" binding:"required,email" example:"test@gmail.com"`
}

type RegisterForEmailBody struct {
	Email     string `json:"email" binding:"required,email" example:"test@gmail.com"`             // 信箱
	Password  string `json:"password" binding:"required,min=8,max=16" example:"12345678"`         // 密碼 (8~16字元)
	Nickname  string `json:"nickname" binding:"required,min=1,max=16" example:"henry"`            // 暱稱 (1~16字元)
	EmailOTP  string `json:"email_otp" binding:"required,max=16" example:"531476"`               // 信箱驗證碼
}

type ValidateNicknameDupBody struct {
	Nickname string `json:"nickname" binding:"required,min=1,max=16" example:"henry"` // 暱稱 (1~16字元)
}