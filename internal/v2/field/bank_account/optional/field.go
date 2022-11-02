package optional

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" gorm:"column:user_id" binding:"omitempty" example:"10001"` //用戶id
}
type AccountNameField struct {
	AccountName *string `json:"account_name,omitempty" form:"account_name" gorm:"column:account_name" binding:"omitempty,max=40" example:"王小明"` // 戶名
}
type AccountField struct {
	Account *string `json:"account,omitempty" form:"account" gorm:"column:account" binding:"omitempty,min=6,max=16" example:"123456789"` // 帳戶
}
type AccountImageField struct {
	AccountImage *string `json:"account_image,omitempty" form:"account_image" gorm:"column:account_image" binding:"omitempty" example:"123.png"` // 帳戶照片
}
type BankCodeField struct {
	BankCode *string `json:"bank_code,omitempty" form:"bank_code" gorm:"column:bank_code" binding:"omitempty,max=40" example:"009"` // 銀行代號
}
type BranchField struct {
	Branch *string `json:"branch,omitempty" form:"branch" gorm:"column:branch" binding:"omitempty,max=40" example:"南京分行"` // 分行
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
