package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"2"` // 管理員 id
}
type EmailField struct {
	Email *string `json:"email,omitempty" gorm:"column:email" binding:"omitempty,email,max=255" example:"henry@gmail.com"` // 信箱
}
type PasswordField struct {
	Password *string `json:"password,omitempty" gorm:"column:password" binding:"omitempty,min=6,max=18" example:"12345678"` // 密碼
}
type NicknameField struct {
	Nickname *string `json:"nickname,omitempty" gorm:"column:nickname" binding:"omitempty,max=20" example:"Henry"` // 暱稱
}
type LvField struct {
	Lv *int `json:"lv,omitempty" gorm:"column:lv" binding:"omitempty" example:"2"` // 身份(普通管理員:1/超級管理員:2)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 更新時間
}
type LastLoginField struct {
	LastLogin *string `json:"last_login,omitempty" gorm:"column:last_login" binding:"omitempty" example:"2022-06-12 00:00:00"` // 最後登入時間
}
type StatusField struct {
	Status *int `json:"status,omitempty" gorm:"column:status" binding:"omitempty,oneof=0 1" example:"1"` //動作狀態(0:下架/1:上架)
}
