package workout

type IDOptional struct {
	ID *int64 `json:"id,omitempty" example:"1"` // 訓練 id
}
type PlanIDOptional struct {
	PlanID *int64 `json:"plan_id,omitempty" example:"1"` // 計畫 id
}
type NameOptional struct {
	Name *string `json:"name,omitempty" example:"腿部訓練"` // 訓練名稱
}
type EquipmentOptional struct {
	Equipment *string `json:"equipment,omitempty" example:"2,3,6"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
type StartAudioOptional struct {
	StartAudio *string `json:"start_audio,omitempty" example:"123.mp3"` // 前導語音
}
type EndAudioOptional struct {
	EndAudio *string `json:"end_audio,omitempty" example:"123.mp3"` // 結束語音
}
