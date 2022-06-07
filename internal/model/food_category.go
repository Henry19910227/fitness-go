package model

type FoodCategory struct {
	ID       int64  `json:"id" gorm:"column:id"`               //主鍵id
	Tag      int    `json:"tag" gorm:"column:tag"`             //食物六大類Tag(1:全穀雜糧/2:蛋豆魚肉/3:水果/4:蔬菜/5:乳製品/6:油脂堅果)
	Title    string `json:"title" gorm:"column:title"`         //類別名稱
	CreateAt string `json:"create_at" gorm:"column:create_at"` //創建日期
	UpdateAt string `json:"update_at" gorm:"column:update_at"` //更新日期
}

func (FoodCategory) TableName() string {
	return "food_categories"
}
