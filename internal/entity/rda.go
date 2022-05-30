package entity

type RDA struct {
	ID int64 `gorm:"column:id primaryKey"` // id
	UserID int64 `gorm:"column:user_id"` // 用戶id
	TDEE int `gorm:"column:tdee"` // TDEE
	Calorie int `gorm:"column:calorie"` // 目標熱量
	Protein int `gorm:"column:protein"` // 蛋白質(克)
	Fat int `gorm:"column:fat"` // 脂肪(克)
	Carbs int `gorm:"column:carbs"` // 碳水化合物(克)
	Grain int `gorm:"column:grain"` // 穀物類(份)
	Vegetable int `gorm:"column:vegetable"` // 蔬菜類(份)
	Fruit int `gorm:"column:fruit"` // 水果類(份)
	Meat int `gorm:"column:meat"` // 蛋豆魚肉類(份)
	Dairy int `gorm:"column:dairy"` // 乳製品類(份)
	Nut int `gorm:"column:nut"` // 堅果類(份)
	CreateAt string `gorm:"create_at"` // 創建時間
}

func (RDA) TableName() string {
	return "rdas"
}
