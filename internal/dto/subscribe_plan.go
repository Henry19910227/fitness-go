package dto

type SubscribePlan struct {
	ID   int64 `json:"id" example:"1"`  // 銷售id
	Period int `json:"period" example:"12"`  // 週期(1:一個月/2:二個月/3:三個月/6:六個月/12:一年/99:永久)
	Name string `json:"name" example:"銅級課表"` // 銷售名稱
	Twd  int `json:"twd" example:"330"` // 台幣價格
	ProductID string `json:"product_id" example:"com.fitness.xxx"` // 產品ID
}