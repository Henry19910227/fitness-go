package model

type WorkoutSetLog struct {
	ID           int64       `gorm:"column:id"`                               // 訓練組歷史 id
	WorkoutLogID int64       `gorm:"column:workout_log_id"`                   // 訓練歷史id
	WorkoutSetID int64       `gorm:"column:workout_set_id"`                   // 訓練組id
	Weight       float64     `gorm:"column:weight"`                           // 重量(公斤)
	Reps         int         `gorm:"column:reps"`                             // 次數
	Distance     float64     `gorm:"column:distance"`                         // 距離(公里)
	Duration     int         `gorm:"column:duration"`                         // 時長(秒)
	Incline      float64     `gorm:"column:incline"`                          // 坡度
	WorkoutSet   *WorkoutSet `gorm:"foreignKey:id;references:workout_set_id"` // 訓練組
}

type WorkoutSetLogParam struct {
	ID           int64   // 訓練組歷史 id
	WorkoutLogID int64   // 訓練歷史id
	WorkoutSetID int64   // 訓練組id
	Weight       float64 // 重量(公斤)
	Reps         int     // 次數
	Distance     float64 // 距離(公里)
	Duration     int     // 時長(秒)
	Incline      float64 // 坡度
}
