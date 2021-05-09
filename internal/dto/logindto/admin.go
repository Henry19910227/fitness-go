package logindto

type Admin struct {
	ID          int64  `json:"id" gorm:"column:id" example:"1"`                       // 帳戶id
	Email       string `json:"email" gorm:"column:email" example:"henry@gmail.com"`   // 信箱
	Nickname    string `json:"nickname" gorm:"column:nickname" example:"henry" `      // 暱稱
	Lv          int    `json:"lv" gorm:"column:lv" example:"1"`                       // 身份 (1:一般管理員)
}