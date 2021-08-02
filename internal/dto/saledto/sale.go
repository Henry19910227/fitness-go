package saledto

type SaleSummary struct {
	ID   int64 `json:"id" example:"1"`    // 銷售id
	Type int64  `json:"type" example:"3"`   // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Name string `json:"name" example:"銅級課表"`   // 銷售名稱
	Price  float64 `json:"price" example:"330"`   // 台幣價格
}