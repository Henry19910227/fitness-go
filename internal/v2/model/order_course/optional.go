package order_course

type OrderIDOptional struct {
	OrderID *string `json:"order_id,omitempty" example:"20220215104747115283"` //訂單id
}
type SaleItemIDOptional struct {
	SaleItemID *int64 `json:"sale_item_id,omitempty" example:"1"` //銷售項目 id
}
type CourseIDOptional struct {
	CourseID *int64 `json:"course_id,omitempty" example:"10"` //課表id
}
