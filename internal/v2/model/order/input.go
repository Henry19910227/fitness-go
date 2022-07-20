package order

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/v2/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/v2/model/order_course"
	"github.com/Henry19910227/fitness-go/internal/v2/model/paging"
	"github.com/Henry19910227/fitness-go/internal/v2/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	IDOptional
	PreloadInput
}

type ListInput struct {
	IDOptional
	UserIDOptional
	OrderTypeOptional
	OrderStatusOptional
	order_course.CourseIDOptional
	OrderByInput
	PagingInput
	PreloadInput
}

// APICreateCourseOrderInput /v2/course_order
type APICreateCourseOrderInput struct {
	UserIDRequired
	Body APICreateCourseOrderBody

}
type APICreateCourseOrderBody struct {
	order_course.CourseIDRequired
}

// APIGetCMSOrdersInput /v2/cms/orders [GET]
type APIGetCMSOrdersInput struct {
	Form APIGetCMSOrdersForm
}
type APIGetCMSOrdersForm struct {
	IDOptional
	UserIDOptional
	OrderTypeOptional
	OrderStatusOptional
	PagingInput
	OrderByInput
}
