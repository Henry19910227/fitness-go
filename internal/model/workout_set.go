package model

type WorkoutSet struct {
	ID int64 `gorm:"column:id"` //訓練組id
	WorkoutID int64 `gorm:"column:workout_id"` //訓練id
	ActionID *int64 `gorm:"column:action_id"` //動作id
	Action *Action `gorm:"foreignkey:id;references:action_id"` //動作
	Type int `gorm:"column:type"` //動作類別(1:動作/2:休息)
	AutoNext string `gorm:"column:auto_next"` //自動下一組(Y:是/N:否)
	StartAudio string `gorm:"column:start_audio"` //前導語音
	ProgressAudio string `gorm:"column:progress_audio"` //進行中語音
	Remark string `gorm:"column:remark"` //備註
	Weight float64 `gorm:"column:weight"` //重量(公斤)
	Reps int `gorm:"column:reps"` //次數
	Distance float64 `gorm:"column:distance"` //距離(公里)
	Duration int `gorm:"column:duration"` //時長(秒)
	Incline float64 `gorm:"column:incline"` //坡度
}

func (WorkoutSet) TableName() string {
	return "workout_sets"
}

type WorkoutSetOrder struct {
	WorkoutID int64 `gorm:"column:workout_id"` //訓練id
	WorkoutSetID int64 `gorm:"column:workout_set_id"` //訓練組id
	Seq int `gorm:"column:seq"` //排列序號
}

func (WorkoutSetOrder) TableName() string {
	return "workout_set_orders"
}

type UpdateWorkoutSetParam struct {
	AutoNext *string `gorm:"column:auto_next"` //自動下一組(Y:是/N:否)
	StartAudio *string `gorm:"column:start_audio"` //前導語音
	ProgressAudio *string `gorm:"column:progress_audio"` //進行中語音
	Remark *string `gorm:"column:remark"` //備註
	Weight *float64 `gorm:"column:weight"` //重量(公斤)
	Reps *int `gorm:"column:reps"` //次數
	Distance *float64 `gorm:"column:distance"` //距離(公里)
	Duration *int `gorm:"column:duration"` //時長(秒)
	Incline *float64 `gorm:"column:incline"` //坡度
	UpdateAt *string `gorm:"column:update_at"` //更新時間
}