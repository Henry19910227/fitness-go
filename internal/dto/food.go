package dto

type Food struct {
	ID        int64  `json:"id" example:"1"`         //主鍵id
	UserID    *int64  `json:"user_id,omitempty" example:"10001"`    //創建用戶id
	Source int `json:"source" example:"1"` //來源(1:系統創建食物/2:用戶創建食物)
	Name string  `json:"name" example:"蕃薯"` //食物名稱
	Calorie int `json:"calorie" example:"70"` //食物熱量
	AmountDesc string `json:"amount_desc" example:"一份地瓜100克70卡"` //份量描述
	FoodCategory *FoodCategory `json:"food_category,omitempty"`  //食物分類
}

type CreateFoodParam struct {
	UserID  int64 //用戶id
	FoodCategoryID  int64 //食物類別id
	Source int //來源(1:系統創建食物/2:用戶創建食物)
	Name string //食物名稱
	AmountDesc string //份量描述
}