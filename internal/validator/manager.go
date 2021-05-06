package validator

type AdminLoginBody struct {
	Email     string `json:"email" binding:"required,email" example:"henry@gmail.com"`           // 信箱
	Password  string `json:"password" binding:"required,min=1,max=16" example:"12345678"`        // 密碼 (8~16字元)
}
