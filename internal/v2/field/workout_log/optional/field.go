package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` // 訓練 log id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` //用戶id
}
type WorkoutIDField struct {
	WorkoutID *int64 `json:"workout_id,omitempty" uri:"workout_id" gorm:"column:workout_id" binding:"omitempty" example:"1"` //訓練id
}
type DurationField struct {
	Duration *int `json:"duration,omitempty" gorm:"column:duration" binding:"omitempty" example:"30"` //時長(秒)
}
type IntensityField struct {
	Intensity *int `json:"intensity,omitempty" gorm:"column:intensity" binding:"omitempty,oneof=0 1 2 3 4" example:"1"` // 訓練強度(0:未指定/1:輕鬆/2:適中/3:稍難/4:很累)
}
type PlaceField struct {
	Place *int `json:"place,omitempty" gorm:"column:place" binding:"omitempty,oneof=0 1 2 3 4 5" example:"1"` // 地點(0:未指定/1:住家/2:健身房/3:戶外/4:空地/5:其他)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-12 00:00:00"` // 創建時間
}
