package model

type Workout struct {
	ID  int64  `gorm:"column:id"`  // 訓練 id
	PlanID int64 `gorm:"column:plan_id"` // 計畫 id
	Name string `gorm:"column:name"` // 訓練名稱
	Equipment string `gorm:"column:equipment"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	StartAudio string `gorm:"column:start_audio"` // 前導語音
	EndAudio string `gorm:"column:end_audio"` // 結束語音
	WorkoutSetCount int `gorm:"column:workout_set_count"` // 動作組數
	CreateAt string `gorm:"column:create_at"` // 創建時間
	UpdateAt string `gorm:"column:update_at"` // 更新時間
}

func (Workout) TableName() string {
	return "workouts"
}

