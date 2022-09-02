package workout_set

type IDOptional struct {
	ID *int64 `json:"id,omitempty" uri:"workout_set_id" example:"2"` // 訓練 id
}
type WorkoutIDOptional struct {
	WorkoutID *int64 `json:"workout_id,omitempty" uri:"workout_id" example:"1"` //訓練id
}
type ActionIDOptional struct {
	ActionID *int64 `json:"action_id,omitempty" example:"1"` //動作id
}
type TypeOptional struct {
	Type *int `json:"type,omitempty" example:"1"` //動作類別(1:動作/2:休息)
}
type AutoNextOptional struct {
	AutoNext *string `json:"auto_next,omitempty" form:"auto_next" binding:"omitempty,oneof=Y N" example:"Y"` //自動下一組(Y:是/N:否)
}
type StartAudioOptional struct {
	StartAudio *string `json:"start_audio,omitempty" example:"1234.mp3"` //前導語音
}
type ProgressAudioOptional struct {
	ProgressAudio *string `json:"progress_audio,omitempty" example:"1234.mp3"` //進行中語音
}
type RemarkOptional struct {
	Remark *string `json:"remark,omitempty" form:"remark" binding:"omitempty,max=40" example:"注意呼吸不可憋氣"` //備註
}
type WeightOptional struct {
	Weight *float64 `json:"weight,omitempty" form:"weight" binding:"omitempty,min=0.01,max=999.99" example:"50.5"` //重量(公斤)
}
type RepsOptional struct {
	Reps *int `json:"reps,omitempty" form:"reps" binding:"omitempty,min=1,max=999" example:"2"` //次數
}
type DistanceOptional struct {
	Distance *float64 `json:"distance,omitempty" form:"distance" binding:"omitempty,min=0.01,max=999.99" example:"2.5"` //距離(公里)
}
type DurationOptional struct {
	Duration *int `json:"duration,omitempty" form:"duration" binding:"omitempty,min=1,max=38439" example:"30"` //時長(秒)
}
type InclineOptional struct {
	Incline *float64 `json:"incline,omitempty" form:"incline" binding:"omitempty,min=0.01,max=999.99" example:"10.5"` //坡度
}
