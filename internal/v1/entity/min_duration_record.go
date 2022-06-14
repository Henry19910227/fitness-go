package entity

type MinDurationRecord struct {
	ID       int64  `gorm:"column:id"`        // 紀錄id
	UserID   int64  `gorm:"column:user_id"`   // 用戶id
	ActionID int64  `gorm:"column:action_id"` // 動作 id
	Duration int    `gorm:"column:duration"`  // 時長(秒)
	UpdateAt string `gorm:"column:update_at"` // 更新時間
}

func (MinDurationRecord) TableName() string {
	return "min_duration_records"
}
