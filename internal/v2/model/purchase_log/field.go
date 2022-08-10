package purchase_log

type IDField struct {
	ID *int64 ` json:"id,omitempty" gorm:"column:id" example:"1"` // 購買記錄id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` //用戶id
}
type OrderIDField struct {
	OrderID *string `json:"order_id,omitempty" gorm:"column:order_id" example:"20220215104747115283"` //訂單id
}
type TypeField struct {
	Type *int `json:"type,omitempty" gorm:"column:type" example:"1"` //訂單類型(1:購買/2:退費)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}

type Table struct {
	IDField
	UserIDField
	OrderIDField
	TypeField
	CreateAtField
}

func (Table) TableName() string {
	return "purchase_logs"
}
