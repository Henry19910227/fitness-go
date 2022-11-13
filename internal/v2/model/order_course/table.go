package order_course

import "github.com/Henry19910227/fitness-go/internal/v2/field/order_course/optional"

type Table struct {
	optional.OrderIDField
	optional.SaleItemIDField
	optional.CourseIDField
}

func (Table) TableName() string {
	return "order_courses"
}
