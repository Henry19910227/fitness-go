package required

type OrderIDField struct {
	OrderID string `json:"order_id" gorm:"column:order_id" binding:"required" example:"20220215104747115283"` //訂單id
}
type SaleItemIDField struct {
	SaleItemID int64 `json:"sale_item_id" gorm:"column:sale_item_id" binding:"required" example:"1"` //銷售項目 id
}
type CourseIDField struct {
	CourseID int64 `json:"course_id" gorm:"column:course_id" binding:"required" example:"10"` //課表id
}
