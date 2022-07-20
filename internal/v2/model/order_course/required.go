package order_course

type OrderIDRequired struct {
	OrderID string `json:"order_id" binding:"required" example:"20220215104747115283"` //訂單id
}
type SaleItemIDRequired struct {
	SaleItemID int64 `json:"sale_item_id" binding:"required" example:"1"` //銷售項目 id
}
type CourseIDRequired struct {
	CourseID int64 `json:"course_id" binding:"required" example:"10"` //課表id
}
