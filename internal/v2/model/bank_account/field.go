package bank_account

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type AccountNameField struct {
	AccountName *string `json:"account_name,omitempty" gorm:"column:account_name" example:"王小明"` // 戶名
}
type AccountImageField struct {
	AccountImage *string `json:"account_image,omitempty" gorm:"column:account_image" example:"123.png"` // 帳戶照片
}
type BackCodeField struct {
	BackCode *string `json:"bank_code,omitempty" gorm:"column:bank_code" example:"009"` // 銀行代號
}
type BranchField struct {
	Branch *string `json:"branch,omitempty" gorm:"column:branch" example:"南京分行"` // 分行
}
type AccountField struct {
	Account *string `json:"account,omitempty" gorm:"column:account" example:"南京分行"` // 分行
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	UserIDField
	AccountNameField
	AccountImageField
	BackCodeField
	BranchField
	AccountField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "bank_accounts"
}
