package entity

type WorkoutSetLog struct {
	ID           int64   `gorm:"column:id"`             // 課表統計報表id
	WorkoutLogID int64   `gorm:"column:workout_log_id"` // 訓練歷史id
	WorkoutSetID int64   `gorm:"column:workout_set_id"` // 訓練組id
	Weight       float64 `gorm:"column:weight"`         // 重量(公斤)
	Reps         int     `gorm:"column:reps"`           // 次數
	Distance     float64 `gorm:"column:distance"`       // 距離(公里)
	Duration     int     `gorm:"column:duration"`       // 時長(秒)
	Incline      float64 `gorm:"column:incline"`        // 坡度
	CreateAt     string  `gorm:"column:create_at"`      // 更新日期
}

func (WorkoutSetLog) TableName() string {
	return "workout_set_logs"
}
