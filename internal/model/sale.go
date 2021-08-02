package model

type SaleSummaryEntity struct {
	ID   int64   // 銷售id
	Type int64   // 銷售類型(1:免費課表/2:訂閱課表/3:付費課表)
	Name string  // 銷售名稱
	Twd  float64 // 台幣價格
}
