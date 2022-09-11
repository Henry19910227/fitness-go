package max_distance_record

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` // 訓練 id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" example:"10001"` // 用戶id
}
type ActionIDField struct {
	ActionID *int64 `json:"action_id,omitempty" gorm:"column:action_id"  example:"1"` //動作id
}
type DistanceField struct {
	Distance *float64 `json:"distance,omitempty" gorm:"column:distance"  example:"1"` //距離(公里)
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	UserIDField
	ActionIDField
	DistanceField
	UpdateAtField
}

func (Table) TableName() string {
	return "max_distance_records"
}
