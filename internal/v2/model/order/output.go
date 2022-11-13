package order

import (
	courseOptional "github.com/Henry19910227/fitness-go/internal/v2/field/course/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/field/order/optional"
	productLabelOptional "github.com/Henry19910227/fitness-go/internal/v2/field/product_label/optional"
	saleItemOptional "github.com/Henry19910227/fitness-go/internal/v2/field/sale_item/optional"
	subscribePlanOptional "github.com/Henry19910227/fitness-go/internal/v2/field/subscribe_plan/optional"
	"github.com/Henry19910227/fitness-go/internal/v2/model/base"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_subscribe_plan"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
)

type Output struct {
	Table
	OrderCourse        *order_course.Output         `json:"order_course,omitempty" gorm:"foreignKey:order_id;references:id"`
	OrderSubscribePlan *order_subscribe_plan.Output `json:"order_subscribe_plan,omitempty" gorm:"foreignKey:order_id;references:id"`
}

func (Output) TableName() string {
	return "orders"
}

func (o Output) OrderCourseOnSafe() order_course.Output {
	if o.OrderCourse != nil {
		return *o.OrderCourse
	}
	return order_course.Output{}
}

func (o Output) OrderSubscribePlanOnSafe() order_subscribe_plan.Output {
	if o.OrderSubscribePlan != nil {
		return *o.OrderSubscribePlan
	}
	return order_subscribe_plan.Output{}
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
			courseOptional.IDField
			courseOptional.NameField
		} `json:"course,omitempty"`
		SaleItem *struct {
			saleItemOptional.IDField
			saleItemOptional.TypeField
			ProductLabel *struct {
				productLabelOptional.IDField
				productLabelOptional.NameField
				productLabelOptional.ProductIDField
				productLabelOptional.TwdField
			} `json:"product_label,omitempty"`
		} `json:"sale_item,omitempty"`
	} `json:"order_course,omitempty"`
}

// APICreateSubscribeOrderOutput /v2/subscribe_order [POST]
type APICreateSubscribeOrderOutput struct {
	base.Output
	Data *APICreateSubscribeOrderData `json:"data,omitempty"`
}
type APICreateSubscribeOrderData struct {
	Table
	OrderSubscribePlan *struct {
		SubscribePlan *struct {
			subscribePlanOptional.IDField
			subscribePlanOptional.PeriodField
			subscribePlanOptional.NameField
			ProductLabel *struct {
				productLabelOptional.IDField
				productLabelOptional.NameField
				productLabelOptional.ProductIDField
				productLabelOptional.TwdField
			} `json:"product_label,omitempty"`
		} `json:"subscribe_plan,omitempty"`
	} `json:"order_subscribe_plan,omitempty"`
}

// APIGetCMSOrdersOutput /v2/cms/orders [GET]
type APIGetCMSOrdersOutput struct {
	base.Output
	Data   APIGetCMSOrdersData `json:"data"`
	Paging *paging.Output      `json:"paging,omitempty"`
}
type APIGetCMSOrdersData []*struct {
	optional.IDField
	optional.UserIDField
	optional.QuantityField
	optional.OrderTypeField
	optional.OrderStatusField
	optional.CreateAtField
	optional.UpdateAtField
	OrderCourse *struct {
		Course *struct {
			courseOptional.IDField
			courseOptional.NameField
		} `json:"course,omitempty"`
		SaleItem *struct {
			saleItemOptional.IDField
			saleItemOptional.TypeField
			ProductLabel *struct {
				productLabelOptional.IDField
				productLabelOptional.NameField
				productLabelOptional.ProductIDField
				productLabelOptional.TwdField
			} `json:"product_label,omitempty"`
		} `json:"sale_item,omitempty"`
	} `json:"order_course,omitempty"`
	OrderSubscribePlan *struct {
		SubscribePlan *struct {
			subscribePlanOptional.IDField
			subscribePlanOptional.PeriodField
			subscribePlanOptional.NameField
			ProductLabel *struct {
				productLabelOptional.IDField
				productLabelOptional.NameField
				productLabelOptional.ProductIDField
				productLabelOptional.TwdField
			} `json:"product_label,omitempty"`
		} `json:"subscribe_plan,omitempty"`
	} `json:"order_subscribe_plan,omitempty"`
}

// APIVerifyAppleReceiptOutput /v2/verify_apple_receipt [POST]
type APIVerifyAppleReceiptOutput struct {
	base.Output
}

// APIVerifyGoogleReceiptOutput /v2/verify_google_receipt [POST]
type APIVerifyGoogleReceiptOutput struct {
	base.Output
}

// APIAppStoreNotificationOutput /v2/app_store_notification/v2 [POST]
type APIAppStoreNotificationOutput struct {
	base.Output
}

// APIGooglePlayNotificationOutput /v2/google_play_notification [POST]
type APIGooglePlayNotificationOutput struct {
	base.Output
}

// APIVerifyAppleSubscribeOutput /v2/verify_apple_payment [POST]
type APIVerifyAppleSubscribeOutput struct {
	base.Output
}
