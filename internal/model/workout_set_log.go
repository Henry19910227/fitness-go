package model

type WorkoutSetLog struct {
	ID int64 // 訓練組歷史 id
	WorkoutLogID int64   // 訓練歷史id
	WorkoutSetID int64   // 訓練組id
	Weight       float64 // 重量(公斤)
	Reps         int     // 次數
	Distance     float64 // 距離(公里)
	Duration     int     // 時長(秒)
	Incline      float64 // 坡度
}
