package validator

type CMSLoginByEmailBody struct {
	Email     string `json:"email" binding:"required,email" example:"test@gmail.com"`             // 信箱
	Password  string `json:"password" binding:"required,min=6,max=18" example:"12345678"`         // 密碼 (6~18字元)
}
