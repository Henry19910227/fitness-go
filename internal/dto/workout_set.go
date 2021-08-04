package dto

type WorkoutSetID struct {
	ID int64 `json:"workout_set_id" example:"10"` //訓練組id
}

type WorkoutSet struct {
	ID int64                 `json:"id" example:"10"`                          //訓練組id
	Type int                 `json:"type" example:"2"`                         //動作類別(1:動作/2:休息)
	Action *WorkoutSetAction `json:"action"`                                   //動作
	AutoNext string          `json:"auto_next" example:"N"`                    //自動下一組(Y:是/N:否)
	StartAudio string        `json:"start_audio" example:""`                   //前導語音
	ProgressAudio string     `json:"progress_audio" example:"1d2w3e51d3w.mp3"` //進行中語音
	Remark string            `json:"remark" example:""`                        //備註
	Weight float64           `json:"weight" example:"0"`                       //重量(公斤)
	Reps int                 `json:"reps" example:"0"`                         //次數
	Distance float64         `json:"distance" example:"0"`                     //距離(公尺)
	Duration int             `json:"duration" example:"30"`                    //時長(秒)
	Incline float64          `json:"incline" example:"0"`                      //坡度
}

type WorkoutSetAction struct {
	ID int64  `json:"id" example:"1"` //動作id
	Name string `json:"name" example:"槓鈴臥推"` //動作名稱
	Source int `json:"source" example:"2"` //動作來源(1:系統動作/2:教練自創動作)
	Type int `json:"type" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Intro string `json:"intro" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹
	Cover string `json:"cover" example:"32as1d5f13e4.png"` //封面
	Video string `json:"video" example:"11d547we1d4f8e.mp4"` //動作影片
}

type UpdateWorkoutSetParam struct {
	AutoNext *string `gorm:"column:auto_next"` //自動下一組(Y:是/N:否)
	StartAudio *string `gorm:"column:start_audio"` //前導語音
	Remark *string `gorm:"column:remark"` //備註
	Weight *float64 `gorm:"column:weight"` //重量(公斤)
	Reps *int `gorm:"column:reps"` //次數
	Distance *float64 `gorm:"column:distance"` //距離(公尺)
	Duration *int `gorm:"column:duration"` //時長(秒)
	Incline *float64 `gorm:"column:incline"` //坡度
}

type WorkoutSetOrder struct {
	WorkoutSetID int64 `gorm:"column:workout_set_id"` //訓練組id
	Seq int `gorm:"column:seq"` //排列序號
}