package workout_log

type IDRequired struct {
	ID int64 `json:"id" example:"1"` // 訓練 log id
}
type UserIDRequired struct {
	UserID int64 `json:"user_id" example:"10001"` //用戶id
}
type WorkoutIDRequired struct {
	WorkoutID int64 `json:"workout_id" uri:"workout_id" gorm:"column:workout_id" example:"1"` //訓練id
}
type DurationRequired struct {
	Duration int `json:"duration" binding:"required" example:"30"` //時長(秒)
}
type IntensityRequired struct {
	Intensity int `json:"intensity" binding:"omitempty,oneof=0 1 2 3 4" example:"1"` // 訓練強度(0:未指定/1:輕鬆/2:適中/3:稍難/4:很累)
}
type PlaceRequired struct {
	Place int `json:"place" binding:"omitempty,oneof=0 1 2 3" example:"1"` // 地點(0:未指定/1:住家/2:健身房/3:戶外)
}
