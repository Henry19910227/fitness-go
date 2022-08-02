package bank_account

type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" example:"10001"` //用戶id
}
type AccountNameOptional struct {
	AccountName *string `json:"account_name,omitempty" example:"王小明"` // 戶名
}
type AccountImageOptional struct {
	AccountImage *string `json:"account_image,omitempty" example:"123.png"` // 帳戶照片
}
type BackCodeOptional struct {
	BackCode *string `json:"bank_code,omitempty" example:"009"` // 銀行代號
}
type BranchOptional struct {
	Branch *string `json:"branch,omitempty" example:"南京分行"` // 分行
}
type AccountOptional struct {
	Account *string `json:"account,omitempty" example:"南京分行"` // 分行
}
type CreateAtOptional struct {
	CreateAt *string `json:"create_at,omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtOptional struct {
	UpdateAt *string `json:"update_at,omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
