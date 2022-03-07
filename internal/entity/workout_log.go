package entity

type WorkoutLog struct {
	ID        int64  `gorm:"column:id"`         // 課表統計報表id
	UserID    int64  `gorm:"column:user_id"`    // 用戶id
	WorkoutID int64  `gorm:"column:workout_id"` // 訓練id
	Duration  int    `gorm:"column:duration"`   // 訓練時長
	Intensity int    `gorm:"column:intensity"`  // 訓練強度(1:輕鬆/2:適中/3:稍難/4:很累)
	Place     int    `gorm:"column:place"`      // 地點(1:住家/2:健身房/3:戶外)
	CreateAt  string `gorm:"column:create_at"`  // 更新日期
}

func (WorkoutLog) TableName() string {
	return "workout_logs"
}
