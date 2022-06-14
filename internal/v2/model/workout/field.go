package workout

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id"` // 訓練 id
}
type PlanIDField struct {
	PlanID *int64 `json:"plan_id,omitempty" gorm:"column:plan_id"` // 計畫 id
}
type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name"` // 訓練名稱
}
type EquipmentField struct {
	Equipment *string `json:"equipment,omitempty" gorm:"column:equipment"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
type StartAudioField struct {
	StartAudio *string `json:"start_audio,omitempty" gorm:"column:start_audio"` // 前導語音
}
type EndAudioField struct {
	EndAudio *string `json:"end_audio,omitempty" gorm:"column:end_audio"` // 結束語音
}
type WorkoutSetCountField struct {
	WorkoutSetCount *int `json:"workout_set_count,omitempty" gorm:"column:workout_set_count"` // 動作組數
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}
