package optional

type TotalCountField struct {
	TotalCount *int `json:"total_count,omitempty" binding:"omitempty" example:"100"` // 總筆數
}
type TotalPageField struct {
	TotalPage *int `json:"total_page,omitempty" binding:"omitempty" example:"10"` // 總頁數
}
type PageField struct {
	Page *int `json:"page,omitempty" form:"page" binding:"omitempty,min=1" example:"1"` // 當前頁數
}
type SizeField struct {
	Size *int `json:"size,omitempty" form:"size" binding:"omitempty,min=1,max=100" example:"5"` // 一頁筆數
}
