package model

type MaxRepsRecord struct {
	ID       int64  `gorm:"column:id"`        // 紀錄id
	UserID   int64  `gorm:"column:user_id"`   // 用戶id
	ActionID int64  `gorm:"column:action_id"` // 動作 id
	Reps     int    `gorm:"column:reps"`      // 次數
	UpdateAt string `gorm:"column:update_at"` // 更新時間
}

func (MaxRepsRecord) TableName() string {
	return "max_reps_record"
}

type SaveMaxRepsRecord struct {
	UserID   int64 // 用戶id
	ActionID int64 // 動作 id
	Reps     int   // 次數
}
