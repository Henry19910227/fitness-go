package dto

type Paging struct {
	TotalCount int `json:"total_count" example:"100"` // 總筆數
	TotalPage  int `json:"total_page" example:"10"`   // 總頁數
	Page       int `json:"page" example:"1"`          // 當前頁數
	Size       int `json:"size" example:"5"`          // 一頁筆數
}

type PagingParam struct {
	Page int
	Size int
}

type OrderByParam struct {
	OrderField *string
	OrderType  *string
}
