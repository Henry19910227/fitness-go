package workout_set_log

type IDRequired struct {
	ID int64 `json:"id" binding:"required" example:"1"` //log id
}
type WorkoutLogIDRequired struct {
	WorkoutLogID int64 `json:"workout_set_id" binding:"required" example:"1"` // 訓練歷史id
}
type WorkoutSetIDRequired struct {
	WorkoutSetID int64 `json:"workout_set_id" binding:"required" example:"1"` //訓練組id
}
type WeightRequired struct {
	Weight float64 `json:"weight" binding:"required,min=0.01,max=999.99" example:"50.5"` //體重(公斤)
}
type RepsRequired struct {
	Reps int `json:"reps" binding:"required,min=1,max=999" example:"2"` //次數
}
type DistanceRequired struct {
	Distance float64 `json:"distance" binding:"required,min=0.01,max=999.99" example:"2.5"` //距離(公里)
}
type DurationRequired struct {
	Duration int `json:"duration" binding:"required,min=1,max=38439" example:"30"` //時長(秒)
}
type InclineRequired struct {
	Incline float64 `json:"incline" binding:"required,min=0.01,max=999.99" example:"10.5"` //坡度
}
