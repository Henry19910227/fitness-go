package workoutdto

type Workout struct {
	ID  int64  `json:"column:id"`  // 訓練 id
	PlanID int64 `json:"column:plan_id"` // 計畫 id
	Name string `json:"column:name"` // 訓練名稱
	Equipment string `json:"column:equipment"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	StartAudio string `json:"column:start_audio"` // 前導語音
	EndAudio string `json:"column:end_audio"` // 結束語音
	WorkoutSetCount int `json:"column:workout_set_count"` // 動作組數
	CreateAt string `json:"column:create_at"` // 創建時間
	UpdateAt string `json:"column:update_at"` // 更新時間
}

type WorkoutID struct {
	ID int64  `json:"workout_id" example:"1"`  // 訓練 id
}