package required

type TotalCountField struct {
	TotalCount int `json:"total_count" binding:"required" example:"100"` // 總筆數
}
type TotalPageField struct {
	TotalPage int `json:"total_page" binding:"required" example:"10"` // 總頁數
}
type PageField struct {
	Page int `json:"page" form:"page" binding:"required,min=1" example:"1"` // 當前頁數
}
type SizeField struct {
	Size int `json:"size" form:"size" binding:"required,min=1,max=100" example:"5"` // 一頁筆數
}
