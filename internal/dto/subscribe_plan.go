package dto

import "github.com/Henry19910227/fitness-go/internal/model"

type SubscribePlan struct {
	ID        int64  `json:"id" example:"1"`                       // 銷售id
	Period    int    `json:"period" example:"12"`                  // 週期(1:一個月/2:二個月/3:三個月/6:六個月/12:一年/99:永久)
	Name      string `json:"name" example:"銅級課表"`                  // 銷售名稱
	Twd       int    `json:"twd" example:"330"`                    // 台幣價格
	ProductID string `json:"product_id" example:"com.fitness.xxx"` // 產品ID
}

func NewSubscribePlan(data *model.SubscribePlan) SubscribePlan {
	subscribePlan := SubscribePlan{
		ID:     data.ID,
		Period: data.Period,
	}
	if data.ProductLabel != nil {
		subscribePlan.ProductID = data.ProductLabel.ProductID
		subscribePlan.Name = data.ProductLabel.Name
		subscribePlan.Twd = data.ProductLabel.Twd
	}
	return subscribePlan
}
