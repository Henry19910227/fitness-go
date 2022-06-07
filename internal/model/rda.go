package model

type RDA struct {
	ID        int64  `json:"id" gorm:"column:id"`               // id
	UserID    int64  `json:"user_id" gorm:"column:user_id"`     // 用戶id
	TDEE      int    `json:"tdee" gorm:"column:tdee"`           // TDEE
	Calorie   int    `json:"calorie" gorm:"column:calorie"`     // 目標熱量
	Protein   int    `json:"protein" gorm:"column:protein"`     // 蛋白質(克)
	Fat       int    `json:"fat" gorm:"column:fat"`             // 脂肪(克)
	Carbs     int    `json:"carbs" gorm:"column:carbs"`         // 碳水化合物(克)
	Grain     int    `json:"grain" gorm:"column:grain"`         // 穀物類(份)
	Vegetable int    `json:"vegetable" gorm:"column:vegetable"` // 蔬菜類(份)
	Fruit     int    `json:"fruit" gorm:"column:fruit"`         // 水果類(份)
	Meat      int    `json:"meat" gorm:"column:meat"`           // 蛋豆魚肉類(份)
	Dairy     int    `json:"dairy" gorm:"column:dairy"`         // 乳製品類(份)
	Nut       int    `json:"nut" gorm:"column:nut"`             // 堅果類(份)
	CreateAt  string `json:"create_at" gorm:"create_at"`        // 創建時間
}

func (RDA) TableName() string {
	return "rdas"
}

type CreateRDAParam struct {
	TDEE      int `gorm:"column:tdee"`      // TDEE
	Calorie   int `gorm:"column:calorie"`   // 目標熱量
	Protein   int `gorm:"column:protein"`   // 蛋白質(克)
	Fat       int `gorm:"column:fat"`       // 脂肪(克)
	Carbs     int `gorm:"column:carbs"`     // 碳水化合物(克)
	Grain     int `gorm:"column:grain"`     // 穀物類(份)
	Vegetable int `gorm:"column:vegetable"` // 蔬菜類(份)
	Fruit     int `gorm:"column:fruit"`     // 水果類(份)
	Meat      int `gorm:"column:meat"`      // 蛋豆魚肉類(份)
	Dairy     int `gorm:"column:dairy"`     // 乳製品類(份)
	Nut       int `gorm:"column:nut"`       // 堅果類(份)
}

type FindRDAParam struct {
	ID     *int64
	UserID *int64
}
