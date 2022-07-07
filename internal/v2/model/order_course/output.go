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
