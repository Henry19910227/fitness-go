package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" binding:"omitempty" example:"1"` //log id
}
type WorkoutLogIDField struct {
	WorkoutLogID *int64 `json:"workout_log_id,omitempty" gorm:"column:workout_log_id" binding:"omitempty" example:"1"` // 訓練歷史id
}
type WorkoutSetIDField struct {
	WorkoutSetID *int64 `json:"workout_set_id,omitempty" gorm:"column:workout_set_id" binding:"omitempty" example:"1"` //訓練組id
}
type WeightField struct {
	Weight *float64 `json:"weight,omitempty" gorm:"column:weight" binding:"omitempty,min=0.01,max=999.99" example:"50.5"` //體重(公斤)
}
type RepsField struct {
	Reps *int `json:"reps,omitempty" gorm:"column:reps" binding:"omitempty,min=1,max=999" example:"2"` //次數
}
type DistanceField struct {
	Distance *float64 `json:"distance,omitempty" gorm:"column:distance" binding:"omitempty,min=0.01,max=999.99" example:"2.5"` //距離(公里)
}
type DurationField struct {
	Duration *int `json:"duration,omitempty" gorm:"column:duration" binding:"omitempty,min=1,max=38439" example:"30"` //時長(秒)
}
type InclineField struct {
	Incline *float64 `json:"incline,omitempty" gorm:"column:incline" binding:"omitempty,min=0.01,max=999.99" example:"10.5"` //坡度
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
