package workoutdto

type WorkoutSet struct {
	ID int64 `json:"id" example:"10"` //訓練組id
	Type int `json:"type" example:"2"` //動作類別(1:動作/2:休息)
	AutoNext string `json:"auto_next" example:"N"` //自動下一組(Y:是/N:否)
	StartAudio string `json:"start_audio" example:""` //前導語音
	ProgressAudio string `json:"progress_audio" example:"1d2w3e51d3w.mp3"` //進行中語音
	Remark string `json:"remark" example:""` //備註
	Weight float64 `json:"weight" example:"0"` //重量(公斤)
	Reps int `json:"reps" example:"0"` //次數
	Distance float64 `json:"distance" example:"0"` //距離(公尺)
	Duration int `json:"duration" example:"30"` //時長(秒)
	Incline float64 `json:"incline" example:"0"` //坡度
}