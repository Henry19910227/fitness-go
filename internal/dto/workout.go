package dto

type Workout struct {
	ID              int64  `json:"id" example:"1"`                       // 訓練 id
	Name            string `json:"name" example:"第一天胸肌訓練"`               // 訓練名稱
	Equipment       string `json:"equipment" example:"2,3,7"`            // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	StartAudio      string `json:"start_audio" example:"e6d2131w5q.mp3"` // 前導語音
	EndAudio        string `json:"end_audio" example:"d2e15qwe42dw.mp3"` // 結束語音
	WorkoutSetCount int    `json:"workout_set_count" example:"1"`        // 動作組數
}

type WorkoutProduct struct {
	ID              int64  `json:"id" example:"1"`                // 訓練 id
	Name            string `json:"name" example:"第一天胸肌訓練"`        // 訓練名稱
	Equipment       string `json:"equipment" example:"2,3,7"`     // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	WorkoutSetCount int    `json:"workout_set_count" example:"1"` // 動作組數
	Finish          int    `json:"finish" example:"1"`            // 是否完成(0:未完成/2:已完成)
}

type WorkoutAsset struct {
	ID              int64  `json:"id" example:"1"`                // 訓練 id
	Name            string `json:"name" example:"第一天胸肌訓練"`        // 訓練名稱
	Equipment       string `json:"equipment" example:"2,3,7"`     // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
	WorkoutSetCount int    `json:"workout_set_count" example:"1"` // 動作組數
	Finish          int    `json:"finish" example:"1"`            // 是否完成(0:未完成/2:已完成)
}

type WorkoutAudio struct {
	Named string `json:"audio" example:"e6d2131w5q.mp3"` // 語音檔案名
}

type WorkoutID struct {
	ID int64 `json:"workout_id" example:"1"` // 訓練 id
}

type UpdateWorkoutParam struct {
	Name      *string `json:"name" example:"第一天胸肌訓練"`    // 訓練名稱
	Equipment *string `json:"equipment" example:"2,3,7"` // 所需器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
