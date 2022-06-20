package workout_set

type IDField struct {
	ID *int64 `json:"id,omitempty" uri:"workout_set_id" gorm:"column:id" example:"2"` // 訓練 id
}
type WorkoutIDField struct {
	WorkoutID *int64 `json:"workout_id,omitempty" uri:"workout_id" gorm:"column:workout_id" example:"1"` //訓練id
}
type ActionIDField struct {
	ActionID *int64 `json:"action_id,omitempty" gorm:"column:action_id" example:"1"` //動作id
}
type TypeField struct {
	Type *int `json:"type,omitempty" gorm:"column:type" example:"1"` //動作類別(1:動作/2:休息)
}
type AutoNextField struct {
	AutoNext *string `json:"auto_next,omitempty" gorm:"column:auto_next" example:"Y"` //自動下一組(Y:是/N:否)
}
type StartAudioField struct {
	StartAudio *string `json:"start_audio,omitempty" gorm:"column:start_audio" example:"1234.mp3"` //前導語音
}
type ProgressAudioField struct {
	ProgressAudio *string `json:"progress_audio,omitempty" gorm:"column:progress_audio" example:"1234.mp3"` //進行中語音
}
type RemarkField struct {
	Remark *string `json:"remark,omitempty" gorm:"column:remark" example:"注意呼吸不可憋氣"` //備註
}
type WeightField struct {
	Weight *float64 `json:"weight,omitempty" gorm:"column:weight" example:"50.5"` //重量(公斤)
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
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	WorkoutIDField
	ActionIDField
	TypeField
	AutoNextField
	StartAudioField
	ProgressAudioField
	RemarkField
	WeightField
	RepsField
	DistanceField
	DurationField
	InclineField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "workout_sets"
}
