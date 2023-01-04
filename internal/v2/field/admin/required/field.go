package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"2"` // 管理員 id
}
type EmailField struct {
	Email string `json:"email" gorm:"column:email" binding:"required,email,max=255" example:"henry@gmail.com"` // 信箱
}
type PasswordField struct {
	Password string `json:"password" gorm:"column:password" binding:"required,min=6,max=18" example:"12345678"` // 密碼
}
type NicknameField struct {
	Nickname string `json:"nickname" gorm:"column:nickname" binding:"required,max=20" example:"Henry"` // 暱稱
}
type LvField struct {
	Lv string `json:"lv" gorm:"column:lv" binding:"required" example:"Henry"` // 身份(普通管理員:1/超級管理員:2)
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-12 00:00:00"` // 更新時間
}
type LastLoginField struct {
	LastLogin string `json:"last_login" gorm:"column:last_login" binding:"required" example:"2022-06-12 00:00:00"` // 最後登入時間
}
type StatusField struct {
	Status int `json:"status" gorm:"column:status" binding:"required,oneof=0 1" example:"1"` //動作狀態(0:下架/1:上架)
}
