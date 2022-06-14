package model

type ActionBestPR struct {
	ActionID            int64   `gorm:"column:action_id"`              //動作id
	MaxRM               float64 `gorm:"column:max_rm"`                 //最大反覆重量(公斤)
	MaxWeight           float64 `gorm:"column:max_weight"`             //最大重量(公斤)
	MaxReps             int     `gorm:"column:max_reps"`               //最多次數
	MinDuration         int     `gorm:"column:min_duration"`           //最小時長(秒)
	MaxSpeed            float64 `gorm:"column:max_speed"`              //最高速率
	MaxDistance         float64 `gorm:"column:max_distance"`           //最長距離(公里)
	MaxRMUpdateAt       string  `gorm:"column:max_rm_update_at"`       //最大反覆重量紀錄更新時間
	MaxWeightUpdateAt   string  `gorm:"column:max_weight_update_at"`   //最大重量紀錄更新時間
	MaxRepsUpdateAt     string  `gorm:"column:max_reps_update_at"`     //最多次數紀錄更新時間
	MinDurationUpdateAt string  `gorm:"column:min_duration_update_at"` //最小時長紀錄更新時間
	MaxSpeedUpdateAt    string  `gorm:"column:max_speed_update_at"`    //最高速率紀錄更新時間
	MaxDistanceUpdateAt string  `gorm:"column:max_distance_update_at"` //最長距離紀錄更新時間
}
