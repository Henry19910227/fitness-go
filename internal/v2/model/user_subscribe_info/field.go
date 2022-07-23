package user_subscribe_info

type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type OrderIDField struct {
	OrderID *string `json:"order_id,omitempty" gorm:"column:order_id" example:"20220215104747115283"` //訂單id
}
type StatusField struct {
	Status *int `json:"status,omitempty" gorm:"column:status" example:"1"` // 會員狀態(0:無會員狀態/1:付費會員狀態)
}
type StartDateField struct {
	StartDate *string `json:"start_date,omitempty" gorm:"column:start_date" example:"2022-07-11 11:00:00"` // 訂閱開始日期
}
type ExpiresDateField struct {
	ExpiresDate *string `json:"end_date,omitempty" gorm:"column:expires_date" example:"2023-07-11 11:00:00"` // 訂閱結束日期
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	UserIDField
	OrderIDField
	StatusField
	StartDateField
	ExpiresDateField
	UpdateAtField
}
func (Table) TableName() string {
	return "user_subscribe_infos"
}
