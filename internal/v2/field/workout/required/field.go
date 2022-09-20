package required

type IDField struct {
	ID int64 `json:"id" uri:"workout_id" gorm:"column:id" example:"1"` // 訓練 id
}
type PlanIDField struct {
	PlanID int64 `json:"plan_id" uri:"plan_id" gorm:"column:plan_id" example:"1"` // 計畫 id
}
type NameField struct {
	Name string `json:"name" form:"name" gorm:"column:name" binding:"omitempty,min=1,max=40" example:"腿部訓練"` // 訓練名稱
}
type EquipmentField struct {
	Equipment string `json:"equipment" form:"equipment" gorm:"column:equipment" binding:"omitempty,equipment,min=0,max=10" example:"2,3,6"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
type StartAudioField struct {
	StartAudio string `json:"start_audio" gorm:"column:start_audio" example:"123.mp3"` // 前導語音
}
type EndAudioField struct {
	EndAudio string `json:"end_audio" gorm:"column:end_audio" example:"123.mp3"` // 結束語音
}
type WorkoutSetCountField struct {
	WorkoutSetCount int `json:"workout_set_count" gorm:"column:workout_set_count" example:"10"` // 動作組數
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}
