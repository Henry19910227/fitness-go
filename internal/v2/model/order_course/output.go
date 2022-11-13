package order_course

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
)

type Output struct {
	Table
	Course   *course.Output    `json:"course,omitempty" gorm:"foreignKey:id;references:course_id"`       // 課表
	SaleItem *sale_item.Output `json:"sale_item,omitempty" gorm:"foreignKey:id;references:sale_item_id"` // 銷售項目
}

func (Output) TableName() string {
	return "order_courses"
}

func (o Output) CourseOnSafe() course.Output {
	if o.Course != nil {
		return *o.Course
	}
	return course.Output{}
}

func (o Output) SaleItemOnSafe() sale_item.Output {
	if o.SaleItem != nil {
		return *o.SaleItem
	}
	return sale_item.Output{}
}
