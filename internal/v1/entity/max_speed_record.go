package entity

type MaxSpeedRecord struct {
	ID       int64   `gorm:"column:id"`        // 紀錄id
	UserID   int64   `gorm:"column:user_id"`   // 用戶id
	ActionID int64   `gorm:"column:action_id"` // 動作 id
	Speed    float64 `gorm:"column:speed"`     // 每小時速率
	UpdateAt string  `gorm:"column:update_at"` // 更新時間
}

func (MaxSpeedRecord) TableName() string {
	return "max_speed_records"
}
