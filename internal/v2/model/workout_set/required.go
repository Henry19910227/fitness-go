package workout_set

type IDRequired struct {
	ID int64 `json:"id" uri:"workout_set_id" example:"2"` // 訓練 id
}
type WorkoutIDRequired struct {
	WorkoutID int64 `json:"workout_id" uri:"workout_id" example:"1"` //訓練id
}
type ActionIDRequired struct {
	ActionID int64 `json:"action_id" example:"1"` //動作id
}
type TypeRequired struct {
	Type int `json:"type" example:"1"` //動作類別(1:動作/2:休息)
}
type AutoNextRequired struct {
	AutoNext string `json:"auto_next" example:"Y"` //自動下一組(Y:是/N:否)
}
type StartAudioRequired struct {
	StartAudio string `json:"start_audio" example:"1234.mp3"` //前導語音
}
type ProgressAudioRequired struct {
	ProgressAudio string `json:"progress_audio" example:"1234.mp3"` //進行中語音
}
type RemarkRequired struct {
	Remark string `json:"remark" example:"注意呼吸不可憋氣"` //備註
}
type WeightRequired struct {
	Weight float64 `json:"weight" example:"50.5"` //重量(公斤)
}
type RepsRequired struct {
	Reps int `json:"reps" example:"2"` //次數
}
type DistanceRequired struct {
	Distance float64 `json:"distance" example:"2.5"` //距離(公里)
}
type DurationRequired struct {
	Duration int `json:"duration" example:"30"` //時長(秒)
}
type InclineRequired struct {
	Incline float64 `json:"incline" example:"10.5"` //坡度
}
