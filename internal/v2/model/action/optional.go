package action

type IDOptional struct {
	ID *int64 `json:"id,omitempty" example:"1"` //動作id
}
type NameOptional struct {
	Name *string `json:"name,omitempty" form:"name" binding:"omitempty,min=1,max=20" example:"划船機"` //動作名稱
}
type TypeOptional struct {
	Type *int `json:"type,omitempty" binding:"omitempty,oneof=1 2 3 4 5" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
}
type SourceOptional struct {
	Source *int `json:"source,omitempty"` //動作來源(1:系統動作/2:教練動作/2:學員動作
}
type CategoryOptional struct {
	Category *int `json:"category,omitempty" binding:"omitempty,oneof=1 2 3 4 5" example:"1234.mp4"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
}
type BodyOptional struct {
	Body *int `json:"body,omitempty" binding:"omitempty,oneof=1 2 3 4 5 6 7 8" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
}
type EquipmentOptional struct {
	Equipment *int `json:"equipment,omitempty" binding:"omitempty,oneof=1 2 3 4 5 6 7 8" example:"1"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
type IntroOptional struct {
	Intro *string `json:"intro,omitempty" form:"intro" binding:"omitempty,min=1,max=400" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹(1~400字元)
}
type StatusOptional struct {
	Status *int `json:"status,omitempty" form:"status" binding:"omitempty,oneof=0 1" example:"1"` //動作狀態(0:下架/1:上架)
}
type IsDeletedOptional struct {
	IsDeleted *int `json:"is_deleted,omitempty" form:"is_deleted" binding:"omitempty,oneof=0 1" example:"0"` //是否刪除(0:否/1:是)
}
