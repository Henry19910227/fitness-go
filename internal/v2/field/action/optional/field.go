package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" uri:"action_id" gorm:"column:id" binding:"omitempty" example:"1"` //動作id
}
type UserIDField struct {
	UserID *int64 `json:"user_id,omitempty" gorm:"column:user_id" binding:"omitempty" example:"10001"` //用戶id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" uri:"course_id" gorm:"column:course_id" binding:"omitempty" example:"10"` //課表id
}
type NameField struct {
	Name *string `json:"name,omitempty" form:"name" gorm:"column:name" binding:"omitempty,min=1,max=20" example:"划船機"` //動作名稱
}
type TypeField struct {
	Type *int `json:"type,omitempty" form:"type" gorm:"column:type" binding:"omitempty,oneof=1 2 3 4 5" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
}
type SourceField struct {
	Source *int `json:"source,omitempty" form:"source" gorm:"column:source" binding:"omitempty,oneof=1 2 3" example:"1"` //動作來源(1:系統動作/2:教練動作/2:學員動作)
}
type CategoryField struct {
	Category *int `json:"category,omitempty" form:"category" form:"category" gorm:"column:category" binding:"omitempty,oneof=1 2 3 4 5" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
}
type BodyField struct {
	Body *int `json:"body,omitempty" form:"body" gorm:"column:body" binding:"omitempty,oneof=1 2 3 4 5 6 7 8" example:"8"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
}
type EquipmentField struct {
	Equipment *int `json:"equipment,omitempty" form:"equipment" gorm:"column:equipment" binding:"omitempty,oneof=1 2 3 4 5 6 7 8 9" example:"1"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
type IntroField struct {
	Intro *string `json:"intro,omitempty" form:"intro" gorm:"column:intro" binding:"omitempty,min=1,max=400" example:"槓鈴胸推是很多人在健身房都會訓練的動作，是胸大肌強化最常見的訓練動作"` //動作介紹(1~400字元)
}
type CoverField struct {
	Cover *string `json:"cover,omitempty" form:"cover" gorm:"column:cover" binding:"omitempty" example:"1234.jpg"` //封面
}
type VideoField struct {
	Video *string `json:"video,omitempty" form:"video" gorm:"column:video" binding:"omitempty" example:"1234.mp4"` //動作影片
}
type StatusField struct {
	Status *int `json:"status,omitempty" form:"status" gorm:"column:status" binding:"omitempty,oneof=0 1" example:"1"` //動作狀態(0:下架/1:上架)
}
type IsDeletedField struct {
	IsDeleted *int `json:"is_deleted,omitempty" form:"is_deleted" gorm:"column:is_deleted" binding:"omitempty,oneof=0 1" example:"0"` //是否刪除(0:否/1:是)
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" binding:"omitempty" example:"2022-06-14 00:00:00"` //更新時間
}
