package required

type IDField struct {
	ID int64 `json:"id" uri:"workout_set_id" gorm:"column:id" binding:"required" example:"2"` // 訓練 id
}
type WorkoutIDField struct {
	WorkoutID int64 `json:"workout_id" uri:"workout_id" gorm:"column:workout_id" binding:"required" example:"1"` //訓練id
}
type ActionIDField struct {
	ActionID int64 `json:"action_id" uri:"action_id" gorm:"column:action_id" binding:"required" example:"1"` //動作id
}
type TypeField struct {
	Type int `json:"type" form:"type" gorm:"column:type" binding:"required,oneof=1 2" example:"1"` //動作類別(1:動作/2:休息)
}
type AutoNextField struct {
	AutoNext string `json:"auto_next" form:"auto_next" gorm:"column:auto_next" binding:"required,oneof=Y N" example:"Y"` //自動下一組(Y:是/N:否)
}
type StartAudioField struct {
	StartAudio string `json:"start_audio" form:"start_audio" gorm:"column:start_audio" binding:"required" example:"1234.mp3"` //前導語音
}
type ProgressAudioField struct {
	ProgressAudio string `json:"progress_audio" form:"progress_audio" gorm:"column:progress_audio" binding:"required" example:"1234.mp3"` //進行中語音
}
type RemarkField struct {
	Remark string `json:"remark" form:"remark" gorm:"column:remark" binding:"required,max=40" example:"注意呼吸不可憋氣"` //備註
}
type WeightField struct {
	Weight float64 `json:"weight" form:"weight" gorm:"column:weight" binding:"required,min=0.01,max=999.99" example:"50.5"` //重量(公斤)
}
type RepsField struct {
	Reps int `json:"reps" form:"reps" gorm:"column:reps" binding:"required,min=1,max=999" example:"2"` //次數
}
type DistanceField struct {
	Distance float64 `json:"distance" form:"distance" gorm:"column:distance" binding:"required,min=0.01,max=999.99" example:"2.5"` //距離(公里)
}
type DurationField struct {
	Duration int `json:"duration" form:"duration" gorm:"column:duration" binding:"required,min=1,max=38439" example:"30"` //時長(秒)
}
type InclineField struct {
	Incline float64 `json:"incline" form:"incline" gorm:"column:incline" binding:"required,min=0.01,max=999.99" example:"10.5"` //坡度
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" binding:"required" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" binding:"required" example:"2022-06-14 00:00:00"` //更新時間
}
