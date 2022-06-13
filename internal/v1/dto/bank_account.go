package dto

type BankAccount struct {
	AccountName  string `json:"account_name" example:"王小明"`                    // 戶名
	AccountImage string `json:"account_image" example:"sd2lkd1e23w54dw3e.png"` // 帳戶照片
	BankCode     string `json:"bank_code" example:"009"`                       // 銀行代號
	Branch       string `json:"branch" example:"南京分行"`                         // 分行
	Account      string `json:"account" example:"8441236"`                     // 銀行帳戶
}

func (Trainer) TableName() string {
	return "bank_accounts"
}
