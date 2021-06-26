package validator

type ActionIDUri struct {
	ActionID int64 `uri:"action_id" binding:"required" example:"1"`
}

type CreateActionBody struct {
	Name string `json:"name" binding:"required,min=1,max=20" example:"槓鈴臥推"` //動作名稱(1~20字元)
	Type int `json:"type" binding:"required,oneof=1 2 3 4 5" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
	Category int `json:"category" binding:"required,oneof=1 2 3 4 5" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body int `json:"body" binding:"required,oneof=1 2 3 4 5 6 7 8" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment int `json:"equipment" binding:"required,oneof=1 2 3 4 5 6 7 8" example:"1"` //器材(1:槓鈴/2:啞鈴/3:長凳/4:機械/5:壺鈴/6:彎曲槓/7:自體體重運動/8:其他)
	Intro string `json:"intro" binding:"required,min=1,max=400" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹(1~400字元)
}

type UpdateActionBody struct {
	Name *string `json:"name" binding:"omitempty,min=1,max=20" example:"槓鈴臥推"` //動作名稱(1~20字元)
	Category *int `json:"category" binding:"omitempty,oneof=1 2 3 4 5" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
	Body *int `json:"body" binding:"omitempty,oneof=1 2 3 4 5 6 7 8" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
	Equipment *int `json:"equipment" binding:"omitempty,oneof=1 2 3 4 5 6 7 8" example:"1"` //器材(1:槓鈴/2:啞鈴/3:長凳/4:機械/5:壺鈴/6:彎曲槓/7:自體體重運動/8:其他)
	Intro *string `json:"intro" binding:"omitempty,min=1,max=400" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹(1~400字元)
}
