package entity

type Food struct {
	ID        int64  `gorm:"column:id"`         //主鍵id
	UserID    *int64  `gorm:"column:user_id"`    //用戶id
	FoodCategoryID  int64 `gorm:"column:food_category_id"` //食物類別id
	Source int `gorm:"column:source"` //來源(1:系統創建食物/2:用戶創建食物)
	Name string  `gorm:"column:name"` //食物名稱
	Calorie int `gorm:"column:calorie"` //食物熱量
	AmountDesc string `gorm:"column:amount_desc"` //份量描述
	IsDeleted int    `gorm:"column:is_deleted"` //是否刪除
	CreateAt  string `gorm:"column:create_at"`  //創建日期
	UpdateAt  string `gorm:"column:update_at"`  //更新日期
}

func (Food) TableName() string {
	return "foods"
}