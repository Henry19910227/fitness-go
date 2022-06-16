package food_category

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id" example:"1"` //主鍵id
}
type TagField struct {
	Tag *int `json:"tag,omitempty" form:"tag" gorm:"column:tag" example:"2"` //食物六大類Tag(1:全穀雜糧/2:蛋豆魚肉/3:水果/4:蔬菜/5:乳製品/6:油脂堅果)
}
type TitleField struct {
	Title *string `json:"title,omitempty" gorm:"column:title" example:"米麥類"` //類別名稱
}
type IsDeletedField struct {
	IsDeleted *int `json:"is_deleted,omitempty" gorm:"column:is_deleted" example:"0"` //是否刪除
}
type CreateAtField struct {
	CreateAt *string `json:"create_at,omitempty" gorm:"column:create_at" example:"2022-06-14 00:00:00"` //創建時間
}
type UpdateAtField struct {
	UpdateAt *string `json:"update_at,omitempty" gorm:"column:update_at" example:"2022-06-14 00:00:00"` //更新時間
}
