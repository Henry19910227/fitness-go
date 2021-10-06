package validator

type GetCourseProductsQuery struct {
	OrderType string `form:"order_type" binding:"required,oneof=latest popular" example:"latest"` // 排序類型(latest:最新/popular:熱門)
	Page int `form:"page" binding:"required,min=1" example:"1"` // 頁數
	Size int `form:"size" binding:"required,min=1" example:"5"` // 筆數
}
