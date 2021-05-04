package model

type Admin struct {
	ID          int64  `json:"id" gorm:"column:id" example:"1"`                       // 帳戶id
	Email       string `json:"email" gorm:"column:email" example:"henry@gmail.com"`   // 信箱
	Password    string `json:"password" gorm:"column:password"`                       // 密碼
	Nickname    string `json:"nickname" gorm:"column:nickname" example:"henry" `      // 暱稱
	Lv          int    `json:"lv" gorm:"column:lv" example:"1"`                   // 身份: 超級管理員 = 1
	CreateAt    string `json:"create_at" gorm:"column:create_at"`                     // 創建時間
	UpdateAt    string `json:"update_at" gorm:"column:update_at"`                     // 更新時間
	Status      int    `json:"status" gorm:"column:status" example:"1"`               // 帳號狀態
}

func (Admin) TableName() string {
	return "admins"
}
