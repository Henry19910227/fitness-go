package dto

type FoodCategory struct {
	ID    int64  `json:"id" example:"1"`      //主鍵id
	Tag   int    `json:"tag" example:"1"`     //食物六大類Tag(1:全穀雜糧/2:蛋豆魚肉/3:水果/4:蔬菜/5:乳製品/6:油脂堅果)
	Title string `json:"title" example:"米麥類"` //類別名稱
}

