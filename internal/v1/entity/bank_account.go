package entity

type BankAccount struct {
	UserID       int64  `gorm:"column:user_id"`       // 用戶id
	AccountName  string `gorm:"column:account_name"`  // 戶名
	AccountImage string `gorm:"column:account_image"` // 帳戶照片
	BackCode     string `gorm:"column:bank_code"`     // 銀行代號
	Branch       string `gorm:"column:branch"`        // 分行
	Account      string `gorm:"column:account"`       // 銀行帳戶
	CreateAt     string `gorm:"column:create_at"`     // 創建日期
	UpdateAt     string `gorm:"column:update_at"`     // 更新時間
}

func (BankAccount) TableName() string {
	return "bank_accounts"
}
