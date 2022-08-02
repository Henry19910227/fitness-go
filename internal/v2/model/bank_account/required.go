package bank_account

type UserIDRequired struct {
	UserID int64 `json:"user_id,omitempty" example:"10001"` //用戶id
}
type AccountNameRequired struct {
	AccountName string `json:"account_name,omitempty" example:"王小明"` // 戶名
}
type AccountImageRequired struct {
	AccountImage string `json:"account_image,omitempty" example:"123.png"` // 帳戶照片
}
type BackCodeRequired struct {
	BackCode string `json:"bank_code,omitempty" example:"009"` // 銀行代號
}
type BranchRequired struct {
	Branch string `json:"branch,omitempty" example:"南京分行"` // 分行
}
type AccountRequired struct {
	Account string `json:"account,omitempty" example:"南京分行"` // 分行
}
type CreateAtRequired struct {
	CreateAt string `json:"create_at,omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtRequired struct {
	UpdateAt string `json:"update_at,omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
