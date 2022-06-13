package entity

type FoodCategory struct {
	ID        int64  `gorm:"column:id"`         //主鍵id
	Tag       int    `gorm:"column:tag"`        //食物六大類Tag(1:全穀雜糧/2:蛋豆魚肉/3:水果/4:蔬菜/5:乳製品/6:油脂堅果)
	Title     string `gorm:"column:title"`      //類別名稱
	IsDeleted int    `gorm:"column:is_deleted"` //是否刪除
	CreateAt  string `gorm:"column:create_at"`  //創建日期
	UpdateAt  string `gorm:"column:update_at"`  //更新日期
}

func (FoodCategory) TableName() string {
	return "food_categories"
}
