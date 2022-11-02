package required

type UserIDField struct {
	UserID int64 `json:"user_id" gorm:"column:user_id" gorm:"column:user_id" binding:"required" example:"10001"` //用戶id
}
type AccountNameField struct {
	AccountName string `json:"account_name" form:"account_name" gorm:"column:account_name" binding:"required,max=40" example:"王小明"` // 戶名
}
type AccountField struct {
	Account string `json:"account" form:"account" gorm:"column:account" binding:"required,min=6,max=16" example:"123456789"` // 帳戶
}
type AccountImageField struct {
	AccountImage string `json:"account_image" form:"account_image" gorm:"column:account_image" binding:"required" example:"123.png"` // 帳戶照片
}
type BankCodeField struct {
	BankCode string `json:"bank_code" form:"bank_code" gorm:"column:bank_code" binding:"required,max=40" example:"009"` // 銀行代號
}
type BranchField struct {
	Branch string `json:"branch" form:"branch" gorm:"column:branch" binding:"required,max=40" example:"南京分行"` // 分行
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}