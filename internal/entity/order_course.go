package entity

type OrderCourse struct {
	OrderID    string `gorm:"column:order_id"`     // 訂單id
	SaleItemID *int64 `gorm:"column:sale_item_id"` // 銷售項目 id
	CourseID   int64  `gorm:"column:course_id"`    // 課表id
}

func (OrderCourse) TableName() string {
	return "order_courses"
}

type OrderCourseTemplate struct {
	OrderID    string          `gorm:"column:order_id"`                    // 訂單id
	SaleItemID *int64          `gorm:"column:sale_item_id"`                // 銷售項目 id
	CourseID   int64           `gorm:"column:course_id"`                   // 課表id
	Course     *CourseTemplate `gorm:"foreignKey:id;references:course_id"` // 課表id
}

func (OrderCourseTemplate) TableName() string {
	return "order_courses"
}
