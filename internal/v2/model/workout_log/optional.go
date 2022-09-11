package workout_log

type IDOptional struct {
	ID *int64 `json:"id,omitempty" example:"1"` // 訓練 log id
}
type UserIDOptional struct {
	UserID *int64 `json:"user_id,omitempty" example:"10001"` //用戶id
}
type WorkoutIDOptional struct {
	WorkoutID *int64 `json:"workout_id,omitempty" uri:"workout_id" example:"1"` //訓練id
}
type DurationOptional struct {
	Duration *int `json:"duration,omitempty" example:"30"` //時長(秒)
}
type IntensityOptional struct {
	Intensity *int `json:"intensity,omitempty" binding:"omitempty,oneof=0 1 2 3 4" example:"1"` // 訓練強度(0:未指定/1:輕鬆/2:適中/3:稍難/4:很累)
}
type PlaceOptional struct {
	Place *int `json:"place,omitempty" binding:"omitempty,oneof=0 1 2 3" example:"1"` // 地點(0:未指定/1:住家/2:健身房/3:戶外)
}
