package action

type IDRequired struct {
	ID int64 `json:"id" uri:"action_id" binding:"required" example:"1"` //動作id
}
type NameRequired struct {
	Name string `json:"name" form:"name" binding:"required,min=1,max=20" example:"划船機"` //動作名稱
}
type TypeRequired struct {
	Type int `json:"type" form:"type" binding:"required,oneof=1 2 3 4 5" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
}
type CategoryRequired struct {
	Category int `json:"category" form:"category" binding:"required,oneof=1 2 3 4 5" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
}
type BodyRequired struct {
	Body int `json:"body" form:"body"  binding:"required,oneof=1 2 3 4 5 6 7 8" example:"1"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
}
type EquipmentRequired struct {
	Equipment int `json:"equipment" form:"equipment" binding:"required,oneof=1 2 3 4 5" binding:"required,oneof=1 2 3 4 5 6 7 8 9" example:"1"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
type IntroRequired struct {
	Intro string `json:"intro" form:"intro" binding:"required,min=1,max=400" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹(1~400字元)
}
