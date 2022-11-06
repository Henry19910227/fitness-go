package required

type IDField struct {
	ID int64 `json:"id" gorm:"column:id" binding:"required" example:"1"` //log id
}
type WorkoutLogIDField struct {
	WorkoutLogID int64 `json:"workout_log_id" gorm:"column:workout_log_id" binding:"required" example:"1"` // 訓練歷史id
}
type WorkoutSetIDField struct {
	WorkoutSetID int64 `json:"workout_set_id" gorm:"column:workout_set_id" binding:"required" example:"1"` //訓練組id
}
type WeightField struct {
	Weight float64 `json:"weight" gorm:"column:weight" binding:"required,min=0.01,max=999.99" example:"50.5"` //體重(公斤)
}
type RepsField struct {
	Reps int `json:"reps" gorm:"column:reps" binding:"required,min=1,max=999" example:"2"` //次數
}
type DistanceField struct {
	Distance float64 `json:"distance" gorm:"column:distance" binding:"required,min=0.01,max=999.99" example:"2.5"` //距離(公里)
}
type DurationField struct {
	Duration int `json:"duration" gorm:"column:duration" binding:"required,min=1,max=38439" example:"30"` //時長(秒)
}
type InclineField struct {
	Incline float64 `json:"incline" gorm:"column:incline" binding:"required,min=0.01,max=999.99" example:"10.5"` //坡度
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}