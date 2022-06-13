package entity

type MaxDistanceRecord struct {
	ID       int64   `gorm:"column:id"`        // 紀錄id
	UserID   int64   `gorm:"column:user_id"`   // 用戶id
	ActionID int64   `gorm:"column:action_id"` // 動作 id
	Distance float64 `gorm:"column:distance"`  // 距離(公里)
	UpdateAt string  `gorm:"column:update_at"` // 更新時間
}

func (MaxDistanceRecord) TableName() string {
	return "max_distance_records"
}
