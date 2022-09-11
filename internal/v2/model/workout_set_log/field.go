package workout_set_log

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //log id
}
type WorkoutLogIDField struct {
	WorkoutLogID *int64 `json:"workout_log_id,omitempty" gorm:"column:workout_log_id" example:"1"` // 訓練歷史id
}
type WorkoutSetIDField struct {
	WorkoutSetID *int64 `json:"workout_set_id,omitempty" gorm:"column:workout_set_id" example:"1"` //訓練組id
}
type WeightField struct {
	Weight *float64 `json:"weight,omitempty" gorm:"weight:value" example:"50.5"` //體重(公斤)
}
type RepsField struct {
	Reps *int `json:"reps,omitempty" gorm:"column:reps" example:"2"` //次數
}
type DistanceField struct {
	Distance *float64 `json:"distance,omitempty" gorm:"column:distance" example:"2.5"` //距離(公里)
}
type DurationField struct {
	Duration *int `json:"duration,omitempty" gorm:"column:duration" example:"30"` //時長(秒)
}
type InclineField struct {
	Incline *float64 `json:"incline,omitempty" gorm:"column:incline" example:"10.5"` //坡度
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at;default:2022-06-14 00:00:00" example:"2022-06-14 00:00:00"` //創建時間
}

type Table struct {
	IDField
	WorkoutLogIDField
	WorkoutSetIDField
	WeightField
	RepsField
	DistanceField
	DurationField
	InclineField
	CreateAtField
}

func (Table) TableName() string {
	return "workout_set_logs"
}
