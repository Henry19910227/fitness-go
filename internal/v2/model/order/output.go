package order

import (
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/product_label"
	"github.com/Henry19910227/fitness-go/internal/v2/model/sale_item"
	"github.com/Henry19910227/fitness-go/internal/v2/model/subscribe_plan"
)

type Output struct {
	Table
	OrderCourse        *order_course.Output         `json:"order_course,omitempty" gorm:"foreignKey:order_id;references:id"`
	OrderSubscribePlan *order_subscribe_plan.Output `json:"order_subscribe_plan,omitempty" gorm:"foreignKey:order_id;references:id"`
}

func (Output) TableName() string {
	return "orders"
}

// APICreateCourseOrderOutput /v2/course_order [POST]
type APICreateCourseOrderOutput struct {
	base.Output
	Data *APICreateCourseOrderData `json:"data,omitempty"`
}
type APICreateCourseOrderData struct {
	Table
	OrderCourse *struct {
		Course *struct {
			course.IDField
			course.NameField
		} `json:"course,omitempty"`
		SaleItem *struct {
			sale_item.IDField
			sale_item.TypeField
			ProductLabel *struct {
				product_label.IDField
				product_label.NameField
				product_label.ProductIDField
				product_label.TwdField
			} `json:"product_label,omitempty"`
		} `json:"sale_item,omitempty"`
	} `json:"order_course,omitempty"`
}

// APIGetCMSOrdersOutput /v2/cms/orders [GET]
type APIGetCMSOrdersOutput struct {
	base.Output
	Data   APIGetCMSOrdersData `json:"data"`
	Paging *paging.Output      `json:"paging,omitempty"`
}
type APIGetCMSOrdersData []*struct {
	IDField
	UserIDField
	QuantityField
	OrderTypeField
	OrderStatusField
	CreateAtField
	UpdateAtField
	OrderCourse *struct {
		Course *struct {
			course.IDField
			course.NameField
		} `json:"course,omitempty"`
		SaleItem *struct {
			sale_item.IDField
			sale_item.TypeField
			ProductLabel *struct {
				product_label.IDField
				product_label.NameField
				product_label.ProductIDField
				product_label.TwdField
			} `json:"product_label,omitempty"`
		} `json:"sale_item,omitempty"`
	} `json:"order_course,omitempty"`
	OrderSubscribePlan *struct {
		SubscribePlan *struct {
			subscribe_plan.IDField
			subscribe_plan.PeriodField
			subscribe_plan.NameField
			ProductLabel *struct {
				product_label.IDField
				product_label.NameField
				product_label.ProductIDField
				product_label.TwdField
			} `json:"product_label,omitempty"`
		} `json:"subscribe_plan,omitempty"`
	} `json:"order_subscribe_plan,omitempty"`
}
