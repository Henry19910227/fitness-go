package workout_set_log

type IDOptional struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` //log id
}
type WorkoutLogIDOptional struct {
	WorkoutLogID *int64 `json:"workout_log_id,omitempty" binding:"omitempty" example:"1"` // 訓練歷史id
}
type WorkoutSetIDOptional struct {
	WorkoutSetID *int64 `json:"workout_set_id,omitempty" binding:"omitempty" example:"1"` //訓練組id
}
type WeightOptional struct {
	Weight *float64 `json:"weight,omitempty" binding:"omitempty,min=0.01,max=999.99" example:"50.5"` //體重(公斤)
}
type RepsOptional struct {
	Reps *int `json:"reps,omitempty" binding:"omitempty,min=1,max=999" example:"2"` //次數
}
type DistanceOptional struct {
	Distance *float64 `json:"distance,omitempty" binding:"omitempty,min=0.01,max=999.99" example:"2.5"` //距離(公里)
}
type DurationOptional struct {
	Duration *int `json:"duration,omitempty" binding:"omitempty,min=1,max=38439" example:"30"` //時長(秒)
}
type InclineOptional struct {
	Incline *float64 `json:"incline,omitempty" binding:"omitempty,min=0.01,max=999.99" example:"10.5"` //坡度
}
