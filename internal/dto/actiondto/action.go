package actiondto

type Action struct {
	ID int64  `json:"id" example:"1"` //動作id
	Name string `json:"name" example:"槓鈴臥推"` //動作名稱
	Source int `json:"source" example:"2"` //動作來源(1:系統動作/2:教練自創動作)
	Type int `json:"type" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Category int `json:"category" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body int `json:"body" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment int `json:"equipment" example:"1"` //器材(1:槓鈴/2:啞鈴/3:長凳/4:機械/5:壺鈴/6:彎曲槓/7:自體體重運動/8:其他)
	Intro string `json:"intro" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹
	Cover string `json:"cover" example:"32as1d5f13e4.png"` //封面
	Video string `json:"video" example:"11d547we1d4f8e.mp4"` //動作影片
}

type CreateActionParam struct {
	Name string `json:"name" example:"槓鈴臥推"` //動作名稱
	Type int `json:"type" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Category int `json:"category" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body int `json:"body" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment int `json:"equipment" example:"1"` //器材(1:槓鈴/2:啞鈴/3:長凳/4:機械/5:壺鈴/6:彎曲槓/7:自體體重運動/8:其他)
	Intro string `json:"intro" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹
}