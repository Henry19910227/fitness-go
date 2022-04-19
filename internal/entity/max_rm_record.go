package entity

type MaxRmRecord struct {
	ID       int64   `gorm:"column:id"`        // 紀錄id
	UserID   int64   `gorm:"column:user_id"`   // 用戶id
	ActionID int64   `gorm:"column:action_id"` // 動作 id
	RM       float64 `gorm:"column:rm"`        // 最大重量
	UpdateAt string  `gorm:"column:update_at"` // 更新時間
}

func (MaxRmRecord) TableName() string {
	return "max_rm_records"
}
