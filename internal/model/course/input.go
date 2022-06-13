package course

import (
	orderBy "github.com/Henry19910227/fitness-go/internal/model/order_by"
	"github.com/Henry19910227/fitness-go/internal/model/paging"
	"github.com/Henry19910227/fitness-go/internal/model/preload"
)

type PagingInput = paging.Input
type PreloadInput = preload.Input
type OrderByInput = orderBy.Input

type FindInput struct {
	IDField
	PreloadInput
}

type ListInput struct {
	IDField
	NameField
	CourseStatusField
	SaleTypeField
	PagingInput
	PreloadInput
	OrderByInput
}

type APIGetCMSCoursesInput struct {
	IDField
	NameField
	CourseStatusField
	SaleTypeField
	PagingInput
	OrderByInput
}

type APIGetCMSCourseInput struct {
	IDField
}
