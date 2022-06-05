package model

type Food struct {
	ID        int64  `json:"id" gorm:"column:id"`         //主鍵id
	UserID    *int64  `json:"user_id,omitempty" gorm:"column:user_id"`    //用戶id
	FoodCategoryID  int64 `json:"food_category_id" gorm:"column:food_category_id"` //食物類別id
	Source int `json:"source" gorm:"column:source"` //來源(1:系統創建食物/2:用戶創建食物)
	Name string  `json:"name" gorm:"column:name"` //食物名稱
	Calorie int `json:"calorie" gorm:"column:calorie"` //食物熱量
	AmountDesc string `json:"amount_desc" gorm:"column:amount_desc"` //份量描述
	IsDeleted int    `json:"is_deleted" gorm:"column:is_deleted"` //是否刪除
	CreateAt  string `json:"create_at" gorm:"column:create_at"`  //創建日期
	UpdateAt  string `json:"update_at" gorm:"column:update_at"`  //更新日期
	FoodCategory *FoodCategory `json:"food_category,omitempty" gorm:"foreignkey:id;references:food_category_id"`  //食物分類
}

func (Food) TableName() string {
	return "foods"
}

type CreateFoodParam struct {
	UserID  *int64 //用戶id
	FoodCategoryID  int64 //食物類別id
	Source int //來源(1:系統創建食物/2:用戶創建食物)
	Name string //食物名稱
	Calorie int //食物熱量
	AmountDesc string //份量描述
}

type FindFoodsParam struct {
	DeletedParam
	PreloadParam
	UserID *int64
	Tag *int
}

type UpdateFoodParam struct {
	FoodID    int64
	IsDeleted *int `gorm:"column:is_deleted"`
	UpdateAt  *string `gorm:"column:update_at"`  //更新日期
}

