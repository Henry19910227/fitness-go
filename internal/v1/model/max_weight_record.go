package model

type MaxWeightRecord struct {
	ID       int64   `gorm:"column:id"`        // 紀錄id
	UserID   int64   `gorm:"column:user_id"`   // 用戶id
	ActionID int64   `gorm:"column:action_id"` // 動作 id
	Weight   float64 `gorm:"column:weight"`    // 重量
	UpdateAt string  `gorm:"column:update_at"` // 更新時間
}

func (MaxWeightRecord) TableName() string {
	return "max_weight_record"
}

type SaveMaxWeightRecord struct {
	UserID   int64   // 用戶id
	ActionID int64   // 動作 id
	Weight   float64 // 重量
}
