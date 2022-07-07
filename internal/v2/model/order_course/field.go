package order_course

type OrderIDField struct {
	OrderID *string `json:"order_id,omitempty" gorm:"column:order_id" example:"20220215104747115283"` //訂單id
}
type SaleItemIDField struct {
	SaleItemID *int64 `json:"sale_item_id,omitempty" gorm:"column:sale_item_id" example:"1"` //銷售項目 id
}
type CourseIDField struct {
	CourseID *int64 `json:"course_id,omitempty" gorm:"column:course_id" example:"10"` //課表id
}

type Table struct {
	OrderIDField
	SaleItemIDField
	CourseIDField
}

func (Table) TableName() string {
	return "order_courses"
}
