package food_category

type IDOptional struct {
	ID *int64 `json:"id,omitempty" binding:"omitempty" example:"1"` //主鍵id
}
type TagOptional struct {
	Tag *int `json:"tag,omitempty" binding:"omitempty" example:"2"` //食物六大類Tag(1:全穀雜糧/2:蛋豆魚肉/3:水果/4:蔬菜/5:乳製品/6:油脂堅果)
}
type TitleOptional struct {
	Title *string `json:"title,omitempty" binding:"omitempty" example:"米麥類"` //類別名稱
}
type IsDeletedOptional struct {
	IsDeleted *int `json:"is_deleted,omitempty" binding:"omitempty" example:"0"` //是否刪除
}
