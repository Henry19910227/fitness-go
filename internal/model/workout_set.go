package model

type WorkoutSet struct {
	ID int64 `gorm:"column:id"` //訓練組id
	WorkoutID int64 `gorm:"column:workout_id"` //訓練id
	ActionID *int64 `gorm:"column:action_id"` //動作id
	Type int `gorm:"column:type"` //動作類別(1:動作/2:休息)
	AutoNext string `gorm:"column:auto_next"` //自動下一組(Y:是/N:否)
	StartAudio string `gorm:"column:start_audio"` //前導語音
	ProgressAudio string `gorm:"column:progress_audio"` //進行中語音
	Remark string `gorm:"column:remark"` //備註
	Weight float64 `gorm:"column:weight"` //重量(公斤)
	Reps int `gorm:"column:reps"` //次數
	Distance float64 `gorm:"column:distance"` //距離(公尺)
	Duration int `gorm:"column:duration"` //時長(秒)
	Incline float64 `gorm:"column:incline"` //坡度
	CreateAt string `gorm:"column:create_at"` // 創建時間
	UpdateAt string `gorm:"column:update_at"` // 更新時間
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

type WorkoutSetEntity struct {
	ID int64 `gorm:"column:id"` //訓練組id
	WorkoutID int64 `gorm:"column:workout_id"` //訓練id
	Action *WorkoutSetAction `gorm:"column:-"` //動作
	Type int `gorm:"column:type"` //動作類別(1:動作/2:休息)
	AutoNext string `gorm:"column:auto_next"` //自動下一組(Y:是/N:否)
	StartAudio string `gorm:"column:start_audio"` //前導語音
	ProgressAudio string `gorm:"column:progress_audio"` //進行中語音
	Remark string `gorm:"column:remark"` //備註
	Weight float64 `gorm:"column:weight"` //重量(公斤)
	Reps int `gorm:"column:reps"` //次數
	Distance float64 `gorm:"column:distance"` //距離(公尺)
	Duration int `gorm:"column:duration"` //時長(秒)
	Incline float64 `gorm:"column:incline"` //坡度
}

type WorkoutSetAction struct {
	ID int64  `gorm:"column:id"` //動作id
	Name string `gorm:"column:name"` //課表名稱
	Source int `gorm:"column:source"` //動作來源(1:系統動作/2:教練自創動作)
	Type int `gorm:"column:type"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Intro string `gorm:"column:intro"` //動作介紹
	Cover string `gorm:"column:cover"` //封面
	Video string `gorm:"column:video"` //動作影片
}

type UpdateWorkoutSetParam struct {
	AutoNext *string `gorm:"column:auto_next"` //自動下一組(Y:是/N:否)
	StartAudio *string `gorm:"column:start_audio"` //前導語音
	ProgressAudio *string `gorm:"column:progress_audio"` //進行中語音
	Remark *string `gorm:"column:remark"` //備註
	Weight *float64 `gorm:"column:weight"` //重量(公斤)
	Reps *int `gorm:"column:reps"` //次數
	Distance *float64 `gorm:"column:distance"` //距離(公尺)
	Duration *int `gorm:"column:duration"` //時長(秒)
	Incline *float64 `gorm:"column:incline"` //坡度
	UpdateAt *string `gorm:"column:update_at"` //更新時間
}