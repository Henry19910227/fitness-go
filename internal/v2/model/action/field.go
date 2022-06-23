package action

type IDField struct {
	ID int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //動作id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" example:"10"` //課表id
}
type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name" example:"划船機"` //動作名稱
}
type SourceField struct {
	Source *int `json:"source,omitempty" gorm:"column:source" example:"2"` //動作來源(1:系統動作/2:教練自創動作)
}
type TypeField struct {
	Type *int `json:"type,omitempty" gorm:"column:type" example:"1"` //紀錄類型(1:重訓/2:時間長度/3:次數/4:次數與時間/5:有氧)
}
type CategoryField struct {
	Category *int `json:"category,omitempty" gorm:"column:category" example:"1"` //分類(1:重量訓練/2:有氧/3:HIIT/4:徒手訓練/5:其他)
}
type BodyField struct {
	Body *int `json:"body,omitempty" gorm:"column:body" example:"4"` //身體部位(1:全身/2:核心/3:手臂/4:背部/5:臀部/6:腿部/7:肩膀/8:胸部)
}
type EquipmentField struct {
	Equipment *int `json:"equipment,omitempty" gorm:"column:equipment" example:"4"` //器材(1:無需任何器材/2:啞鈴/3:槓鈴/4:固定式器材/5:彈力繩/6:壺鈴/7:訓練椅/8:瑜珈墊/9:其他)
}
type IntroField struct {
	Intro *string `json:"intro,omitempty" gorm:"column:intro" example:"全靠雙腳的力量將身體後蹬，雙手依舊輕鬆握住拉桿，上半身利用下背力量稍微挺直，直到雙腳完全伸直為止"` //動作介紹
}
type CoverField struct {
	Cover *string `json:"cover,omitempty" gorm:"column:cover" example:"1234.jpg"` //封面
}
type VideoField struct {
	Video *string `json:"video,omitempty" gorm:"column:video" example:"1234.mp4"` //動作影片
}
type IsDeletedField struct {
	IsDeleted *string `json:"is_deleted,omitempty" gorm:"column:is_deleted" example:"0"` //是否刪除
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}

type Table struct {
	IDField
	CourseIDField
	NameField
	SourceField
	TypeField
	CategoryField
	BodyField
	EquipmentField
	IntroField
	CoverField
	VideoField
	IsDeletedField
	CreateAtField
	UpdateAtField
}

func (Table) TableName() string {
	return "actions"
}