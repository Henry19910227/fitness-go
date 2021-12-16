package dto

type SaleItem struct {
	ID   int64 `json:"id" example:"1"`  // 銷售id
	Type int `json:"type" example:"3"`  // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Name string `json:"name" example:"銅級課表"` // 銷售名稱
	Twd  int `json:"twd" example:"330"` // 台幣價格
	Identifier string `json:"identifier" example:"com.fitness.xxx"` // 銷售識別碼
}