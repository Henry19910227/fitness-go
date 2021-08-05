package dto

type Action struct {
	ID int64  `json:"id" example:"1"` //動作id
	Name string `json:"name" example:"槓鈴臥推"` //動作名稱
	Source int `json:"source" example:"2"` //動作來源(1:系統動作/2:教練自創動作)
	Type int `json:"type" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Category int `json:"category" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body int `json:"body" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment int `json:"equipment" example:"1"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:其他)
	Intro string `json:"intro" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹
	Cover string `json:"cover" example:"32as1d5f13e4.png"` //封面
	Video string `json:"video" example:"11d547we1d4f8e.mp4"` //動作影片
}

type CreateActionParam struct {
	Name string `json:"name" example:"槓鈴臥推"` //動作名稱
	Type int `json:"type" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Category int `json:"category" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body int `json:"body" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment int `json:"equipment" example:"1"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:其他)
	Intro string `json:"intro" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹
}

type UpdateActionParam struct {
	Name *string `gorm:"column:name"` //課表名稱
	Category *int `gorm:"column:category"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body *int `gorm:"column:body"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment *int `gorm:"column:equipment"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:其他)
	Intro *string `gorm:"column:intro"` //動作介紹
}

type FindActionsParam struct {
	Name *string //課表名稱
	Source *string //動作來源(1:平台動作/2:教練動作)
	Category *string //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body *string //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment *string //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:其他)
}

type ActionID struct {
	ID int64 `json:"action_id" example:"1"` //動作id
}

type ActionCover struct {
	Cover string `json:"cover" example:"kd3kf54ew5.png"` //封面圖片
}

type ActionVideo struct {
	Video string `json:"video" example:"f5e23q5e45fe32.mp4"` //動作影片
}